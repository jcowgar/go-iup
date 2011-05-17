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
void goIupSetMapFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_MAP_CB", f);
	IupSetCallback(ih, "MAP_CB", (Icallback) goIupMapCB);
}

extern int goIupUnmapCB(void *);
void goIupSetUnmapFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_UNMAP_CB", f);
	IupSetCallback(ih, "UNMAP_CB", (Icallback) goIupUnmapCB);
}

extern int goIupDestroyCB(void *);
void goIupSetDestroyFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_DESTROY_CB", f);
	IupSetCallback(ih, "DESTROY_CB", (Icallback) goIupDestroyCB);
}

extern int goIupGetFocusCB(void *);
void goIupSetGetFocusFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_GETFOCUS_CB", f);
	IupSetCallback(ih, "GETFOCUS_CB", (Icallback) goIupGetFocusCB);
}

extern int goIupKillFocusCB(void *);
void goIupSetKillFocusFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_KILLFOCUS_CB", f);
	IupSetCallback(ih, "KILLFOCUS_CB", (Icallback) goIupKillFocusCB);
}

extern int goIupEnterWindowCB(void *);
void goIupSetEnterWindowFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_ENTERWINDOW_CB", f);
	IupSetCallback(ih, "ENTERWINDOW_CB", (Icallback) goIupEnterWindowCB);
}

extern int goIupLeaveWindowCB(void *);
void goIupSetLeaveWindowFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_LEAVEWINDOW_CB", f);
	IupSetCallback(ih, "LEAVEWINDOW_CB", (Icallback) goIupLeaveWindowCB);
}

extern int goIupKAnyCB(void *);
void goIupSetKAnyFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_K_ANY", f);
	IupSetCallback(ih, "K_ANY", (Icallback) goIupKAnyCB);
}

extern int goIupHelpCB(void *);
void goIupSetHelpFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_HELP_CB", f);
	IupSetCallback(ih, "HELP_CB", (Icallback) goIupHelpCB);
}

extern int goIupButtonCB(void *, int button, int pressed, int x, int y, void *status);
void goIupSetButtonFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_BUTTON_CB", f);
	IupSetCallback(ih, "BUTTON_CB", (Icallback) goIupButtonCB);
}

extern int goIupDropFilesCB(void *, void *filename, int num, int x, int y);
void goIupSetDropFilesFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_DROPFILES_CB", f);
	IupSetCallback(ih, "DROPFILES_CB", (Icallback) goIupDropFilesCB);	
}

extern int goIupActionCB(void *);
void goIupSetActionFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_ACTION", f);
	IupSetCallback(ih, "ACTION", (Icallback) goIupActionCB);
}

extern int goIupListActionCB(void *, void *text, int item, int state);
void goIupSetListActionFunc(Ihandle *ih, void *f) {	
	IupSetCallback(ih, "_GO_LIST_ACTION", f);
	IupSetCallback(ih, "ACTION", (Icallback) goIupListActionCB);
}

extern int goIupCaretCB(void *, int lin, int col, int pos);
void goIupSetCaretFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_CARET_CB", f);
	IupSetCallback(ih, "CARET_CB", (Icallback) goIupCaretCB);
}

extern int goIupDblclickCB(void *, int item, void *text);
void goIupSetDblclickFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_DBLCLICK_CB", f);
	IupSetCallback(ih, "DBLCLICK_CB", (Icallback) goIupDblclickCB);
}

extern int goIupEditCB(void *, int c, void *new_value);
void goIupSetEditFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_EDIT_CB", f);
	IupSetCallback(ih, "EDIT_CB", (Icallback) goIupEditCB);
}

extern int goIupMotionCB(void *, int x, int y, void *status);
void goIupSetMotionFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_MOTION_CB", f);
	IupSetCallback(ih, "MOTION_CB", (Icallback) goIupMotionCB);
}

extern int goIupMultiselectCB(void *, void *text);
void goIupSetMultiselectFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_MULTISELECT_CB", f);
	IupSetCallback(ih, "MULTISELECT_CB", (Icallback) goIupMultiselectCB);
}

extern int goIupValuechangedCB(void *);
void goIupSetValuechangedFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_VALUECHANGED_CB", f);
	IupSetCallback(ih, "VALUECHANGED_CB", (Icallback) goIupValuechangedCB);
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

func (ih *Ihandle) SetMapFunc(f MapFunc) {
	C.goIupSetMapFunc(ih.h, unsafe.Pointer(&f))
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

func (ih *Ihandle) SetUnmapFunc(f UnmapFunc) {
	C.goIupSetUnmapFunc(ih.h, unsafe.Pointer(&f))
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

func (ih *Ihandle) SetDestroyFunc(f DestroyFunc) {
	C.goIupSetDestroyFunc(ih.h, unsafe.Pointer(&f))
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

func (ih *Ihandle) SetGetFocusFunc(f GetFocusFunc) {
	C.goIupSetGetFocusFunc(ih.h, unsafe.Pointer(&f))
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

func (ih *Ihandle) SetKillFocusFunc(f KillFocusFunc) {
	C.goIupSetKillFocusFunc(ih.h, unsafe.Pointer(&f))
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

func (ih *Ihandle) SetEnterWindowFunc(f EnterWindowFunc) {
	C.goIupSetEnterWindowFunc(ih.h, unsafe.Pointer(&f))
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

func (ih *Ihandle) SetLeaveWindowFunc(f LeaveWindowFunc) {
	C.goIupSetLeaveWindowFunc(ih.h, unsafe.Pointer(&f))
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

func (ih *Ihandle) SetKAnyFunc(f KAnyFunc) {
	C.goIupSetKAnyFunc(ih.h, unsafe.Pointer(&f))
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

func (ih *Ihandle) SetHelpFunc(f HelpFunc) {
	C.goIupSetHelpFunc(ih.h, unsafe.Pointer(&f))
}

type ButtonFunc func(*Ihandle, int, int, int, int, string) int

//export goIupButtonCB
func goIupButtonCB(ih unsafe.Pointer, button, pressed, x, y int, status unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_BUTTON_CB")
	defer C.free(unsafe.Pointer(cName))
	
	f := *(*ButtonFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	goStatus := C.GoString((*C.char)(status))	
	return f(&Ihandle{h}, button, pressed, x, y, goStatus)
}

func (ih *Ihandle) SetButtonFunc(f ButtonFunc) {
	C.goIupSetButtonFunc(ih.h, unsafe.Pointer(&f))
}

type DropFilesFunc func(*Ihandle, string, int, int, int) int

//export goIupDropFilesCB
func goIupDropFilesCB(ih, filename unsafe.Pointer, num, x, y int) int {
	h := (*C.Ihandle)(ih)
	
	cName := C.CString("_GO_DROPFILES_CB")
	defer C.free(unsafe.Pointer(cName))
	
	f := *(*DropFilesFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	
	goFilename := C.GoString((*C.char)(filename))
	
	return f(&Ihandle{h}, goFilename, int(num), int(x), int(y))
}

func (ih *Ihandle) SetDropFilesFunc(f DropFilesFunc) {
	C.goIupSetDropFilesFunc(ih.h, unsafe.Pointer(&f))
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

type ListActionFunc func(ih *Ihandle, text string, item, state int) int

//export goIupListActionCB
func goIupListActionCB(ih, text unsafe.Pointer, item, state int) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_CARET_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*ListActionFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	goText := C.GoString((*C.char)(text))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, goText, int(item), int(state))
}

func (ih *Ihandle) SetListActionFunc(f ListActionFunc) {
	C.goIupSetListActionFunc(ih.h, unsafe.Pointer(&f))
}

type CaretFunc func(ih *Ihandle, lin, col, pos int) int

//export goIupCaretCB
func goIupCaretCB(ih unsafe.Pointer, lin, col, pos int) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_CARET_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*CaretFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, int(lin), int(col), int(pos))
}

func (ih *Ihandle) SetCaretFunc(f CaretFunc) {
	C.goIupSetCaretFunc(ih.h, unsafe.Pointer(&f))
}

type DblclickFunc func(ih *Ihandle, item int, text string) int

//export goIupDblclickCB
func goIupDblclickCB(ih unsafe.Pointer, item int, text unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_DBLCLICK_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*DblclickFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	goText := C.GoString((*C.char)(text))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, int(item), goText)
}

func (ih *Ihandle) SetDblclickFunc(f DblclickFunc) {
	C.goIupSetDblclickFunc(ih.h, unsafe.Pointer(&f))
}

type EditFunc func(ih *Ihandle, item int, text string) int

//export goIupEditCB
func goIupEditCB(ih unsafe.Pointer, item int, text unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_EDIT_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*EditFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	goText := C.GoString((*C.char)(text))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, item, goText)
}

func (ih *Ihandle) SetEditFunc(f EditFunc) {
	C.goIupSetEditFunc(ih.h, unsafe.Pointer(&f))
}

type MotionFunc func(ih *Ihandle, x, y int, status string) int

//export goIupMotionCB
func goIupMotionCB(ih unsafe.Pointer, x, y int, status unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_MOTION_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*MotionFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	goStatus := C.GoString((*C.char)(status))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, x, y, goStatus)
}

func (ih *Ihandle) SetMotionFunc(f MotionFunc) {
	C.goIupSetMotionFunc(ih.h, unsafe.Pointer(&f))
}

type MultiselectFunc func(ih *Ihandle, text string) int

//export goIupMultiselectCB
func goIupMultiselectCB(ih, text unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_MULTISELECT_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*MultiselectFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	goText := C.GoString((*C.char)(text))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, goText)
}

func (ih *Ihandle) SetMultiselectFunc(f MultiselectFunc) {
	C.goIupSetMultiselectFunc(ih.h, unsafe.Pointer(&f))
}

type ValuechangedFunc func(ih *Ihandle) int

//export goIupValuechangedCB
func goIupValuechangedCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_VALUECHANGED_CB")
	defer C.free(unsafe.Pointer(cName))

	f := *(*ValuechangedFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetValuechangedFunc(f ValuechangedFunc) {
	C.goIupSetValuechangedFunc(ih.h, unsafe.Pointer(&f))
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
