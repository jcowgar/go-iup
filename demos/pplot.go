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

func main() {
	iup.Open()
	defer iup.Close()
	
	pplot := iup.PPlot("EXPAND=YES")
	pplot.StoreAttribute("TITLE", "Bar Mode")
	pplot.StoreAttribute("TITLEFONTSIZE", "16")
	pplot.StoreAttribute("MARGINTOP", "40")
	pplot.StoreAttribute("MARGINLEFT", "30")
	pplot.StoreAttribute("MARGINBOTTOM","65")
	
	labels := []string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun", 
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	values := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	
	pplot.PlotBegin(1)
	for i := 0; i < len(labels); i++ {
		pplot.PlotAddStr(labels[i], values[i])
	}
	pplot.PlotEnd()
	pplot.StoreAttribute("DS_MODE", "BAR")
	pplot.StoreAttribute("DS_COLOR", "100 100 200")
	
	dlg := iup.Dialog(pplot, "SIZE=200x200,TITLE=\"PPlot Example\"")
	dlg.Show()
	
	pplot.StoreAttribute("REDRAW", "")
	
	iup.MainLoop()
}
