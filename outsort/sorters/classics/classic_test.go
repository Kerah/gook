package classics

import (
	"bufio"
	"io"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"strings"
	"sort"
)

func getWorkDir(t *testing.T) (workDirPath string) {
	workDirPath = path.Join(os.TempDir(), "gook_merger")
	require.NoError(t, os.Mkdir(workDirPath, os.ModePerm))
	return
}

func putToWorkDir(t *testing.T, workDir, filename string, body string) {
	f, err := os.Create(path.Join(workDir, filename))
	require.NoError(t, err)
	defer f.Close()
	_, err = f.Write([]byte(body))
	require.NoError(t, err)
}

const SplitedFile0 = `0000
1000
4001
6000`

const SplitedFile1 = `1001
2001
3100
7000`

const SplitedFile2 = `1101
2301
3400
9600`

const SplitedFile3 = `0096
1090
2200
9999`

func detachedRead(reader io.Reader) chan []string {
	ch := make(chan []string, 1)
	go func() {
		defer close(ch)
		buff := make([]string, 0, 8)
		sc := bufio.NewScanner(reader)
		for sc.Scan() {
			buff = append(buff, sc.Text())
		}
		ch <- buff
	}()
	return ch
}

func readMerged(t *testing.T, wd, fileName string) []string {
	f, err := ioutil.ReadFile(path.Join(wd, fileName))
	require.NoError(t, err)
	return strings.Split(string(f), "\n")
}

func TestSorter_Process(t *testing.T) {
	t.Run("last iteration (with 2 total files) on 0 step", func(t *testing.T) {
		wd := getWorkDir(t)

		putToWorkDir(t, wd, "0_0.chunk", SplitedFile0)
		putToWorkDir(t, wd, "0_1.chunk", SplitedFile1)

		rdr, wrt := io.Pipe()
		srt := &sorter{
			workDir:     wd,
			total:       2,
			step:        0,
			destination: wrt,
		}
		drd := detachedRead(rdr)
		ttl, err := srt.Process()
		wrt.Close()
		require.NoError(t, err)
		assert.Equal(t, 1, ttl)

		exp := []string{
			"0000",
			"1000",
			"1001",
			"2001",
			"3100",
			"4001",
			"6000",
			"7000",
		}

		assert.Equal(t, exp, <-drd)
		os.RemoveAll(wd)
	})
	t.Run("merge 3 total input files", func(t *testing.T) {
		wd := getWorkDir(t)
		putToWorkDir(t, wd, "0_0.chunk", SplitedFile0)
		putToWorkDir(t, wd, "0_1.chunk", SplitedFile1)
		putToWorkDir(t, wd, "0_2.chunk", SplitedFile2)
		_, wrt := io.Pipe()
		srt := &sorter{
			workDir:     wd,
			total:       3,
			step:        0,
			destination: wrt,
		}
		ttl, err := srt.Process()
		require.NoError(t, err)
		assert.Equal(t, 2, ttl)

		exp := append(strings.Split(SplitedFile0, "\n"), strings.Split(SplitedFile1, "\n")...)
		sort.Strings(exp)
		assert.Equal(t, exp, readMerged(t, wd, "1_0.chunk"))

		exp = strings.Split(SplitedFile2, "\n")
		sort.Strings(exp)
		assert.Equal(t, exp, readMerged(t, wd, "1_1.chunk"))
		os.RemoveAll(wd)

	})
	t.Run("first merge step with 4 total input files", func(t *testing.T) {
		wd := getWorkDir(t)
		putToWorkDir(t, wd, "0_0.chunk", SplitedFile0)
		putToWorkDir(t, wd, "0_1.chunk", SplitedFile1)
		putToWorkDir(t, wd, "0_2.chunk", SplitedFile2)
		putToWorkDir(t, wd, "0_3.chunk", SplitedFile3)

		_, wrt := io.Pipe()
		srt := &sorter{
			workDir:     wd,
			total:       4,
			step:        0,
			destination: wrt,
		}
		ttl, err := srt.Process()
		require.NoError(t, err)
		assert.Equal(t, 2, ttl)

		exp := append(strings.Split(SplitedFile0, "\n"), strings.Split(SplitedFile1, "\n")...)
		sort.Strings(exp)
		assert.Equal(t, exp, readMerged(t, wd, "1_0.chunk"))

		exp = append(strings.Split(SplitedFile2, "\n"), strings.Split(SplitedFile3, "\n")...)
		sort.Strings(exp)
		assert.Equal(t, exp, readMerged(t, wd, "1_1.chunk"))
		os.RemoveAll(wd)
	})

}
