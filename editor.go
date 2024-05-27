package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gobuffalo/envy"
)

func getEditor() {
	var err error
	editor, err = envy.MustGet("EDITOR")
	if err != nil {
		fmt.Printf(eviError, "no editor specified in $EDITOR")
		os.Exit(1)
	}
}

func openEditor() {
	cmd := exec.Command(editor, os.Args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
