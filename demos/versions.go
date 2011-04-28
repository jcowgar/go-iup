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

// Display the iup.go and Iup library version numbers.
package main

import "github.com/jcowgar/iup.go"
import "fmt"

func main() {
	iup.Open()
	defer iup.Close()

	fmt.Printf("iup.go Version=%s\n", iup.IupGoVersion)
	fmt.Printf("   Iup Version=%s\n", iup.Version())
}
