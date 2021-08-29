package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/MickLuypaerts/search/search"
)

const (
	defaultBufSize = 4096
)

func main() {
	feFlag := flag.Bool("fe", false, "automatically makes a regex for finding file extensions: .go => \\.go")
	lnFlag := flag.Bool("ln", false, "print line number for files")
	bufFlag := flag.Int("buf", defaultBufSize, "sets the buffer size for printing")

	flag.Parse()
	if len(os.Args) < 3 {
		usage()
		return
	}
	pattern := os.Args[len(os.Args)-2]
	file := os.Args[len(os.Args)-1]
	if *feFlag {
		pattern = fmt.Sprintf("(\\%s$)", pattern)
	}
	err := search.Search(pattern, file, lnFlag, bufFlag)
	if err != nil {
		log.Println(err)
		return
	}
}

func usage() {
	fmt.Printf("Usage: %s [OPTION] [PATTERN] [FILE]\n", os.Args[0])
	fmt.Printf("Use %s -help for a list of flags.\n", os.Args[0])
}
