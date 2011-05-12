# Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>
# 
# This file is part of go-iup.
#
# go-iup is free software: you can redistribute it and/or modify
# it under the terms of the GNU Lesser General Public License as
# published by the Free Software Foundation, either version 3 of
# the License, or (at your option) any later version.
# 
# go-iup is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
# 
# You should have received a copy of the GNU Lesser General Public
# License along with go-iup.  If not, see <http://www.gnu.org/licenses/>.

# Use "gomake install" to build and install

include $(GOROOT)/src/Make.inc

TARG=github.com/jcowgar/go-iup

CGOFILES=\
	constants.go \
	common.go \
	core.go \
	attributes.go \
	events.go \
	callback.go \
	layout.go \
	dialog.go \
	controls.go \
	resources.go \
	pplot.go

GOFILES=\
	doc.go

CGO_LDFLAGS_BASE=-liupcontrols -liupcd -liup_pplot -liupgl -liuptuio -liupim

ifeq ($(GOOS),windows)
CGO_LDFLAGS=-liup $(CGO_LDFLAGS_BASE) -lgdi32 -lole32 -lcomdlg32 -lcomctl32
else
CGO_CFLAGS=-I/usr/include/iup
CGO_LDFLAGS=-liupgtk $(CGO_LDFLAGS_BASE) -lXm
endif

include $(GOROOT)/src/Make.pkg
