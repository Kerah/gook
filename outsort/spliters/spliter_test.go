package spliters

import (
	"testing"
	"bytes"
	"os"
	"github.com/stretchr/testify/require"
	"path"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
)

const testText = `9999
4444
0000
1001
0001
1002
0003
0002`

func compareWith(t *testing.T, exp []string, filePath string ) {
	data, err := ioutil.ReadFile(filePath)
	require.NoError(t, err, filePath)
	assert.Equal(t, exp, strings.Split(string(data), "\n"), filePath)
}

func TestSpliter_Split(t *testing.T) {
	data := []byte(testText)
	wd := path.Join(os.TempDir(), "gook_spliter")
	err := os.MkdirAll(wd, os.ModePerm)
	require.NoError(t, err)

	spliter := spliter{
		source: bytes.NewReader(data),
		chunkSize: 16,
		workDirPath: wd,
	}
	cnt, err := spliter.Process()
	require.NoError(t, err)
	assert.Equal(t, 2, cnt)
	assert.Equal(t, 2, spliter.total)

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
