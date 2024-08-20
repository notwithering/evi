//go:build windows
// +build windows

package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/gobuffalo/envy"
)

func getEditor() {
	var err error
	editor, err = envy.MustGet("EDITOR")
	if err != nil {
		editor = "notepad"
	}
}

func openEditor() {
	var cmd *exec.Cmd
	if strings.HasPrefix(filename, "-") {
		cmd = exec.Command("cmd", "/C", "start", "/WAIT", editor, "--", filename)
	} else {
		cmd = exec.Command("cmd", "/C", "start", "/WAIT", editor, filename)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
