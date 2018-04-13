package generator

import (
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
)


const greatesDEPoem = `
maine klaine porosenok
vdol po shtrasse
pobezhal
`

func newTestReader() io.Reader {
	return bytes.NewReader([]byte(greatesDEPoem))
}

func readFromGenerator(g Generator, attempts int) []string {
	out := make([]string,attempts)

	for i:=0; i<attempts; i++ {
		out[i] = string(g.Generate())
	}
	return out
}

func TestGenerator_Generate(t *testing.T) {
	g := fromReader(newTestReader(), 8)

	assert.NotEmpty(t, g.lines)
	expects := []string{
		"maine kl",
		"aine por",
		"osenok",
		"vdol po ",
		"shtrasse",
		"pobezhal",
		"maine kl",
	}
	out := readFromGenerator(g, len(g.lines)+1)

	assert.Equal(t, expects, out)
}
