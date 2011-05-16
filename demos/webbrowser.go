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
	"github.com/jcowgar/go-iup"
)

var urlentry, browser *iup.Ihandle

func onGo(ih *iup.Ihandle) int {
	browser.StoreAttribute("VALUE", urlentry.GetAttribute("VALUE"))
	
	return iup.IGNORE
}

func main() {
	iup.Open()
	defer iup.Close()
	
	urlentry  = iup.Text("EXPAND=HORIZONTAL")
	browser   = iup.WebBrowser("EXPAND=YES")	
	gobutton := iup.Button("Go...", (iup.ActionFunc)(onGo))
	toolbar  := iup.Hbox(urlentry, gobutton).SetAttrs("MARGIN","5x5", "GAP", "3")
	dialog   := iup.Dialog(iup.Vbox(toolbar,browser), 
		"SIZE=500x300,TITLE=\"go-iup Web Browser\"")
	dialog.Show()
	
	browser.StoreAttribute("VALUE", "http://github.com/jcowgar/go-iup")

	iup.MainLoop()
}
