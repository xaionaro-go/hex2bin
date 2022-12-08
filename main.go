package main

import (
	"encoding/hex"
	"flag"
	"io"
	"log"
	"os"
)

func assertNoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() > 1 {
		log.Fatalf("syntax: hex2bin [file path]")
	}
	input := os.Stdin
	if flag.NArg() == 1 {
		filePath := flag.Arg(0)
		f, err := os.Open(filePath)
		assertNoError(err)
		defer f.Close()
		input = f
	}

	trimmed := NewTrimmer(input)

	decoder := hex.NewDecoder(trimmed)
	_, err := io.Copy(os.Stdout, decoder)
	assertNoError(err)
}
