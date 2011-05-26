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

	return (*Ihandle)(C.IupCreate(cClassname))
}

func Createv(classname string, args []string) *Ihandle {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))

	cArgs := StringArrayToC(args)
	defer FreeCStringArray(cArgs)

	return (*Ihandle)(C.IupCreatev(cClassname, (*unsafe.Pointer)(unsafe.Pointer(&cArgs[0]))))
}

func Createp(classname string, args ...string) *Ihandle {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))

	cArgs := StringArrayToC(args)
	defer FreeCStringArray(cArgs)

	return (*Ihandle)(C.IupCreatev(cClassname, (*unsafe.Pointer)(unsafe.Pointer(&cArgs[0]))))
}

func Destroy(ih *Ihandle) {
	C.IupDestroy(ih.C())
}

func Map(ih *Ihandle) int {
	return int(C.IupMap(ih.C()))
}

func Unmap(ih *Ihandle) {
	C.IupUnmap(ih.C())
}

func GetClassName(ih *Ihandle) string {
	return C.GoString(C.IupGetClassName(ih.C()))
}

func GetClassType(ih *Ihandle) string {
	return C.GoString(C.IupGetClassType(ih.C()))
}

func ClassMatch(ih *Ihandle, classname string) bool {
	cClassname := C.CString(classname)
	defer C.free(unsafe.Pointer(cClassname))

	return int(C.IupClassMatch(ih.C(), cClassname)) == 1
}

func SaveClassAttributes(ih *Ihandle) {
	C.IupSaveClassAttributes(ih.C())
}

func CopyClassAttributes(ih, dest *Ihandle) {
	C.IupCopyClassAttributes(ih.C(), dest.C())
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
	ih := (*Ihandle)(C.IupFill())

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			SetAttributes(ih, v)
		}
	}

	return ih
}

func hboxv(ihs []*C.Ihandle) *Ihandle {
	return (*Ihandle)(C.IupHboxv(&ihs[0]))
}

func Hbox(args ...*Ihandle) *Ihandle {
	return hboxv(IHandleArrayToC(args))
}

func Hboxv(args []*Ihandle) *Ihandle {
	return hboxv(IHandleArrayToC(args))
}

func vboxv(ihs []*C.Ihandle) *Ihandle {
	return (*Ihandle)(C.IupVboxv(&ihs[0]))
}

func Vbox(args ...*Ihandle) *Ihandle {
	return vboxv(IHandleArrayToC(args))
}

func Vboxv(args []*Ihandle) *Ihandle {
	return vboxv(IHandleArrayToC(args))
}

func zboxv(ihs []*C.Ihandle) *Ihandle {
	return (*Ihandle)(C.IupZboxv(&ihs[0]))
}

func Zbox(args ...*Ihandle) *Ihandle {
	return zboxv(IHandleArrayToC(args))
}

func Zboxv(args []*Ihandle) *Ihandle {
	return zboxv(IHandleArrayToC(args))
}

func Radio(child *Ihandle) *Ihandle {
	return (*Ihandle)(C.IupRadio(child.C()))
}

func normalizerv(ihs []*C.Ihandle) *Ihandle {
	return (*Ihandle)(C.IupNormalizerv(&ihs[0]))
}

func Normalizer(args ...*Ihandle) *Ihandle {
	return normalizerv(IHandleArrayToC(args))
}

func Normalizerv(args []*Ihandle) *Ihandle {
	return normalizerv(IHandleArrayToC(args))
}

func cboxv(ihs []*C.Ihandle) *Ihandle {
	return (*Ihandle)(C.IupCboxv(&ihs[0]))
}

func Cbox(args ...*Ihandle) *Ihandle {
	return cboxv(IHandleArrayToC(args))
}

func Cboxv(args []*Ihandle) *Ihandle {
	return cboxv(IHandleArrayToC(args))
}

func Sbox(child *Ihandle) *Ihandle {
	return (*Ihandle)(C.IupSbox(child.C()))
}

func Split(child1, child2 *Ihandle, opts ...interface{}) *Ihandle {
	ih := (*Ihandle)(C.IupSplit(child1.C(), child2.C()))

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			SetAttributes(ih, v)
		}
	}

	return ih
}

/*******************************************************************************
** 
** Hierarchy
**
*******************************************************************************/

func Append(ih, new_child *Ihandle) *Ihandle {
	result := C.IupAppend(ih.C(), new_child.C())
	if result == nil {
		return nil
	}

	return ih
}

func Detach(ih *Ihandle) {
	C.IupDetach(ih.C())
}

func Insert(ih, ref_child, new_child *Ihandle) *Ihandle {
	result := C.IupInsert(ih.C(), ref_child.C(), new_child.C())
	if result == nil {
		return nil
	}

	return ih
}

func Reparent(child, new_parent, ref_child *Ihandle) int {
	return int(C.IupReparent(child.C(), new_parent.C(), ref_child.C()))
}

func GetParent(ih *Ihandle) *Ihandle {
	return (*Ihandle)(C.IupGetParent(ih.C()))
}

func GetChild(ih *Ihandle, pos int) *Ihandle {
	return (*Ihandle)(C.IupGetChild(ih.C(), C.int(pos)))
}

func GetChildPos(ih *Ihandle, child *Ihandle) int {
	return int(C.IupGetChildPos(ih.C(), child.C()))
}

func GetChildCount(ih *Ihandle) int {
	return int(C.IupGetChildCount(ih.C()))
}

func GetNextChild(ih, child *Ihandle) *Ihandle {
	return (*Ihandle)(C.IupGetNextChild(ih.C(), child.C()))
}

func GetBrother(ih *Ihandle) *Ihandle {
	return (*Ihandle)(C.IupGetBrother(ih.C()))
}

func GetDialog(ih *Ihandle) *Ihandle {
	return (*Ihandle)(C.IupGetDialog(ih.C()))
}

func GetDialogChild(ih *Ihandle, name string) *Ihandle {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return (*Ihandle)(C.IupGetDialogChild(ih.C(), cName))
}

/*******************************************************************************
** 
** Utilities
**
*******************************************************************************/

func Refresh(ih *Ihandle) {
	C.IupRefresh(ih.C())
}

func RefreshChildren(ih *Ihandle) {
	C.IupRefreshChildren(ih.C())
}

func Update(ih *Ihandle) {
	C.IupUpdate(ih.C())
}

func UpdateChildren(ih *Ihandle) {
	C.IupUpdateChildren(ih.C())
}

func Redraw(ih *Ihandle, children bool) {
	updateChildren := 0
	if children {
		updateChildren = 1
	}

	C.IupRedraw(ih.C(), C.int(updateChildren))
}

func ConvertXYToPos(ih *Ihandle, x, y int) int {
	return int(C.IupConvertXYToPos(ih.C(), C.int(x), C.int(y)))
}
