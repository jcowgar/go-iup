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

package main

import (
	"fmt"

	"github.com/jcowgar/go-iup"
)

func main() {
	iup.Open()
	defer iup.Close()

	abool := true
	aint := 10
	afloat := 1293.33
	aint2 := 192
	afloat2 := 0.5
	aangle := 90.0
	astr := "string text"
	aopt := 1  // opt/list defined in format string
	alist := 1 // this is the selected item
	afile := "hello.txt"
	acolor := "255 122 255"
	afont := "Courier, 24"
	astr2 := "hello\nworld"

	result := iup.GetParam("Sample",
		"Boolean: %b[No,Yes]\n"+
			"Integer 1: %i\n"+
			"Real 1: %r\n"+
			"Sep1 %t\n"+
			"Integer 2: %i[0,255]\n"+
			"Real 2: %r[-1.5,1.5,0.05]\n"+
			"Sep2 %t\n"+
			"Angle %a[0,360]\n"+
			"String: %s\n"+
			"Options: %o|item0|item1|item2|item3\n"+
			"List: %l|item0|item1|item2|item3\n"+
			"File: %f[OPEN|*.txt;*.asc|CURRENT|NO|NO]\n"+
			"Color: %c{Color Tip}\n"+
			"Font: %n\n"+
			"Sep3 %t\n"+
			"Multiline: %m\n",
		&abool, &aint, &afloat, &aint2, &afloat2, &aangle, &astr, &aopt, &alist, &afile,
		&acolor, &afont, &astr2)

	if result {
		fmt.Printf(
			"Boolean: %v\n"+
				"Integer 1: %v\n"+
				"Real 1: %v\n"+
				"Integer 2: %v\n"+
				"Real 2: %v\n"+
				"Angle %v\n"+
				"String: %v\n"+
				"Options: %v\n"+
				"List: %v\n"+
				"File: %v\n"+
				"Color: %v\n"+
				"Font: %v\n"+
				"Multiline: %v\n",
			abool, aint, afloat, aint2, afloat2, aangle, astr, aopt, alist, afile,
			acolor, afont, astr2)
	}
}
