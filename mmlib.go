package goexamples

/*
#cgo LDFLAGS: mmlib.so

#include "mmlib/mmlib.h"

// int GetValue() {
// 	return 234879;
// }

*/
import "C"

import "fmt"

func Run() {
	v := C.GetValue(12)
	fmt.Printf("hello! %v\n", v) //C.GetValue())
}
