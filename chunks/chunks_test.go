package chunks

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSplitByLimit(t *testing.T) {
	t.Run("check split empty slice", func(t *testing.T) {
		in := []byte{}
		exp := make([][]byte, 0)

		out := SplitByLimit(in, 100)
		assert.Equal(t, exp, out)
	})
	t.Run("check split nil slice", func(t *testing.T) {
		exp := make([][]byte, 0)

		out := SplitByLimit(nil, 100)
		assert.Equal(t, exp, out)
	})
	t.Run("check correct split slice", func(t *testing.T) {
		in := []byte("test")
		exp := [][]byte{
			[]byte("t"),
			[]byte("e"),
			[]byte("s"),
			[]byte("t"),
		}

		out := SplitByLimit(in, 1)
		assert.Equal(t, exp, out)
	})
	t.Run("check correct split small slice", func(t *testing.T) {
		in := []byte("test")
		exp := [][]byte{
			in,
		}

		out := SplitByLimit(in, 100)
		assert.Equal(t, exp, out)
	})
}
