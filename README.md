# Hello Go Modules

A starter project for experimenting with Go modules (new in Go v1.11)

### Prerequisites

A working Go environment (Go v1.11+) on Linux

[Install Go](https://golang.org/doc/install)

Update system paths (`~/.bashrc, ~/.zshrc etc`)
```sh
# Go
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$(go env GOPATH)/bin
```
Then logout (or source the system file)

### Setup

Add this repo anywhere on your system (outside of $GOPATH)

e.g. ~/Projects/Go/github.com/markbrownsword/hello-go-modules

### Run
Execute from project root

```sh
$ go run ./cmd/speak/speak.go
```

### Install
Execute from project root, installs to $GOPATH/bin

```sh
$ go install ./cmd/speak/speak.go
$ speak
```

### Test
```sh
$ go test ./...
```

### References
[Go 1.11 Modules](https://github.com/golang/go/wiki/Modules)
