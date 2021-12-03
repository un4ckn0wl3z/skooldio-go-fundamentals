package main

import (
	"errors"
	"fmt"
)

func validateLength(s string) error {
	if len([]rune(s)) < 8 {
		return errors.New("too short string")
	}
	return nil
}

var ageError = errors.New("your age is negative!")

type ErrTooYoung int

func (err ErrTooYoung) Error() string {
	return fmt.Sprintf("%d is too young", err)
}

func validateAge(n int) error {
	if n < 0 {
		return ageError
	}
	if n < 18 {
		return ErrTooYoung(n)
	}
	return nil
}

func main() {
	fmt.Println(validateAge(17))
	fmt.Println(validateLength("Hello"))

}
