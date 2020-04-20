package main

import "fmt"

type simpler interface {
	get() int
	set(int)
}

type simple struct {
	v int
}

func (s *simple) get() int {
	return s.v
}

func (s *simple) set(val int) {
	s.v = val
}

func main() {
	var s simpler
	s = &simple{2}
	fmt.Printf("get() = %d \n", s.get())
	s.set(3)
	fmt.Printf("get() = %d \n", s.get())
}
