package main

import (
	"encoding/json"
	"practice/urlstice/urls"
	"fmt"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type Data struct {
	Data []User `json:"data"`
}

type Users struct {
	User User `json:"data"`
}

func main() {
	var users Users

	urls.SetBaseUrl("https://reqres.in/api/")
	data, _ := urls.Get("users/1")

	//fmt.Printf("%s", data)
	if err := json.Unmarshal(data, &users); err != nil {
		panic(err)
	}

	fmt.Println(users.User.FirstName)


	/*for i, users := range users.User {
		fmt.Println(i, user.FirstName, user.LastName)
	}*/


}