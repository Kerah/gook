package spliters

type container []string

func (c container) Len() int {
	return len(c)
}

func (c container) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c container) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}


