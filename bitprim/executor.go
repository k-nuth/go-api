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
	return NewExecutorWithStd(path, syscall.Stdin, syscall.Stdout, syscall.Stderr)
}

func (x *Executor) Close() {
	ExecutorDestruct(x.ptr)
	x.ptr = nil
}

func (x Executor) Run() int {
	return ExecutorRun(x.ptr)
}

func (x Executor) RunAndWait() int {
	return ExecutorRunAndWait(x.ptr)
}

func (x Executor) Initchain() int {
	return ExecutorInitchain(x.ptr)
}

func (x Executor) GetLastHeight() (int, int) {
	return GetLastHeight(x.ptr)
}

func (x Executor) GetLastHeightAsync() (chan int, chan int) {
	ce := make(chan int)
	ch := make(chan int)
	go func() {
		te, th := x.GetLastHeight()
		ce <- te
		ch <- th
	}()
	return ce, ch
}

func (x Executor) GetBlockHeight(hash HashT) (int, int) {
	return GetBlockHeight(x.ptr, hash)
}

func (x Executor) GetBlockHeightAsync(hash HashT) (chan int, chan int) {
	ce := make(chan int)
	ch := make(chan int)
	go func() {
		te, th := x.GetBlockHeight(hash)
		ce <- te
		ch <- th
	}()
	return ce, ch
}

//TODO: Error management!
func (x Executor) GetBlockHeader(height int) *Header {
	_, ptr, h := GetBlockHeader(x.ptr, height)
	res := NewHeader(ptr, h)
	return res
}

//TODO: Error management!
func (x Executor) GetBlockHeaderAsync(height int) chan *Header {
	ch := make(chan *Header)
	go func() {
		th := x.GetBlockHeader(height)
		ch <- th
	}()
	return ch
}

//TODO: Error management!
func (x Executor) GetBlockHeaderByHash(hash HashT) *Header {
	_, ptr, h := GetBlockHeaderByHash(x.ptr, hash)
	res := NewHeader(ptr, h)
	return res
}

//TODO: Error management!
func (x Executor) GetBlockHeaderByHashAsync(hash HashT) chan *Header {
	ch := make(chan *Header)
	go func() {
		th := x.GetBlockHeaderByHash(hash)
		ch <- th
	}()
	return ch
}

//TODO: Error management!
func (x Executor) GetBlock(height int) *Block {
	_, ptr, h := GetBlock(x.ptr, height)
	res := NewBlock(ptr, h)
	return res
}

//TODO: Error management!
func (x Executor) GetBlockAsync(height int) chan *Block {
	ch := make(chan *Block)
	go func() {
		th := x.GetBlock(height)
		ch <- th
	}()
	return ch
}

//TODO: Error management!
func (x Executor) GetBlockByHash(hash HashT) *Block {
	_, ptr, h := GetBlockByHash(x.ptr, hash)
	res := NewBlock(ptr, h)
	return res
}

//TODO: Error management!
func (x Executor) GetBlockByHashAsync(hash HashT) chan *Block {
	ch := make(chan *Block)
	go func() {
		th := x.GetBlockByHash(hash)
		ch <- th
	}()
	return ch
}

//TODO: Error management!
func (x Executor) GetTransaction(hash HashT, requiredConfirmed bool) *Transaction {
	_, ptr, h, i := GetTransaction(x.ptr, hash, requiredConfirmed)
	res := NewTransaction(ptr, h, i)
	return res
}

//TODO: Error management!
func (x Executor) GetTransactionAsync(hash HashT, requiredConfirmed bool) chan *Transaction {
	ch := make(chan *Transaction)
	go func() {
		th := x.GetTransaction(hash, requiredConfirmed)
		ch <- th
	}()
	return ch
}

//TODO: Error management!
func (x Executor) GetOutput(hash HashT, index int, requiredConfirmed bool) *Output {
	_, ptr := GetOutput(x.ptr, hash, index, requiredConfirmed)
	res := NewOutput(ptr)
	return res
}

//TODO: Error management!
func (x Executor) GetOutputAsync(hash HashT, index int, requiredConfirmed bool) chan *Output {
	ch := make(chan *Output)
	go func() {
		th := x.GetOutput(hash, index, requiredConfirmed)
		ch <- th
	}()
	return ch
}
