/*
arauto (herald)
*/

package main

/*
	#cgo LDFLAGS: -lstdc++ edison.a
	#cgo pkg-config: upm-i2clcd
	#include "edison.hpp"
*/
import "C"

import (
	"time"
	"unsafe"
)

func main() {
	C.initLCD()
	C.setCursor(0, 0)

	cs := C.CString("Hello from stdio\n")
	C.writeLCD(cs)
	C.free(unsafe.Pointer(cs))

	time.Sleep(5 * time.Second)
	C.removeLCD()
}
