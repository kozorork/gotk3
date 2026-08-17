// +build !glib_2_34

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
// #include "glib_since_2_40.go.h"
import "C"
import "unsafe"

// VariantFromString is a wrapper around g_variant_new_string/g_variant_new_take_string.
// Uses g_variant_new_take_string to reduce memory allocations if possible.
func VariantFromString(value string) *Variant {
	cstr := (*C.gchar)(C.CString(value))
	// g_variant_new_take_string takes owhership of the cstring and will call free() on it when done.
	// Do NOT free this string in this function!
	return takeVariant(C.g_variant_new_take_string(cstr))
}

