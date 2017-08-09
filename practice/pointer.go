package main

import (
	"fmt"
)

func main() {
	x := 0

	incr(&x)
	incr(&x)
	incr(&x)
	incr(&x)
	incr(&x)

	fmt.Println(x)
}

func incr(y *int) {
	*y++
}
