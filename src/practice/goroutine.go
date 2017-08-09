package main

import (
	"fmt"
)

func hello(x int) {
	fmt.Printf("%s", x)
}

func main() {
	//hello()
	for x := 1; x <= 10; x++ {
		go hello(x)
	}
}
