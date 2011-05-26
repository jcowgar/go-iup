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

package pplot

/*
#cgo LDFLAGS: -liupcontrols -liuptuio
#cgo linux LDFLAGS: -liupgtk
#cgo windows LDFLAGS: -liup -lgdi32 -lole32 -lcomdlg32 -lcomctl32

#include <stdlib.h>
#include <iup.h>
#include <iupcontrols.h>
#include <iuptuio.h>
*/
import "C"
import . "github.com/jcowgar/go-iup"

var tuioLibOpened = false

func TuioClient(port int) *Ihandle {
	OpenControlLib()
	
	if tuioLibOpened == false {
		C.IupTuioOpen()
		tuioLibOpened = true
	}
	
	return (*Ihandle)(C.IupTuioClient(C.int(port)))
}
