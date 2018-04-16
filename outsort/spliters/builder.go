package spliters

import (
	"github.com/Kerah/gook/outsort"
	"io"
)

func New() outsort.SpliterBuilder {
	return &builder{}
}

type builder struct {
	chunkLimit int
	workDir string
	source io.Reader
}

func (b *builder) ChunkLimit(size int) outsort.SpliterBuilder {
	b.chunkLimit = size
	return b
}

func (b *builder) WorkDir(path string) outsort.SpliterBuilder {
	b.workDir = path
	return b
}

func (b *builder) Source(source io.Reader) outsort.SpliterBuilder {
	b.source = source
	return b
}

func (b *builder) Build() outsort.Processor {
	return &spliter{
		chunkSize: b.chunkLimit,
		workDirPath: b.workDir,
		source: b.source,
	}
}

