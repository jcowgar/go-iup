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
	
	var age int
	var salary float64
	var wantRaise bool
	
	age = 10
	salary = 1293.55
	wantRaise = true
	
	if iup.GetParam("Age", "Age: %i\nSalary: %r\nWant Raise: %b[No,Yes]\n", &age, &salary, &wantRaise) {
		fmt.Printf("age is %d\nsalary is %f\nwant raise %v\n", age, salary, wantRaise)
	}
}
