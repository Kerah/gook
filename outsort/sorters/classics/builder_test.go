package classics

import (
	"testing"
	"io"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"sort"
)

func TestBuilder_Build(t *testing.T) {
	wd := getWorkDir(t)

	putToWorkDir(t, wd, "0_0.chunk", SplitedFile0)
	putToWorkDir(t, wd, "0_1.chunk", SplitedFile1)

	rdr, wrt := io.Pipe()

	prc := New().
		WorkDir(wd).
		Destination(wrt).
		SplitTotal(2).
		Step(0).
		Build()
	drd := detachedRead(rdr)

	total, err := prc.Process()
	wrt.Close()

	require.NoError(t, err)
	assert.Equal(t, 1, total)

	exp := append(strings.Split(SplitedFile0, "\n"), strings.Split(SplitedFile1, "\n")...)
	sort.Strings(exp)
	assert.Equal(t, exp, <-drd)
	os.RemoveAll(wd)
}
