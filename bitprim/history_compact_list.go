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
// HistoryCompactList Golang idiomatic Interface
// --------------------------------

package bitprim

import (
	"C"
	"unsafe"
)

type HistoryCompactList struct {
	ptr unsafe.Pointer
}

func NewHistoryCompactList(ptr unsafe.Pointer) *HistoryCompactList {
	x := new(HistoryCompactList)
	x.ptr = ptr
	return x
}

func (x *HistoryCompactList) Close() {
	// fmt.Printf("Go.HistoryCompactList.Close() - ptr: %p\n", x.ptr)
	historyCompactListDestruct(x.ptr)
	x.ptr = nil
}

func (x *HistoryCompactList) Count() int {
	return historyCompactListCount(x.ptr)
}

func (x *HistoryCompactList) Nth(n int) *HistoryCompact {
	return NewHistoryCompact(historyCompactListNth(x.ptr, n))
}
