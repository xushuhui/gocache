package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	c := NewCache(0)

	c.Set("key1", ByteValue{b: []byte("111")})
	c.Set("key2", ByteValue{b: []byte("22")})
	c.Set("key2", ByteValue{b: []byte("33")})
	t.Log(c.lru.Len())

	t.Log(c.lru.Get("key1"))
	t.Log(c.lru.Get("key2"))
	c.lru.Clear()
	t.Log(c.lru.Len())
}
