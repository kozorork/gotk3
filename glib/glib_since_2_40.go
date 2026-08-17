// Same copyright and license as the rest of the files in this project

// +build !glib_2_34

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
// #include "glib_since_2_40.go.h"
import "C"
import "unsafe"

/*
 * Application
 */

// MarkBusy is a wrapper around g_application_mark_busy().
func (v *Application) MarkBusy() {
	C.g_application_mark_busy(v.native())
}

// UnmarkBusy is a wrapper around g_application_unmark_busy().
func (v *Application) UnmarkBusy() {
	C.g_application_unmark_busy(v.native())
}

// SendNotification is a wrapper around g_application_send_notification().
func (v *Application) SendNotification(id string, notification *Notification) {
	cstr1 := (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(cstr1))

	C.g_application_send_notification(v.native(), cstr1, notification.native())
}

// WithdrawNotification is a wrapper around g_application_withdraw_notification().
func (v *Application) WithdrawNotification(id string) {
	cstr1 := (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(cstr1))

	C.g_application_withdraw_notification(v.native(), cstr1)
}

