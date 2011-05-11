Iup Go Wrapper
==============

iup.go is a [Go][1] wrapper around the [Iup][2] GUI toolkit. The project
was started on April 27, 2011.

Building on Linux
-----------------

    $ git clone https://github.com/jcowgar/iup.go.git
    $ cd iup.go
    $ ./all.bash
    
Building on Windows
-------------------

To (easily) build on Windows you need to install MinGW and MSYS. Once installed, you must
then install the Iup related libraries. You should download the 'dllw4' lib.zip files for
Iup, Cd and Im. Links to all of these files can be found on the [Iup Download Tips][3] page.

* .a (library) files should go into your MinGW/lib directory
* .h (header) files should go into your MinGW/include directory
* .dll (dynamically linked library) files should go somewhere in your path. Depending on
  which controls you use, some of these files will need to be shipped with your application.

You can now launch a mSYS shell and follow the instructions for Building on Linux.

[1]: http://golang.org                                       "Go"
[2]: http://www.tecgraf.puc-rio.br/iup/                      "Iup"
[3]: http://www.tecgraf.puc-rio.br/iup/en/download_tips.html "Iup Download Tips"
