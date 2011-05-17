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
	"github.com/jcowgar/go-iup/pplot"
)

func main() {
	iup.Open()
	defer iup.Close()
	
	p := pplot.PPlot("EXPAND=YES")
	p.StoreAttribute("TITLE", "Bar Mode")
	p.StoreAttribute("TITLEFONTSIZE", "16")
	p.StoreAttribute("MARGINTOP", "40")
	p.StoreAttribute("MARGINLEFT", "30")
	p.StoreAttribute("MARGINBOTTOM","65")
	
	labels := []string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun", 
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	values := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	
	pplot.PlotBegin(p, 1)
	for i := 0; i < len(labels); i++ {
		pplot.PlotAddStr(p, labels[i], values[i])
	}
	pplot.PlotEnd(p)
	p.StoreAttribute("DS_MODE", "BAR")
	p.StoreAttribute("DS_COLOR", "100 100 200")
	
	dlg := iup.Dialog(p, "SIZE=200x200,TITLE=\"PPlot Example\"")
	dlg.Show()
	
	p.StoreAttribute("REDRAW", "")
	
	iup.MainLoop()
}
