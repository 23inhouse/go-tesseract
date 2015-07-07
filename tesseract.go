package tesseract

/*
#cgo LDFLAGS: -llept -ltesseract
#include "tess.h"
*/
import "C"
import (
	"regexp"
	"strings"
)

func Process(filename, language string) string {
	p := C.CString(filename)
	l := C.CString(language)
	w := C.CString("")
	s := C.simple(p, w, l)
	str := C.GoString(s)

	r := regexp.MustCompile(`\s+`)
	str = string(r.ReplaceAll([]byte(str), []byte(" ")))
	return strings.Trim(str, " ")
}
