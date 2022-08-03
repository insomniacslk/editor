# editor

`editor` is a Go package to call the default editor configured in the user's
environment on their OS.

Currently Linux, Windows and macOS are supported. The default editors are:
* on Linux: determined by the environment variable `EDITOR`
* on Windows: `notepad.exe`
* on macOS: `TextEdit.app`
