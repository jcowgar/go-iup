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

const DEFAULT      = C.IUP_DEFAULT
const CLOSE        = C.IUP_CLOSE
const IGNORE       = C.IUP_IGNORE
const CONTINUE     = C.IUP_CONTINUE

const LEFT         = C.IUP_LEFT
const CENTER       = C.IUP_CENTER
const RIGHT        = C.IUP_RIGHT
const MOUSEPOS     = C.IUP_MOUSEPOS
const CENTERPARENT = C.IUP_CENTERPARENT
const CURRENT      = C.IUP_CURRENT
