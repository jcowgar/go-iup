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
#include <iup.h>
#include <iupweb.h>
*/
import "C"

func WebBrowser(opts ...interface{}) *Ihandle {
	if webBrowserLibOpened == false {
		C.IupWebBrowserOpen()
		webBrowserLibOpened = true
	}

	ih := &Ihandle{h: C.IupWebBrowser()}

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
			
		case CompletedFunc:
			ih.SetCompletedFunc(v)
			
		case NavigateFunc:
			ih.SetNavigateFunc(v)
		}
	}

	return ih
}
