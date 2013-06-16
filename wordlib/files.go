// files
package wordlib

import (
	"io/ioutil"
	"net/http"
	"os"
)

func ReadFile(word string) string {
	re, err := ioutil.ReadFile("e:/words/" + word + ".htm")
	if err != nil {
		panic(err)
	}
	return string(re)
}

func WriteFile(word string, text string) {
	err := ioutil.WriteFile("e:/words/"+word+".htm", []byte(text), 0644)
	if err != nil {
		panic(err)
	}
}
func FileExists(filename string) bool {
	f, err := os.Open(filename)
	if err != nil && os.IsNotExist(err) {
		//fmt.Printf("file not exist!\n")         
		return false
	} else {
		defer f.Close()
	}

	return true

}

func httprequest(s string) string {
	response, _ := http.Get("http://www.iciba.com/" + s)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}
