package main

import "fmt"

type shaper interface {
	area() int
	getName() string
}

type circle struct {
	r    int
	name string
}

func (c *circle) area() int {
	return c.r * c.r
}

func (c *circle) getName() string {
	return c.name
}

type rectangular struct {
	w, h int
	name string
}

func (r *rectangular) area() int {
	return r.w * r.h
}

func (r *rectangular) getName() string {
	return r.name
}

func main() {

	c := &circle{2, "circle"}
	r := &rectangular{2, 3, "rectangular"}
	shapes := []shaper{c, r}

	//接口的多肽
	for _, item := range shapes {
		fmt.Printf("%s area = %d \n", item.getName(), item.area())
	}

}
