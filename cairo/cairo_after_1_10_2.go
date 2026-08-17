// +build !cairo_1_9,!cairo_1_10

package cairo

// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"
import (
	"unsafe"
)

const (
	ANTIALIAS_FAST     Antialias = C.CAIRO_ANTIALIAS_FAST // (since 1.12)
	ANTIALIAS_GOOD     Antialias = C.CAIRO_ANTIALIAS_GOOD // (since 1.12)
	ANTIALIAS_BEST     Antialias = C.CAIRO_ANTIALIAS_BEST // (since 1.12)
)


