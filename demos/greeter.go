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

var someone *iup.Ihandle

func nameKeyEntry(ih *iup.Ihandle, ch int, newValue string) int {
	fmt.Printf("Char Pressed: %d, New Value: %s\n", ch, newValue)
	return 0
}

func sayHello(ih *iup.Ihandle) int {
	fmt.Printf("Hello, %s!\n", ih.GetAttribute("TO_WHO"))
	return 0
}

func sayHelloToSomeone(ih *iup.Ihandle) int {
	fmt.Printf("Hello, %s!\n", someone.GetAttribute("VALUE"))
	return 0
}

func main() {
	iup.Open()
	defer iup.Close()
	
	// Line one contains a name entry box and a hello button
	someone = iup.Text((iup.TextActionFunc)(nameKeyEntry))	
	helloSomeone := iup.Button("Say Hello", 
		(iup.ActionFunc)(sayHelloToSomeone))
	
	line1 := iup.Hbox(iup.Label("Name:"), someone, helloSomeone)
	line1.SetAttributes("ALIGNMENT=ACENTER,GAP=5")
	
	// Line two contains two pre-defined hello buttons
	helloJohn := iup.Button("Hello John", 
		(iup.ActionFunc)(sayHello),
		"TO_WHO=\"John Doe\"")
	
	helloJim := iup.Button("Hello Jim", 
		(iup.ActionFunc)(sayHello),
		"TO_WHO=\"Jim Doe\"")
	
	line2 := iup.Hbox(iup.Label("Predefined greeters:"), helloJohn, helloJim)
	line2.SetAttributes("ALIGNMENT=ACENTER,GAP=5")
	
	form := iup.Vbox(line1, line2)
	form.SetAttributes("GAP=5,MARGIN=3x3")
	
	dlg := iup.Dialog(form, "TITLE=Greeter")
	iup.Show(dlg)
	
	iup.MainLoop()
}
