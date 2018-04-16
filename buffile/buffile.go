package buffile

import (
	"bufio"
	"io"
)

type Bufferer interface {
	UnBuf() (result string)
	Read() (result string, err error)
}

func NewFromReadCloser(rc io.ReadCloser) Bufferer {
	return &bufferedFile{
		file: rc,
		scanner: bufio.NewScanner(rc),
	}
}

type bufferedFile struct {
	scanner *bufio.Scanner
	file io.ReadCloser
	data []byte
}

func (bf *bufferedFile) UnBuf() (result string) {
	result = string(bf.data)
	bf.data = nil
	return
}

func (bf *bufferedFile) Read() (result string, err error) {
	if bf.data != nil {
		result = string(bf.data)
		return
	}

	if !bf.scanner.Scan() {
		err = io.EOF
		return
	}
	bf.data = bf.scanner.Bytes()
	result =  string(bf.data)
	return
}
