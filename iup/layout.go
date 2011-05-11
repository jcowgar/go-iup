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

func toC(ihs []*Ihandle) []*C.Ihandle {
	max := len(ihs)
	result := make([]*C.Ihandle, max + 1)
	
	for k, v := range ihs {
		result[k] = v.h
	}	
	result[max] = nil
	
	return result
}

func hboxv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{h: C.IupHboxv(&ihs[0])}
}

func Hbox(args ...*Ihandle) *Ihandle {
	return hboxv(toC(args))
}

func Hboxv(args []*Ihandle) *Ihandle {
	return hboxv(toC(args))
}

func vboxv(ihs []*C.Ihandle) *Ihandle {
	return &Ihandle{h: C.IupVboxv(&ihs[0])}
}

func Vbox(args ...*Ihandle) *Ihandle {
	return vboxv(toC(args))
}

func Vboxv(args []*Ihandle) *Ihandle {
	return vboxv(toC(args))
}
