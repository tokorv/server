package main

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
)


type Book struct {
	Title string `json:"title"`
	Filename string `json:"filename"`
	Path string `json:"path"`
}

func FetchBooks() ([]Book, error) {
	var books = make([]Book, 0)
	file, err := ioutil.ReadFile("data/book_data.json")
	if err != nil {
		return books, err
	}
	err = json.Unmarshal([]byte(file), &books)
	return books, err
}
