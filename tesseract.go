package tesseract

/*
#cgo LDFLAGS: -llept -ltesseract
#include "clib/tess.h"
*/
import "C"

func Process(filename, language string) (string, error) {
	p := C.CString(filename)
	w := C.CString("")
	s := C.simple(p, w)
	return C.GoString(s)
}
