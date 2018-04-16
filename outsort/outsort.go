package outsort

import (
	"io"
)

type Sorter interface {
	Sort(reader io.Reader, writer io.Writer) (err error)
}

type Processor interface {
	Process() (total int, err error)
}

type SorterBuilder interface {
	WorkDir(path string) SorterBuilder
	SplitTotal(total int) SorterBuilder
	Step(step int) SorterBuilder
	Destination(dest io.Writer) SorterBuilder
	Build() Processor
}

type SpliterBuilder interface {
	Source(source io.Reader) SpliterBuilder
	ChunkLimit(size int) SpliterBuilder
	WorkDir(path string) SpliterBuilder
	Build() Processor
}

type Builder interface {
	ChunkLimit(size int) Builder
	WorkDir(path string) Builder
	SpliterBuilder(builder SpliterBuilder) Builder
	SorterBuilder(builder SorterBuilder) Builder
	Build() Sorter
}

func New() Builder {
	return &builder{}
}

type builder struct {
	spliter   SpliterBuilder
	sorter    SorterBuilder
	workDir   string
	chunkSize int
}

func (b *builder) ChunkLimit(size int) Builder {
	b.chunkSize = size
	return b
}

func (b *builder) WorkDir(path string) Builder {
	b.workDir = path
	return b
}

func (b *builder) SpliterBuilder(builder SpliterBuilder) Builder {
	b.spliter = builder
	return b
}

func (b *builder) SorterBuilder(builder SorterBuilder) Builder {
	b.sorter = builder
	return b
}

func (b *builder) Build() Sorter {
	return &sorter{
		workDir:        b.workDir,
		chunkSize:      b.chunkSize,
		sorterBuilder:  b.sorter,
		spliterBuilder: b.spliter,
	}
}

type sorter struct {
	workDir        string
	chunkSize      int
	sorterBuilder  SorterBuilder
	spliterBuilder SpliterBuilder
}

func (srt *sorter) Sort(reader io.Reader, writer io.Writer) (err error) {
	total, err := srt.spliterBuilder.
		ChunkLimit(srt.chunkSize).
		WorkDir(srt.workDir).
		Source(reader).
		Build().
		Process()
	if err != nil {
		return
	}
	var step int
	for total > 1 {
		total, err = srt.sorterBuilder.
			SplitTotal(total).
			WorkDir(srt.workDir).
			Step(step).
			Destination(writer).
			Build().
			Process()
		if err != nil {
			return
		}
		step++
	}
	return
}
