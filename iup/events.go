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

func MainLoop() int {
	return int(C.IupMainLoop())
}

func MainLoopLevel() int {
	return int(C.IupMainLoopLevel())
}

func LoopStep() int {
	return int(C.IupLoopStep())
}

func LoopStepWait() int {
	return int(C.IupLoopStepWait())
}

func ExitLoop() {
	C.IupExitLoop()
}

func Flush() {
	C.IupFlush()
}

func RecordInput(filename string, mode int) int {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	return int(C.IupRecordInput(cFilename, C.int(mode)))
}

func PlayInput(filename string) int {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	return int(C.IupPlayInput(cFilename))
}
