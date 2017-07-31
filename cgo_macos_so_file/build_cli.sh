#!/bin/bash

cp mmlib/mmlib.so .
go build -o run_example github.com/bigflood/goexamples/cgo_macos_so_file
