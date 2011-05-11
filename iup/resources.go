/* 
	Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>
	
	This file is part of iup.go.

	iup.go is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as
	published by the Free Software Foundation, either version 3 of
	the License, or (at your option) any later version.
	
	iup.go is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.
	
	You should have received a copy of the GNU Lesser General Public
	License along with iup.go.  If not, see <http://www.gnu.org/licenses/>.
*/

package iup

/*
#include <stdlib.h>
#include <iup.h>
*/
import "C"

import "unsafe"

/*******************************************************************************
**
** Menus
**
*******************************************************************************/

func Item(title string, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	
	ih := &Ihandle{h: C.IupItem(cTitle, nil)}
	
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
	return &Ihandle{h: C.IupMenuv(&ihs[0])}
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
	return &Ihandle{h: C.IupSeparator()}
}

func Submenu(title string, menu *Ihandle, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	
	ih := &Ihandle{h: C.IupSubmenu(cTitle, menu.h)}
	
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

func Help(url string) int {
	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))
	
	return int(C.IupHelp(cUrl))
}

