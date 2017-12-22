package main

import (
	"net/http"
	"fmt"
)

func get(link string) {

	resp, err := http.Get(link)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var res string = formatResponse(resp)
	fmt.Println(res)
}