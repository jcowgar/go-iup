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
#include <iup.h>
*/
import "C"

const (
	// iup.go version string.
	//
	// iup.go version string is based off the built-against version
	// code of Iup with the addition of a `iup.go' version code
	// as the forth digit. i.e. 3.5.0.1 means that this version of
	// iup.go was built with Iup 3.5.0 in mind and is the `.1' release
	// of iup.go against Iup 3.5.0.
	IupGoVersion = "3.5.0.1"
)

// Primary widget handle type.
type Ihandle C.Ihandle
