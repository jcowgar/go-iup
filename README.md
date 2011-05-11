Iup Go Wrapper
==============

iup.go is a [Go][1] wrapper around the [Iup][2] GUI toolkit. The project
was started on April 27, 2011.

Installing the Iup library
--------------------------

Iup.go deals with three projects from [Tecgraf][4]:

* Iup - Cross platform native control GUI library
* Cd - Cross platform canvas drawing library
* Im - Cross platform image library

To use iup.go you must install all three libraries. Download the appropriate archive files
from the [Iup Download Tips][3] page. You should then place the .a (library) files in
your lib/ directory and .h (header) files in your include/ directory. If compiling on 
Windows, .dll (dynamically linked library) files in your %PATH%.

Building Iup.go
---------------

    $ git clone https://github.com/jcowgar/iup.go.git
    $ cd iup.go
    $ ./all.bash
    
Special notes for Windows
-------------------------

You should download the 'dllw4' variant of Iup. This is the dynamically linked version of
the libraries for MinGW. For easy compliation you should have already installed a working
version of MinGW and MSYS. Building will be the same as any other platform but started from
the MSYS Bash shell.

[1]: http://golang.org                                       "Go"
[2]: http://www.tecgraf.puc-rio.br/iup/                      "Iup"
[3]: http://www.tecgraf.puc-rio.br/iup/en/download_tips.html "Iup Download Tips"
[4]: http://www.tecgraf.puc-rio.br/                          "Tecgraf"
