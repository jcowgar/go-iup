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

// Hello World using a custom dialog.
package main

import "github.com/jcowgar/iup.go"

func main() {
	iup.Open()
	defer iup.Close()

	lb := iup.Label("Press button below...")
	ok := iup.Button("Say Hello", "SAY_HELLO")
	iup.SetButtonCallback(ok)
	
	bx := iup.Vbox(lb, ok)
	dg := iup.Dialog(bx)
	
	iup.Show(dg)
	iup.MainLoop()
}
