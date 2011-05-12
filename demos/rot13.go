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

// *****************************************************************************
//
// NOTICE: rot13char and rot13 functions borrowed from Rosetta Code. No
// original author information was given.
//
// *****************************************************************************

// Hello World using a custom dialog.
package main

import (
	"strings"
	"io/ioutil"
	
	"github.com/jcowgar/go-iup"
)

var mainDlg, text *iup.Ihandle

func rot13char(c int) int {
	if c >= 'a' && c <= 'm' || c >= 'A' && c <= 'M' {
		return c + 13
	} else if c >= 'n' && c <= 'z' || c >= 'N' && c <= 'Z' {
		return c - 13
	}
	
	return c
}
 
func rot13(s string) string {
	return strings.Map(rot13char, s)
}

func onLoadFile(ih *iup.Ihandle) int {
	dlg := iup.FileDlg("ALLOWNEW=NO,DIALOGTYPE=Open,TITLE=Open")
	dlg.SetAttributeHandle("PARENTDIALOG", mainDlg)
	iup.Popup(dlg, iup.CENTER, iup.CENTER)
	if dlg.GetInt("STATUS") == -1 {
		return iup.IGNORE
	}
	
	filename := dlg.GetAttribute("VALUE")
	dlg.Destroy()
	
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		iup.Message("Error", "Error: " + err.String())
		return iup.IGNORE
	}
	
	text.StoreAttribute("VALUE", string(content))
	
	return iup.IGNORE
}

func onRotate(ih *iup.Ihandle) int {
	content := text.GetAttribute("VALUE")
	content = rot13(content)
	text.StoreAttribute("VALUE", string(content))
	
	return iup.DEFAULT
}

func onQuit(ih *iup.Ihandle) int {
	return iup.CLOSE
}

func onAboutIup(ih *iup.Ihandle) int {
	iup.Help("http://www.tecgraf.puc-rio.br/iup/")
	
	return iup.DEFAULT
}

func onAboutIupGo(ih *iup.Ihandle) int {
	iup.Help("http://github.com/jcowgar/go-iup")
	
	return iup.DEFAULT
}

func main() {
	iup.Open()
	defer iup.Close()
	
	menu := iup.Menu(
		iup.Submenu("File",
			iup.Menu(
				iup.Item("Load File", (iup.ActionFunc)(onLoadFile)),
				iup.Item("Rotate", (iup.ActionFunc)(onRotate)),
				iup.Item("Quit", (iup.ActionFunc)(onQuit)))),
		iup.Submenu("Help",
			iup.Menu(
				iup.Item("About Iup", (iup.ActionFunc)(onAboutIup)),
				iup.Item("About go-iup", (iup.ActionFunc)(onAboutIupGo)))))
	
	text = iup.Text("MULTILINE=YES,EXPAND=YES,WORDWRAP=YES,SIZE=250x100,SCROLLBAR=YES")
	mainBox := iup.Vbox(
		iup.Label("Text to be rotated:"),
		text,
		iup.Hbox(
			iup.Button("Load File", (iup.ActionFunc)(onLoadFile)),
			iup.Button("Rotate", (iup.ActionFunc)(onRotate)),
			iup.Button("Quit", (iup.ActionFunc)(onQuit))))
	mainBox.SetAttributes("GAP=5,MARGIN=5x5")
	
	mainDlg = iup.Dialog(mainBox, "TITLE=\"Rot 13\"")
	mainDlg.SetAttributeHandle("MENU", menu)
	mainDlg.Show()
	
	iup.MainLoop()
}
