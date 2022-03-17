package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	c := NewCache(8)

	c.Set("k1", true)
	c.Set("k2", 1)
	s, _ := c.Get("k1")
	t.Log(s)

}
