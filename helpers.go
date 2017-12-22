package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
)

func Init() {
	get("http://naasongs.com")
}

func LoadConfig(conf string) {

	fmt.Println(conf)

}

func formatResponse(resp *http.Response) string {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
	
}

func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}