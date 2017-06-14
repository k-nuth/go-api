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

// --------------------------------------------------------------------------------

/*
#cgo CFLAGS: -IC:/development/bitprim/bitprim-node-cint/include -IC:/development/bitprim/bitprim-core/include
#cgo LDFLAGS: -LC:/development/bitprim/bitprim-node-cint/cmake-build-debug -lbitprim-node-cint

#include <stdio.h>
#include <stdlib.h>
#include <bitprim/nodecint/executor_c.h>
*/
import "C"

import (
	"fmt" // or "runtime"
	"syscall"
	"unsafe"
)

func ExecutorConstruct(path string, sin syscall.Handle, sout syscall.Handle, serr syscall.Handle) unsafe.Pointer {
	fmt.Println("ExecutorConstruct 1")
	path_c := C.CString(path)
	fmt.Println("ExecutorConstruct 2")
	defer C.free(unsafe.Pointer(path_c))
	fmt.Println("ExecutorConstruct 3")

	// xxx := int(C.fileno((*C.FILE)(unsafe.Pointer(sin))))
	// fmt.Printf("xxx = %d.\n", xxx)

	exec := C.executor_construct_handles(path_c,
		unsafe.Pointer(sin),
		unsafe.Pointer(sout),
		unsafe.Pointer(serr))

	fmt.Println("ExecutorConstruct 4")
	fmt.Printf("exec address = %p.\n", unsafe.Pointer(exec))
	fmt.Println("ExecutorConstruct 5")
	return unsafe.Pointer(exec)
}

func NewExecutorWithStd(path string, sin_fd syscall.Handle, sout_fd syscall.Handle, serr_fd syscall.Handle) *Executor {
	x := new(Executor)
	x.native_ptr = ExecutorConstruct(path, sin_fd, sout_fd, serr_fd)
	return x
}