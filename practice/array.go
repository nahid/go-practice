package main

import (
	"fmt"
)

func main() {

	//var name = [...]string{"nahid", "sumi", "hasan"}

	/*for _, v := range name {
		fmt.Printf("%s \n", v)
	}*/

	/*var age = [...]int{3, 4, 5}

	age = [3]int{2, 3, 4}

	fmt.Printf("%x", age)*/

	var y, x [2]int

	x[1] = 3

	y = x

	y[1] = 7

	fmt.Println(x[1])

}
