# extender

Go toolchain subcommand extender. 

[![reportcard](https://goreportcard.com/badge/github.com/gomatic/extender)](https://goreportcard.com/report/github.com/gomatic/extender)
[![build](https://travis-ci.org/gomatic/extender.svg?branch=master)](https://travis-ci.org/gomatic/extender)

`extender` provides a `go` executable to precede `GOROOT/bin/go`.

This allows you to extend the `go` toolchain or replace native commands. 

`extender`'s `go` provides the ability to call subcommands through the `go`
command but the subcommands are implemented as standalone executables instead of
natively by `GOROOT/bin/go`. If the standalone executable doesn't exist, it
falls back to `GOROOT/bin/go`.

For example,

    go ex

will (by default) execute

    go-ex


A consequence: allows pointing the go command to separate versions of `go`.

    go build
    GOROOT=/go1.7.1 go build

# Installation

:warning: Installing this adds a `go` executable to your `GOBIN`
 and overrides `GOROOT/bin` in the `PATH` :warning:  

    go get github.com/gomatic/extender/...
    eval $(extender)

Additional, you can specify the prefixes that'll be considered extensions:

    eval $(extender go go-)

will first look for, e.g., `goex` and if not found, `go-ex`, then try `GOROOT/bin/go ex`.
