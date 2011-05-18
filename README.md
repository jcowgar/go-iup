Iup Go Wrapper
==============

go-iup is a [Go][1] wrapper around the [Iup][2] GUI toolkit. The project
was started on April 27, 2011.

Changes in go-iup vs. Iup in C
------------------------------

Documentation is minimal with go-iup because Iup's documentation is very good and valid since 
go-iup strives for a 1-to-1 mapping. However, there are some general changes to better fit into
the Go language.

1. Iup has been dropped from the function names. go-iup functions are already accessed by the 
   iup package name. So, IupOpen becomes iup.Open(), IupVersion() becomes iup.Version(), etc...
2. IUP_ has been dropped from the constant names for the same reason as #1. Thus, IUP_IGNORE becomes
   iup.IGNORE, IUP_DEFAULT becomes iup.DEFAULT, etc...
3. Anything as of Iup 3.5 that has been marked as deprecated has not and will not be wrapped in
   go-iup. No sense in wrapping it and then next release removing it. Just don't even start with it.
4. The old ACTION name has been replaced by a SetCallback method in Iup. Thus, any widget that
   expects an ACTION name on control creation in Iup C does not in go-iup. For example, 
   IupButton("Press Me", "PRESS_ME_ACTION") no longer takes the "PRESS_ME_ACTION" parameter.
5. All widgets can accept a variable number of optional parameters. These parameters, if a string,
   are considered attributes to be set on the newly created widget. If a valid Icallback type, then
   they are considered callbacks to be set on the newly created widget.

### Callbacks, Actions ... Old and New ###

```c
button = IupButton("Say Hello", "SAY_HELLO")
IupSetAttributes(button, "FLAT=YES,ALIGNMENT=ALEFT")
IupSetFunction("SAY_HELLO", printHello)
```

has been replaced with

```go
button := iup.Button("Say Hello", 
    "FLAT=YES,ALIGNMENT=ALEFT",
    (iup.ActionFunc)(printHello))
```

The optional parameters can appear in any order. For example the following is the same as the
prior:

```go
button := iup.Button("Say Hello", 
    "FLAT=YES", 
    (iup.ActionFunc)(printHello)), 
    "ALIGNMENT=ALEFT")
```

Installing the Iup library
--------------------------

go-iup deals with three projects from [Tecgraf][4]:

* Iup - Cross platform native control GUI library
* Cd - Cross platform canvas drawing library
* Im - Cross platform image library

To use go-iup you must install all three libraries. Download the appropriate archive files
from the [Iup Download Tips][3] page. You should then place the .a (library) files in
your lib/ directory and .h (header) files in your include/ directory. If compiling on 
Windows, .dll (dynamically linked library) files in your %PATH%.

Building go-iup
---------------

    $ git clone https://github.com/jcowgar/go-iup.git
    $ cd go-iup
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
