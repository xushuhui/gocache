package main

import (
	"fmt"
	"gocache/lru"
	"sync"
)

type Cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func NewCache(cacheBytes int64) *Cache {
	if cacheBytes <= 0 {
		panic("cacheBytes must positive")
	}
	return &Cache{
		mu:         sync.Mutex{},
		lru:        lru.New(cacheBytes),
		cacheBytes: cacheBytes,
	}
}
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.lru.Add(key, ByteValue{b: value})
}

func (c *Cache) Get(key string) (value interface{}, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		fmt.Println(v.(ByteValue))
		return v.(ByteValue), ok
	}

	return
}
func (c *Cache) Del(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.lru.Remove(key)
	return true
}

// 检测⼀个值 是否存在
func (c *Cache) Exists(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return false
	}

	if _, ok := c.lru.Get(key); ok {
		return ok
	}
	return false
}

// 情况所有值
func (c *Cache) Flush() bool {
	c.lru.Clear()
	return true
}

// 返回所有的key 多少
func (c *Cache) Keys() int64 {
	return int64(c.lru.Len())
}

var memory = map[string]int64{"1KB": 1024, "100KB": 1024 * 100, "1MB": 1024 * 1024, "2MB": 2 * 1024 * 1024, "1GB": 1024 * 1024 * 1024}

func (c *Cache) SetMaxMemory(size string) bool {
	capacity, ok := memory[size]
	if !ok {
		panic("unsupport parmater")
	}
	c.cacheBytes = capacity
	c.lru.Resize(capacity)

	return true
}
