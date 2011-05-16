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
#include <stdio.h>
#include <iup.h>

int _IupGetParam(const char *title, const char *format, void **args) {
	return IupGetParam(title, NULL, 0, format, 
		args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9],
		args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17], args[18], 
		args[19]
	);
}
*/
import "C"
import "unsafe"

// Differs from IupDialog in that any number of parameters may be passed after the child
// widget. Strings will be interpreted as attributes to set on the newly created dialog.
func Dialog(child *Ihandle, opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupDialog(child.h)}

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}

	return ih
}

func (ih *Ihandle) Show() int {
	return int(C.IupShow(ih.h))
}

func (ih *Ihandle) ShowXY(x, y int) int {
	return int(C.IupShowXY(ih.h, C.int(x), C.int(y)))
}

func (ih *Ihandle) Hide() int {
	return int(C.IupHide(ih.h))
}

func Popup(ih *Ihandle, x, y int) int {
	return int(C.IupPopup(ih.h, C.int(x), C.int(y)))
}

/*******************************************************************************
**
** Predefined
**
*******************************************************************************/

func FileDlg(opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupFileDlg()}

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}

	return ih
}

func MessageDlg(opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupMessageDlg()}

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}

	return ih
}

func ColorDlg(opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupColorDlg()}

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}

	return ih
}

func FontDlg(opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupFontDlg()}

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}

	return ih
}

func Alarm(t, m string, buttons ...string) int {
	var b1, b2, b3 *C.char

	cT := C.CString(t)
	defer C.free(unsafe.Pointer(cT))

	cM := C.CString(m)
	defer C.free(unsafe.Pointer(cM))

	switch len(buttons) {
	case 1:
		b1 = C.CString(buttons[0])
		b2 = nil
		b3 = nil

		defer C.free(unsafe.Pointer(b1))

	case 2:
		b1 = C.CString(buttons[0])
		b2 = C.CString(buttons[1])
		b3 = nil

		defer C.free(unsafe.Pointer(b1))
		defer C.free(unsafe.Pointer(b2))

	case 3:
		b1 = C.CString(buttons[0])
		b2 = C.CString(buttons[1])
		b3 = C.CString(buttons[2])

		defer C.free(unsafe.Pointer(b1))
		defer C.free(unsafe.Pointer(b2))
		defer C.free(unsafe.Pointer(b3))

	default:
		panic("iup.Alarm() requires at least 1 button and at most 3 buttons")
	}

	return int(C.IupAlarm(cT, cM, b1, b2, b3))
}

// Warning: This allocates 2048 for an incoming buffer size for the filename but if 
// the user selects a filename more than 2048 characters a buffer over run WILL 
// occur. This problem has been brought to the attention of the Iup maintainers as 
// there is no way to tell GetFile how large the receiving buffer actually is.
func GetFile(filename string) (string, int) {
	cFilename := make([]C.char, 2048)
	for i, v := range filename {
		cFilename[i] = (C.char)(v)
	}

	result := C.IupGetFile(&cFilename[0])
	return C.GoString(&cFilename[0]), int(result)
}

// Return value is r, g, b. Each will contain -1 if cancel was pressed
func GetColor(x, y, r, g, b int) (int, int, int) {
	cR := (C.uchar)(r)
	cG := (C.uchar)(g)
	cB := (C.uchar)(b)

	result := C.IupGetColor(C.int(x), C.int(y), &cR, &cG, &cB)
	if int(result) == 1 {
		return int(cR), int(cG), int(cB)
	}

	return -1, -1, -1
}

// Differs in that size is derrived from the list array
func ListDialog(typ int, title string, list []string, opt, max_col, max_lin int, marks []int) (int, []int) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	cList := stringArrayToC(list)
	defer freeCStringArray(cList)

	cMarks := intArrayToC(marks)

	result := C.IupListDialog(C.int(typ), cTitle, C.int(len(list)), &cList[0], C.int(opt), C.int(max_col),
		C.int(max_lin), &cMarks[0])

	for k, v := range cMarks {
		marks[k] = int(v)
	}

	return int(result), marks
}

const maxArgs = 20

// Warning: Incompelete, only supports int64, float64 and bool variable types. This
// method is a work in progress
func GetParam(title string, format string, args ...interface{}) bool {
	if len(args) > maxArgs {
		panic("too many args")
	}

	cargs := new([maxArgs]unsafe.Pointer)

	// copy values in
	for i, a := range args {
		switch v := a.(type) {
		case *int:
			p := new(C.int)
			*p = C.int(*v)
			cargs[i] = unsafe.Pointer(p)

		case *float64:
			p := new(C.float)
			*p = C.float(*v)
			cargs[i] = unsafe.Pointer(p)

		case *bool:
			p := new(C.int)
			if *v {
				*p = 1
			} else {
				*p = 0
			}
			cargs[i] = unsafe.Pointer(p)

		case *string:
			p := make([]byte, 4096)
			copy(p, *v)
			p[len(*v)] = 0
			cargs[i] = unsafe.Pointer(&p[0])

		default:
			panic("unknown type")
		}
	}

	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))

	result := int(C._IupGetParam(cTitle, cFormat, (*unsafe.Pointer)(unsafe.Pointer(cargs))))

	if result == 0 {
		return false
	}

	for i, a := range args {
		switch v := a.(type) {
		case *int:
			*v = int(*(*C.int)(cargs[i]))

		case *float64:
			*v = float64(*(*C.float)(cargs[i]))

		case *bool:
			*v = int(*(*C.int)(cargs[i])) == 1

		case *string:
			*v = C.GoString((*C.char)(cargs[i]))
		}
	}

	return true
}

// Warning: This allocates 4096 for an incoming buffer size but if the user enters
// more than 4096 characters a buffer over run WILL occur. This problem has been
// brought to the attention of the Iup maintainers as there is no way to tell
// IupGetText how large the receiving buffer actually is.
func GetText(title, text string) (string, int) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	cText := make([]C.char, 4096)
	for i, v := range text {
		cText[i] = (C.char)(v)
	}

	result := C.IupGetText(cTitle, &cText[0])

	return C.GoString(&cText[0]), int(result)

}

func Message(title, message string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))

	C.IupMessage(cTitle, cMessage)
}

func LayoutDialog(ih *Ihandle) *Ihandle {
	return &Ihandle{h: C.IupLayoutDialog(ih.h)}
}

func ElementPropertiesDialog(ih *Ihandle) *Ihandle {
	return &Ihandle{h: C.IupElementPropertiesDialog(ih.h)}
}
