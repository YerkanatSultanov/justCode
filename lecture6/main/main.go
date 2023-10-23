package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"justCode/lecture6"
)

func main() {
	jsonFile, err := ioutil.ReadFile("./lecture6/json.json")
	if err != nil {
		fmt.Printf("Error reading file %s", err)
	}

	jsonString := string(jsonFile)
	BookInfoFromJson(jsonString)
}

func BookInfoFromJson(jsonString string) {
	var books []lecture6.Book
	if err := json.Unmarshal([]byte(jsonString), &books); err != nil {
		fmt.Printf("Error unmarshall: %v", err)
	}

	for _, book := range books {
		fmt.Printf("Title: %s\nAuthor: %s\nGenre: %s\nSubGenre: %s\n\n", book.Title, book.Author, book.Genre.GenreName, book.Genre.Subgenre)
	}
}
