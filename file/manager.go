package file

import (
	"io/ioutil"
)

var Path string

func GetContents(filename string) (string, error) {

	content, err := GetContentsByte(filename)


	if err != nil {
		return "", err
	}

	return string(content), nil
}

func GetContentsByte(filename string) ([]byte, error) {
	fullpath := Path + "/" + filename
	content, err := ioutil.ReadFile(fullpath)


	if err != nil {
		return []byte(""), err
	}

	return content, nil
}
