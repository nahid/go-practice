package main

import (
	"fmt"
)

func main() {
	data := []int{5, 2, 4, 7, 5, 2, 3, 4, 3, 6, 7}

	res :=0
	for i:=0; i<len(data); i++ {
		newres := res
		res ^= data[i]
		fmt.Printf("%d(%b) \n%d(%b)\n--------- \n %d \n\n", newres, newres, data[i], data[i], res)

	}

	//fmt.Println(res)
}
