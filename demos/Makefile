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

include $(GOROOT)/src/Make.inc

ifeq ($(GOOS),windows)
EXE=.exe
else
EXE=
endif

ALL = \
	test$(EXE) \
	versions$(EXE) \
	hello$(EXE) \
	greeter$(EXE) \
	rot13$(EXE) \
	getparam$(EXE) \
	alarm$(EXE) \
	dlgs$(EXE) \
	webbrowser$(EXE)

INC_DIRS=-I../iup/_obj
LIB_DIRS=-L../iup/_obj

all: $(ALL)

clean:
	rm -rf *.[68] $(ALL)

%$(EXE): %.go
	$(GC) $(INC_DIRS) $*.go
	$(LD) $(LIB_DIRS) -o $@ $*.$O
