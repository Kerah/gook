package generator

import (
	"io"
	"bufio"
	"github.com/Kerah/gook/chunks"
)

type generator struct {
	line  int
	lines [][]byte
}

func fromReader(reader io.Reader, maxLineSize int) *generator {
	g := &generator{}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		g.lines = append(g.lines, chunks.SplitByLimit(scanner.Bytes(), maxLineSize)...)
	}
	return g
}

func (g *generator) Generate() []byte {
	data := g.lines[g.line]
	g.line++
	if g.line >= len(g.lines) {
		g.line = 0
	}
	return data
}
