package main

type ByteValue struct {
	b interface{}
}

func (v ByteValue) Len() int {
	return 1
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
