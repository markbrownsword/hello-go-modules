# Hello Go Modules

A starter project for experimenting with Go modules (new in Go v1.11)

### Prerequisites

A working Go environment (Go v1.11+) on Linux

[Install Go](https://golang.org/doc/install)

Update system paths (add at the bottom of `~/.profile`)
```sh
export GOPATH=$HOME/Go # Add this Directory `$ mkdir ~/Go`
export PATH=$PATH:$GOPATH/bin # Add this Directory `mkdir ~/Go/bin`
export PATH=$PATH:/usr/local/go/bin
```
Then
```sh
$ source ~/.profile
```

### Setup

Add this repo anywhere on your system (outside of $GOPATH)

e.g. ~/Projects/Go/github.com/markbrownsword/hello-go-modules

### Run
Execute from project root

```sh
$ go build -o $GOPATH/bin/hello-cli cmd/cli/main.go
$ hello-cli
```
The -o flag outputs the compiled binary to the specified location

### Test
```sh
$ go test ./...
```

### References
[Go 1.11 Modules](https://github.com/golang/go/wiki/Modules)
