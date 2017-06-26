package main

import (
	"flag"

	amdsi "github.com/xianggong/amdsi/modules"
)

func main() {
	// Commandline flags
	filePtr := flag.String("i", "", "Input file to be parsed")
	flag.Parse()

	if *filePtr != "" {
		amdsi.Parse(*filePtr)
	}
}
