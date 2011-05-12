Iup Go Wrapper
==============

iup.go is a [Go][1] wrapper around the [Iup][2] GUI toolkit. The project
was started on April 27, 2011.

Changes in iup.go vs. Iup in C
------------------------------

Documentation is minimal with Iup.go because Iup's documentation is very good and valid since 
iup.go strives for a 1-to-1 mapping. However, there are some general changes to better fit into
the Go language.

1. Iup has been dropped from the function names. iup.go functions are already accessed by the 
   iup package name. So, IupOpen becomes iup.Open(), IupVersion() becomes iup.Version(), etc...
2. IUP_ has been dropped from the constant names for the same reason as #1. Thus, IUP_IGNORE becomes
   iup.IGNORE, IUP_DEFAULT becomes iup.DEFAULT, etc...
3. Functions that work directly on an Ihandle pointer are referenced as a receiver. For example
   IupSetAttribute(textWidget, "VALUE", "Hello, World!") becomes 
   textWidget.SetAttribute("VALUE", "Hello, World!"). IupShow(dialog) becomes dialog.Show(), 
   etc...
4. Anything as of Iup 3.5 that has been marked as deprecated has not and will not be wrapped in
   iup.go. No sense in wrapping it and then next release removing it. Just don't even start with it.
5. The old ACTION name has been replaced by a SetCallback method in Iup. Thus, any widget that
   expects an ACTION name on control creation in Iup C does not in Iup.go. For example, 
   IupButton("Press Me", "PRESS_ME_ACTION") no longer takes the "PRESS_ME_ACTION" parameter.
6. All widgets can accept a variable number of optional parameters. These parameters, if a string,
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

Wrapper Status
==============

System
------

* **DONE** IupOpen
* **DONE** IupClose
* iuplua_open
* **DONE** IupVersion
* **DONE** IupLoad
* **DONE** IupLoadBuffer
* **DONE** IupSetLanguage
* **DONE** IupGetLanguage

Attributes
----------

* **DONE** IupStoreAttribute
* **DONE** IupStoreAttributeId
* **DONE** IupSetAttribute
* **DONE** IupSetAttributeId
* **DONE** IupSetfAttribute
* **DONE** IupSetfAttributeId
* **DONE** IupSetfAttributeId2
* **DONE** IupSetAttributes
* **DONE** IupResetAttribute
* **DONE** IupSetAtt see also iup.SetAttrs
* **DONE** IupSetAttributeHandle
* **DONE** IupGetAttributeHandle
* **DONE** IupGetAttribute
* **DONE** IupGetAttributeId
* **DONE** IupGetIntId - IupList additional function
* **DONE** IupGetFloatId - IupList additional function
* IupGetAllAttributes
* **DONE** IupGetAttributes
* **DONE** IupGetFloat
* **DONE** IupGetInt
* **DONE** IupStoreGlobal
* **DONE** IupSetGlobal
* **DONE** IupGetGlobal

Events
------

* **DONE** IupMainLoop
* **DONE** IupMainLoopLevel
* **DONE** IupLoopStep
* **DONE** IupLoopStepWait
* **DONE** IupExitLoop
* **DONE** IupFlush
* IupGetCallback
* IupSetCallback
* IupSetCallbacks
* IupGetActionName
* IupGetFunction
* IupSetFunction
* **DONE** IupRecordInput
* **DONE** IupPlayInput

Layout/Construction
-------------------

* **DONE** IupCreate
* **DONE** IupCreatev
* **DONE** IupCreatep
* **DONE** IupDestroy
* **DONE** IupMap
* **DONE** IupUnmap
* IupGetAllClasses
* **DONE** IupGetClassName
* **DONE** IupGetClassType
* **DONE** IupClassMatch
* IupGetClassAttributes
* IupGetClassCallbacks
* **DONE** IupSaveClassAttributes
* **DONE** IupCopyClassAttributes
* **DONE** IupSetClassDefaultAttribute

Layout/Composition
------------------

* **DONE** IupFill
* **DONE** IupHbox
* **DONE** IupHboxv
* **DONE** IupVbox
* **DONE** IupVboxv
* **DONE** IupZbox
* **DONE** IupZboxv
* **DONE** IupRadio
* **DONE** IupNormalizer
* **DONE** IupNormalizerv
* **DONE** IupCbox
* **DONE** IupCboxv
* **DONE** IupSbox
* **DONE** IupSplit

Layout/Hierarchy
----------------

* **DONE** IupAppend
* **DONE** IupDetach
* **DONE** IupInsert
* **DONE** IupReparent
* **DONE** IupGetParent
* **DONE** IupGetChild
* **DONE** IupGetChildPos
* **DONE** IupGetChildCount
* **DONE** IupGetNextChild
* **DONE** IupGetBrother
* **DONE** IupGetDialog
* **DONE** IupGetDialogChild

Layout/Utilities
----------------

* **DONE** IupRefresh
* **DONE** IupRefreshChildren
* **DONE** IupUpdate
* **DONE** IupUpdateChildren
* **DONE** IupRedraw
* **DONE** IupConvertXYToPos

Dialogs/Reference
-----------------

* **DONE** IupDialog
* **DONE** IupPopup
* **DONE** IupShow
* **DONE** IupShowXY
* **DONE** IupHide

Dialogs/Predefined
------------------

* **DONE** IupFileDlg
* **DONE** IupMessageDlg
* **DONE** IupColorDlg
* **DONE** IupFontDlg
* **DONE** IupAlarm
* IupGetFile
* IupGetColor
* *In Progress* IupGetParam
* IupGetText
* IupListDialog
* **DONE** IupMessage
* **DONE** IupLayoutDialog
* **DONE** IupElementPropertiesDialog

Controls/Standard
-----------------

* **DONE** IupButton
* **DONE** IupCanvas
* **DONE** IupFrame
* **DONE** IupLabel
* **DONE** IupList
* **DONE** IupProgressBar
* **DONE** IupSpin
* **DONE** IupTabs
* **DONE** IupTabsv
* **DONE** IupText
    * **DONE** IupTextConvertLinColToPos
    * **DONE** IupTextConvertPosToLinCol
* **DONE** IupToggle
* **DONE** IupTree
* **DONE** IupVal

Controls/Additional
-------------------

* **DONE** IupCells
* **DONE** IupColorbar
* **DONE** IupColorBrowser
* **DONE** IupDial
* **DONE** IupMatrix
    * **DONE** IupMatSetAttribute
    * **DONE** IupMatStoreAttribute
    * **DONE** IupMatGetAttribute
    * **DONE** IupMatGetInt
    * **DONE** IupMatGetFloat
    * **DONE** IupMatSetfAttribute
* **DONE** IupGLCanvas
* **DONE** IupPPlot
    * **DONE** IupPlotBegin
    * **DONE** IupPlotAdd
    * **DONE** IupPlotAddStr
    * **DONE** IupPlotEnd
    * **DONE** IupPlotInsert
    * **DONE** IupPlotInsertStr
    * **DONE** IupPlotInsertPoints
    * **DONE** IupPlotInsertStrPoints
    * **DONE** IupPlotAddPoints
    * **DONE** IupPlotAddStrPoints
    * **DONE** IupPlotTransform
    * IupPlotPaintTo
* IupOleControl
* IupWebBrowser

Resources/Fonts
---------------

* **DONE** IupMapFont
* **DONE** IupUnMapFont

Resources/Images
----------------

* IupImage
* IupImageRGB
* IupImageRGBA
* IupImageLibOpen
* IupLoadImage
* IupSaveImage
* IupGetNativeHandleImage
* IupGetImageNativeHandle
* IupSaveImageAsText

Resources/Keyboard
------------------

* **DONE** IupNextField
* **DONE** IupPreviousField
* **DONE** IupGetFocus
* **DONE** IupSetFocus

Resources/Menus
---------------

* **DONE** IupItem
* **DONE** IupMenu
* **DONE** IupSeparator
* **DONE** IupSubmenu

Resources/Names
---------------

* IupSetHandle
* IupGetHandle
* IupGetName
* IupGetAllNames
* IupGetAllDialogs

Resources/Miscellaneous
-----------------------

* **DONE** IupClipboard
* **DONE** IupTimer
* **DONE** IupTuioClient
* **DONE** IupUser
* **DONE** IupHelp

[1]: http://golang.org                                       "Go"
[2]: http://www.tecgraf.puc-rio.br/iup/                      "Iup"
[3]: http://www.tecgraf.puc-rio.br/iup/en/download_tips.html "Iup Download Tips"
[4]: http://www.tecgraf.puc-rio.br/                          "Tecgraf"
