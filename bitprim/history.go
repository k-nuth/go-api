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
// History Golang idiomatic Interface
// --------------------------------

package bitprim

import (
	"encoding/hex"
)

type OutputPoint struct {
	Hash  string
	Index uint32
}

type History struct {
	Output       OutputPoint
	OutputHeight uint64
	Value        uint64
	Spend        OutputPoint

	// TemporaryChecksum if Output
	// SpendHeight if Spend
	Data uint64
	// SpendHeight       uint64
	// TemporaryChecksum uint64
}

type HistoryList []History

func Expand(list HistoryCompactList) HistoryList {
	var result HistoryList

	count := list.Count()
	//Process the outputs
	for n := 0; n < count; n++ {
		output := list.Nth(n)

		if output.PointKind() == 0 { // point_kind::output
			var row History

			hash := ReverseHash(output.Point().Hash())
			hashStr := hex.EncodeToString(hash[:])

			row.Output = OutputPoint{Hash: hashStr, Index: output.Point().Index()}
			row.OutputHeight = uint64(output.Height())
			row.Value = output.ValueOrPreviousChecksum()
			row.Spend = OutputPoint{Hash: "", Index: 4294967295}
			row.Data = output.Point().Checksum()

			result = append(result, row)

			// output = compact.erase(output);
		}
	}

	//Process the spends
	for n := 0; n < count; n++ {
		spend := list.Nth(n)

		if spend.PointKind() == 1 { // point_kind::spend

			// Update outputs with the corresponding spends.
			found := false
			for i := range result {

				if result[i].Data == spend.ValueOrPreviousChecksum() &&
					result[i].Spend.Hash == "" {
					hash := ReverseHash(spend.Point().Hash())
					hashStr := hex.EncodeToString(hash[:])
					result[i].Spend = OutputPoint{Hash: hashStr, Index: spend.Point().Index()}
					result[i].Data = uint64(spend.Height())
					found = true
					break
				}
			}

			// This will only happen if the history height cutoff comes between
			// an output and its spend. In this case we return just the spend.
			if !found {
				var row History

				// row.Output = output_point{null_hash, max_uint32}
				row.Output = OutputPoint{Hash: "", Index: 4294967295}

				row.OutputHeight = 18446744073709551615
				row.Value = 18446744073709551615

				// row.spend = spend.point
				hash := ReverseHash(spend.Point().Hash())
				hashStr := hex.EncodeToString(hash[:])
				row.Spend = OutputPoint{Hash: hashStr, Index: spend.Point().Index()}

				row.Data = uint64(spend.Height())
				result = append(result, row)
			}

		}
	}

	// compact.clear();

	// // Clear all remaining checksums from unspent rows.
	for i := range result {
		if result[i].Spend.Hash == "" {
			result[i].Data = 18446744073709551615
		}
	}

	// TODO: sort by height and index of output, spend or both in order.
	return result
}
