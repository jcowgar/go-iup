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

const (
	// go-iup version string.
	//
	// go-iup version string is based off the built-against version
	// code of Iup with the addition of a `go-iup' version code
	// as the forth digit. i.e. 3.5.0.1 means that this version of
	// go-iup was built with Iup 3.5.0 in mind and is the `.1' release
	// of go-iup against Iup 3.5.0.
	IupGoVersion = "3.5.0.1"
)

// Primary widget handle type.
type Ihandle struct {
	h *C.Ihandle
}

func iHandleArrayToC(ihs []*Ihandle) []*C.Ihandle {
	max := len(ihs)
	result := make([]*C.Ihandle, max+1)

	for k, v := range ihs {
		result[k] = v.h
	}
	result[max] = nil

	return result
}

func stringArrayToC(strs []string) []*C.char {
	max := len(strs)
	result := make([]*C.char, max+1)

	for k, v := range strs {
		result[k] = C.CString(v)
	}
	result[max] = nil

	return result
}

func freeCStringArray(strs []*C.char) {
	for _, v := range strs {
		if v != nil {
			C.free(unsafe.Pointer(v))
		}
	}
}

func float64ArrayToC(nums []float64) []C.float {
	result := make([]C.float, len(nums))

	for k, v := range nums {
		result[k] = C.float(v)
	}

	return result
}

func byteArrayToCUCharArray(content []byte) []C.uchar {
	cContent := make([]C.uchar, len(content))
	for i, v := range content {
		cContent[i] = (C.uchar)(v)
	}
	cContent[len(content)] = 0

	return cContent
}

func intArrayToC(nums []int) []C.int {
	result := make([]C.int, len(nums))

	for k, v := range nums {
		result[k] = C.int(v)
	}

	return result
}
