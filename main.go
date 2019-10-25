package main

import (
	"fmt"
	"reflect"
)

var (
	name, course string
	module float64
)

var (
	a, b, c = "xxx", 1.2, 3
)

type Abc interface {
	Area() float64
}

func main() {

	var a Abc

	fmt.Println("value", a)
	fmt.Println("type of ", reflect.TypeOf(a))


}%
