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

extern int buttonCB(void*);

void _IupSetButtonCallback(Ihandle *ih, void *f) {
	IupSetCallback(ih, "GO_ACTION_FUNC", f);
	IupSetCallback(ih, "ACTION", (Icallback) buttonCB);
}
*/
import "C"

import "unsafe"

//export buttonCB
func buttonCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("GO_ACTION_FUNC")
	defer C.free(unsafe.Pointer(cName))
	
	f := *(*ButtonActionCallback)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	f(&Ihandle{h: (*C.Ihandle)(ih)})
		
	return C.IUP_DEFAULT
}

type ButtonActionCallback func(*Ihandle) int

func Button(title, action string) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	
	cAction := C.CString(action)
	defer C.free(unsafe.Pointer(cAction))
	
	return &Ihandle{h: C.IupButton(cTitle, cAction)}
}

func (ih *Ihandle) SetButtonCallback(f ButtonActionCallback) {
	C._IupSetButtonCallback(ih.h, unsafe.Pointer(&f))
}

func Label(title string) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	
	return &Ihandle{h: C.IupLabel(cTitle)}
}
