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
*/
import "C"
import "unsafe"

func Create(classname string) *Ihandle {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))

	return &Ihandle{H: C.IupCreate(cClassname)}
}

func Createv(classname string, args []string) *Ihandle {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))

	cArgs := stringArrayToC(args)
	defer freeCStringArray(cArgs)

	return &Ihandle{H: C.IupCreatev(cClassname, (*unsafe.Pointer)(unsafe.Pointer(&cArgs[0])))}
}

func Createp(classname string, args ...string) *Ihandle {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))

	cArgs := stringArrayToC(args)
	defer freeCStringArray(cArgs)

	return &Ihandle{H: C.IupCreatev(cClassname, (*unsafe.Pointer)(unsafe.Pointer(&cArgs[0])))}
}

func (ih *Ihandle) Destroy() {
	C.IupDestroy(ih.H)
}

func (ih *Ihandle) Map() int {
	return int(C.IupMap(ih.H))
}

func (ih *Ihandle) Unmap() {
	C.IupUnmap(ih.H)
}

func (ih *Ihandle) GetClassName() string {
	return C.GoString(C.IupGetClassName(ih.H))
}

func (ih *Ihandle) GetClassType() string {
	return C.GoString(C.IupGetClassType(ih.H))
}

func (ih *Ihandle) ClassMatch(classname string) bool {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))

	return int(C.IupClassMatch(ih.H, cClassname)) == 1
}

func (ih *Ihandle) SaveClassAttributes() {
	C.IupSaveClassAttributes(ih.H)
}

func (ih *Ihandle) CopyClassAttributes(dest *Ihandle) {
	C.IupCopyClassAttributes(ih.H, dest.H)
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
	ih := &Ihandle{H: C.IupFill()}

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}

	return ih
}

func hboxv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{H: C.IupHboxv(&ihs[0])}
}

func Hbox(args ...*Ihandle) *Ihandle {
	return hboxv(iHandleArrayToC(args))
}

func Hboxv(args []*Ihandle) *Ihandle {
	return hboxv(iHandleArrayToC(args))
}

func vboxv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{H: C.IupVboxv(&ihs[0])}
}

func Vbox(args ...*Ihandle) *Ihandle {
	return vboxv(iHandleArrayToC(args))
}

func Vboxv(args []*Ihandle) *Ihandle {
	return vboxv(iHandleArrayToC(args))
}

func zboxv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{H: C.IupZboxv(&ihs[0])}
}

func Zbox(args ...*Ihandle) *Ihandle {
	return zboxv(iHandleArrayToC(args))
}

func Zboxv(args []*Ihandle) *Ihandle {
	return zboxv(iHandleArrayToC(args))
}

func Radio(child *Ihandle) *Ihandle {
	return &Ihandle{H: C.IupRadio(child.H)}
}

func normalizerv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{H: C.IupNormalizerv(&ihs[0])}
}

func Normalizer(args ...*Ihandle) *Ihandle {
	return normalizerv(iHandleArrayToC(args))
}

func Normalizerv(args []*Ihandle) *Ihandle {
	return normalizerv(iHandleArrayToC(args))
}

func cboxv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{H: C.IupCboxv(&ihs[0])}
}

func Cbox(args ...*Ihandle) *Ihandle {
	return cboxv(iHandleArrayToC(args))
}

func Cboxv(args []*Ihandle) *Ihandle {
	return cboxv(iHandleArrayToC(args))
}

func Sbox(child *Ihandle) *Ihandle {
	return &Ihandle{H: C.IupSbox(child.H)}
}

func Split(child1, child2 *Ihandle, opts ...interface{}) *Ihandle {
	ih := &Ihandle{H: C.IupSplit(child1.H, child2.H)}

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
** Hierarchy
**
*******************************************************************************/

func (ih *Ihandle) Append(new_child *Ihandle) *Ihandle {
	result := C.IupAppend(ih.H, new_child.H)
	if result == nil {
		return nil
	}

	return ih
}

func (ih *Ihandle) Detach() {
	C.IupDetach(ih.H)
}

func (ih *Ihandle) Insert(ref_child, new_child *Ihandle) *Ihandle {
	result := C.IupInsert(ih.H, ref_child.H, new_child.H)
	if result == nil {
		return nil
	}

	return ih
}

func (child *Ihandle) Reparent(new_parent, ref_child *Ihandle) int {
	return int(C.IupReparent(child.H, new_parent.H, ref_child.H))
}

func (ih *Ihandle) GetParent() *Ihandle {
	return &Ihandle{H: C.IupGetParent(ih.H)}
}

func (ih *Ihandle) GetChild(pos int) *Ihandle {
	return &Ihandle{H: C.IupGetChild(ih.H, C.int(pos))}
}

func (ih *Ihandle) GetChildPos(child *Ihandle) int {
	return int(C.IupGetChildPos(ih.H, child.H))
}

func (ih *Ihandle) GetChildCount() int {
	return int(C.IupGetChildCount(ih.H))
}

func (ih *Ihandle) GetNextChild(child *Ihandle) *Ihandle {
	return &Ihandle{H: C.IupGetNextChild(ih.H, child.H)}
}

func (ih *Ihandle) GetBrother() *Ihandle {
	return &Ihandle{H: C.IupGetBrother(ih.H)}
}

func (ih *Ihandle) GetDialog() *Ihandle {
	return &Ihandle{H: C.IupGetDialog(ih.H)}
}

func (ih *Ihandle) GetDialogChild(name string) *Ihandle {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return &Ihandle{H: C.IupGetDialogChild(ih.H, cName)}
}

/*******************************************************************************
** 
** Utilities
**
*******************************************************************************/

func (ih *Ihandle) Refresh() {
	C.IupRefresh(ih.H)
}

func (ih *Ihandle) RefreshChildren() {
	C.IupRefreshChildren(ih.H)
}

func (ih *Ihandle) Update() {
	C.IupUpdate(ih.H)
}

func (ih *Ihandle) UpdateChildren() {
	C.IupUpdateChildren(ih.H)
}

func (ih *Ihandle) Redraw(children bool) {
	updateChildren := 0
	if children {
		updateChildren = 1
	}

	C.IupRedraw(ih.H, C.int(updateChildren))
}

func (ih *Ihandle) ConvertXYToPos(x, y int) int {
	return int(C.IupConvertXYToPos(ih.H, C.int(x), C.int(y)))
}
