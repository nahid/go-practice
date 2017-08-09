package urls

import (
	"io/ioutil"
	"net/http"
)

var url string

func SetBaseUrl(baseUrl string) {
	url = baseUrl

}

func Get(uri string) ([]byte, bool) {
	resp, _ := http.Get(url + uri)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body, true
}
