package generator

import "io"

type Generator interface {
	Generate() []byte
}

type Builder interface {
	Source(source io.Reader) Builder
	MaxLineSize(size int) Builder
	Build() (g Generator)
}
