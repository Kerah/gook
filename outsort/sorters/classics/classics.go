package classics

import (
	"io"
	"os"
	"path"
	"strconv"

	"github.com/Kerah/gook/buffile"
)

type sorter struct {
	workDir     string
	total       int
	step        int
	destination io.Writer
}

func (srt *sorter) processPair(i int) (err error) {
	firstFile, err := srt.readFrom(i * 2)
	if err != nil {
		return
	}
	defer firstFile.Close()
	first := buffile.NewFromReadCloser(firstFile)
	secondFile, err := srt.readFrom(i*2 + 1)
	if err != nil {
		return
	}
	defer secondFile.Close()
	second := buffile.NewFromReadCloser(secondFile)

	dest := srt.destination
	if srt.total > 2 {
		var destFile *os.File
		destFile, err = os.Create(srt.getDestFilePath(i))
		if err != nil {
			return
		}
		defer destFile.Close()
		dest = destFile
	}

	for j := 0; j >= 0; j++ {
		left, leftErr := first.Read()
		right, rightErr := second.Read()
		if leftErr == io.EOF && rightErr == io.EOF {
			break
		}
		if leftErr != io.EOF && leftErr != nil {
			err = leftErr
			return
		}
		if rightErr != io.EOF && rightErr != nil {
			err = rightErr
			return
		}
		var toWrite string

		if leftErr == io.EOF {
			toWrite = second.UnBuf()
		} else if rightErr == io.EOF {
			toWrite = first.UnBuf()
		} else {
			if left < right {
				toWrite = first.UnBuf()
			} else {
				toWrite = second.UnBuf()
			}
		}
		if j > 0 {
			toWrite = "\n" + toWrite
		}

		if _, err = dest.Write([]byte(toWrite)); err != nil {
			return
		}
	}
	if err = os.Remove(srt.getSpliterFilePath(i * 2)); err != nil {
		return
	}
	if err = os.Remove(srt.getSpliterFilePath(i*2 + 1)); err != nil {
		return
	}
	return
}

func (srt *sorter) getDestFilePath(i int) string {
	return path.Join(srt.workDir, strconv.Itoa(srt.step+1)+"_"+strconv.Itoa(i)+".chunk")
}

func (srt *sorter) Process() (total int, err error) {
	total = srt.total / 2
	for i := 0; i < total; i++ {
		if err = srt.processPair(i); err != nil {
			return
		}
	}
	if total*2 < srt.total {
		oph := srt.getSpliterFilePath(srt.total - 1)
		nph := srt.getDestFilePath(total)
		err = os.Rename(oph, nph)
		total++
	}
	return
}

func (srt *sorter) getSpliterFilePath(ind int) string {
	return path.Join(srt.workDir, strconv.Itoa(srt.step)+"_"+strconv.Itoa(ind)+".chunk")

}

func (srt *sorter) readFrom(ind int) (*os.File, error) {
	fpa := srt.getSpliterFilePath(ind)
	return os.Open(fpa)
}
