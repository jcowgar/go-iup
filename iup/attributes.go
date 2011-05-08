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

func (ih *Ihandle) StoreAttribute(name, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	
	C.IupStoreAttribute(ih.h, cName, cValue)
}

func (ih *Ihandle) StoreAttributeId(name string, id int, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	
	C.IupStoreAttributeId(ih.h, cName, C.int(id), cValue)
}

func (ih *Ihandle) GetAttribute(name string) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	cVal := C.IupGetAttribute(ih.h, cName)
	
	return C.GoString(cVal)
}
