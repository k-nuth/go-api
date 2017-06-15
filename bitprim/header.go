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
// Header Golang idiomatic Interface
// --------------------------------

package bitprim

import (
	"fmt" // or "runtime"
	"unsafe"
)

type Header struct {
	ptr    unsafe.Pointer
	height int
}

func NewHeader(ptr unsafe.Pointer, height int) *Header {
	x := new(Header)
	x.ptr = ptr
	x.height = height
	return x
}

func (x *Header) Close() {
	fmt.Printf("Go.Header.Close() - ptr: %p\n", x.ptr)
	headerDestruct(x.ptr)
	x.ptr = nil
}

func (x Header) IsValid() bool {
	return headerIsValid(x.ptr)
}

func (x Header) Version() int {
	return headerVersion(x.ptr)
}

func (x *Header) SetVersion(version int) {
	headerSetVersion(x.ptr, version)
}

func (x Header) Timestamp() int {
	return headerTimestamp(x.ptr)
}

func (x *Header) SetTimestamp(timestamp int) {
	headerSetTimestamp(x.ptr, timestamp)
}

func (x Header) Bits() int {
	return headerBits(x.ptr)
}

func (x *Header) SetBits(bits int) {
	headerSetBits(x.ptr, bits)
}

func (x Header) Nonce() int {
	return headerNonce(x.ptr)
}

func (x *Header) SetNonce(nonce int) {
	headerSetNonce(x.ptr, nonce)
}

func (x Header) PreviousBlockHash() HashT {
	return headerPreviousBlockHash(x.ptr)
}

func (x Header) Merkle() HashT {
	return headerMerkle(x.ptr)
}

func (x Header) Hash() HashT {
	return headerHash(x.ptr)
}
