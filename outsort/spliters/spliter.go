package spliters

import (
	"bufio"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
)

type spliter struct {
	source      io.Reader
	total       int
	chunkSize   int
	workDirPath string
}

func (spl *spliter) Process() (result int, err error) {
	var total int
	chunk := make(container, 0, 512)

	scanner := bufio.NewScanner(spl.source)
	buf := make([]byte, spl.chunkSize)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	for scanner.Scan() {
		line := scanner.Bytes()
		cnt := total + len(line)
		if cnt < spl.chunkSize {
			chunk = append(chunk, string(line))
			total = cnt
			continue
		}
		if cnt == spl.chunkSize {
			chunk = append(chunk, string(line))
			total = 0
		} else {
			total = len(line)
		}

		sort.Sort(chunk)
		if err = spl.save(chunk); err != nil {
			return
		}
		chunk = make(container, 0, 512)

		if total == cnt {
			chunk = append(chunk, string(line))
		}
	}
	if chunk.Len() > 0 {
		sort.Sort(chunk)
		spl.save(chunk)
	}
	result = spl.total
	return
}

func (spl *spliter) save(chunk container) (err error) {
	fh, err := os.Create(path.Join(spl.workDirPath, "0_"+strconv.Itoa(spl.total))+".chunk")
	if err != nil {
		return
	}
	defer fh.Close()

	for i, line := range chunk {
		prefix := "\n"
		if i == 0 {
			prefix = ""
		}
		_, err = fh.Write([]byte(prefix+line))
		if err != nil {
			return
		}
	}
	spl.total++
	return
}
