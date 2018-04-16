package classics

import (
	"io"

	"github.com/Kerah/gook/outsort"
)

func New() outsort.SorterBuilder {
	return &builder{}
}

type builder struct {
	total       int
	step        int
	destination io.Writer
	workDir     string
}

func (b *builder) WorkDir(path string) outsort.SorterBuilder {
	b.workDir = path
	return b
}

func (b *builder) SplitTotal(total int) outsort.SorterBuilder {
	b.total = total
	return b
}

func (b *builder) Step(step int) outsort.SorterBuilder {
	b.step = step
	return b
}

func (b *builder) Destination(dest io.Writer) outsort.SorterBuilder {
	b.destination = dest
	return b
}

func (b *builder) Build() outsort.Processor {
	return &sorter{
		destination: b.destination,
		workDir:     b.workDir,
		step:        b.step,
		total:       b.total,
	}
}
