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

type Item struct {
	Hash   string `json:"hash"`
	Height uint64 `json:"height"`
	Index  uint32 `json:"index"`
}

type Transfer struct {
	Received *Item  `json:"received,omitempty"`
	Spent    *Item  `json:"spent,omitempty"`
	Value    uint64 `json:"value"`
}

type Transfers []Transfer

// func (this Transfers) MarshalJSON() ([]byte, error) {
// 	buffer := bytes.NewBufferString("{")
// 	length := len(this)
// 	count := 0
// 	for key, value := range this {
// 		jsonValue, err := json.Marshal(value)
// 		if err != nil {
// 			return nil, err
// 		}
// 		buffer.WriteString(fmt.Sprintf("\"%d\":%s", key, string(jsonValue)))
// 		count++
// 		if count < length {
// 			buffer.WriteString(",")
// 		}
// 	}
// 	buffer.WriteString("}")
// 	return buffer.Bytes(), nil
// }

// func (this Transfer) MarshalJSON() ([]byte, error) {
// 	buffer := bytes.NewBufferString("{")
// 	length := len(this)
// 	count := 0
// 	for key, value := range this {
// 		jsonValue, err := json.Marshal(value)
// 		if err != nil {
// 			return nil, err
// 		}
// 		buffer.WriteString(fmt.Sprintf("\"%d\":%s", key, string(jsonValue)))
// 		count++
// 		if count < length {
// 			buffer.WriteString(",")
// 		}
// 	}
// 	buffer.WriteString("}")
// 	return buffer.Bytes(), nil
// }

func ToJsonStruct(list HistoryList) Transfers {
	// transfers := make([]Transfer, 1)
	transfers := make([]Transfer, 0)

	for _, row := range list {

		tr := Transfer{Value: row.Value}

		if row.Output.Hash != "" {
			tr.Received = new(Item)
			tr.Received.Hash = row.Output.Hash

			// zeroized received.height implies output unconfirmed (in mempool)
			if row.OutputHeight != 0 {
				tr.Received.Height = row.OutputHeight
			}
			tr.Received.Index = row.Output.Index
		}

		// missing input implies unspent
		if row.Spend.Hash != "" {
			tr.Spent = new(Item)
			tr.Spent.Hash = row.Output.Hash

			// zeroized input.height implies spend unconfirmed (in mempool)
			if row.Data != 0 {
				tr.Spent.Height = row.Data //spend_height
			}

			tr.Spent.Index = row.Output.Index
		}

		transfers = append(transfers, tr)

	}
	return transfers
}
