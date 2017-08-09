package main

import (
	"fmt"
)

func main() {
	sum(2, 3.4, 5)
	//fmt.Println(sum(1, 2.3, 3, 4, 5.2))
}

func sum(vars ...interface{}) {

	fmt.Println(vars)
}
