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
	cPixels := ByteArrayToCUCharArray(pixels)

	return (*Ihandle)(C.IupImage(C.int(width), C.int(height), &cPixels[0]))
}

func ImageRGB(width, height int, pixels []byte) *Ihandle {
	cPixels := ByteArrayToCUCharArray(pixels)

	return (*Ihandle)(C.IupImageRGB(C.int(width), C.int(height), &cPixels[0]))
}

func ImageRGBA(width, height int, pixels []byte) *Ihandle {
	cPixels := ByteArrayToCUCharArray(pixels)

	return (*Ihandle)(C.IupImageRGBA(C.int(width), C.int(height), &cPixels[0]))
}

func LoadImage(file_name string) *Ihandle {
	cFile_Name := C.CString(file_name)
	defer C.free(unsafe.Pointer(cFile_Name))

	return (*Ihandle)(C.IupLoadImage(cFile_Name))
}

func SaveImage(ih *Ihandle, file_name, format string) int {
	cFile_Name := C.CString(file_name)
	defer C.free(unsafe.Pointer(cFile_Name))

	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))

	return int(C.IupSaveImage(ih.C(), cFile_Name, cFormat))
}

func SaveImageAsText(ih *Ihandle, file_name, format, name string) int {
	cFile_Name := C.CString(file_name)
	defer C.free(unsafe.Pointer(cFile_Name))

	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return int(C.IupSaveImageAsText(ih.C(), cFile_Name, cFormat, cName))
}

/*******************************************************************************
**
** Keyboard
**
*******************************************************************************/

func NextField(ih *Ihandle) *Ihandle {
	return (*Ihandle)(C.IupNextField(ih.C()))
}

func PreviousField(ih *Ihandle) *Ihandle {
	return (*Ihandle)(C.IupPreviousField(ih.C()))
}

func GetFocus() *Ihandle {
	return (*Ihandle)(C.IupGetFocus())
}

func SetFocus(ih *Ihandle) *Ihandle {
	return (*Ihandle)(C.IupSetFocus(ih.C()))
}

/*******************************************************************************
**
** Menus
**
*******************************************************************************/

func Item(title string, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	ih := (*Ihandle)(C.IupItem(cTitle, nil))

	for _, o := range opts {
		switch v := o.(type) {
		case ActionFunc:
			SetActionFunc(ih, v)
			
		default:
			Decorate(ih, v)
		}
	}

	return ih
}

func menuv(ihs []*C.Ihandle) *Ihandle {
	return (*Ihandle)(C.IupMenuv(&ihs[0]))
}

func Menu(args ...*Ihandle) *Ihandle {
	return menuv(IHandleArrayToC(args))
}

func Menuv(args []*Ihandle, opts ...interface{}) *Ihandle {
	ih := menuv(IHandleArrayToC(args))

	for _, o := range opts {
		switch v := o.(type) {
		default:
			Decorate(ih, v)
		}
	}

	return ih
}

func Separator() *Ihandle {
	return (*Ihandle)(C.IupSeparator())
}

func Submenu(title string, menu *Ihandle, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	ih := (*Ihandle)(C.IupSubmenu(cTitle, menu.C()))

	for _, o := range opts {
		switch v := o.(type) {
		default:
			Decorate(ih, v)
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
	return (*Ihandle)(C.IupClipboard())
}

func Timer() *Ihandle {
	return (*Ihandle)(C.IupTimer())
}

func User() *Ihandle {
	return (*Ihandle)(C.IupUser())
}

func Help(url string) int {
	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))

	return int(C.IupHelp(cUrl))
}
