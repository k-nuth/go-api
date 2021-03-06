// +build windows

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

//cgo windows CFLAGS: -IC:/development/bitprim/bitprim-node-cint/include -IC:/development/bitprim/bitprim-core/include
//cgo windows LDFLAGS: -LC:/development/bitprim/bitprim-node-cint/build -lbitprim-node-cint

// cgo CFLAGS: -I./bitprim_c/include -I./bitprim_c/include
// cgo LDFLAGS: -L./bitprim_c/lib -lbitprim-node-cint

// cgo CFLAGS: -IC:/Users/Fernando/go/src/github.com/bitprim/bitprim-go/bitprim_c/include -IC:/Users/Fernando/go/src/github.com/bitprim/bitprim-go/bitprim_c/include
// cgo LDFLAGS: -LC:/Users/Fernando/go/src/github.com/bitprim/bitprim-go/bitprim_c/lib -lbitprim-node-cint

// #cgo LDFLAGS: -LC:/Users/Fernando/go/src/github.com/bitprim/bitprim-go/bitprim_c/lib/libbitprim-node-cint.a

// conan install . --profile mingw -o bitprim-node-cint:shared=True

/*
#cgo CFLAGS: -IC:/Users/Fernando/go/src/github.com/bitprim/bitprim-go/bitprim_c/include -IC:/Users/Fernando/go/src/github.com/bitprim/bitprim-go/bitprim_c/include
#cgo LDFLAGS: -LC:/Users/Fernando/go/src/github.com/bitprim/bitprim-go/bitprim_c/lib -lbitprim-node-cint

#include <stdio.h>
#include <stdlib.h>
#include <bitprim/nodecint/executor_c.h>
*/
import "C"

import (
	// or "runtime"
	"syscall"
	"unsafe"
)

func executorConstruct(path string, sout syscall.Handle, serr syscall.Handle) unsafe.Pointer {
	path_c := C.CString(path)
	defer C.free(unsafe.Pointer(path_c))

	exec := C.executor_construct_handles(path_c,
		unsafe.Pointer(sout),
		unsafe.Pointer(serr))

	// fmt.Printf("exec address = %p.\n", unsafe.Pointer(exec))
	return unsafe.Pointer(exec)
}

func newExecutor(path string) *Executor {
	x := new(Executor)
	x.ptr = executorConstruct(path, 0, 0)
	return x
}

func newExecutorWithStd(path string, sout_fd syscall.Handle, serr_fd syscall.Handle) *Executor {
	x := new(Executor)
	x.ptr = executorConstruct(path, sout_fd, serr_fd)
	return x
}
