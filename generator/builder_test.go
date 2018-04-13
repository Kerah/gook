package generator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestBuilder_Build(t *testing.T) {
	g := New().
		MaxLineSize(8).
		Source(newTestReader()).
		Build()

	out := readFromGenerator(g, 3)
	expects := []string{
		"maine kl",
		"aine por",
		"osenok",
	}
	assert.Equal(t, expects, out)

}
