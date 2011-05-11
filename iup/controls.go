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

// Create a new button widget. Differs from Iup implementation in that ACTION
// is not the second parameter. Instead, any number of optional parameters can
// be supplied. The only types understood are:
//
//   1. string: sent to the newly created widget as an attribute to be set
//   2. ButtonActionFunc: sets the ACTION callback to the value
//
func Button(title string, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	
	ih := &Ihandle{h: C.IupButton(cTitle, nil)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
			
		case ButtonActionFunc:
			ih.SetButtonActionFunc(v)
				
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

// Differs from IupLabel in that any number of optional parameters may be passed. Strings
// will be interpreted as attributes to set on the newly created widget.
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

// Differs from IupText in that any number of optional parameters may be passed. Strings
// will be interpreted as attributes that will be set on the newly created widget. Any
// Callback functions supplied will be assigned to the correct callback attribute.
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
