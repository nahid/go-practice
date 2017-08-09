package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func main() {
	//var name string
	student := map[string]string{"name": "Nahid", "age": "28"}

	data := make(map[string]user)

	student["company"] = "Pathao"

	// name = "Nahid"

	data["nahid"] = user{"Nahid Bin Azhar", 33}

	fmt.Println(student["company"])
}
