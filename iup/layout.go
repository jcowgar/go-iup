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

func iHandleArrayToC(ihs []*Ihandle) []*C.Ihandle {
	max := len(ihs)
	result := make([]*C.Ihandle, max + 1)
	
	for k, v := range ihs {
		result[k] = v.h
	}	
	result[max] = nil
	
	return result
}

func stringArrayToC(strs []string) []*C.char {
	max := len(strs)
	result := make([]*C.char, max + 1)
	
	for k, v := range strs {
		result[k] = C.CString(v)
	}
	result[max] = nil
	
	return result
}

func freeCStringArray(strs []*C.char) {
	for _, v := range strs {
		if v != nil {
			C.free(unsafe.Pointer(v))
		}
	}
}

func Create(classname string) *Ihandle {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))
	
	return &Ihandle{h: C.IupCreate(cClassname)}
}

func (ih *Ihandle) Destroy() {
	C.IupDestroy(ih.h)
}

func (ih *Ihandle) Map() int {
	return int(C.IupMap(ih.h))
}

func (ih *Ihandle) Unmap() {
	C.IupUnmap(ih.h)
}

func (ih *Ihandle) GetClassName() string {
	return C.GoString(C.IupGetClassName(ih.h))
}

func (ih *Ihandle) GetClassType() string {
	return C.GoString(C.IupGetClassType(ih.h))
}

func (ih *Ihandle) ClassMatch(classname string) bool {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))
	
	return int(C.IupClassMatch(ih.h, cClassname)) == 1
}

func (ih *Ihandle) SaveClassAttributes() {
	C.IupSaveClassAttributes(ih.h)
}

func (ih *Ihandle) CopyClassAttributes(dest *Ihandle) {
	C.IupCopyClassAttributes(ih.h, dest.h)
}

func SetClassDefaultAttribute(classname, name, value string) {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))
	
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.IupSetClassDefaultAttribute(cClassname, cName, cValue)
}

func hboxv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{h: C.IupHboxv(&ihs[0])}
}

func Hbox(args ...*Ihandle) *Ihandle {
	return hboxv(iHandleArrayToC(args))
}

func Hboxv(args []*Ihandle) *Ihandle {
	return hboxv(iHandleArrayToC(args))
}

func vboxv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{h: C.IupVboxv(&ihs[0])}
}

func Vbox(args ...*Ihandle) *Ihandle {
	return vboxv(iHandleArrayToC(args))
}

func Vboxv(args []*Ihandle) *Ihandle {
	return vboxv(iHandleArrayToC(args))
}
