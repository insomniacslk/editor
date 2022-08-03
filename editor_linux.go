package editor

import "os"

var (
	DefaultEditorPath = os.ExpandEnv("$EDITOR")
	DefaultEditorArgs = []string{}
)
