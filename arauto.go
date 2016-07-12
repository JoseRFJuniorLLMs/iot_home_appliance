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

	c := C.CString(s)
	defer C.free(unsafe.Pointer(c))
	C.writeLCD(c)

	time.Sleep(5 * time.Second)
	C.removeLCD()
}
