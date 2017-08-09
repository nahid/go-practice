package compose

import (
	"file"
	"encoding/json"
	"log"
)

type ComposeFile struct {
	Name string `json:name`
	Apps map[string]App `json:apps`
}

type App struct {
	Run string `json:run`
	DependsOn []string `json:dependsOn`
	Shell string `json:shell`
}

func GetAppConfig() ComposeFile {
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





