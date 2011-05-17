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
#include <iuptuio.h>
#include <iupim.h>
*/
import "C"

import "unsafe"

/*******************************************************************************
**
** Fonts
**
*******************************************************************************/

func MapFont(iupfont string) string {
	cIupfont := C.CString(iupfont)
	defer C.free(unsafe.Pointer(cIupfont))

	return C.GoString(C.IupMapFont(cIupfont))
}

func UnMapFont(driverfont string) string {
	cDriverfont := C.CString(driverfont)
	defer C.free(unsafe.Pointer(cDriverfont))

	return C.GoString(C.IupUnMapFont(cDriverfont))
}


/*******************************************************************************
**
** Images
**
*******************************************************************************/

func Image(width, height int, pixels []byte) *Ihandle {
	cPixels := byteArrayToCUCharArray(pixels)

	return &Ihandle{H: C.IupImage(C.int(width), C.int(height), &cPixels[0])}
}

func ImageRGB(width, height int, pixels []byte) *Ihandle {
	cPixels := byteArrayToCUCharArray(pixels)

	return &Ihandle{H: C.IupImageRGB(C.int(width), C.int(height), &cPixels[0])}
}

func ImageRGBA(width, height int, pixels []byte) *Ihandle {
	cPixels := byteArrayToCUCharArray(pixels)

	return &Ihandle{H: C.IupImageRGBA(C.int(width), C.int(height), &cPixels[0])}
}

func LoadImage(file_name string) *Ihandle {
	cFile_Name := C.CString(file_name)
	defer C.free(unsafe.Pointer(cFile_Name))

	return &Ihandle{H: C.IupLoadImage(cFile_Name)}
}

func (ih *Ihandle) SaveImage(file_name, format string) int {
	cFile_Name := C.CString(file_name)
	defer C.free(unsafe.Pointer(cFile_Name))

	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))

	return int(C.IupSaveImage(ih.H, cFile_Name, cFormat))
}

func (ih *Ihandle) SaveImageAsText(file_name, format, name string) int {
	cFile_Name := C.CString(file_name)
	defer C.free(unsafe.Pointer(cFile_Name))

	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return int(C.IupSaveImageAsText(ih.H, cFile_Name, cFormat, cName))
}

/*******************************************************************************
**
** Keyboard
**
*******************************************************************************/

func (ih *Ihandle) NextField() *Ihandle {
	return &Ihandle{H: C.IupNextField(ih.H)}
}

func (ih *Ihandle) PreviousField() *Ihandle {
	return &Ihandle{H: C.IupPreviousField(ih.H)}
}

func GetFocus() *Ihandle {
	return &Ihandle{H: C.IupGetFocus()}
}

func (ih *Ihandle) SetFocus() *Ihandle {
	return &Ihandle{H: C.IupSetFocus(ih.H)}
}

/*******************************************************************************
**
** Menus
**
*******************************************************************************/

func Item(title string, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	ih := &Ihandle{H: C.IupItem(cTitle, nil)}

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)

		case ActionFunc:
			ih.SetActionFunc(v)
		}
	}

	return ih
}

func menuv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{H: C.IupMenuv(&ihs[0])}
}

func Menu(args ...*Ihandle) *Ihandle {
	return menuv(iHandleArrayToC(args))
}

func Menuv(args []*Ihandle, opts ...interface{}) *Ihandle {
	ih := menuv(iHandleArrayToC(args))

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}

	return ih
}

func Separator() *Ihandle {
	return &Ihandle{H: C.IupSeparator()}
}

func Submenu(title string, menu *Ihandle, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	ih := &Ihandle{H: C.IupSubmenu(cTitle, menu.H)}

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}

	return ih
}

/*******************************************************************************
**
** Miscellaneous
**
*******************************************************************************/

func Clipboard() *Ihandle {
	return &Ihandle{H: C.IupClipboard()}
}

func Timer() *Ihandle {
	return &Ihandle{H: C.IupTimer()}
}

func TuioClient(port int) *Ihandle {
	return &Ihandle{H: C.IupTuioClient(C.int(port))}
}

func User() *Ihandle {
	return &Ihandle{H: C.IupUser()}
}

func Help(url string) int {
	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))

	return int(C.IupHelp(cUrl))
}
