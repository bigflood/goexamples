package main

/*
#cgo LDFLAGS: mmlib.a

__declspec(dllimport) void PrintValue(int v);

*/
import "C"

import "fmt"

func main() {
	fmt.Printf("win64 dll cpp\n")
	C.PrintValue(11)
}
