package main

import (
	"github.com/notwithering/sgr"
)

const (
	eviError          string = sgr.FgHiRed + "error:" + sgr.Reset + " %s\n"
	eviInfo           string = sgr.FgHiBlue + "::" + sgr.Reset + " %s\n"
	eviInfoPair       string = sgr.FgHiMagenta + "::" + sgr.Reset + " %s " + sgr.FgHiBlack + ":" + sgr.Reset + " %s\n"
	eviChoice         string = "   %d) %s\n"
	eviChoiceSelected string = " " + sgr.FgHiRed + ">" + sgr.Reset + " %d) %s\n"
	eviInput          string = ">> "
)

var (
	key              []byte
	editor, filename string
)
