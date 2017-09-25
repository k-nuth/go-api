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

/*
#include <bitprim/nodecint/chain/block.h>
*/
import "C"

import (
	"unsafe"
)

func blockDestruct(block unsafe.Pointer) {
	ptr := (C.block_t)(block)
	C.chain_block_destruct(ptr)
}

func blockIsValid(block unsafe.Pointer) bool {
	ptr := (C.block_t)(block)
	res := C.chain_block_is_valid(ptr)
	return res == 0
}

func blockHash(block unsafe.Pointer) HashT {
	ptr := (C.block_t)(block)
	return CHashToGo(C.chain_block_hash(ptr))
}

func blockHeader(block unsafe.Pointer) unsafe.Pointer {
	ptr := (C.block_t)(block)
	return unsafe.Pointer(C.chain_block_header(ptr))
}

func blockTransactionCount(block unsafe.Pointer) int {
	ptr := (C.block_t)(block)
	return (int)(C.chain_block_transaction_count(ptr))
}

func blockTransactionNth(block unsafe.Pointer, n int) unsafe.Pointer {
	ptr := (C.block_t)(block)
	res := C.chain_block_transaction_nth(ptr, C.uint64_t(n))
	return unsafe.Pointer(res)
}

// -----------------------

func blockSerializedSize(block unsafe.Pointer, version uint32) uint64 {
	return uint64(C.chain_block_serialized_size((C.block_t)(block), C.uint32_t(version)))
}

/*static*/
func blockSubsidy(height uint64) uint64 {
	return uint64(C.chain_block_subsidy(C.uint64_t(height)))
}

func blockFees(block unsafe.Pointer) uint64 {
	return uint64(C.chain_block_fees((C.block_t)(block)))
}

func blockClaim(block unsafe.Pointer) uint64 {
	return uint64(C.chain_block_claim((C.block_t)(block)))
}

func blockReward(block unsafe.Pointer, height uint64) uint64 {
	return uint64(C.chain_block_reward((C.block_t)(block), C.uint64_t(height)))
}

func blockGenerateMerkleRoot(block unsafe.Pointer) HashT {
	return CHashToGo(C.chain_block_generate_merkle_root((C.block_t)(block)))
}

func blockSignatureOperations(block unsafe.Pointer) uint64 {
	return uint64(C.chain_block_signature_operations((C.block_t)(block)))
}

func blockSignatureOperationsBip16Active(block unsafe.Pointer, bip16_active bool) uint64 {
	return uint64(C.chain_block_signature_operations_bip16_active((C.block_t)(block), boolToC(bip16_active)))
}

func blockTotalInputs(block unsafe.Pointer, with_coinbase bool) uint64 {
	return uint64(C.chain_block_total_inputs((C.block_t)(block), boolToC(with_coinbase)))
}

func blockIsExtraCoinbases(block unsafe.Pointer) bool {
	return CToBool(C.chain_block_is_extra_coinbases((C.block_t)(block)))
}

func blockIsFinal(block unsafe.Pointer, height uint64, blockTime uint32) bool {
	return CToBool(C.chain_block_is_final((C.block_t)(block), C.uint64_t(height), C.uint32_t(blockTime)))
}

func blockIsDistinctTransactionSet(block unsafe.Pointer) bool {
	return CToBool(C.chain_block_is_distinct_transaction_set((C.block_t)(block)))
}

func blockIsValidCoinbaseClaim(block unsafe.Pointer, height uint64) bool {
	return CToBool(C.chain_block_is_valid_coinbase_claim((C.block_t)(block), C.uint64_t(height)))
}

func blockIsValidCoinbaseScript(block unsafe.Pointer, height uint64) bool {
	return CToBool(C.chain_block_is_valid_coinbase_script((C.block_t)(block), C.uint64_t(height)))
}

func blockIsInternalDoubleSpend(block unsafe.Pointer) bool {
	return CToBool(C.chain_block_is_internal_double_spend((C.block_t)(block)))
}

func blockIsValidMerkleRoot(block unsafe.Pointer) bool {
	return CToBool(C.chain_block_is_valid_merkle_root((C.block_t)(block)))
}
