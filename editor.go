package editor

import (
	"os"
	"os/exec"
)

// Editor contains the information to run the editor
type Editor struct {
	path string
	args []string
}

// New creates a new Editor object with default editor path and args.
func New() *Editor {
	return &Editor{
		path: DefaultEditorPath,
		args: DefaultEditorArgs,
	}
}

// DefaultEditor is the default editor.
var DefaultEditor = New()

// Open is like Editor.Open for the default editor.
func Open(filename string) error {
	return DefaultEditor.Open(filename)
}

// Set is like Editor.Set for the default editor.
func Set(path string, args ...string) {
	DefaultEditor.Set(path, args...)
}

// Get is like Editor.Get for the default editor.
func Get() (string, []string) {
	return DefaultEditor.Get()
}

// Set changes the default editor to the one specified by the user.
func (e *Editor) Set(path string, args ...string) {
	e.path = path
	e.args = args
}

// Get returns editor path and args.
func (e *Editor) Get() (string, []string) {
	return e.path, e.args
}

// Open opens the requested file in the editor.
func (e *Editor) Open(filename string) error {
	cmd := exec.Command(e.path, append(e.args, filename)...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}
