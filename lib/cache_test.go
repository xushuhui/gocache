package lib

import (
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	var m sync.Map

	m.Store(`foo`, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {

			m.Store(`foo`, 1)
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {

			m.Store(`foo`, 1)
		}
	}()
	wg.Wait()
}
func TestMap(t *testing.T) {
	m := make(map[string]int, 1)
	m[`foo`] = 1
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			m[`foo`]++
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			m[`foo`]++
		}
	}()
	wg.Wait()
}
func TestCurrent(t *testing.T) {
	var wg sync.WaitGroup
	// 准备一系列的网站地址
	cache := NewLocal()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		// 开启一个并发
		go func(i int) {
			// 使用defer, 表示函数完成时将等待组值减1
			defer wg.Done()
			cache.Set("int", i)
			printMemStats()
			res, ok := cache.Get("int")
			if !ok {
				t.Log(res)
			}

		}(i)
	}

	wg.Wait()
}

func TestLocal_Del(t *testing.T) {
	cache := NewLocal()

	printMemStats()

	cache.Set("data", map[string]interface{}{"a": 1})
	printMemStats()
	cache.Set("data1", map[string]interface{}{"a": 1, "B": 2})
	printMemStats()
	cache.Set("data2", map[string]interface{}{"a": 11, "B": 2})
	printMemStats()
	cache.Set("data3", map[string]interface{}{"a": 12, "B": 2})
	printMemStats()
}
