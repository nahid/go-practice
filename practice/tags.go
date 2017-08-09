package main

import (
	"fmt"
	"reflect"
)

func main() {
	user := User{Name: "Nahid Bin Azhar"}

	fmt.Println(user.Name)

	t := reflect.TypeOf(user)

	// Get the type and kind of our user variable
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())
	field := t.Field(0)

	fmt.Println("tag:", field.Tag.Get("tag"))

}

type User struct {
	Name string `tag:"name"`
}
