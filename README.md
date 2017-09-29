# Bitcoin/Litecoin Full-Node Library written in Go Language

## Instructions to use the go-library. (Ubuntu)
*Note: the go-library is currently being developed, when it's finished the instructions to use it will be simplified.*

### Install go-lang 1.8
```sh
sudo add-apt-repository ppa:gophers/archive
sudo apt update
sudo apt-get install golang-1.8-go
```
### Clone `bitprim-go`
```sh
git clone https://github.com/bitprim/bitprim-go
```

### Export `GO` variables:
```sh
export PATH=$PATH:$(/usr/lib/go-1.8/bin/go env GOROOT)/bin
export GOPATH=$(go env GOPATH)
```

### Export library path: (`bitprim-node-cint` build folder)
```sh 
export LD_LIBRARY_PATH=/home/bitprim/bitprim/build/bitprim-node-cint:$LD_LIBRARY_PATH
```

### Execute `go get` to copy the project to ~/go
```sh
go get github.com/bitprim/bitprim-go/rest-api
```

### In the file `bitprim/executor_native.go` set the `c-api` location:
```
#cgo linux CFLAGS: -I/home/bitprim/bitprim/bitprim-node-cint/include 
#cgo linux LDFLAGS: -L/home/bitprim/bitprim/build/bitprim-node-cint -lbitprim-node-cint
```

### Run `go install`:
```sh
go install github.com/bitprim/bitprim-go/rest-api
```

### Run the server:
```sh
~/go/bin/rest-api
```
