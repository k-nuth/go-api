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
export LD_LIBRARY_PATH=/home/fernando/dev/bitprim/bitprim-node-cint/build:$LD_LIBRARY_PATH

go install github.com/bitprim/bitprim-go/rest-api
go get github.com/bitprim/bitprim-go/rest-api
go get go get github.com/gorilla/mux

$GOPATH/bin/rest-api

cd C:\Users\Fernando\go\bin
*/

package main

import (
	"encoding/json"
	"fmt" // or "runtime"
	// "html"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bitprim/bitprim-go/bitprim"
	"github.com/gorilla/mux"
)

func startHttpServer(c *bitprim.Chain) *http.Server {

	router := mux.NewRouter().StrictSlash(true)

	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// http://127.0.0.1:8088/last-height
	router.HandleFunc("/last-height", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Called last-height")

		_, height := c.GetLastHeight()

		fmt.Println("Called last-height AFTER")

		// fmt.Printf("Last Height %d\n", height)
		fmt.Fprintf(w, "Last Height: %d\n", height)
	})

	// ./bx-linux-x64-qrcode fetch-history 134HfD2fdeBTohfx8YANxEpsYXsv5UoWyz
	// ./bx-linux-x64-qrcode fetch-history 1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa
	// ./bx-linux-x64-qrcode fetch-history 1MLVpZC2CTFHheox8SCEnAbW5NBdewRTdR
	// ./bx-linux-x64-qrcode fetch-history 15Z5YJaaNSxeynvr6uW6jQZLwq3n1Hu6RX

	// Ejemplo BX - 247683
	// http://127.0.0.1:8088/history/134HfD2fdeBTohfx8YANxEpsYXsv5UoWyz

	// Satoshi - 123723
	// http://127.0.0.1:8088/history/1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa

	// Juan - 262421
	// http://127.0.0.1:8088/history/1MLVpZC2CTFHheox8SCEnAbW5NBdewRTdR

	// xxx - 50026
	// http://127.0.0.1:8088/history/15Z5YJaaNSxeynvr6uW6jQZLwq3n1Hu6RX

	router.HandleFunc("/history/{paymentAddress}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paymentAddress := vars["paymentAddress"]
		// fmt.Fprintln(w, "history:", paymentAddress)
		// fmt.Println("history:", paymentAddress)

		list := c.GetHistoryExpanded(paymentAddress, 0, 0)
		transfers := bitprim.ToJsonStruct(list)
		json.MarshalIndent(transfers, "", "    ")
		json.NewEncoder(w).Encode(transfers)
	})

	srv := &http.Server{Addr: ":8088", Handler: router}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	return srv
}

func main() {

	e := bitprim.NewExecutor("/pepe")
	// e := bitprim.NewExecutorWithStd("/pepe")
	// defer e.Close()

	res := e.Initchain()
	// fmt.Printf("%d\n", res)

	if res != 1 {
		fmt.Println("Blockchain was previously initialized")
	}

	res = e.Run()

	if res != 0 {
		fmt.Println("Error in Bitprim...")
		e.Close()
		return
	}

	srv := startHttpServer(e.Chain())

	// running := true

	c := make(chan os.Signal, 1)
	exitC := make(chan int, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for sig := range c {
			// for _ = range c {
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			fmt.Printf("captured %v\n", sig)

			// if running {
			// 	running = false
			// 	exitC <- 1
			// }

			exitC <- 1

		}
	}()

	fmt.Println("Bitprim Server running...")

	<-exitC

	fmt.Println("closing...")

	e.Close()
	err := srv.Shutdown(nil)

	// if err := srv.Shutdown(nil); err != nil {
	// 	panic(err) // failure/timeout shutting down the server gracefully
	// }

	if err != nil {
		// panic(err) // failure/timeout shutting down the server gracefully
	}

	fmt.Println("exiting...")
}
