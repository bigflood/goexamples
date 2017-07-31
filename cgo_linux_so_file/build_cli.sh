#!/bin/bash

cp mmlib/mmlib.so .
go build -o run_example github.com/bigflood/goexamples/cgo_linux_so_file
