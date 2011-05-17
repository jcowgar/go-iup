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

package pplot

/*
#cgo LDFLAGS: -liupcontrols -liupcd -liup_pplot
#cgo linux LDFLAGS: -liupgtk
#cgo windows LDFLAGS: -liup -lgdi32 -lole32 -lcomdlg32 -lcomctl32

#include <stdlib.h>
#include <iup.h>
#include <iupcontrols.h>
#include <iup_pplot.h>
*/
import "C"
import "unsafe"
import . "github.com/jcowgar/go-iup"

var pplotLibOpened = false

func PPlot(opts ...interface{}) *Ihandle {
	OpenControlLib()
	
	if pplotLibOpened == false {
		C.IupPPlotOpen()
		pplotLibOpened = true
	}

	ih := &Ihandle{H: C.IupPPlot()}
	
	for _, o := range opts {
		switch v := o.(type) {
		default:
			//iup.Decorate(ih, o)
		}
	}

	return ih
}

func PlotBegin(ih *Ihandle, strXdata int) {
	C.IupPlotBegin(ih.H, C.int(strXdata))
}

func PlotAdd(ih *Ihandle, x, y float64) {
	C.IupPlotAdd(ih.H, C.float(x), C.float(y))
}

func PlotAddStr(ih *Ihandle, x string, y float64) {
	cX := C.CString(x)
	defer C.free(unsafe.Pointer(cX))

	C.IupPlotAddStr(ih.H, cX, C.float(y))
}

func PlotEnd(ih *Ihandle) {
	C.IupPlotEnd(ih.H)
}

func PlotInsert(ih *Ihandle, index, sample_index int, x, y float64) {
	C.IupPlotInsert(ih.H, C.int(index), C.int(sample_index), C.float(x), C.float(y))
}

func PlotInsertStr(ih *Ihandle, index, sample_index int, x string, y float64) {
	cX := C.CString(x)
	defer C.free(unsafe.Pointer(cX))

	C.IupPlotInsertStr(ih.H, C.int(index), C.int(sample_index), cX, C.float(y))
}

/*
// Differs from IupPlotInsertPoints as `count' is determined automatically in this case
func PlotInsertPoints(ih *iup.Ihandle, index, sample_index int, x, y []float64) {
	count := len(x)
	cX := float64ArrayToC(x)
	cY := float64ArrayToC(y)

	C.IupPlotInsertPoints(ih.H, C.int(index), C.int(sample_index), &cX[0], &cY[0],
		C.int(count))
}

// Differs from IupPlotInsertPoints as `count' is determined automatically in this case
func PlotInsertStrPoints(ih *iup.Ihandle, index, sample_index int, x []string, y []float64) {
	count := len(x)

	cX := stringArrayToC(x)
	defer freeCStringArray(cX)

	cY := float64ArrayToC(y)

	C.IupPlotInsertStrPoints(ih.H, C.int(index), C.int(sample_index), &cX[0], &cY[0], C.int(count))
}

// Differs from IupPlotInsertPoints as `count' is determined automatically in this case
func PlotAddPoints(ih *iup.Ihandle, index int, x, y []float64) {
	count := len(x)
	cX := float64ArrayToC(x)
	cY := float64ArrayToC(y)

	C.IupPlotAddPoints(ih.H, C.int(index), &cX[0], &cY[0], C.int(count))
}

// Differs from IupPlotInsertPoints as `count' is determined automatically in this case
func PlotAddStrPoints(ih *iup.Ihandle, index int, x []string, y []float64) {
	count := len(x)

	cX := stringArrayToC(x)
	defer freeCStringArray(cX)

	cY := float64ArrayToC(y)

	C.IupPlotAddStrPoints(ih.H, C.int(index), &cX[0], &cY[0], C.int(count))
}

func PlotTransform(ih *iup.Ihandle, x, y float64) (int, int) {
	cIX := new(C.int)
	cIY := new(C.int)

	C.IupPlotTransform(ih.H, C.float(x), C.float(y), cIX, cIY)

	return int(*cIX), int(*cIY)
}
*/

