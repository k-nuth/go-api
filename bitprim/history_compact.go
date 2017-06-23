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
// HistoryCompact Golang idiomatic Interface
// --------------------------------

package bitprim

import (
	"unsafe"
)

// func historyCompactGetPointKind(historyCompact unsafe.Pointer) C.point_kind_t {
// func historyCompactGetPoint(historyCompact unsafe.Pointer) unsafe.Pointer {
// func historyCompactGetHeight(historyCompact unsafe.Pointer) uint64 {
// func historyCompactGetValueOrSpend(historyCompact unsafe.Pointer) uint64_t

type HistoryCompact struct {
	ptr unsafe.Pointer
}

func NewHistoryCompact(ptr unsafe.Pointer) *HistoryCompact {
	x := new(HistoryCompact)
	x.ptr = ptr
	return x
}

func (x *HistoryCompact) PointKind() int {
	return int(historyCompactGetPointKind(x.ptr))
}

func (x *HistoryCompact) Point() *Point {
	return NewPoint(historyCompactGetPoint(x.ptr))
}

func (x *HistoryCompact) Height() uint32 {
	return historyCompactGetHeight(x.ptr)
}

func (x *HistoryCompact) ValueOrSpend() uint64 {
	return historyCompactGetValueOrSpend(x.ptr)
}
