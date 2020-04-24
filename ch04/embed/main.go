package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spoke int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v\n", w)
	fmt.Println(reflect.TypeOf(w))
}
