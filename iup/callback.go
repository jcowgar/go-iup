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

// IUP definitions not defined
#define IUP_UNMAP_CB        "UNMAP_CB"
#define IUP_DESTROY_CB      "DESTROY_CB"
#define IUP_CARET_CB        "CARET_CB"
#define IUP_DBLCLICK_CB     "DBLCLICK_CB"
#define IUP_EDIT_CB         "EDIT_CB"
#define IUP_MULTISELECT_CB  "MULTISELECT_CB"
#define IUP_VALUECHANGED_CB "VALUECHANGED_CB"
#define IUP_TABCHANGE_CB    "TABCHANGE_CB"
#define IUP_TABCHANGEPOS_CB "TABCHANGEPOS_CB"
#define IUP_SPIN_CB         "SPIN_CB"

#define GO_PREFIX "_GO_"

const char *GO_MAP_CB = GO_PREFIX IUP_MAP_CB;
extern int goIupMapCB(void *);

void goIupSetMapFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_MAP_CB, f);
	IupSetCallback(ih, IUP_MAP_CB, (Icallback) goIupMapCB);
}

const char *GO_UNMAP_CB = GO_PREFIX IUP_UNMAP_CB;
extern int goIupUnmapCB(void *);

void goIupSetUnmapFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_UNMAP_CB, f);
	IupSetCallback(ih, IUP_UNMAP_CB, (Icallback) goIupUnmapCB);
}

const char *GO_DESTROY_CB = GO_PREFIX IUP_DESTROY_CB;
extern int goIupDestroyCB(void *);

void goIupSetDestroyFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_DESTROY_CB, f);
	IupSetCallback(ih, IUP_DESTROY_CB, (Icallback) goIupDestroyCB);
}

const char *GO_GETFOCUS_CB = GO_PREFIX IUP_GETFOCUS_CB;
extern int goIupGetFocusCB(void *);

void goIupSetGetFocusFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_GETFOCUS_CB, f);
	IupSetCallback(ih, IUP_GETFOCUS_CB, (Icallback) goIupGetFocusCB);
}

const char *GO_KILLFOCUS_CB = GO_PREFIX IUP_KILLFOCUS_CB;
extern int goIupKillFocusCB(void *);

void goIupSetKillFocusFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_KILLFOCUS_CB, f);
	IupSetCallback(ih, IUP_KILLFOCUS_CB, (Icallback) goIupKillFocusCB);
}

const char *GO_ENTERWINDOW_CB = GO_PREFIX IUP_ENTERWINDOW_CB;
extern int goIupEnterWindowCB(void *);

void goIupSetEnterWindowFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_ENTERWINDOW_CB, f);
	IupSetCallback(ih, IUP_ENTERWINDOW_CB, (Icallback) goIupEnterWindowCB);
}

const char *GO_LEAVEWINDOW_CB = GO_PREFIX IUP_LEAVEWINDOW_CB;
extern int goIupLeaveWindowCB(void *);

void goIupSetLeaveWindowFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_LEAVEWINDOW_CB, f);
	IupSetCallback(ih, IUP_LEAVEWINDOW_CB, (Icallback) goIupLeaveWindowCB);
}

const char *GO_K_ANY_CB = GO_PREFIX IUP_K_ANY;
extern int goIupKAnyCB(void *);

void goIupSetKAnyFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_K_ANY_CB, f);
	IupSetCallback(ih, IUP_K_ANY, (Icallback) goIupKAnyCB);
}

const char *GO_HELP_CB = GO_PREFIX IUP_HELP_CB;
extern int goIupHelpCB(void *);

void goIupSetHelpFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_HELP_CB, f);
	IupSetCallback(ih, IUP_HELP_CB, (Icallback) goIupHelpCB);
}

const char *GO_BUTTON_CB = GO_PREFIX IUP_BUTTON_CB;
extern int goIupButtonCB(void *, int button, int pressed, int x, int y, void *status);

void goIupSetButtonFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_BUTTON_CB, f);
	IupSetCallback(ih, IUP_BUTTON_CB, (Icallback) goIupButtonCB);
}

const char *GO_DROPFILES_CB = GO_PREFIX IUP_DROPFILES_CB;
extern int goIupDropFilesCB(void *, void *filename, int num, int x, int y);

void goIupSetDropFilesFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_DROPFILES_CB, f);
	IupSetCallback(ih, IUP_DROPFILES_CB, (Icallback) goIupDropFilesCB);	
}

extern int goIupActionCB(void *);
const char *GO_ACTION = GO_PREFIX IUP_ACTION;

void goIupSetActionFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_ACTION, f);
	IupSetCallback(ih, IUP_ACTION, (Icallback) goIupActionCB);
}

void goSetFunc(Ihandle *ih, char *goName, void *gof, char *cName, void *cf) {
	IupSetCallback(ih, goName, gof);
	IupSetCallback(ih, cName, cf);
}

const char *GO_LIST_ACTION = GO_PREFIX IUP_ACTION;
extern int goIupListActionCB(void *, void *text, int item, int state);

void goIupSetListActionFunc(Ihandle *ih, void *f) {	
	IupSetCallback(ih, GO_LIST_ACTION, f);
	IupSetCallback(ih, IUP_ACTION, (Icallback) goIupListActionCB);
}

const char *GO_CARET_CB = GO_PREFIX IUP_CARET_CB;
extern int goIupCaretCB(void *, int lin, int col, int pos);

void goIupSetCaretFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_CARET_CB, f);
	IupSetCallback(ih, IUP_CARET_CB, (Icallback) goIupCaretCB);
}

const char *GO_DBLCLICK_CB = GO_PREFIX IUP_DBLCLICK_CB;
extern int goIupDblclickCB(void *, int item, void *text);

void goIupSetDblclickFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_DBLCLICK_CB, f);
	IupSetCallback(ih, IUP_DBLCLICK_CB, (Icallback) goIupDblclickCB);
}

const char *GO_EDIT_CB = GO_PREFIX IUP_EDIT_CB;
extern int goIupEditCB(void *, int c, void *new_value);

void goIupSetEditFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_EDIT_CB, f);
	IupSetCallback(ih, IUP_EDIT_CB, (Icallback) goIupEditCB);
}

const char *GO_MOTION_CB = GO_PREFIX IUP_MOTION_CB;
extern int goIupMotionCB(void *, int x, int y, void *status);

void goIupSetMotionFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_MOTION_CB, f);
	IupSetCallback(ih, IUP_MOTION_CB, (Icallback) goIupMotionCB);
}

const char *GO_MULTISELECT_CB = GO_PREFIX IUP_MULTISELECT_CB;
extern int goIupMultiselectCB(void *, void *text);

void goIupSetMultiselectFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_MULTISELECT_CB, f);
	IupSetCallback(ih, IUP_MULTISELECT_CB, (Icallback) goIupMultiselectCB);
}

const char *GO_VALUECHANGED_CB = GO_PREFIX IUP_VALUECHANGED_CB;
extern int goIupValueChangedCB(void *);

void goIupSetValueChangedFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_VALUECHANGED_CB, f);
	IupSetCallback(ih, IUP_VALUECHANGED_CB, (Icallback) goIupValueChangedCB);
}

extern int goIupTextActionCB(void *ih, int ch, void *newValue);
void goIupSetTextActionFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_ACTION, f);
	IupSetCallback(ih, IUP_ACTION, (Icallback) goIupTextActionCB);
}

extern int goIupToggleActionCB(void *, int state);
void goIupSetToggleActionFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_ACTION, f);
	IupSetCallback(ih, IUP_ACTION, (Icallback) goIupToggleActionCB);
}

const char *GO_TABCHANGE_CB = GO_PREFIX IUP_TABCHANGE_CB;
extern int goIupTabChangeCB(void *ih, void *new_tab, void *old_tab);
void goIupSetTabChangeFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_TABCHANGE_CB, f);
	IupSetCallback(ih, IUP_TABCHANGE_CB, (Icallback) goIupTabChangeCB);
}

const char *GO_TABCHANGEPOS_CB = GO_PREFIX IUP_TABCHANGEPOS_CB;
extern int goIupTabChangePosCB(void *ih, int old_pos, int new_pos);
void goIupSetTabChangePosFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_TABCHANGEPOS_CB, f);
	IupSetCallback(ih, IUP_TABCHANGEPOS_CB, (Icallback) goIupTabChangePosCB);
}

const char *GO_SPIN_CB = GO_PREFIX IUP_SPIN_CB;
extern int goIupSpinCB(void *ih, int inc);
void goIupSetSpinFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_SPIN_CB, f);
	IupSetCallback(ih, IUP_SPIN_CB, (Icallback) goIupSpinCB);
}
*/
import "C"
import "unsafe"

// Common callbacks

type MapFunc func(*Ihandle) int

//export goIupMapCB
func goIupMapCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*MapFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_MAP_CB)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetMapFunc(f MapFunc) {
	C.goIupSetMapFunc(ih.h, unsafe.Pointer(&f))
}

type UnmapFunc func(*Ihandle) int

//export goIupUnmapCB
func goIupUnmapCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*UnmapFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_UNMAP_CB)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetUnmapFunc(f UnmapFunc) {
	C.goIupSetUnmapFunc(ih.h, unsafe.Pointer(&f))
}

type DestroyFunc func(*Ihandle) int

//export goIupDestroyCB
func goIupDestroyCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*DestroyFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_DESTROY_CB)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetDestroyFunc(f DestroyFunc) {
	C.goIupSetDestroyFunc(ih.h, unsafe.Pointer(&f))
}

type GetFocusFunc func(*Ihandle) int

//export goIupGetFocusCB
func goIupGetFocusCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*GetFocusFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_GETFOCUS_CB)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetGetFocusFunc(f GetFocusFunc) {
	C.goIupSetGetFocusFunc(ih.h, unsafe.Pointer(&f))
}

type KillFocusFunc func(*Ihandle) int

//export goIupKillFocusCB
func goIupKillFocusCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*KillFocusFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_KILLFOCUS_CB)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetKillFocusFunc(f KillFocusFunc) {
	C.goIupSetKillFocusFunc(ih.h, unsafe.Pointer(&f))
}

type EnterWindowFunc func(*Ihandle) int

//export goIupEnterWindowCB
func goIupEnterWindowCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*EnterWindowFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_ENTERWINDOW_CB)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetEnterWindowFunc(f EnterWindowFunc) {
	C.goIupSetEnterWindowFunc(ih.h, unsafe.Pointer(&f))
}

type LeaveWindowFunc func(*Ihandle) int

//export goIupLeaveWindowCB
func goIupLeaveWindowCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*LeaveWindowFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_LEAVEWINDOW_CB)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetLeaveWindowFunc(f LeaveWindowFunc) {
	C.goIupSetLeaveWindowFunc(ih.h, unsafe.Pointer(&f))
}

type KAnyFunc func(*Ihandle, int) int

//export goIupKAnyCB
func goIupKAnyCB(ih unsafe.Pointer, c C.int) int {
	h := (*C.Ihandle)(ih)
	f := *(*KAnyFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_K_ANY_CB)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, int(c))
}

func (ih *Ihandle) SetKAnyFunc(f KAnyFunc) {
	C.goIupSetKAnyFunc(ih.h, unsafe.Pointer(&f))
}

type HelpFunc func(*Ihandle) int

//export goIupHelpCB
func goIupHelpCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*HelpFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_HELP_CB)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetHelpFunc(f HelpFunc) {
	C.goIupSetHelpFunc(ih.h, unsafe.Pointer(&f))
}

type ButtonFunc func(*Ihandle, int, int, int, int, string) int

//export goIupButtonCB
func goIupButtonCB(ih unsafe.Pointer, button, pressed, x, y int, status unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*ButtonFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_BUTTON_CB)))
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
	f := *(*DropFilesFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_DROPFILES_CB)))
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
	f := *(*ActionFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_ACTION)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetActionFunc(f ActionFunc) {
	C.goIupSetActionFunc(ih.h, unsafe.Pointer(&f))
}

type ListActionFunc func(ih *Ihandle, text string, item, state int) int

//export goIupListActionCB
func goIupListActionCB(ih, text unsafe.Pointer, item, state int) int {
	h := (*C.Ihandle)(ih)
	f := *(*ListActionFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_LIST_ACTION)))
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
	f := *(*CaretFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_CARET_CB)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, int(lin), int(col), int(pos))
}

func (ih *Ihandle) SetCaretFunc(f CaretFunc) {
	C.goIupSetCaretFunc(ih.h, unsafe.Pointer(&f))
}

type DblclickFunc func(ih *Ihandle, item int, text string) int

//export goIupDblclickCB
func goIupDblclickCB(ih unsafe.Pointer, item int, text unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*DblclickFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_DBLCLICK_CB)))
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
	f := *(*EditFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_EDIT_CB)))
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
	f := *(*MotionFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_MOTION_CB)))
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
	f := *(*MultiselectFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_MULTISELECT_CB)))
	goText := C.GoString((*C.char)(text))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, goText)
}

func (ih *Ihandle) SetMultiselectFunc(f MultiselectFunc) {
	C.goIupSetMultiselectFunc(ih.h, unsafe.Pointer(&f))
}

type ValueChangedFunc func(ih *Ihandle) int

//export goIupValueChangedCB
func goIupValueChangedCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*ValueChangedFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_VALUECHANGED_CB)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetValueChangedFunc(f ValueChangedFunc) {
	C.goIupSetValueChangedFunc(ih.h, unsafe.Pointer(&f))
}

type TextActionFunc func(ih *Ihandle, ch int, newValue string) int

//export goIupTextActionCB
func goIupTextActionCB(ih unsafe.Pointer, ch int, newValue unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*TextActionFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_ACTION)))
	goNewValue := C.GoString((*C.char)(newValue))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, ch, goNewValue)
}

func (ih *Ihandle) SetTextActionFunc(f TextActionFunc) {
	C.goIupSetTextActionFunc(ih.h, unsafe.Pointer(&f))
}

type ToggleActionFunc func(ih *Ihandle, state int) int

//export goIupToggleActionCB
func goIupToggleActionCB(ih unsafe.Pointer, state int) int {
	h := (*C.Ihandle)(ih)
	f := *(*ToggleActionFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_ACTION)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, state)
}

func (ih *Ihandle) SetToggleActionFunc(f ToggleActionFunc) {
	C.goIupSetToggleActionFunc(ih.h, unsafe.Pointer(&f))
}

type TabChangeFunc func(ih, new_tab, old_tab *Ihandle) int

//export goIupTabChangeCB
func goIupTabChangeCB(ih, new_tab, old_tab unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*TabChangeFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_TABCHANGE_CB)))
	return f(&Ihandle{h}, &Ihandle{(*C.Ihandle)(new_tab)}, &Ihandle{(*C.Ihandle)(old_tab)})
}

func (ih *Ihandle) SetTabChangeFunc(f TabChangeFunc) {
	C.goIupSetTabChangeFunc(ih.h, unsafe.Pointer(&f))
}

type TabChangePosFunc func(ih *Ihandle, new_pos, old_pos int) int

//export goIupTabChangePosCB
func goIupTabChangePosCB(ih unsafe.Pointer, new_pos, old_pos int) int {
	h := (*C.Ihandle)(ih)
	f := *(*TabChangePosFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_TABCHANGEPOS_CB)))
	return f(&Ihandle{h}, new_pos, old_pos)
}

func (ih *Ihandle) SetTabChangePosFunc(f TabChangePosFunc) {
	C.goIupSetTabChangePosFunc(ih.h, unsafe.Pointer(&f))
}

type SpinFunc func(ih *Ihandle, inc int) int

//export goIupSpinCB
func goIupSpinCB(ih unsafe.Pointer, inc int) int {
	h := (*C.Ihandle)(ih)
	f := *(*SpinFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_SPIN_CB)))
	return f(&Ihandle{h}, inc)
}

func (ih *Ihandle) SetSpinFunc(f SpinFunc) {
	C.goIupSetSpinFunc(ih.h, unsafe.Pointer(&f))
}
