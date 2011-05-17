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
*/
import "C"
import "unsafe"
import "fmt"

var webBrowserLibOpened = false

/*******************************************************************************
 * Supporting Methods
 */

func decorate(ih *Ihandle, opt interface{}) {
	switch v := opt.(type) {
	case string:
		ih.SetAttributes(v)
		
	case MapFunc:
		ih.SetMapFunc(v)
		
	case UnmapFunc:
		ih.SetUnmapFunc(v)
		
	case DestroyFunc:
		ih.SetDestroyFunc(v)
		
	case GetFocusFunc:
		ih.SetGetFocusFunc(v)
		
	case KillFocusFunc:
		ih.SetKillFocusFunc(v)
		
	case EnterWindowFunc:
		ih.SetEnterWindowFunc(v)
		
	case LeaveWindowFunc:
		ih.SetLeaveWindowFunc(v)
		
	case KAnyFunc:
		ih.SetKAnyFunc(v)
		
	case HelpFunc:
		ih.SetHelpFunc(v)
		
	case ButtonFunc:
		ih.SetButtonFunc(v)
	}
}

/*******************************************************************************
**
** Basic Controls
**
*******************************************************************************/

func Button(title string, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	ih := &Ihandle{H: C.IupButton(cTitle, nil)}
	
	for _, o := range opts {
		switch v := o.(type) {
		case ActionFunc:
			ih.SetActionFunc(v)
			
		default:
			decorate(ih, o)
		}
	}
	
	return ih
}

func Canvas(opts ...interface{}) *Ihandle {
	ih := &Ihandle{H: C.IupCanvas(nil)}

	for _, o := range opts {
		switch v := o.(type) {
		default:
			decorate(ih, o)
		}
	}

	return ih
}

func Frame(child *Ihandle, opts ...interface{}) *Ihandle {
	ih := &Ihandle{H: C.IupFrame(child.H)}

	for _, o := range opts {
		switch v := o.(type) {
		default:
			decorate(ih, o)
		}
	}

	return ih
}

func Label(title string, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	ih := &Ihandle{H: C.IupLabel(cTitle)}

	for _, o := range opts {
		switch v := o.(type) {
		case DropFilesFunc:
			ih.SetDropFilesFunc(v)
			
		default:
			decorate(ih, o)
		}
	}

	return ih
}

func List(opts ...interface{}) *Ihandle {
	ih := &Ihandle{H: C.IupList(nil)}

	for _, o := range opts {
		switch v := o.(type) {
		case ListActionFunc:
			ih.SetListActionFunc(v)
			
		case CaretFunc:
			ih.SetCaretFunc(v)
			
		case DblclickFunc:
			ih.SetDblclickFunc(v)
			
		case EditFunc:
			ih.SetEditFunc(v)
			
		case MotionFunc:
			ih.SetMotionFunc(v)
			
		case MultiselectFunc:
			ih.SetMultiselectFunc(v)
			
		case ValueChangedFunc:
			ih.SetValueChangedFunc(v)
			
		default:
			decorate(ih, o)
		}
	}

	return ih
}

func ProgressBar(opts ...interface{}) *Ihandle {
	ih := &Ihandle{H: C.IupProgressBar()}

	for _, o := range opts {
		switch v := o.(type) {
		default:
			decorate(ih, o)
		}
	}
	
	return ih
}

func Spin(opts ...interface{}) *Ihandle {
	ih := &Ihandle{H: C.IupSpin()}

	for _, o := range opts {
		switch v := o.(type) {
		default:
			decorate(ih, o)
		}
	}
	
	return ih
}

func tabsv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{H: C.IupTabsv(&ihs[0])}
}

func Tabs(args ...*Ihandle) *Ihandle {
	return tabsv(iHandleArrayToC(args))
}

func Tabsv(args []*Ihandle, opts ...interface{}) *Ihandle {
	ih := tabsv(iHandleArrayToC(args))

	for _, o := range opts {
		switch v := o.(type) {
		case TabChangeFunc:
			ih.SetTabChangeFunc(v)
			
		case TabChangePosFunc:
			ih.SetTabChangePosFunc(v)
			
		default:
			decorate(ih, o)
		}
	}

	return ih
}

func Text(opts ...interface{}) *Ihandle {
	ih := &Ihandle{H: C.IupText(nil)}

	for _, o := range opts {
		switch v := o.(type) {
		case TextActionFunc:
			ih.SetTextActionFunc(v)

		default:
			decorate(ih, o)
		}
	}

	return ih
}

func (ih *Ihandle) TextConvertLinColToPos(lin, col int) int {
	pos := new(C.int)
	C.IupTextConvertLinColToPos(ih.H, C.int(lin), C.int(col), pos)
	return int(*pos)
}

func (ih *Ihandle) TextConvertPosToLinCol(pos int) (int, int) {
	lin := new(C.int)
	col := new(C.int)

	C.IupTextConvertPosToLinCol(ih.H, C.int(pos), lin, col)

	return int(*lin), int(*col)
}

func Toggle(title string, opts ...interface{}) *Ihandle {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	ih := &Ihandle{H: C.IupToggle(cTitle, nil)}

	for _, o := range opts {
		switch v := o.(type) {
		case ToggleActionFunc:
			ih.SetToggleActionFunc(v)
			
		case ValueChangedFunc:
			ih.SetValueChangedFunc(v)
			
		default:
			decorate(ih, o)
		}
	}

	return ih
}

func Tree(opts ...interface{}) *Ihandle {
	ih := &Ihandle{H: C.IupTree()}

	for _, o := range opts {
		switch v := o.(type) {
		default:
			decorate(ih, o)
		}
	}

	return ih
}

func Val(orientation string, opts ...interface{}) *Ihandle {
	cOrientation := C.CString(orientation)
	defer C.free(unsafe.Pointer(cOrientation))

	ih := &Ihandle{H: C.IupVal(cOrientation)}

	for _, o := range opts {
		switch v := o.(type) {
		case ValueChangedFunc:
			ih.SetValueChangedFunc(v)
			
		default:
			decorate(ih, o)
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
	OpenControlLib()

	ih := &Ihandle{H: C.IupCells()}

	for _, o := range opts {
		switch v := o.(type) {
		default:
			decorate(ih, o)
		}
	}

	return ih
}

func Colorbar(opts ...interface{}) *Ihandle {
	OpenControlLib()

	ih := &Ihandle{H: C.IupColorbar()}

	for _, o := range opts {
		switch v := o.(type) {
		default:
			decorate(ih, o)
		}
	}

	return ih
}

func ColorBrowser(opts ...interface{}) *Ihandle {
	OpenControlLib()

	ih := &Ihandle{H: C.IupColorBrowser()}

	for _, o := range opts {
		switch v := o.(type) {
		default:
			decorate(ih, o)
		}
	}

	return ih
}

func Dial(orientation string, opts ...interface{}) *Ihandle {
	OpenControlLib()

	cOrientation := C.CString(orientation)
	defer C.free(unsafe.Pointer(cOrientation))

	ih := &Ihandle{H: C.IupDial(cOrientation)}

	for _, o := range opts {
		switch v := o.(type) {
		default:
			decorate(ih, o)
		}
	}

	return ih
}

func Matrix(opts ...interface{}) *Ihandle {
	OpenControlLib()

	ih := &Ihandle{H: C.IupMatrix(nil)}

	for _, o := range opts {
		switch v := o.(type) {
		default:
			decorate(ih, o)
		}
	}

	return ih
}

func (ih *Ihandle) MatSetAttribute(name string, lin int, col int, value unsafe.Pointer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.IupMatSetAttribute(ih.H, cName, C.int(lin), C.int(col), (*C.char)(value))
}

func (ih *Ihandle) MatStoreAttribute(name string, lin int, col int, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.IupMatStoreAttribute(ih.H, cName, C.int(lin), C.int(col), cValue)
}

func (ih *Ihandle) MatGetAttribute(name string, lin int, col int) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return C.GoString(C.IupMatGetAttribute(ih.H, cName, C.int(lin), C.int(col)))
}

func (ih *Ihandle) MatGetInt(name string, lin int, col int) int64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return int64(C.IupMatGetInt(ih.H, cName, C.int(lin), C.int(col)))
}

func (ih *Ihandle) MatGetFloat(name string, lin int, col int) float64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return float64(C.IupMatGetFloat(ih.H, cName, C.int(lin), C.int(col)))
}

func (ih *Ihandle) MatSetfAttribute(name string, lin int, col int, format string, args ...interface{}) {
	ih.MatStoreAttribute(name, lin, col, fmt.Sprintf(format, args...))
}

func OleControl(opts ...interface{}) *Ihandle {
	panic("OleControl is not yet implemented")

	return nil
}
