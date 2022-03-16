package main

type ByteValue struct {
	b []byte
}

func (v ByteValue) Len() int {
	return len(v.b)
}

func (v ByteValue) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
