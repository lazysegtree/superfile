# Splay
## Plan 
Get the setup and understand the product.
Fix a few issues.
## Basic Action Items
- [x] Complete local build
- [x] Understand usage 
  
  - [x] Remove Question mark
    - https://github.com/yorukot/superfile/issues/112
    - Install nerdfont and change font to nerd font in Iterm.
  - [x] Try themes
  - [x] Try multi panels
  - [x] Try config files

- [ ] Unsupported file text shows with black color
- [ ] cd on quit
- [ ] Fix file deletion 

## Issues - Priority sorted

- [x] ASCII control characters breaking layout
- [ ] Windows Operations
  - [ ] Copy/Paste
  - [ ] Deletion
  - [ ] 'e' and 'E' keys
- [ ] Use stat instead of exiftool
- [ ] Open list of dir
- [ ] Testcases

## PRs

:pencil2: Draft

:mega: Raised(Pending Reviews)

:electric_plug: Pending Review Fix

:hourglass: Pending Merge (No AI on me)

:white_check_mark: Merged

:no_entry_sign: PR declined/abandoned

| PR  | Status | Remarks |
| ------------- | ------------- | ------------- |
| https://github.com/yorukot/superfile/pull/555 | :white_check_mark: | Minor fix in log statements |
| https://github.com/yorukot/superfile/pull/556 | :no_entry_sign: | Fix in isTextFile check(Changes will go in #557) |
| https://github.com/yorukot/superfile/pull/557 | :white_check_mark: | Fix layout breaking and more improvements |
| https://github.com/yorukot/superfile/pull/558 | :white_check_mark: | Fix in delete operation |

## Issues helped with (No PRs)

| Issue  | Summary | Status | Remarks |
| ------------- | ------------- | ------------- | ------------- |
| https://github.com/yorukot/superfile/issues/538 | Cant run on Termux | Resolved(False positive) |  |
| https://github.com/yorukot/superfile/issues/434 | Timezone

 issue with version file | Resolved(False positive) |  |

---------------------------------------------------------------------------------------------
=============================================================================================
---------------------------------------------------------------------------------------------
# Windows Issues

### Action Items
- [x] Get windows running in a VM.
- [ ] Text stuff on windows
  - [ ] Config and hotkeys and file operations
  - [ ] Image preview file preview
  - [ ] performance of file operations
  - [ ] does exiftools works on windows ?

## Issue with Hyper terminal on windows 
https://github.com/yorukot/superfile/issues/332
- [ ] Reproduce

## Ctrl D not working
https://github.com/yorukot/superfile/issues/380
- [x] Reproduce
- [ ] Fix
  - Better to move file to recylce bin, not to local trash 
    - https://github.com/hymkor/trash-go/blob/master/main_windows.go
    - Or making https://github.com/rkoesters/xdg windows compatible .
    - https://github.com/trubitsyn/go-recyclebin/blob/master/recyclebin_windows.go

## display issue
https://github.com/yorukot/superfile/issues/383
- [ ] Reproduce

## 'e' or 'E' does not works.
https://github.com/yorukot/superfile/issues/437
- [x] Reproduce
- [ ] Fix

## Paste not working
https://github.com/yorukot/superfile/issues/514
https://github.com/yorukot/superfile/issues/552
- [x] Reproduce
- [ ] Fix

## Powershell and debian display issue
https://github.com/yorukot/superfile/issues/544
- [ ] Reproduce

## Crash on wsl
https://github.com/yorukot/superfile/issues/551

---------------------------------------------------------------------------------------------
# Non-Windows-specific Issues

## Background color issue
https://github.com/yorukot/superfile/issues/74
- [ ] Reproduce

## Focus_on_metadata hotkey doesn't work
https://github.com/yorukot/superfile/issues/384
- [x] Reproduce
- [ ] See if it can be fixed
- [ ] Check issue in windows and linux
- [ ] Do doc update

## Freeze 
https://github.com/yorukot/superfile/issues/455
- Asked about reproduction of issue

## Option menu with two panels 
https://github.com/yorukot/superfile/issues/497
- [ ] Reproduce
- Arch linux user with kitty terminal. 
- This is low priority

## Bulk delete crash
https://github.com/yorukot/superfile/issues/542
- [ ] Reproduce
  - Couldnt' - waiting for responses.

---------------------------------------------------------------------------------------------
# Setup issues

## JSON breaks UI
https://github.com/yorukot/superfile/issues/457
- Cant reproduce. Asked for more info.

## Command line not working AND back button misbehaving
https://github.com/yorukot/superfile/issues/506
- Win setup issue.
- Asked if its fixed or not

## No color in MacOS
https://github.com/yorukot/superfile/issues/508

---------------------------------------------------------------------------------------------
# Other low priority issues

## Double ANSI 'esc' characters 
https://github.com/yorukot/superfile/issues/449
- I dont get it why would there be a single esc character
---------------------------------------------------------------------------------------------
# Other ideas ( Not issues)

## Load multiple directories in separate panes
https://github.com/yorukot/superfile/issues/286

## Use built-in unix command stat instead of exiftool 
https://github.com/yorukot/superfile/issues/306



## Light and Dark theme option
https://github.com/yorukot/superfile/issues/468

## Sixel support
https://github.com/yorukot/superfile/issues/480
- [ ] Discuss amount of work needed and is this relevant 

## Paste from wl-clipboard
https://github.com/yorukot/superfile/issues/492


---------------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------------

# Issueless stuff

## Documentation
- [ ] Build steps to build from source

## Other bugs
- [ ] Layout breaking in ~/Applications
- [ ] Directories ending with .localized

## No way to toggle the sidebar. 

## Better error handling
- [ ] os.Stat() call in checkFirstUse, writeConfigFile
- [ ] err during InitTrash is not even logged. !!!!
- [ ] default case in type switch in model.Update()
- [ ] model operations on handle_file_operations.go dont return error, keep executing after failure


## Comment typos
- [ ] "Handle message exchanging whithin the application"
- [ ] "Write data to the path file if it does exists"
- [ ] "check if the file is unsipported"
- [ ] "Paste item function move file to trash can error"per

## Misc corner cases
- [ ] string_funcion.isTextFile read a hard coded value of buffer 1024, this should be a defined constant
- [ ] ranger is able to detec a binary file renamed to .txt and not preview it.
  - `file` command can too.
  - `https://github.com/file-go/fil` - This go packet can do that too. Cross platform.
- [ ] Ctrl + p behaviour in case we have multiple files selected via 'v' mode. Right now, only one file's path gets copied.
- [ ] spf used to freeze for .pdf and  .pptx, any other extension it could freeze ?
- [ ] Processes sometimes just get stuck forever (ex : copy and paste in windows at this point of commit).

## Refractoring
- [ ] Code duplication in model_render.filePreviewPanelRender for text files.
- [ ] Constants like "\n --- " + icon.Error + " Error open file ---" should be defined at one place
- [ ] Code duplication in model.View()
- [ ] key_functions.containsKey() better variable names
- [ ] key_functions.mainKey() indentation fix

- [ ] Better logging. And adding more logs in verbose mode to help debugging.
- [ ] .pdf hardcoded as unsupported extension
- [ ] float64(fileInfo.Size())/(1024*1024) < 250
- [ ] Hardcoded buffer size 1024 at many places
- [ ] Duplicacy in file_operations.pasteDir and file_operations.moveElement
  - [ ] pastDir is basicall .moveElements + directory support.

## Testing
- [ ] Use testify library

## Performance Optimizations
- [ ] model_render.filePreviewPanelRender being called 2-3 times per file
- [ ] File copying via os syscalls not via clipboard
- [ ] Needless clipboard usage in case of multi file copy

# Usage and Setup

## Build in windows
- Install golang
- Build
```
go build -o ./bin/spf.exe
# Execute

./bin/spf.exe
```

## Build in mac
```
➜  ~/Workspace/kuknitin/superfile git:(test) ✗ [12:44:51] ./build.sh
➜  ~/Workspace/kuknitin/superfile git:(test) ✗ [12:45:02] ls -ltr ./bin/spf
-rwxr-xr-x  1 kuknitin  staff  28244050 Jan 23 00:45 ./bin/spf
➜  ~/Workspace/kuknitin/superfile git:(test) ✗ [12:45:04]
```

# Implementation Notes

## Questions
- [ ] What is TrashDirectories used for, and why not Mac has them.
  
- [x] log.Fatalln bugs.
  - PR raised - https://github.com/yorukot/superfile/pull/555

- [ ] Fix logging - setLogOutput early

# References
- Superfile Todo list - https://github.com/users/yorukot/projects/4


time: fix example for RFC3339 and RFC3339Nano time layout
Added test cases that validates the correct parsing of example
Fixes #71378

# Old Stuff 
Information about stuff that is not needed anymore 


## ✅ Display issue due to text Files with control characters
https://github.com/yorukot/superfile/issues/495
- [x] Reproduce
- [ ] Fix
  - Removing non Graphic characters is one way
  - ranger seems to convert `\v` to newline , and print `\b` correctly 
  - cat prints `\v` and `\b` correctly and doesn't cause any issues
  - lf seems to just remove `\v` or `\b`
    - Not always, see all_ctrl_char.txt (ctrl+j is four spaces, and ctrl+k is newline)
  - mc fucking converts them to `.`. Damn.
    - Not always, see all_ctrl_char.txt
  - yazi replaces them to `^K` and `^H`.
    - Not always. (all_ctrl_char.txt.go) 
  - we can use `%q` -> But it adds quotes, and results in ugly output.
  - nerd font does not have icons for ascii control characters
  - Loss of info is okay. See cat for all_ctrl_char.txt

- Caution
  - Cant ignore \x09 - horizontal tab
  - Cant ignore \x0a - line feed
  - Ignore all else 
  - Indistinguishable from vertical tab via unicode
- References
  -  Unicode IsGraphic and IsPrint, IsControl, IsSpace, https://pkg.go.dev/unicode
  - https://www3.rocketsoftware.com/bluezone/help/v42/en/bzadmin/APPENDIX_C/ASCII_Character_Set.htm
  - https://vimdoc.sourceforge.net/htmldoc/digraph.html#digraph-table 


## ✅ Not working in termux
https://github.com/yorukot/superfile/issues/538
- Not a bug. It works.

## ✅ Timezone issue
https://github.com/yorukot/superfile/issues/434
- Not a bug

## ✅ Icon not working
https://github.com/yorukot/superfile/issues/465
- [x] Fix your setup and reply to this too
