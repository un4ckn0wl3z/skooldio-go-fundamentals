package main

import "fmt"

func main() {
	var a interface{}
	a = 10
	fmt.Printf("%T %v\n", a, a)

	var i int
	i = a.(int)
	fmt.Printf("%T %v\n", i, i)

	if i, ok := a.(float32); ok {
		fmt.Printf("%T %v\n", i, i)
	}

	var x, y Obj
	x = rectangle{w: 4, l: 4}
	y = triangle{w: 4, h: 4}

	fmt.Println(x.Area())
	fmt.Println(y.Area())

	if y, ok := y.(triangle); ok {
		y.h = 5
	}
	fmt.Println(y.Area())

}

type Obj interface {
	Area() float64
}

type rectangle struct {
	w, l float64
}

func (rec rectangle) Area() float64 {
	return rec.l * rec.w
}

type triangle struct {
	w, h float64
}

func (tri triangle) Area() float64 {
	return tri.w * tri.h * 0.5
}
