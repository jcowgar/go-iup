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
#cgo LDFLAGS: -liupcontrols -liupcd -liup_pplot -liupgl -liuptuio -liupim
#cgo linux LDFLAGS: -liupgtk -lXm
#cgo windows LDFLAGS: -liup -liupweb -lgdi32 -lole32 -lcomdlg32 -lcomctl32

#include <stdlib.h>
#include <iup.h>
*/
import "C"

import (
	"os"
	"unsafe"
)

func Open() int {
	return int(C.IupOpen(nil, nil))
}

func Close() {
	C.IupClose()
}

func Version() string {
	return C.GoString(C.IupVersion())
}

func Load(filename string) (err os.Error) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	cResult := C.IupLoad(cFilename)
	if cResult != nil {
		err = os.NewError(C.GoString(cResult))
	}

	return
}

func LoadBuffer(buffer string) (err os.Error) {
	cBuffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(cBuffer))

	cResult := C.IupLoadBuffer(cBuffer)
	if cResult != nil {
		err = os.NewError(C.GoString(cResult))
	}

	return
}

func SetLanguage(lng string) {
	cLng := C.CString(lng)
	defer C.free(unsafe.Pointer(cLng))

	C.IupSetLanguage(cLng)
}

func GetLanguage() string {
	return C.GoString(C.IupGetLanguage())
}
