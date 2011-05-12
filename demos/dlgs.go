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

/* TODO: Make this program a single dialog that provides buttons for various
predefined dialog types */

import (
	"fmt"
	
	"github.com/jcowgar/go-iup"
)

func main() {
	iup.Open()
	defer iup.Close()
	
	//filename, code := iup.GetFile("./*.go")	
	//fmt.Printf("code=%d, filename='%s'\n", code, filename)
	
	//r,g,b := iup.GetColor(0,0,150,180,200)
	//fmt.Printf("r=%d, g=%d, b=%d\n", r, g, b)
	
	value, code := iup.GetText("Address", "123 Main St.\nSmall Town, TN 12345")
	fmt.Printf("code=%d value='%s'\n", code, value)
}
