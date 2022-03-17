package lru

import (
	"container/list"
)

type Cache struct {
	maxBytes int64
	useBytes int64
	ll       *list.List
	data     map[string]*list.Element
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
		data:     make(map[string]*list.Element),
	}
}
func (c *Cache) Resize(maxBytes int64) {
	c.maxBytes = maxBytes

}

//func (c *Cache)String() string {
//	var buffer bytes.Buffer
//	buffer.WriteString("[")
//	for k,v := range c.data {
//		buffer.WriteString(fmt.Sprint(k))
//		buffer.WriteString(": ")
//		ele := v.Value.(*entry)
//
//		fmt.Printf("%T\n", ele.value.String())
//		fmt.Println(ele.value.String())
//		buffer.WriteString(ele.value.String())
//		buffer.WriteString(", ")
//	}
//	buffer.WriteString("]")
//
//
//	return buffer.String()
//}
func (c *Cache) Add(key string, value Value) {

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
	for c.maxBytes != 0 && c.maxBytes < c.useBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key string) (value interface{}, ok bool) {
	if ele, ok := c.data[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}
func (c *Cache) Remove(key string) {
	if ele, ok := c.data[key]; ok {
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
		delete(c.data, kv.key)
		c.useBytes -= int64(len(kv.key)) + int64(kv.value.Len())

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
