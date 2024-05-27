package main

import (
	"github.com/notwithering/sgr"
)

const (
	eviError    string = sgr.FgHiRed + "error:" + sgr.Reset + " %s\n"
	eviInfo     string = sgr.FgHiBlue + "::" + sgr.Reset + " %s\n"
	eviInfoPair string = sgr.FgHiMagenta + "::" + sgr.Reset + " %s " + sgr.FgHiBlack + ":" + sgr.Reset + " %s\n"
	eviChoice   string = "   %d) %s\n"
	eviInput    string = ">> "
)

var (
	algorithm, mode  int
	key              []byte
	editor, filename string
)
