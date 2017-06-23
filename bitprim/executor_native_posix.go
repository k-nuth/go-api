// +build linux

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

import (
	"C" // or "runtime"
	"unsafe"
)

func executorConstruct(path string, sout_fd int, serr_fd int) unsafe.Pointer {
	path_c := C.CString(path)
	defer C.free(unsafe.Pointer(path_c))

	exec := C.executor_construct_fd(path_c, C.int(sout_fd), C.int(serr_fd))
	// fmt.Printf("exec address = %p.\n", unsafe.Pointer(exec))
	return unsafe.Pointer(exec)

}

func newExecutor(path string) *Executor {
	x := new(Executor)
	x.ptr = executorConstruct(path, -1, -1)
	return x
}

func newExecutorWithStd(path string, sout_fd int, serr_fd int) *Executor {
	x := new(Executor)
	x.ptr = executorConstruct(path, sout_fd, serr_fd)
	return x
}
