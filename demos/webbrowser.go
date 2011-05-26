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
	"github.com/jcowgar/go-iup/webbrowser"
)

var urlentry, browser *iup.Ihandle

func updateUrlEntry(ih *iup.Ihandle, url string) int {
	iup.StoreAttribute(urlentry, "VALUE", url)
	
	return iup.DEFAULT
}

func onGo(ih *iup.Ihandle) int {
	iup.StoreAttribute(browser, "VALUE", iup.GetAttribute(urlentry, "VALUE"))
	
	return iup.DEFAULT
}

func onUrlEntryKey(ih *iup.Ihandle, ch int) int {
	if ch == 13 {
		onGo(ih)
		return iup.IGNORE
	}
	
	return iup.DEFAULT
}

func main() {
	iup.Open()
	defer iup.Close()
	
	browserFunc := func(attr, val string) iup.ActionFunc {
		return (iup.ActionFunc)(func(ih *iup.Ihandle) int {
				iup.StoreAttribute(browser, attr, val)
				return iup.DEFAULT
		})
	}

	urlentry    = iup.Text("EXPAND=HORIZONTAL", (iup.KAnyFunc)(onUrlEntryKey))
	browser     = webbrowser.WebBrowser("EXPAND=YES", 
		(webbrowser.NavigateFunc)(updateUrlEntry), 
		(webbrowser.CompletedFunc)(updateUrlEntry))
	backbutton := iup.Button("<<",     browserFunc("BACKFORWARD", "-1"))
	forbutton  := iup.Button(">>",     browserFunc("BACKFORWARD", "1"))
	relbutton  := iup.Button("Reload", browserFunc("RELOAD", "1"))
	gobutton   := iup.Button("Go...",  (iup.ActionFunc)(onGo))
	stopbutton := iup.Button("Stop",   browserFunc("STOP", "1"))
	toolbar    := iup.SetAttrs(iup.Hbox(backbutton,forbutton,relbutton,urlentry,gobutton,stopbutton), "MARGIN","5x5","GAP","3")
	dialog     := iup.Dialog(iup.Vbox(toolbar,browser), "SIZE=500x300,TITLE=\"go-iup Web Browser\"")
	iup.Show(dialog)
	
	iup.StoreAttribute(browser, "VALUE", "http://github.com/jcowgar/go-iup")

	iup.MainLoop()
}
