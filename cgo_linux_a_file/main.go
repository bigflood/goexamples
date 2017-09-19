package main

/*
#cgo LDFLAGS: mmlib/mmlib.a
#include "mmlib/mmlib.h"
*/
import "C"

import "fmt"

func main() {
	v := C.GetValue(12)
	fmt.Printf("hello! %v\n", v)
}
