package compose

import (
	"file"
	"encoding/json"
	"log"
)

type ComposeFile struct {
	Name string `json:"name"`
	Os string `json:"os"`
	Apps map[string]App `json:"apps"`
}

type App struct {
	Run string `json:"run"`
	DependsOn []string `json:"dependsOn"`
	Shell string `json:"shell"`
	After After `json:"after"`
}

type After struct {
	Copy string `json:"copy"`
	Env map[string]string `json:"env"`
	Shell string `json:"shell"`
}

func GetAppCompose() ComposeFile {
	var composeFile ComposeFile
	file.Path = "."
	jsonContents, err := file.GetContentsByte("appiefile.json")

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(jsonContents, &composeFile); err != nil {
		panic(err)
	}


	return composeFile

}





