package main

import (
	"flag"
	"github.com/Kerah/gook/generator"
	"os"
	"log"
)

var maxLineSize int
var readFilePattern string
var attempts int
var outFilePath string

func init() {
	flag.IntVar(&maxLineSize, "max", 100, "max line size for generate new file")
	flag.StringVar(&readFilePattern, "pattern", "/specify/path/to/patter/file", "file with pattern content")
	flag.IntVar(&attempts, "limit", 1000, "output file limit")
	flag.StringVar(&outFilePath, "output", "/specify/path/to/output/file", "file for generated output")

}

func main() {
	flag.Parse()
	fh, err := os.Open(readFilePattern)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()
	gen := generator.New().
		MaxLineSize(maxLineSize).
		Source(fh).
		Build()
	out, err := os.Create(outFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	delim := []byte("\n")
	total := 0
	for total < attempts {
		gen := gen.Generate()
		_, err := out.Write(gen)
		if err != nil {
			log.Fatal(err)
		}

		if _, err = out.Write(delim); err != nil {
			log.Fatal(err)
		}
		total += len(gen)
	}
	out.Sync()
}
