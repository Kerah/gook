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
		if prev > scanner.Text() {
			log.Fatalf("unsorted file")
		}
	}
	fmt.Println("All ok. File is sorted")
}
