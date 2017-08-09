package main

import (
	"fmt"
	"appie/compose"

)


func main() {
	var appie compose.ComposeFile

	appie = compose.GetAppConfig()

	apps := appie.Apps

	fmt.Println(appie.Apps["php"].DependsOn[1])

	for key, val := range apps {
		fmt.Println(key, val)
	}

}
