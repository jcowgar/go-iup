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
#include <iupcontrols.h>
//#include <iupweb.h>
#include <iup_pplot.h>
#include <iupgl.h>
*/
import "C"

import "unsafe"

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
	return tabsv(toC(args))
}

func Tabsv(args []*Ihandle, opts ...interface{}) *Ihandle {
	ih := tabsv(toC(args))
	
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

func Cells(opts ...interface{}) *Ihandle {
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
	ih := &Ihandle{h: C.IupMatrix(nil)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}
	
	return ih
}

func GLCanvas(opts ...interface{}) *Ihandle {
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
	// TODO: FAIL, Not Implemented
	
	return nil
}

func WebBrowser(opts ...interface{}) *Ihandle {
	// TODO: FAIL, Not Implemented
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
