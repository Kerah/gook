package generator

import "io"

func New() Builder {
	return &builder{}
}

type builder struct {
	maxSize int
	source  io.Reader
}
func (b *builder) Source(source io.Reader) Builder {
	b.source = source
	return b
}

func (b *builder) MaxLineSize(size int) Builder {
	b.maxSize = size
	return b
}

func (b *builder) Build() (g Generator) {
	return fromReader(b.source, b.maxSize)
}

