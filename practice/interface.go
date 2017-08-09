package practice

import (
	"fmt"
)

type Human interface {
	speak()
	walk()
}

type User struct {
}

func (u User) speak() {
	fmt.Println("Manoniyo Specker")
}

func (u User) walk() {
	fmt.Println("Walking")
}

func (u User) sing() {
	fmt.Println("Walking")
}

type Animal struct {
}

func (u Animal) speak() {
	fmt.Println("Animal Manoniyo Specker")
}

func (u Animal) walk() {
	fmt.Println("Animal Walking")
}

func main() {

	u := User{}
	a := Animal{}
	logger(u)
	logger(a)

}

func logger(human Human) {
	human.speak()
	human.walk()
}
