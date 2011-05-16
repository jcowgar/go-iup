/* 
	Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>

	This file is part of go-iup.

	go-iup is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as
	published by the Free Software Foundation, either version 3 of
	the License, or (at your option) any later version.

	go-iup is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU Lesser General Public
	License along with go-iup.  If not, see <http://www.gnu.org/licenses/>.
*/

package iup

/*
#include <stdlib.h>
#include <iup.h>

extern int goIupMapCB(void *);
void goIupMapFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_MAP_CB", f);
	IupSetCallback(ih, "MAP_CB", (Icallback) goIupMapCB);
}

extern int goIupUnmapCB(void *);
void goIupUnmapFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_UNMAP_CB", f);
	IupSetCallback(ih, "UNMAP_CB", (Icallback) goIupUnmapCB);
}

extern int goIupDestroyCB(void *);
void goIupDestroyFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_DESTROY_CB", f);
	IupSetCallback(ih, "DESTROY_CB", (Icallback) goIupDestroyCB);
}

extern int goIupGetFocusCB(void *);
void goIupGetFocusFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_GETFOCUS_CB", f);
	IupSetCallback(ih, "GETFOCUS_CB", (Icallback) goIupGetFocusCB);
}

extern int goIupKillFocusCB(void *);
void goIupKillFocusFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_KILLFOCUS_CB", f);
	IupSetCallback(ih, "KILLFOCUS_CB", (Icallback) goIupKillFocusCB);
}

extern int goIupEnterWindowCB(void *);
void goIupEnterWindowFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_ENTERWINDOW_CB", f);
	IupSetCallback(ih, "ENTERWINDOW_CB", (Icallback) goIupEnterWindowCB);
}

extern int goIupLeaveWindowCB(void *);
void goIupLeaveWindowFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_LEAVEWINDOW_CB", f);
	IupSetCallback(ih, "LEAVEWINDOW_CB", (Icallback) goIupLeaveWindowCB);
}

extern int goIupKAnyCB(void *);
void goIupKAnyFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_K_ANY", f);
	IupSetCallback(ih, "K_ANY", (Icallback) goIupKAnyCB);
}

extern int goIupHelpCB(void *);
void goIupHelpFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_HELP_CB", f);
	IupSetCallback(ih, "HELP_CB", (Icallback) goIupHelpCB);
}

extern int goIupActionCB(void *);
void goIupSetActionFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_ACTION", f);
	IupSetCallback(ih, "ACTION", (Icallback) goIupActionCB);
}

extern int goIupTextActionCB(void *ih, int ch, void *newValue);
void goIupSetTextActionFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_ACTION", f);
	IupSetCallback(ih, "ACTION", (Icallback) goIupTextActionCB);
}
*/
import "C"
import "unsafe"


// Common callbacks

type MapFunc func(*Ihandle) int

//export goIupMapCB
func goIupMapCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_MAP_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*MapFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

type UnmapFunc func(*Ihandle) int

//export goIupUnmapCB
func goIupUnmapCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_UNMAP_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*UnmapFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

type DestroyFunc func(*Ihandle) int

//export goIupDestroyCB
func goIupDestroyCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_DESTROY_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*DestroyFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

type GetFocusFunc func(*Ihandle) int

//export goIupGetFocusCB
func goIupGetFocusCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_GETFOCUS_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*GetFocusFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

type KillFocusFunc func(*Ihandle) int

//export goIupKillFocusCB
func goIupKillFocusCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_KILLFOCUS_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*KillFocusFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

type EnterWindowFunc func(*Ihandle) int

//export goIupEnterWindowCB
func goIupEnterWindowCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_ENTERWINDOW_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*EnterWindowFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

type LeaveWindowFunc func(*Ihandle) int

//export goIupLeaveWindowCB
func goIupLeaveWindowCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_LEAVEWINDOW_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*LeaveWindowFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

type KAnyFunc func(*Ihandle, int) int

//export goIupKAnyCB
func goIupKAnyCB(ih unsafe.Pointer, c C.int) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_K_ANY")
	defer C.free(unsafe.Pointer(cName))

	f := *(*KAnyFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, int(c))
}

type HelpFunc func(*Ihandle) int

//export goIupHelpCB
func goIupHelpCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_HELP_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*HelpFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

type ActionFunc func(*Ihandle) int

//export goIupActionCB
func goIupActionCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_ACTION")
	defer C.free(unsafe.Pointer(cName))

	f := *(*ActionFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetActionFunc(f ActionFunc) {
	C.goIupSetActionFunc(ih.h, unsafe.Pointer(&f))
}

type TextActionFunc func(ih *Ihandle, ch int, newValue string) int

//export goIupTextActionCB
func goIupTextActionCB(ih unsafe.Pointer, ch int, newValue unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_ACTION")
	defer C.free(unsafe.Pointer(cName))

	f := *(*TextActionFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	goNewValue := C.GoString((*C.char)(newValue))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, ch, goNewValue)
}

func (ih *Ihandle) SetTextActionFunc(f TextActionFunc) {
	C.goIupSetTextActionFunc(ih.h, unsafe.Pointer(&f))
}
