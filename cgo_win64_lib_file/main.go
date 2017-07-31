package main

/*
// #cgo LDFLAGS: mmlib/mmlib.a
#cgo LDFLAGS: mmlib/Release/mmlib.lib

#include "mmlib/mmlib.h"

// int GetValue() {
// 	return 234879;
// }

*/
import "C"

import "fmt"

func main() {
	v := C.GetValue(13)
	fmt.Printf("hello! %v\n", v)
}
