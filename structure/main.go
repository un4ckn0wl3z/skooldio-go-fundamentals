package main

import (
	"fmt"
)

type rectangle struct {
	w, l float64
}

type Book struct {
	Name   string
	Author string
}

func (rec rectangle) Area() float64 {
	return rec.l * rec.w
}

func (rec *rectangle) SetWidth(w float64) {
	rec.w = w
}

func (book Book) String() string {
	return fmt.Sprintf("%s by %s", book.Name, book.Author)
}

func (book *Book) SetName(name string) {
	book.Name = name
}

func main() {
	rec := rectangle{
		w: 4,
		l: 5,
	}

	rec.w = 6
	fmt.Println(rec.l)
	fmt.Println(rec.w)

	fmt.Println(rec.Area())
	rec.SetWidth(5)
	fmt.Println(rec.Area())

	book := Book{
		Name:   "Harry Potter",
		Author: "J. K. Rowling",
	}

	fmt.Println(book.String())

	book.SetName("Anuwat")
	fmt.Println(book.String())

}
