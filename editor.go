package editor

import (
	"os"
	"os/exec"
)

// Run executes the default editor for the user environment on the OS.
func Run(filename string) error {
	cmd := exec.Command(DefaultEditorPath, append(DefaultEditorArgs, filename)...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}
