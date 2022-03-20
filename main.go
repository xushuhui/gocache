package main

import (
	"fmt"
)

func hello() []string {
	return nil
}
func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

type Person struct {
	age int
}

func main() {
	person := &Person{28}
	// 1.
	defer fmt.Println("1", person.age)
	// 2.
	defer func(p *Person) {
		fmt.Println("2", p.age)
	}(person)
	// 3.
	defer func() {
		fmt.Println("3", person.age)
	}()

	person.age = 29
	defer fmt.Println("4", person.age)
}
