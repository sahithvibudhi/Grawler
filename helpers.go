package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

type Config struct {
	urls []string
}

func Init() {

	docs := make(map[string]*goquery.Document)
	docs["http://naasongs.com"] = getDoc("http://telugua2z.in")
	docs["http://naasongs.com"].Find("a").Each(func(indx int, s *goquery.Selection){
		fmt.Println(s.Text())
	})

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