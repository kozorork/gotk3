// +build !glib_2_34

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

// Explicitly releases the binding between the source and the target property
// expressed by Binding
func (v *Binding) Unbind() {
	C.g_binding_unbind(v.native())
}

