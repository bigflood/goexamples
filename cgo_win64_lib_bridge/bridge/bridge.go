package bridge

/*
// #cgo LDFLAGS: mmlib/mmlib.a
#cgo LDFLAGS: ${SRCDIR}/mmlib/Release/mmlib.lib

#include "mmlib/mmlib.h"

*/
import "C"

import "fmt"

func Run(i int) {
	v := C.GetValue(C.int(i))
	fmt.Printf("hello! %v\n", v)
}
