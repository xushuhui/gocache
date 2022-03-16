package lru

import "container/list"

type Cache struct {
	maxBytes int64
	useBytes int64
	ll       *list.List
	cache    map[string]*list.Element
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}

//todo 拉链法
func (c *Cache) Add(key string, value interface{}) {
	val := value.(Value)
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)

		kv := ele.Value.(*entry)
		c.useBytes += int64(val.Len()) - int64(kv.value.Len())
		kv.value = val
	} else {
		ele := c.ll.PushFront(&entry{key, val})
		c.cache[key] = ele
		c.useBytes += int64(len(key)) + int64(val.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.useBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key string) (value interface{}, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}
func (c *Cache) Remove(key string) {
	if ele, ok := c.cache[key]; ok {
		c.ll.Remove(ele)
		return
	}
	return
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.useBytes -= int64(len(kv.key)) + int64(kv.value.Len())

	}
}
func (c *Cache) Clear() {
	_ = c.ll.Init()
	for i, _ := range c.cache {
		delete(c.cache, i)
	}

	c.useBytes = 0
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
