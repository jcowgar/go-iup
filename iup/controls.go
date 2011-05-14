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
#include <iupcontrols.h>
//#include <iupweb.h>
#include <iup_pplot.h>
#include <iupgl.h>
*/
import "C"
import "unsafe"
import "fmt"

var controlsLibOpened = false

/*******************************************************************************
**
** Basic Controls
**
*******************************************************************************/

func Button(title string, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	
	ih := &Ihandle{h: C.IupButton(cTitle, nil)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
			
		case ActionFunc:
			ih.SetActionFunc(v)
				
		default:
			// TODO: Do something here, runtime error?
		}
	}
	
	return ih
}

func Canvas(opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupCanvas(nil)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func Frame(child *Ihandle, opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupFrame(child.h)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}
	
func Label(title string, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	
	ih := &Ihandle{h: C.IupLabel(cTitle)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func List(opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupList(nil)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func ProgressBar(opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupProgressBar()}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func Spin(opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupSpin()}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func tabsv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{h: C.IupTabsv(&ihs[0])}
}

func Tabs(args ...*Ihandle) *Ihandle {
	return tabsv(iHandleArrayToC(args))
}

func Tabsv(args []*Ihandle, opts ...interface{}) *Ihandle {
	ih := tabsv(iHandleArrayToC(args))
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih	
}

func Text(opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupText(nil)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
			
		case TextActionFunc:
			ih.SetTextActionFunc(v)
			
		default:
			// TODO: Do something here, runtime error?
		}
	}
	
	return ih
}

func (ih *Ihandle) TextConvertLinColToPos(lin, col int) int {
	pos := new(C.int)
	C.IupTextConvertLinColToPos(ih.h, C.int(lin), C.int(col), pos)	
	return int(*pos)
}

func (ih *Ihandle) TextConvertPosToLinCol(pos int) (int, int) {
	lin := new(C.int)
	col := new(C.int)
	
	C.IupTextConvertPosToLinCol(ih.h, C.int(pos), lin, col)
	
	return int(*lin), int(*col)
}

func Toggle(title string, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	
	ih := &Ihandle{h: C.IupToggle(cTitle, nil)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih	
}

func Tree(opts ...interface{}) *Ihandle {
	ih := &Ihandle{h: C.IupTree()}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func Val(orientation string, opts ...interface{}) *Ihandle {
	cOrientation := C.CString(orientation)
	defer C.free(unsafe.Pointer(cOrientation))
	
	ih := &Ihandle{h: C.IupVal(cOrientation)}
	
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
** Additional Controls
**
*******************************************************************************/

func ensureControlLibOpened() {
	if controlsLibOpened == false {
		C.IupControlsOpen()
		controlsLibOpened = true
	}
}

func Cells(opts ...interface{}) *Ihandle {
	ensureControlLibOpened()
	
	ih := &Ihandle{h: C.IupCells()}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func Colorbar(opts ...interface{}) *Ihandle {
	ensureControlLibOpened()
	
	ih := &Ihandle{h: C.IupColorbar()}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func ColorBrowser(opts ...interface{}) *Ihandle {
	ensureControlLibOpened()
	
	ih := &Ihandle{h: C.IupColorBrowser()}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func Dial(orientation string, opts ...interface{}) *Ihandle {
	ensureControlLibOpened()
	
	cOrientation := C.CString(orientation)
	defer C.free(unsafe.Pointer(cOrientation))
	
	ih := &Ihandle{h: C.IupDial(cOrientation)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih	
}

func Matrix(opts ...interface{}) *Ihandle {
	ensureControlLibOpened()
	
	ih := &Ihandle{h: C.IupMatrix(nil)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func (ih *Ihandle) MatSetAttribute(name string, lin int, col int, value unsafe.Pointer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	C.IupMatSetAttribute(ih.h, cName, C.int(lin), C.int(col), (*C.char)(value))
}

func (ih *Ihandle) MatStoreAttribute(name string, lin int, col int, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	
	C.IupMatStoreAttribute(ih.h, cName, C.int(lin), C.int(col), cValue)
}

func (ih *Ihandle) MatGetAttribute(name string, lin int, col int) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	return C.GoString(C.IupMatGetAttribute(ih.h, cName, C.int(lin), C.int(col)))
}

func (ih *Ihandle) MatGetInt(name string, lin int, col int) int64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	return int64(C.IupMatGetInt(ih.h, cName, C.int(lin), C.int(col)))	
}

func (ih *Ihandle) MatGetFloat(name string, lin int, col int) float64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	return float64(C.IupMatGetFloat(ih.h, cName, C.int(lin), C.int(col)))
}

func (ih *Ihandle) MatSetfAttribute(name string, lin int, col int, format string, args ...interface{}) {
	ih.MatStoreAttribute(name, lin, col, fmt.Sprintf(format, args...))
}

func GLCanvas(opts ...interface{}) *Ihandle {
	ensureControlLibOpened()
	
	ih := &Ihandle{h: C.IupGLCanvas(nil)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func PPlot(opts ...interface{}) *Ihandle {
	ensureControlLibOpened()
	
	ih := &Ihandle{h: C.IupPPlot()}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func OleControl(opts ...interface{}) *Ihandle {
	panic("OleControl is not yet implemented")
	
	return nil
}

func WebBrowser(opts ...interface{}) *Ihandle {
	panic("WebBrowser is not yet implemented")
	
	return nil
	/*
	ih := &Ihandle{h: C.IupWebBrowser()}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
	*/
}
