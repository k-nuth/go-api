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
#cgo linux CFLAGS: -I/home/fernando/dev/bitprim/bitprim-node-cint/include -I/home/fernando/dev/bitprim/bitprim-core/include
#cgo linux LDFLAGS: -L/home/fernando/dev/bitprim/bitprim-node-cint/build -lbitprim-node-cint

#include <stdio.h>
#include <stdlib.h>
// #include <bitprim/nodecint/block.h>
#include <bitprim/nodecint/executor_c.h>
#include <bitprim/nodecint/header.h>
#include <bitprim/nodecint/transaction.h>
#include <bitprim/nodecint/payment_address.h>
*/
import "C"

// --------------------------------------------------------------------------------

import (
	"unsafe"
)

func CHashToGo(hashCPtr C.hash_t) HashT {
	hashC := unsafe.Pointer(hashCPtr)

	hashGoSlice := C.GoBytes(hashC, 32)
	var hash HashT
	copy(hash[:], hashGoSlice)
	return hash
}

func boolToC(x bool) C.int {
	if x {
		return 1
	}
	return 0
}

func CToBool(x C.int) bool {
	if x == 0 {
		return false
	}
	return true
}

func ExecutorDestruct(exec unsafe.Pointer) {
	ptr := (*C.struct_executor)(exec)
	C.executor_destruct(ptr)
}

func ExecutorRun(exec unsafe.Pointer) int {
	ptr := (*C.struct_executor)(exec)
	res := C.executor_run_wait(ptr)
	return int(res)
}

func ExecutorInitchain(exec unsafe.Pointer) int {
	ptr := (*C.struct_executor)(exec)
	res := C.executor_initchain(ptr)
	return int(res)
}

// --------------------------------
// GetLastHeight
// --------------------------------

func GetLastHeight(exec unsafe.Pointer) (int, int) {
	ptr := (*C.struct_executor)(exec)

	var outHeight C.size_t
	res := C.get_last_height(ptr, &outHeight)
	return int(res), int(outHeight)
}

// --------------------------------
// GetBlockHeight
// --------------------------------
type HashT [32]byte

func GetBlockHeight(exec unsafe.Pointer, hash HashT) (int, int) {
	ptr := (*C.struct_executor)(exec)

	hashC := C.CBytes(hash[:])
	defer C.free(hashC)

	var outHeight C.size_t
	res := C.get_block_height(ptr, (*C.uint8_t)(hashC), &outHeight)
	return int(res), int(outHeight)
}

// --------------------------------
// GetBlockHeaderByHeight
// --------------------------------

func GetBlockHeaderByHeight(exec unsafe.Pointer, height int) (int, unsafe.Pointer, int) {
	ptr := (*C.struct_executor)(exec)

	var outHeight C.size_t
	var headerPtr unsafe.Pointer

	res := C.get_block_header_by_height(ptr, (C.size_t)(height), (*C.header_t)(&headerPtr), &outHeight)

	return int(res), headerPtr, int(outHeight)
}

// --------------------------------
// GetBlockHeaderByHash
// --------------------------------

func GetBlockHeaderByHash(exec unsafe.Pointer, hash HashT) (int, unsafe.Pointer, int) {
	hashC := C.CBytes(hash[:])
	defer C.free(hashC)

	ptr := (*C.struct_executor)(exec)

	var outHeight C.size_t
	var headerPtr unsafe.Pointer

	res := C.get_block_header_by_hash(ptr, (*C.uint8_t)(hashC), (*C.header_t)(&headerPtr), &outHeight)
	return int(res), headerPtr, int(outHeight)
}

// --------------------------------
// GetBlockByHeight
// --------------------------------

func GetBlockByHeight(exec unsafe.Pointer, height int) (int, unsafe.Pointer, int) {
	ptr := (*C.struct_executor)(exec)

	var outHeight C.size_t
	var blockPtr unsafe.Pointer

	res := C.get_block_by_height(ptr, (C.size_t)(height), (*C.block_t)(&blockPtr), &outHeight)
	return int(res), blockPtr, int(outHeight)
}

// --------------------------------
// GetBlockByHash
// --------------------------------

func GetBlockByHash(exec unsafe.Pointer, hash HashT) (int, unsafe.Pointer, int) {
	hashC := C.CBytes(hash[:])
	defer C.free(hashC)

	ptr := (*C.struct_executor)(exec)

	var outHeight C.size_t
	var blockPtr unsafe.Pointer

	res := C.get_block_by_hash(ptr, (*C.uint8_t)(hashC), (*C.block_t)(&blockPtr), &outHeight)
	return int(res), blockPtr, int(outHeight)
}

// --------------------------------
// GetTransaction
// --------------------------------

func GetTransaction(exec unsafe.Pointer, hash HashT, requireConfirmed bool) (int, unsafe.Pointer, int, int) {
	hashC := C.CBytes(hash[:])
	defer C.free(hashC)

	ptr := (*C.struct_executor)(exec)

	var outHeight C.size_t
	var outIndex C.size_t
	var txPtr unsafe.Pointer

	res := C.get_transaction(ptr, (*C.uint8_t)(hashC), boolToC(requireConfirmed), (*C.transaction_t)(&txPtr), &outHeight, &outIndex)
	return int(res), txPtr, int(outHeight), int(outIndex)
}

// --------------------------------
// GetOutput
// --------------------------------
func GetOutput(exec unsafe.Pointer, hash HashT, index int, requireConfirmed bool) (int, unsafe.Pointer) {
	hashC := C.CBytes(hash[:])
	defer C.free(hashC)

	ptr := (*C.struct_executor)(exec)

	var outputPtr unsafe.Pointer
	res := C.get_output(ptr, (*C.uint8_t)(hashC), C.uint32_t(index), boolToC(requireConfirmed), (*C.output_t)(&outputPtr))
	return int(res), outputPtr
}

// --------------------------------
// getHistory
// --------------------------------

// //It is the user's responsibility to release the history returned in the callback
// int get_history(executor_t exec,
//                 payment_address_t address
//                 size_t limit,
//                 size_t from_height,
//                 history_compact_list_t* out_history) {

func getHistory(exec unsafe.Pointer, address string, limit int, fromHeight int) (int, unsafe.Pointer) {
	ptr := (*C.struct_executor)(exec)

	address_c_str := C.CString(address)
	defer C.free(unsafe.Pointer(address_c_str))

	pa := C.payment_address_construct_from_string(address_c_str)
	// fetch_history(exec, pa, py_limit, py_from_height, history_fetch_handler);
	//

	var historyPtr unsafe.Pointer
	res := C.get_history(ptr, pa, C.size_t(limit), C.size_t(fromHeight), (*C.history_compact_list_t)(&historyPtr))

	C.payment_address_destruct(pa)

	return int(res), historyPtr
}
