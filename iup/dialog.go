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

func Dialog(child *Ihandle) (ih *Ihandle) {
	return &Ihandle{h: C.IupDialog(child.h)}
}

func Show(ih *Ihandle) int {
	return int(C.IupShow(ih.h))
}

func Message(title, message string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))
	
	C.IupMessage(cTitle, cMessage)
}
