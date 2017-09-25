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
// Chain Golang idiomatic Interface
// --------------------------------

package bitprim

// --------------------------------------------------------------------------------

import (
	"fmt" // or "runtime"
	"unsafe"
)

//Chain structure
type Chain struct {
	ptr unsafe.Pointer
}

func NewChain(ptr unsafe.Pointer) *Chain {
	x := new(Chain)
	x.ptr = ptr
	return x
}

func (x Chain) GetLastHeight() (int, int) {
	fmt.Println("Called last-height at chain.go")

	return GetLastHeight(x.ptr)
}

func (x Chain) GetLastHeightAsync() (chan int, chan int) {
	ce := make(chan int)
	ch := make(chan int)
	go func() {
		te, th := x.GetLastHeight()
		ce <- te
		ch <- th
	}()
	return ce, ch
}

func (x Chain) GetBlockHeight(hash HashT) (int, int) {
	return GetBlockHeight(x.ptr, hash)
}

func (x Chain) GetBlockHeightAsync(hash HashT) (chan int, chan int) {
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
func (x Chain) GetBlockHeaderByHeight(height int) *Header {
	_, ptr, h := GetBlockHeaderByHeight(x.ptr, height)
	res := NewHeader(ptr, h)
	return res
}

//TODO: Error management!
func (x Chain) GetBlockHeaderByHeightAsync(height int) chan *Header {
	ch := make(chan *Header)
	go func() {
		th := x.GetBlockHeaderByHeight(height)
		ch <- th
	}()
	return ch
}

//TODO: Error management!
func (x Chain) GetBlockHeaderByHash(hash HashT) *Header {
	_, ptr, h := GetBlockHeaderByHash(x.ptr, hash)
	res := NewHeader(ptr, h)
	return res
}

//TODO: Error management!
func (x Chain) GetBlockHeaderByHashAsync(hash HashT) chan *Header {
	ch := make(chan *Header)
	go func() {
		th := x.GetBlockHeaderByHash(hash)
		ch <- th
	}()
	return ch
}

//TODO: Error management!
func (x Chain) GetBlockByHeight(height int) *Block {
	_, ptr, h := GetBlockByHeight(x.ptr, height)
	res := NewBlock(ptr, h)
	return res
}

//TODO: Error management!
func (x Chain) GetBlockByHeightAsync(height int) chan *Block {
	ch := make(chan *Block)
	go func() {
		th := x.GetBlockByHeight(height)
		ch <- th
	}()
	return ch
}

//TODO: Error management!
func (x Chain) GetBlockByHash(hash HashT) *Block {
	_, ptr, h := GetBlockByHash(x.ptr, hash)
	res := NewBlock(ptr, h)
	return res
}

//TODO: Error management!
func (x Chain) GetBlockByHashAsync(hash HashT) chan *Block {
	ch := make(chan *Block)
	go func() {
		th := x.GetBlockByHash(hash)
		ch <- th
	}()
	return ch
}

//TODO: Error management!
func (x Chain) GetTransaction(hash HashT, requiredConfirmed bool) *Transaction {
	_, ptr, h, i := GetTransaction(x.ptr, hash, requiredConfirmed)
	res := NewTransaction(ptr, h, i)
	return res
}

//TODO: Error management!
func (x Chain) GetTransactionAsync(hash HashT, requiredConfirmed bool) chan *Transaction {
	ch := make(chan *Transaction)
	go func() {
		th := x.GetTransaction(hash, requiredConfirmed)
		ch <- th
	}()
	return ch
}

// // ----------------------------------------------------------
// //TODO: Error management!
// func (x Chain) GetOutput(hash HashT, index int, requiredConfirmed bool) *Output {
// 	_, ptr := GetOutput(x.ptr, hash, index, requiredConfirmed)
// 	res := NewOutput(ptr)
// 	return res
// }
// //TODO: Error management!
// func (x Chain) GetOutputAsync(hash HashT, index int, requiredConfirmed bool) chan *Output {
// 	ch := make(chan *Output)
// 	go func() {
// 		th := x.GetOutput(hash, index, requiredConfirmed)
// 		ch <- th
// 	}()
// 	return ch
// }

// ----------------------------------------------------------

func (x Chain) GetHistory(address string, limit int, fromHeight int) *HistoryCompactList {
	_, ptr := getHistory(x.ptr, address, limit, fromHeight)
	res := NewHistoryCompactList(ptr)
	return res
}

func (x Chain) GetHistoryExpanded(address string, limit int, fromHeight int) HistoryList {
	list := x.GetHistory(address, limit, fromHeight)
	return Expand(*list)
}
