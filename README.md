[![Go Report Card](https://goreportcard.com/badge/github.com/tie/tools.go)](https://goreportcard.com/report/github.com/tie/tools.go)
[![GoDoc Card](https://godoc.org/github.com/tie/tools.go?status.svg)](https://godoc.org/github.com/tie/tools.go)

# tools.go

The tools.go command helps you manage [tool dependencies](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module) for a Go module.

Unlike the [retool](https://github.com/twitchtv/retool), it relies on the go command managing versions and downloads. The only files used are `tools.go`, `go.mod` and `go.sum`.

## Usage

Running `tools.go` will install all necessary dependencies to the `GOBIN`  path.
That’s it!

At least for now. I’m planning to add more features.

## Installation

It’s recommended to install executable from [releases](https://github.com/tie/tools.go/releases).

You can, however, build from source.

```sh
git clone git@github.com:tie/tools.go

export GOBIN=$HOME/bin
go install ./tools.go
```

## License

### Public domain

This project is published and distributed under the [Unlicense](LICENSE). Attribution is optional, but desirable.

Rationales for placing software in public domain are listed in [nothings/stb docs](https://github.com/nothings/stb/blob/master/docs/why_public_domain.md).

### Traditional license

Want a traditional copyright-ish license?

>You are granted a perpetual, irrevocable license to copy, modify, publish, and distribute this software as you see fit.
