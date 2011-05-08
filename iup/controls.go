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

void _IupSetButtonCallback(Ihandle *ih, void *p) {
	IupSetAttribute(ih, "GO_ACTION_P", p);
	IupSetCallback(ih, "ACTION", (Icallback)buttonCB);
}
*/
import "C"

import "unsafe"
import "fmt"

//export buttonCB
func buttonCB(ih unsafe.Pointer) int {
	fmt.Printf("Button Callback Got!\n")	
	return 0
}

func Button(title, action string) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	
	cAction := C.CString(action)
	defer C.free(unsafe.Pointer(cAction))
	
	return &Ihandle{h: C.IupButton(cTitle, cAction)}
}

func SetButtonCallback(ih *Ihandle) {
	C._IupSetButtonCallback(ih.h, unsafe.Pointer(ih))
}

func Label(title string) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	
	return &Ihandle{h: C.IupLabel(cTitle)}
}
