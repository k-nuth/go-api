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
#cgo linux LDFLAGS: -L/home/fernando/dev/bitprim/bitprim-node-cint/cmake-build-release -lbitprim-node-cint

#include <stdio.h>
#include <stdlib.h>
// #include <bitprim/nodecint/block.h>
#include <bitprim/nodecint/executor_c.h>
#include <bitprim/nodecint/header.h>
#include <bitprim/nodecint/transaction.h>

// typedef void (*fetch_last_height_handler)(int);
// typedef void (*fetch_block_height_handler)(int);

// void fetchLastHeightGoCallBack_cgo(int in); // Forward declaration.
// void fetchBlockHeightGoCallBack_cgo(int in); // Forward declaration.

// void fetchBlockHeaderGoCallBack_cgo(header_t header, size_t height); // Forward declaration.
// void fetchBlockHeaderByHashGoCallBack_cgo(header_t header, size_t height); // Forward declaration.

// void fetchBlockGoCallBack_cgo(block_t block, size_t height); // Forward declaration.
// void fetchBlockByHashGoCallBack_cgo(block_t block, size_t height); // Forward declaration.

// void fetchTransactionGoCallBack_cgo(transaction_t transaction, size_t height, size_t index); // Forward declaration.

// void fetchOutputGoCallBack_cgo(output_t output); // Forward declaration.

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
	res := C.executor_run(ptr)
	return int(res)
}

func ExecutorRunAndWait(exec unsafe.Pointer) int {
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
// GetBlockHeader
// --------------------------------

func GetBlockHeader(exec unsafe.Pointer, height int) (int, unsafe.Pointer, int) {
	ptr := (*C.struct_executor)(exec)

	var outHeight C.size_t
	var headerPtr unsafe.Pointer

	res := C.get_block_header(ptr, (C.size_t)(height), (*C.header_t)(&headerPtr), &outHeight)

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
// GetBlock
// --------------------------------

func GetBlock(exec unsafe.Pointer, height int) (int, unsafe.Pointer, int) {
	ptr := (*C.struct_executor)(exec)

	var outHeight C.size_t
	var blockPtr unsafe.Pointer

	res := C.get_block(ptr, (C.size_t)(height), (*C.block_t)(&blockPtr), &outHeight)
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
