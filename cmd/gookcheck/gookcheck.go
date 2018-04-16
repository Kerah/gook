package main

import (
	"flag"
	"os"
	"log"
	"bufio"
	"fmt"
)

var pathToCheckFile string

func init() {
	flag.StringVar(&pathToCheckFile, "in", "", "path to check file")
	flag.Parse()
}


func main() {
	fh, err := os.Open(pathToCheckFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	var prev string
	for scanner.Scan() {
		current := scanner.Text()
		if prev >  current {
			log.Fatalf("unsorted file: '%s' >  '%s'", prev, current)
		}
		prev = current

	}
	fmt.Println("All ok. File is sorted")
}
