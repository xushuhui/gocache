package lru

import (
	"container/list"
	"fmt"
	"time"
)

type Cache struct {
	maxBytes int64
	useBytes int64
	ll       *list.List
	data     map[string]*list.Element
	expires  map[string]int64
}
type ByteValue struct {
	B interface{}
}

func (v ByteValue) Len() int {
	return 1
}

type entry struct {
	key   string
	value ByteValue
}

func New(maxBytes int64) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll:       list.New(),
		data:     make(map[string]*list.Element),
		expires:  make(map[string]int64),
	}
}
func (c *Cache) Resize(maxBytes int64) {
	c.maxBytes = maxBytes

}

func (c *Cache) Add(key string, value ByteValue, duration int64) {

	if ele, ok := c.data[key]; ok {
		c.ll.MoveToFront(ele)

		kv := ele.Value.(*entry)
		c.useBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value

	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.data[key] = ele
		c.useBytes += int64(len(key)) + int64(value.Len())

	}
	if duration > 0 {
		c.expires[key] = duration
	} else {
		c.expires[key] = 0
	}
	for c.maxBytes != 0 && c.maxBytes < c.useBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key string) (value ByteValue, ok bool) {
	ele, ok := c.data[key]
	if !ok {
		return
	}
	expire, ok := c.expires[key]
	if !ok {
		return
	}
	fmt.Println("expire", key, expire, time.Now().Unix())
	if expire > 0 && expire <= time.Now().Unix() {

		c.remove(key)
		return
	}
	c.ll.MoveToFront(ele)
	kv := ele.Value.(*entry)
	value = kv.value
	return value, true

}
func (c *Cache) Remove(key string) {
	c.remove(key)
	return
}
func (c *Cache) remove(key string) {
	if ele, ok := c.data[key]; ok {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.data, kv.key)
		c.useBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		delete(c.expires, kv.key)
		return
	}
}
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.data, kv.key)
		c.useBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		delete(c.expires, kv.key)
	}
}
func (c *Cache) Clear() {
	_ = c.ll.Init()
	for i := range c.data {
		delete(c.data, i)
	}

	c.useBytes = 0
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
