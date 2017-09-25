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

// #include <bitprim/nodecint/chain/history_compact.h>
import "C"

import (
	"unsafe"
)

// point_kind_t history_compact_get_point_kind(history_compact_t history);
// point_t history_compact_get_point(history_compact_t history);
// uint32_t history_compact_get_height(history_compact_t history);
// uint64_t history_compact_get_value_or_spend(history_compact_t history);

func historyCompactGetPointKind(historyCompact unsafe.Pointer) C.point_kind_t {
	return C.chain_history_compact_get_point_kind(C.history_compact_t(historyCompact))
}

func historyCompactGetPoint(historyCompact unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.chain_history_compact_get_point(C.history_compact_t(historyCompact)))
}

func historyCompactGetHeight(historyCompact unsafe.Pointer) uint32 {
	return uint32(C.chain_history_compact_get_height(C.history_compact_t(historyCompact)))
}

func historyCompactGetValueOrPreviousChecksum(historyCompact unsafe.Pointer) uint64 {
	return uint64(C.chain_history_compact_get_value_or_previous_checksum(C.history_compact_t(historyCompact)))
}
