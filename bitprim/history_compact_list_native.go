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

// #include <bitprim/nodecint/chain/history_compact_list.h>
import "C"

import (
	"unsafe"
)

func historyCompactListDestruct(historyCompactList unsafe.Pointer) {
	C.chain_history_compact_list_destruct(C.history_compact_list_t(historyCompactList))
}

func historyCompactListCount(block unsafe.Pointer) int {
	ptr := (C.history_compact_list_t)(block)
	return (int)(C.chain_history_compact_list_count(ptr))
}

func historyCompactListNth(block unsafe.Pointer, n int) unsafe.Pointer {
	ptr := (C.history_compact_list_t)(block)
	res := C.chain_history_compact_list_nth(ptr, C.uint64_t(n))
	return unsafe.Pointer(res)
}
