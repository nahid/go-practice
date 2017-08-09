package main

import (
	"fmt"
)

func main() {
	var ages []int

	for i := 20; i <= 50; i++ {
		ages = append(ages, i)
	}

	fmt.Println(ages)

}
