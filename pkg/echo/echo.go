package echo

/*
#include <stdlib.h>
#include "echo.h"
*/
import "C"

import (
	"unsafe"
	"strings"
)

func Echo(args ...string) {
	cs := C.CString(strings.Join(args, " "))
	C.echo(cs)
	C.free(unsafe.Pointer(cs))
}