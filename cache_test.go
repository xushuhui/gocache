package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	c := NewCache(28)

	c.Set("k1", true)

	t.Log(c.Get("k1"))
	c.Set("k3", 1)
	t.Log(c.Get("k3"))
	c.SetEx("k2", 1, 3*time.Second)
	a, _ := c.Get("k2")
	t.Log(a)
	time.Sleep(3 * time.Second)
	a2, _ := c.Get("k2")

	t.Log(a2)
}

var wg sync.WaitGroup

func TestPrint(t *testing.T) {

	catch := make(chan struct{}, 1)
	dogch := make(chan struct{}, 1)
	fishch := make(chan struct{}, 1)
	catch <- struct{}{}
	for i := 0; i < 100; i++ {
		wg.Add(3)
		go cat(catch, dogch)
		go dog(dogch, fishch)
		go fish(fishch, catch)

	}
	wg.Wait()
}
func cat(catch, dogch chan struct{}) {
	<-catch
	fmt.Println("cat")
	dogch <- struct{}{}
	wg.Done()
}
func dog(dogch, fishch chan struct{}) {
	<-dogch
	fmt.Println("dog")
	fishch <- struct{}{}
	wg.Done()
}
func fish(fishch, catch chan struct{}) {
	<-fishch
	fmt.Println("fish")
	catch <- struct{}{}
	wg.Done()
}
