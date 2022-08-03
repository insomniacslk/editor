package main

import (
	"flag"
	"log"

	"github.com/insomniacslk/editor"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatalf("No filename specified")
	}
	filename := flag.Arg(0)
	if err := editor.Open(filename); err != nil {
		log.Fatalf("Editor failed: %v", err)
	}
}
