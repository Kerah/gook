package main

import (
	"flag"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/Kerah/gook/outsort"
	"github.com/Kerah/gook/outsort/sorters/classics"
	"github.com/Kerah/gook/outsort/spliters"
)

var workDirPath string
var chunkLimitSize int
var inputFilePath string
var outPutFilePath string

func init() {
	flag.StringVar(&workDirPath, "workdir", "", "path working directory or used temp dir")
	flag.StringVar(&inputFilePath, "in", "", "input file")
	flag.StringVar(&outPutFilePath, "out", "", "output file")
	flag.IntVar(&chunkLimitSize, "chunk", 8, "allow in-memory chunk size in kb")
	flag.Parse()
}

func main() {
	start := time.Now()
	if workDirPath == "" {
		workDirPath = os.TempDir()
	}

	workDirPath = path.Join(workDirPath, strconv.Itoa(int(start.Unix())))
	err := os.Mkdir(workDirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("problem with creating work dir: %s", err.Error())
	}
	chunkLimitSize *= 1024
	input, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("problem with open input file: %s", err.Error())
	}

	output, err := os.Create(outPutFilePath)
	if err != nil {
		log.Fatalf("problem with create output file: %s", err.Error())
	}
	defer input.Close()
	defer output.Close()

	err = outsort.New().
		SpliterBuilder(spliters.New()).
		SorterBuilder(classics.New()).WorkDir(workDirPath).
		ChunkLimit(chunkLimitSize).
		Build().
		Sort(input, output)
	if err != nil {
		log.Fatalf("problem in sorting process: %s", err.Error())
	}

}
