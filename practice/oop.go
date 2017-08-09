package main

import (
	"fmt"
)

type Calculator struct {
	x int
	y int
}

func (cal Calculator) sum() int {
	return cal.x + cal.y
}

func (cal Calculator) sub() int {
	return cal.x - cal.y
}

func (cal Calculator) mul() int {
	return cal.x * cal.y
}

func (cal Calculator) rat() int {
	return cal.mul() - cal.sum()
}

func (cal Calculator) square(x int) int {
	return x * x
}

func main() {

	var cal Calculator = Calculator{x: 5, y: 6}

	fmt.Println(cal.square(5))
}
