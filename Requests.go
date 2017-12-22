package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

func getDoc(link string) *goquery.Document {

	doc, err := goquery.NewDocument(link)

	if err != nil {
		log.Fatal(err)
	}

	return doc

}