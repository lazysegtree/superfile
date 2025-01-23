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
- [ ] ASCII control characters breaking layout
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
| https://github.com/yorukot/superfile/pull/555 | :mega: | Minor fix in log statements |
| https://github.com/yorukot/superfile/pull/556 | :mega: | Fix in isTextFile check |

# Windows Issues

### Action Items
- [x] Get windows running in a VM.

## Issue with Hyper terminal on windows 
https://github.com/yorukot/superfile/issues/332
- [ ] Reproduce

## Ctrl D not working
https://github.com/yorukot/superfile/issues/380
- [x] Reproduce
- [ ] Fix

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

# Non-Windows-specific Issues

## Background color issue
https://github.com/yorukot/superfile/issues/74
- [ ] Reproduce

## Focus_on_metadata hotkey doesn't work
https://github.com/yorukot/superfile/issues/384
- [x] Reproduce
- [ ] See if it can be fixed

## ✅ Timezone issue
https://github.com/yorukot/superfile/issues/434
- Not a bug

## Freeze 
https://github.com/yorukot/superfile/issues/455
- Asked about reproduction of issue

## Display issue due to text Files with control characters
https://github.com/yorukot/superfile/issues/495
- [x] Reproduce
- [ ] Fix

## Option menu with two panels 
https://github.com/yorukot/superfile/issues/497
- [ ] Reproduce
- Arch linux user with kitty terminal. 
- This is low priority

## Bulk delete crash
https://github.com/yorukot/superfile/issues/542
- [ ] Reproduce
  - Couldnt' - waiting for responses.



# Setup issues

## JSON breaks UI
https://github.com/yorukot/superfile/issues/457
- Cant reproduce. Asked for more info.

## Icon not working
https://github.com/yorukot/superfile/issues/465
- [ ] Fix your setup and reply to this too

## Command line not working AND back button misbehaving
https://github.com/yorukot/superfile/issues/506
- Win setup issue.
- Asked if its fixed or not

## No color in MacOS
https://github.com/yorukot/superfile/issues/508

## ✅ Not working in termux
https://github.com/yorukot/superfile/issues/538
- Not a bug. It works.

# Other low priority issues

## Double ANSI 'esc' characters 
https://github.com/yorukot/superfile/issues/449
- I dont get it why would there be a single esc character

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

# Issueless stuff

## No way to toggle the sidebar. 

## Better error handling
- [ ] os.Stat() call in checkFirstUse, writeConfigFile
- [ ] err during InitTrash is not even logged. !!!!
- [ ] default case in type switch in model.Update()

## Comment typos
- [ ] "Handle message exchanging whithin the application"
- [ ] "Write data to the path file if it does exists"
- [ ] "check if the file is unsipported"

## Misc corner cases
- [ ] string_funcion.isTextFile read a hard coded value of buffer 1024, this should be a defined constant

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