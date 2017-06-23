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

package bitprim

// #include <bitprim/nodecint/point.h>
import "C"

import (
	"unsafe"
)

// hash_t point_get_hash(point_t point);
// int /*bool*/ point_is_valid(point_t point);
// uint32_t point_get_index(point_t point);
// uint64_t point_get_checksum(point_t point);

func pointHash(point unsafe.Pointer) HashT {
	ptr := (C.point_t)(point)
	return CHashToGo(C.point_get_hash(ptr))
}

func pointIsValid(point unsafe.Pointer) bool {
	return CToBool(C.point_is_valid(C.point_t(point)))
}

func pointGetIndex(point unsafe.Pointer) uint32 {
	return uint32(C.point_get_index(C.point_t(point)))
}

func pointGetChecksum(point unsafe.Pointer) uint64 {
	return uint64(C.point_get_checksum(C.point_t(point)))
}
