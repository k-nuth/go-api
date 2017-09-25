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
// Executor Golang idiomatic Interface
// --------------------------------

package bitprim

// --------------------------------------------------------------------------------

import (
	"syscall" // or "runtime"
	"unsafe"
)

//Executor structure
type Executor struct {
	ptr unsafe.Pointer
}

func NewExecutor(path string) *Executor {
	return newExecutor(path)
}

func NewExecutorWithStd(path string) *Executor {
	return newExecutorWithStd(path, syscall.Stdout, syscall.Stderr)
}

func (x *Executor) Close() {
	ExecutorDestruct(x.ptr)
	x.ptr = nil
}

func (x Executor) Run() int {
	return ExecutorRun(x.ptr)
}

func (x Executor) Initchain() int {
	return ExecutorInitchain(x.ptr)
}

func (x Executor) Chain() *Chain {
	return NewChain(ExecutorChain(x.ptr))
}
