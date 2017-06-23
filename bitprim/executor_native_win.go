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

// cgo CFLAGS: -IC:/development/bitprim/bitprim-node-cint/include -IC:/development/bitprim/bitprim-core/include
// cgo LDFLAGS: -LC:/development/bitprim/bitprim-node-cint/cmake-build-release -lbitprim-node-cint

package bitprim

// --------------------------------------------------------------------------------

/*
#cgo CFLAGS: -IC:/development/bitprim2/bitprim-node-cint/include -IC:/development/bitprim2/bitprim-core/include
#cgo LDFLAGS: -LC:/development/bitprim2/bitprim-node-cint/build -lbitprim-node-cint

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
	path_c := C.CString(path)
	defer C.free(unsafe.Pointer(path_c))

	exec := C.executor_construct_handles(path_c,
		unsafe.Pointer(sin),
		unsafe.Pointer(sout),
		unsafe.Pointer(serr))

	fmt.Printf("exec address = %p.\n", unsafe.Pointer(exec))
	return unsafe.Pointer(exec)
}

func NewExecutorWithStd(path string, sin_fd syscall.Handle, sout_fd syscall.Handle, serr_fd syscall.Handle) *Executor {
	x := new(Executor)
	x.ptr = ExecutorConstruct(path, sin_fd, sout_fd, serr_fd)
	return x
}
