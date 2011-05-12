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

package iup

/*
#include <stdlib.h>
#include <iup.h>
#include <iupcontrols.h>
#include <iup_pplot.h>
*/
import "C"
import "unsafe"

func (ih *Ihandle) PlotBegin(strXdata int) {
	C.IupPlotBegin(ih.h, C.int(strXdata))
}

func (ih *Ihandle) PlotAdd(x, y float64) {
	C.IupPlotAdd(ih.h, C.float(x), C.float(y))
}

func (ih *Ihandle) PlotAddStr(x string, y float64) {
	cX := C.CString(x)
	defer C.free(unsafe.Pointer(cX))
	
	C.IupPlotAddStr(ih.h, cX, C.float(y))
}

func (ih *Ihandle) PlotEnd() {
	C.IupPlotEnd(ih.h)
}

func (ih *Ihandle) PlotInsert(index, sample_index int, x, y float64) {
	C.IupPlotInsert(ih.h, C.int(index), C.int(sample_index), C.float(x), C.float(y))
}

func (ih *Ihandle) PlotInsertStr(index, sample_index int, x string, y float64) {	
	cX := C.CString(x)
	defer C.free(unsafe.Pointer(cX))
	
	C.IupPlotInsertStr(ih.h, C.int(index), C.int(sample_index), cX, C.float(y))
}

// Differs from IupPlotInsertPoints as `count' is determined automatically in this case
func (ih *Ihandle) PlotInsertPoints(index, sample_index int, x, y []float64) {
	count := len(x)
	cX := float64ArrayToC(x)
	cY := float64ArrayToC(y)
	
	C.IupPlotInsertPoints(ih.h, C.int(index), C.int(sample_index), &cX[0], &cY[0], 
		C.int(count))
}

// Differs from IupPlotInsertPoints as `count' is determined automatically in this case
func (ih *Ihandle) PlotInsertStrPoints(index, sample_index int, x []string, y []float64) {
	count := len(x)
	
	cX := stringArrayToC(x)
	defer freeCStringArray(cX)
	
	cY := float64ArrayToC(y)
	
	C.IupPlotInsertStrPoints(ih.h, C.int(index), C.int(sample_index), &cX[0], &cY[0], C.int(count))
}

// Differs from IupPlotInsertPoints as `count' is determined automatically in this case
func (ih *Ihandle) IupPlotAddPoints(index int, x, y []float64) {
	count := len(x)
	cX := float64ArrayToC(x)
	cY := float64ArrayToC(y)
	
	C.IupPlotAddPoints(ih.h, C.int(index), &cX[0], &cY[0], C.int(count))
}

// Differs from IupPlotInsertPoints as `count' is determined automatically in this case
func (ih *Ihandle) IupPlotAddStrPoints(index int, x []string, y []float64) {
	count := len(x)
	
	cX := stringArrayToC(x)
	defer freeCStringArray(cX)
	
	cY := float64ArrayToC(y)
	
	C.IupPlotAddStrPoints(ih.h, C.int(index), &cX[0], &cY[0], C.int(count))	
}

func (ih *Ihandle) PlotTransform(x, y float64) (int, int) {
	cIX := new(C.int)
	cIY := new(C.int)
	
	C.IupPlotTransform(ih.h, C.float(x), C.float(y), cIX, cIY)
	
	return int(*cIX), int(*cIY)
}
