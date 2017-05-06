# extender

Go toolchain subcommand extender. 

[![reportcard](https://goreportcard.com/badge/github.com/gomatic/extender)](https://goreportcard.com/report/github.com/gomatic/extender)
[![build](https://travis-ci.org/gomatic/extender.svg?branch=master)](https://travis-ci.org/gomatic/extender)

# for critics

Here's a quick explanation for the shortsighted who just love to criticize
before they take the two seconds needed to grasp the actual purpose of such a tool:

- This isn't about typing `go vend` instead of `govend`.
- This is about being able to replace or augment native `go` toolchain commands. e.g.
  - Do you have specific things you need as part of `go build`, e.g. adding common `-X`? 
  Then just install the augmentations and you've changed the behavior of `go build`.
  - Don't like how `go dep`, `go vet`, `go *` works? Replace it with a different implementation. 

It allows Go developers to continue using the standard toolchain and simple commands but
gain different/customized functionality ... because, thankfully, we don't all like exactly the same
things, so flexibility is useful.

If you're concerned about complexity and fragmentation ... then i'll challenge you to
create a simple toolchain that does everything everyone needs in every situation.
When someone has done that, I'll delete this repository.

# introduction

`extender` provides a `go` executable to precede `GOROOT/bin/go`.

This allows you to extend the `go` toolchain or even replace Go's native subcommands. 

:warning: Do not use this unless you have a good understanding of the go environment
and, more generally, configuring your shell's environment.  

`extender`'s `go` executable provides the ability to call subcommands through
the `go` command. The subcommands are implemented as executables with a recognizable
prefix (default is `go-` and `go`) instead of natively by `GOROOT/bin/go` (and the
`GOTOOLDIR` tools). If such an executable doesn't exist, it falls back to
`GOROOT/bin/go`.

For example,

    go ex

will (by default) execute

    go-ex

or, if that doesn't exist, it tries

    goex


A consequence: allows pointing the go command to separate versions of `go`.

    GOROOT=/go1.8.1 go build
    GOROOT=/go1.7.5 go build

# Installation

:warning: Installing this adds a `go` executable to your `GOBIN` or
`GOPATH[0]/bin` and overrides `GOROOT/bin` in the `PATH`

    go get github.com/gomatic/extender/...
    eval $(extender)

Additional, you can specify the prefixes that'll be considered extensions (the default is `go- go`):

    eval $(extender go go-)

will first look for, e.g., `goex` and if not found, `go-ex`, then try `GOROOT/bin/go ex`.
