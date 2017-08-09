package main

import (
	"fmt"
	//"appie/dispatcher"
	"os/exec"
	"appie/dispatcher"
)


func main() {
	var (
		output []byte
		err error
	)

	if output, err = exec.Command("sh", "-c", "cal").Output(); err != nil {
		fmt.Println(output)
		panic(err)
	}


	fmt.Println(string(output))
	dispatcher.Dispatch()

}
