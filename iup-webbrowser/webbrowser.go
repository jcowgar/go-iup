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

package webbrowser

/*
#cgo LDFLAGS: -liupcontrols -liupweb
#cgo linux LDFLAGS: -liupgtk
#cgo windows LDFLAGS: -liup -lgdi32 -lole32 -lcomdlg32 -lcomctl32

#include <stdlib.h>
#include <iup.h>
#include <iupweb.h>

#define GO_PREFIX        "_GO_"
#define IUP_COMPLETED_CB "COMPLETED_CB"
#define IUP_ERROR_CB     "ERROR_CB"
#define IUP_NAVIGATE_CB  "NAVIGATE_CB"
#define IUP_NEWWINDOW_CB "NEWWINDOW_CB"

const char* GO_COMPLETED_CB = GO_PREFIX IUP_COMPLETED_CB;
const char* GO_ERROR_CB     = GO_PREFIX IUP_ERROR_CB;
const char* GO_NAVIGATE_CB  = GO_PREFIX IUP_NAVIGATE_CB;
const char* GO_NEWWINDOW_CB = GO_PREFIX IUP_NEWWINDOW_CB;

extern int goIupCompletedCB(void *, char *);
void goIupSetCompletedFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_COMPLETED_CB, f);
	IupSetCallback(ih, IUP_COMPLETED_CB, (Icallback) goIupCompletedCB);
}

extern int goIupErrorCB(void *, char *);
void goIupSetErrorFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_ERROR_CB, f);
	IupSetCallback(ih, IUP_ERROR_CB, (Icallback) goIupErrorCB);
}

extern int goIupNavigateCB(void *, char *);
void goIupSetNavigateFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_NAVIGATE_CB, f);
	IupSetCallback(ih, IUP_NAVIGATE_CB, (Icallback) goIupNavigateCB);
}

extern int goIupNewWindowCB(void *, char *);
void goIupSetNewWindowFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_NEWWINDOW_CB, f);
	IupSetCallback(ih, IUP_NEWWINDOW_CB, (Icallback) goIupNewWindowCB);
}
*/
import "C"
import "unsafe"
import . "github.com/jcowgar/go-iup"

var webBrowserLibOpened = false

type CompletedFunc func(*Ihandle, string) int

//export goIupCompletedCB
func goIupCompletedCB(ih unsafe.Pointer, url unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*CompletedFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_COMPLETED_CB)))
	goUrl := C.GoString((*C.char)(url))
	return f((*Ihandle)(ih), goUrl)
}

func SetCompletedFunc(ih *Ihandle, f CompletedFunc) {
	C.goIupSetCompletedFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type ErrorFunc func(*Ihandle, string) int

//export goIupErrorCB
func goIupErrorCB(ih unsafe.Pointer, url unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*ErrorFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_ERROR_CB)))
	goUrl := C.GoString((*C.char)(url))
	return f((*Ihandle)(ih), goUrl)
}

func SetErrorFunc(ih *Ihandle, f ErrorFunc) {
	C.goIupSetErrorFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type NavigateFunc func(*Ihandle, string) int

//export goIupNavigateCB
func goIupNavigateCB(ih unsafe.Pointer, url unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*NavigateFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_NAVIGATE_CB)))
	goUrl := C.GoString((*C.char)(url))
	return f((*Ihandle)(ih), goUrl)
}

func SetNavigateFunc(ih *Ihandle, f NavigateFunc) {
	C.goIupSetNavigateFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type NewWindowFunc func(*Ihandle, string) int

//export goIupNewWindowCB
func goIupNewWindowCB(ih unsafe.Pointer, url unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*NewWindowFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_NEWWINDOW_CB)))
	goUrl := C.GoString((*C.char)(url))
	return f((*Ihandle)(ih), goUrl)
}

func SetNewWindowFunc(ih *Ihandle, f NewWindowFunc) {
	C.goIupSetNewWindowFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

func WebBrowser(opts ...interface{}) *Ihandle {
	if webBrowserLibOpened == false {
		C.IupWebBrowserOpen()
		webBrowserLibOpened = true
	}

	ih := (*Ihandle)(C.IupWebBrowser())

	for _, o := range opts {
		switch v := o.(type) {
		case CompletedFunc:
			SetCompletedFunc(ih, v)
			
		case ErrorFunc:
			SetErrorFunc(ih, v)
			
		case NavigateFunc:
			SetNavigateFunc(ih, v)
			
		case NewWindowFunc:
			SetNewWindowFunc(ih, v)
		
		default:
			Decorate(ih, o)
		}
	}

	return ih
}
