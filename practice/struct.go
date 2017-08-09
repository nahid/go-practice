package main

import (
	"fmt"
)

type Edu struct {
	institue string
	dept     string
	year     int
}

type User struct {
	ID        int
	Name      string
	Age       int
	City      string
	Role      string
	education Edu
}

func (u User) Print() {
	fmt.Println(u.Name, u.Age)
}

func (u *User) setAge(age int) {
	u.Age = age
}

func (u *User) setEdu(inst string, dept string, year int) {
	u.education.institue = inst
	u.education.dept = dept
	u.education.year = year
}

func main() {

	var user User

	// user.setAge(27)

	user.setEdu("BUBT", "CSE", 2013)
	//user.Print()

	fmt.Println(user.education.institue)

	// fmt.Println(user.nahid())
}
