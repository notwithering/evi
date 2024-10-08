//go:build !windows
// +build !windows

package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/gobuffalo/envy"
	"github.com/k0kubun/go-ansi"
)

func getEditor() {
	var err error
	editor, err = envy.MustGet("EDITOR")
	if err != nil {
		ansi.Printf(eviError, "no editor specified in $EDITOR")
		os.Exit(1)
	}
}

func openEditor() {
	var cmd *exec.Cmd
	if strings.HasPrefix(filename, "-") {
		cmd = exec.Command(editor, "--", filename)
	} else {
		cmd = exec.Command(editor, filename)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
