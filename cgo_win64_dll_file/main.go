package main

/*
#cgo LDFLAGS: mmlib.a

__declspec(dllimport) int GetValue(int v);

*/
import "C"

import "fmt"

func main() {
	v := C.GetValue(11)
	fmt.Printf("hello!! %v\n", v)
}
