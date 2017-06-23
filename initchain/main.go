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

/*
export PATH=$PATH:$(/usr/lib/go-1.8/bin/go env GOROOT)/bin
export GOPATH=$(go env GOPATH)
export GODEBUG=cgocheck=0

export LD_LIBRARY_PATH=/home/fernando/dev/bitprim/bitprim-node-cint/cmake-build-debug:$LD_LIBRARY_PATH
go install github.com/bitprim/bitprim-go/initchain
go get github.com/bitprim/bitprim-go/initchain
$GOPATH/bin/bitprim_test

cd C:\Users\Fernando\go\bin
*/

package main

import (
	// or "runtime"
	"github.com/bitprim/bitprim-go/bitprim"
)

func main() {
	e := bitprim.NewExecutor("/pepe")
	//defer e.Close()

	// fmt.Println("before RUN")

	res := e.Initchain()
	// fmt.Printf("%d\n", res)
	e.Close()
}
