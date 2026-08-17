// Same copyright and license as the rest of the files in this project

// +build !glib_2_34

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

/*
 * Menu
 */
// Predefined attribute names for GMenu
var (
	MENU_ATTRIBUTE_ACTION_NAMESPACE string = C.G_MENU_ATTRIBUTE_ACTION_NAMESPACE
	MENU_ATTRIBUTE_ICON             string = C.G_MENU_ATTRIBUTE_ICON
)

// RemoveAll is a wrapper around g_menu_remove_all().
func (v *Menu) RemoveAll() {
	C.g_menu_remove_all(v.native())
}

