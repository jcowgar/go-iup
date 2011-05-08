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

extern int goIupButtonActionCB(void *);

void goIupSetButtonActionFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_ACTION", f);
	IupSetCallback(ih, "ACTION", (Icallback) goIupButtonActionCB);
}
*/
import "C"
import "unsafe"

type ButtonActionFunc func(*Ihandle) int

//export goIupButtonActionCB
func goIupButtonActionCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_ACTION")
	defer C.free(unsafe.Pointer(cName))
	
	f := *(*ButtonActionFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	return f(&Ihandle{h: (*C.Ihandle)(ih)})
}

func (ih *Ihandle) SetButtonActionFunc(f ButtonActionFunc) {
	C.goIupSetButtonActionFunc(ih.h, unsafe.Pointer(&f))
}
