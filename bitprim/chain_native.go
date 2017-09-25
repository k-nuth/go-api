/**
 * Copyright (c) 2017 Bitprim developers (see AUTHORS)
 *
 * This file is part of Bitprim.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

// --------------------------------
// Interface one-to-one with C Interface
// --------------------------------

package bitprim

/*
#include <stdlib.h>
#include <bitprim/nodecint/chain/chain.h>
#include <bitprim/nodecint/chain/payment_address.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// func CHashToGo(hashCPtr C.hash_t) HashT {
// 	hashC := unsafe.Pointer(hashCPtr)

// 	hashGoSlice := C.GoBytes(hashC, 32)
// 	var hash HashT
// 	copy(hash[:], hashGoSlice)
// 	return hash
// }

// func boolToC(x bool) C.int {
// 	if x {
// 		return 1
// 	}
// 	return 0
// }

// func CToBool(x C.int) bool {
// 	if x == 0 {
// 		return false
// 	}
// 	return true
// }

// --------------------------------
// GetLastHeight
// --------------------------------

func GetLastHeight(chain unsafe.Pointer) (int, int) {
	fmt.Println("GetLastHeight - 1")
	ptr := (C.chain_t)(chain)

	fmt.Println("GetLastHeight - 2")

	var outHeight C.uint64_t
	fmt.Println("GetLastHeight - 3")

	res := C.chain_get_last_height(ptr, &outHeight)
	fmt.Println("GetLastHeight - 4")

	return int(res), int(outHeight)
}

// --------------------------------
// GetBlockHeight
// --------------------------------

func GetBlockHeight(chain unsafe.Pointer, hash HashT) (int, int) {
	// ptr := (C.chain_t)(chain)
	// hashC := C.CBytes(hash[:])
	// defer C.free(hashC)
	// var outHeight C.uint64_t
	// res := C.chain_get_block_height(ptr, (*C.uint8_t)(hashC), &outHeight)
	// return int(res), int(outHeight)

	ptr := (C.chain_t)(chain)

	// var hashC C.struct_hash_t
	// hashCTemp := C.CBytes(hash[:])
	// defer C.free(hashCTemp)
	// C.to_hash_t((*C.struct_hash_t)(&hashC), hashCTemp)

	var outHeight C.uint64_t
	res := C.chain_get_block_height(ptr, GoHashToC(hash), &outHeight)
	return int(res), int(outHeight)

}

// --------------------------------
// GetBlockHeaderByHeight
// --------------------------------

func GetBlockHeaderByHeight(chain unsafe.Pointer, height int) (int, unsafe.Pointer, int) {
	ptr := (C.chain_t)(chain)

	var outHeight C.uint64_t
	var headerPtr unsafe.Pointer

	res := C.chain_get_block_header_by_height(ptr, (C.uint64_t)(height), (*C.header_t)(&headerPtr), &outHeight)

	return int(res), headerPtr, int(outHeight)
}

// --------------------------------
// GetBlockHeaderByHash
// --------------------------------

func GetBlockHeaderByHash(chain unsafe.Pointer, hash HashT) (int, unsafe.Pointer, int) {
	ptr := (C.chain_t)(chain)

	var outHeight C.uint64_t
	var headerPtr unsafe.Pointer

	res := C.chain_get_block_header_by_hash(ptr, GoHashToC(hash), (*C.header_t)(&headerPtr), &outHeight)
	return int(res), headerPtr, int(outHeight)
}

// --------------------------------
// GetBlockByHeight
// --------------------------------

func GetBlockByHeight(chain unsafe.Pointer, height int) (int, unsafe.Pointer, int) {
	ptr := (C.chain_t)(chain)

	var outHeight C.uint64_t
	var blockPtr unsafe.Pointer

	res := C.chain_get_block_by_height(ptr, (C.uint64_t)(height), (*C.block_t)(&blockPtr), &outHeight)
	return int(res), blockPtr, int(outHeight)
}

// --------------------------------
// GetBlockByHash
// --------------------------------

func GetBlockByHash(chain unsafe.Pointer, hash HashT) (int, unsafe.Pointer, int) {
	ptr := (C.chain_t)(chain)

	var outHeight C.uint64_t
	var blockPtr unsafe.Pointer

	res := C.chain_get_block_by_hash(ptr, GoHashToC(hash), (*C.block_t)(&blockPtr), &outHeight)
	return int(res), blockPtr, int(outHeight)
}

// --------------------------------
// GetTransaction
// --------------------------------

func GetTransaction(chain unsafe.Pointer, hash HashT, requireConfirmed bool) (int, unsafe.Pointer, int, int) {
	ptr := (C.chain_t)(chain)

	var outHeight C.uint64_t
	var outIndex C.uint64_t
	var txPtr unsafe.Pointer

	res := C.chain_get_transaction(ptr, GoHashToC(hash), boolToC(requireConfirmed), (*C.transaction_t)(&txPtr), &outHeight, &outIndex)
	return int(res), txPtr, int(outHeight), int(outIndex)
}

//Note: removed on v.3.3.0
// // --------------------------------
// // GetOutput
// // --------------------------------
// func GetOutput(chain unsafe.Pointer, hash HashT, index int, requireConfirmed bool) (int, unsafe.Pointer) {
// 	hashC := C.CBytes(hash[:])
// 	defer C.free(hashC)

// 	ptr := (C.chain_t)(chain)

// 	var outputPtr unsafe.Pointer
// 	res := C.chain_get_output(ptr, GoHashToC(hash), C.uint32_t(index), boolToC(requireConfirmed), (*C.output_t)(&outputPtr))
// 	return int(res), outputPtr
// }

// --------------------------------
// getHistory
// --------------------------------

// //It is the user's responsibility to release the history returned in the callback
// int get_history(chainutor_t chain,
//                 payment_address_t address
//                 size_t limit,
//                 size_t from_height,
//                 history_compact_list_t* out_history) {

func getHistory(chain unsafe.Pointer, address string, limit int, fromHeight int) (int, unsafe.Pointer) {
	ptr := (C.chain_t)(chain)

	address_c_str := C.CString(address)
	defer C.free(unsafe.Pointer(address_c_str))

	pa := C.chain_payment_address_construct_from_string(address_c_str)
	// fetch_history(chain, pa, py_limit, py_from_height, history_fetch_handler);
	//

	var historyPtr unsafe.Pointer
	res := C.chain_get_history(ptr, pa, C.uint64_t(limit), C.uint64_t(fromHeight), (*C.history_compact_list_t)(&historyPtr))

	C.chain_payment_address_destruct(pa)

	return int(res), historyPtr
}
