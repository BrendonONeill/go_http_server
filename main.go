package main

import (
	"fmt"
	"go_get_them/data"
)

func main() {

	var a string = "cat"
	b := 21
	c := true
	fmt.Printf("%v %.2f %v \n", a, b, c)
	s := fmt.Sprintf("I am %f years old", 10.523)
	fmt.Println(s)
	g := testing(b, math)
	fmt.Printf("%v %v", b, g)
	dog()
	fmt.Printf("%v is number 3 in the list", data.People[2])
}

func testing(cat int, add func(a int) int) (g int) {
	o := add(cat)
	return o
}

func math(a int) int {
	var l int
	l = a + 20
	return l
}
