package bridge

/*
#cgo LDFLAGS: ${SRCDIR}/mmlib/mmlib.a
#include "mmlib/mmlib.h"
*/
import "C"

import "fmt"

func Run(i int) {
	v := C.GetValue(C.int(i))
	fmt.Printf("hello! %v\n", v)
}
