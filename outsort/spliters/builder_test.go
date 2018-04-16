package spliters

import (
	"testing"
	"path"
	"os"
	"github.com/stretchr/testify/require"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_Build(t *testing.T) {
	data := []byte(testText)
	wd := path.Join(os.TempDir(), "gook_spliter_builder")
	err := os.MkdirAll(wd, os.ModePerm)
	require.NoError(t, err)

	total, err := New().
		ChunkLimit(16).
		Source(bytes.NewReader(data)).
		WorkDir(wd).Build().
		Process()

	require.NoError(t, err)
	assert.Equal(t, 2, total)

	exp1 := []string{
		"0000",
		"1001",
		"4444",
		"9999",
	}
	exp2 := []string {
		"0001",
		"0002",
		"0003",
		"1002",
	}
	compareWith(t, exp1, path.Join(wd, "0_0.chunk"))
	compareWith(t, exp2, path.Join(wd, "0_1.chunk"))
	os.RemoveAll(wd)
}
