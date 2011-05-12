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

Ihandle *_IupCreatev(const char *classname, char *params[]) {
	return IupCreatev(classname, (void *)params);
}
*/
import "C"
import "unsafe"

func Create(classname string) *Ihandle {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))
	
	return &Ihandle{h: C.IupCreate(cClassname)}
}

func Createv(classname string, args []string) *Ihandle {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))
	
	cArgs := stringArrayToC(args)
	defer freeCStringArray(cArgs)
	
	return &Ihandle{h: C._IupCreatev(cClassname, cArgs)}
}

func Createp(classname string, args ...string) *Ihandle {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))
	
	cArgs := stringArrayToC(args)
	defer freeCStringArray(cArgs)
	
	return &Ihandle{h: C._IupCreatev(cClassname, cArgs)}
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

/*******************************************************************************
** 
** Composition
**
*******************************************************************************/

func Fill(opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupFill()}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
		
	return ih
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

func zboxv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{h: C.IupZboxv(&ihs[0])}
}

func Zbox(args ...*Ihandle) *Ihandle {
	return zboxv(iHandleArrayToC(args))
}

func Zboxv(args []*Ihandle) *Ihandle {
	return zboxv(iHandleArrayToC(args))
}

func Radio(child *Ihandle) *Ihandle {
	return &Ihandle{h: C.IupRadio(child.h)}
}

func normalizerv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{h: C.IupNormalizerv(&ihs[0])}
}

func Normalizer(args ...*Ihandle) *Ihandle {
	return normalizerv(iHandleArrayToC(args))
}

func Normalizerv(args []*Ihandle) *Ihandle {
	return normalizerv(iHandleArrayToC(args))
}

func cboxv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{h: C.IupCboxv(&ihs[0])}
}

func Cbox(args ...*Ihandle) *Ihandle {
	return cboxv(iHandleArrayToC(args))
}

func Cboxv(args []*Ihandle) *Ihandle {
	return cboxv(iHandleArrayToC(args))
}

func Sbox(child *Ihandle) *Ihandle {
	return &Ihandle{h: C.IupSbox(child.h)}
}

func Split(child1, child2 *Ihandle) *Ihandle {
	return &Ihandle{h: C.IupSplit(child1.h, child2.h)}
}

/*******************************************************************************
** 
** Hierarchy
**
*******************************************************************************/

func (ih *Ihandle) Append(new_child *Ihandle) *Ihandle {
	result := C.IupAppend(ih.h, new_child.h)
	if result == nil {
		return nil
	}
	
	return ih
}

func (ih *Ihandle) Detach() {
	C.IupDetach(ih.h)
}

func (ih *Ihandle) Insert(ref_child, new_child *Ihandle) *Ihandle {
	result := C.IupInsert(ih.h, ref_child.h, new_child.h)
	if result == nil {
		return nil
	}
	
	return ih
}

func (child *Ihandle) Reparent(new_parent, ref_child *Ihandle) int {
	return int(C.IupReparent(child.h, new_parent.h, ref_child.h))
}

func (ih *Ihandle) GetParent() *Ihandle {
	return &Ihandle{h: C.IupGetParent(ih.h)}
}

func (ih *Ihandle) GetChild(pos int) *Ihandle {
	return &Ihandle{h: C.IupGetChild(ih.h, C.int(pos))}
}

func (ih *Ihandle) GetChildPos(child *Ihandle) int {
	return int(C.IupGetChildPos(ih.h, child.h))
}

func (ih *Ihandle) GetChildCount() int {
	return int(C.IupGetChildCount(ih.h))
}

func (ih *Ihandle) GetNextChild(child *Ihandle) *Ihandle {
	return &Ihandle{h: C.IupGetNextChild(ih.h, child.h)}
}

func (ih *Ihandle) GetBrother() *Ihandle {
	return &Ihandle{h: C.IupGetBrother(ih.h)}
}

func (ih *Ihandle) GetDialog() *Ihandle {
	return &Ihandle{h: C.IupGetDialog(ih.h)}
}

func (ih *Ihandle) GetDialogChild(name string) *Ihandle {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	return &Ihandle{h: C.IupGetDialogChild(ih.h, cName)}
}

/*******************************************************************************
** 
** Utilities
**
*******************************************************************************/

func (ih *Ihandle) Refresh() {
	C.IupRefresh(ih.h)
}

func (ih *Ihandle) RefreshChildren() {
	C.IupRefreshChildren(ih.h)
}

func (ih *Ihandle) Update() {
	C.IupUpdate(ih.h)
}

func (ih *Ihandle) UpdateChildren() {
	C.IupUpdateChildren(ih.h)
}

func (ih *Ihandle) Redraw(children bool) {
	updateChildren := 0
	if children {
		updateChildren = 1
	}
	
	C.IupRedraw(ih.h, C.int(updateChildren))
}

func (ih *Ihandle) ConvertXYToPos(x, y int) int {
	return int(C.IupConvertXYToPos(ih.h, C.int(x), C.int(y)))
}

