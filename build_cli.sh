#!/bin/bash

go build -ldflags -s -o clirun github.com/bigflood/goexamples/cli
# go build -o clirun github.com/bigflood/goexamples/cli
# go build  -ldflags "-linkmode external -extldflags -static"  -o clirun github.com/bigflood/goexamples/cli