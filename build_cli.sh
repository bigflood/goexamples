#!/bin/bash

# go build -o clirun github.com/bigflood/goexamples/cli
# -> "[?]   ???? killed   ./clirun"
# https://github.com/golang/go/issues/19734

go build -ldflags -s -o clirun github.com/bigflood/goexamples/cli


# go build  -ldflags "-linkmode external -extldflags -static"  -o clirun github.com/bigflood/goexamples/cli