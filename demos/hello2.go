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
import "fmt"

func sayHello(ih *iup.Ihandle) int {
	fmt.Printf("Hello, %s!\n", ih.GetAttribute("TO_WHO"))
	
	return 0
}

func main() {
	iup.Open()
	defer iup.Close()

	lb := iup.Label("Press button below...")
	
	helloJohn := iup.Button("Hello John", "SAY_HELLO")
	helloJohn.SetButtonCallback(sayHello)
	helloJohn.StoreAttribute("TO_WHO", "John Doe")
	
	helloJim := iup.Button("Hello Jim", "SAY_HELLO")
	helloJim.SetButtonCallback(sayHello)
	helloJim.StoreAttribute("TO_WHO", "Jim Doe")
	
	helloBx := iup.Hbox(helloJohn, helloJim)
	bx := iup.Vbox(lb, helloBx)
	
	dg := iup.Dialog(bx)
	
	iup.Show(dg)
	iup.MainLoop()
}
