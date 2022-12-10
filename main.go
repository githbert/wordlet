package main

import (
	"flag"

	"github.com/githbert/wordlet/game"
)

type ProgArg struct {
	file    string
	nrwords int
	wait    int
}

func processArgs() *ProgArg {

	tmpFile := flag.String("file", "~/tmp/ordlista.txt", "Full path to dictionary file.")
	tmpNrWords := flag.Int("nrwords", 5, "Amount of words to in session.")
	tmpWait := flag.Int("wait", 3, "Number of seconds to wait before clearing the screen.")

	flag.Parse()

	a := ProgArg{
		file:    *tmpFile,
		nrwords: *tmpNrWords,
		wait:    *tmpWait,
	}

	return &a
}

func main() {
	pa := processArgs()
	game.GameLoop(pa.file, pa.nrwords, pa.wait)
}
