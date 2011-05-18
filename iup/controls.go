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
		SetAttributes(ih, v)
		
	case MapFunc:
		SetMapFunc(ih, v)
		
	case UnmapFunc:
		SetUnmapFunc(ih, v)
		
	case DestroyFunc:
		SetDestroyFunc(ih, v)
		
	case GetFocusFunc:
		SetGetFocusFunc(ih, v)
		
	case KillFocusFunc:
		SetKillFocusFunc(ih, v)
		
	case EnterWindowFunc:
		SetEnterWindowFunc(ih, v)
		
	case LeaveWindowFunc:
		SetLeaveWindowFunc(ih, v)
		
	case KAnyFunc:
		SetKAnyFunc(ih, v)
		
	case HelpFunc:
		SetHelpFunc(ih, v)
		
	case ButtonFunc:
		SetButtonFunc(ih, v)
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
			SetActionFunc(ih, v)
			
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
			SetDropFilesFunc(ih, v)
			
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
			SetListActionFunc(ih, v)
			
		case CaretFunc:
			SetCaretFunc(ih, v)
			
		case DblclickFunc:
			SetDblclickFunc(ih, v)
			
		case EditFunc:
			SetEditFunc(ih, v)
			
		case MotionFunc:
			SetMotionFunc(ih, v)
			
		case MultiselectFunc:
			SetMultiselectFunc(ih, v)
			
		case ValueChangedFunc:
			SetValueChangedFunc(ih, v)
			
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
			SetTabChangeFunc(ih, v)
			
		case TabChangePosFunc:
			SetTabChangePosFunc(ih, v)
			
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
			SetTextActionFunc(ih, v)

		default:
			decorate(ih, o)
		}
	}

	return ih
}

func TextConvertLinColToPos(ih *Ihandle, lin, col int) int {
	pos := new(C.int)
	C.IupTextConvertLinColToPos(ih.H, C.int(lin), C.int(col), pos)
	return int(*pos)
}

func TextConvertPosToLinCol(ih *Ihandle, pos int) (int, int) {
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
			SetToggleActionFunc(ih, v)
			
		case ValueChangedFunc:
			SetValueChangedFunc(ih, v)
			
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
			SetValueChangedFunc(ih, v)
			
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

func MatSetAttribute(ih *Ihandle, name string, lin int, col int, value unsafe.Pointer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.IupMatSetAttribute(ih.H, cName, C.int(lin), C.int(col), (*C.char)(value))
}

func MatStoreAttribute(ih *Ihandle, name string, lin int, col int, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.IupMatStoreAttribute(ih.H, cName, C.int(lin), C.int(col), cValue)
}

func MatGetAttribute(ih *Ihandle, name string, lin int, col int) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return C.GoString(C.IupMatGetAttribute(ih.H, cName, C.int(lin), C.int(col)))
}

func MatGetInt(ih *Ihandle, name string, lin int, col int) int64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return int64(C.IupMatGetInt(ih.H, cName, C.int(lin), C.int(col)))
}

func MatGetFloat(ih *Ihandle, name string, lin int, col int) float64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return float64(C.IupMatGetFloat(ih.H, cName, C.int(lin), C.int(col)))
}

func MatSetfAttribute(ih *Ihandle, name string, lin int, col int, format string, args ...interface{}) {
	MatStoreAttribute(ih, name, lin, col, fmt.Sprintf(format, args...))
}

func OleControl(opts ...interface{}) *Ihandle {
	panic("OleControl is not yet implemented")

	return nil
}
