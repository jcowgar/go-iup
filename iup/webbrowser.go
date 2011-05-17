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
#include <iupweb.h>

extern int goIupCompletedCB(void *, char *);
void goIupSetCompletedFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_COMPLETED_CB", f);
	IupSetCallback(ih, "COMPLETED_CB", (Icallback) goIupCompletedCB);
}

extern int goIupErrorCB(void *, char *);
void goIupSetErrorFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_ERROR_CB", f);
	IupSetCallback(ih, "ERROR_CB", (Icallback) goIupErrorCB);
}

extern int goIupNavigateCB(void *, char *);
void goIupSetNavigateFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_NAVIGATE_CB", f);
	IupSetCallback(ih, "NAVIGATE_CB", (Icallback) goIupNavigateCB);
}

extern int goIupNewWindowCB(void *, char *);
void goIupSetNewWindowFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, "_GO_NEWWINDOW_CB", f);
	IupSetCallback(ih, "NEWWINDOW_CB", (Icallback) goIupNewWindowCB);
}
*/
import "C"
import "unsafe"

type CompletedFunc func(*Ihandle, string) int

//export goIupCompletedCB
func goIupCompletedCB(ih unsafe.Pointer, url unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_COMPLETED_CB")
	defer C.free(unsafe.Pointer(cName))
	
	f := *(*CompletedFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	
	goUrl := C.GoString((*C.char)(url))
	
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, goUrl)
}

func (ih *Ihandle) SetCompletedFunc(f CompletedFunc) {
	C.goIupSetCompletedFunc(ih.h, unsafe.Pointer(&f))
}

type ErrorFunc func(*Ihandle, string) int

//export goIupErrorCB
func goIupErrorCB(ih unsafe.Pointer, url unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_ERROR_CB")
	defer C.free(unsafe.Pointer(cName))
	
	f := *(*ErrorFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	
	goUrl := C.GoString((*C.char)(url))
	
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, goUrl)
}

func (ih *Ihandle) SetErrorFunc(f ErrorFunc) {
	C.goIupSetErrorFunc(ih.h, unsafe.Pointer(&f))
}

type NavigateFunc func(*Ihandle, string) int

//export goIupNavigateCB
func goIupNavigateCB(ih unsafe.Pointer, url unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_NAVIGATE_CB")
	defer C.free(unsafe.Pointer(cName))
	
	f := *(*NavigateFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	
	goUrl := C.GoString((*C.char)(url))
	
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, goUrl)
}

func (ih *Ihandle) SetNavigateFunc(f NavigateFunc) {
	C.goIupSetNavigateFunc(ih.h, unsafe.Pointer(&f))
}

type NewWindowFunc func(*Ihandle, string) int

//export goIupNewWindowCB
func goIupNewWindowCB(ih unsafe.Pointer, url unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	cName := C.CString("_GO_NEWWINDOW_CB")
	defer C.free(unsafe.Pointer(cName))
	
	f := *(*NewWindowFunc)(unsafe.Pointer(C.IupGetAttribute(h, cName)))
	
	goUrl := C.GoString((*C.char)(url))
	
	return f(&Ihandle{h: (*C.Ihandle)(ih)}, goUrl)
}

func (ih *Ihandle) SetNewWindowFunc(f NewWindowFunc) {
	C.goIupSetNewWindowFunc(ih.h, unsafe.Pointer(&f))
}

func WebBrowser(opts ...interface{}) *Ihandle {
	if webBrowserLibOpened == false {
		C.IupWebBrowserOpen()
		webBrowserLibOpened = true
	}

	ih := &Ihandle{h: C.IupWebBrowser()}

	for _, o := range opts {
		switch v := o.(type) {
		case CompletedFunc:
			ih.SetCompletedFunc(v)
			
		case ErrorFunc:
			ih.SetErrorFunc(v)
			
		case NavigateFunc:
			ih.SetNavigateFunc(v)
			
		case NewWindowFunc:
			ih.SetNewWindowFunc(v)
		
		default:
			decorate(ih, o)
		}
	}

	return ih
}
