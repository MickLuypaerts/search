package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"search/search"
)

func usage() {
	fmt.Printf("Usage: %s [OPTION] [PATTERN] [FILE]\n", os.Args[0])
	fmt.Printf("Use %s -help for a list of flags.\n", os.Args[0])
}

func main() {
	feFlag := flag.Bool("fe", false, "automatically makes a regex for finding file extensions: .go => \\.go")
	flag.Parse()
	if len(os.Args) < 3 {
		usage()
		return
	}
	pattern := os.Args[len(os.Args)-2]
	file := os.Args[len(os.Args)-1]
	if *feFlag {
		pattern = fmt.Sprintf("(\\%s)", pattern)
	}
	err := search.Search(pattern, file)
	if err != nil {
		log.Println(err)
		return
	}
}
