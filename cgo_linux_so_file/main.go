package main

/*
#cgo LDFLAGS: mmlib.so
#include "mmlib/mmlib.h"
*/
import "C"

import "fmt"

func main() {
	fmt.Printf("linux so!\n")
	C.PrintValue(13)
}
