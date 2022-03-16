package lib

import (
	"log"
	"runtime"
	"sync"
	"time"
)

type Cache interface {
	//size 是⼀个字符串。⽀持以下参数: 1KB，100KB，1MB，2MB，1GB 等
	SetMaxMemory(size string) bool
	// 设置⼀个缓存项，并且在expire时间之后过期
	Set(key string, val interface{})
	// 获取⼀个值
	Get(key string) (interface{}, bool)
	// 删除⼀个值
	Del(key string) bool
	// 检测⼀个值 是否存在
	Exists(key string) bool
	// 情况所有值
	Flush() bool
	// 返回所有的key 多少
	Keys() int64
}

type Local struct {
	items sync.Map
	count int64
}
type Item struct {
	item       sync.Map
	Expiration int64
}

// 判断数据项是否已经过期
func (item *Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().Unix() > item.Expiration
}

func NewLocal() *Local {
	return &Local{
		items: sync.Map{},
	}
}

func (c *Local) SetMaxMemory(size string) bool {
	//1KB，100KB，1MB，2MB，1GB
	memory := map[string]struct{}{"1KB": {}, "100KB": {}, "1MB": {}, "2MB": {}, "1GB": {}}
	if _, ok := memory[size]; !ok {
		panic("unsupport parm")
	}
	return true
}

func (c *Local) Set(key string, val interface{}) {
	isOk := c.Exists(key)
	if !isOk {

		c.items.Store(key, val)
		c.count++
	}

}
func (c *Local) SetEx(key string, val interface{}, expire time.Duration) {
	c.items.Store(key, val)
	c.count++
	//todo timeexpire
}
func (c *Local) Get(key string) (interface{}, bool) {
	// todo timeexpire
	return c.items.Load(key)
}

func (c *Local) Del(key string) bool {
	c.items.Delete(key)
	c.count--
	return true
}

func (c *Local) Exists(key string) bool {
	_, ok := c.items.Load(key)
	return ok
}

func (c *Local) Flush() bool {
	c.items.Range(func(key, value interface{}) bool {
		c.items.Delete(key)
		return true
	})
	c.count = 0
	return true
}

func (c *Local) Keys() int64 {

	return c.count
}
func printMemStats() {

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)

}
