// https://tutorialedge.net/golang/go-json-tutorial/

package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	BookSales   int  `json:"bookSales"`
	Age         int  `json:"age"`
	IsDeveloper bool `json:"isDeveloper"`
}

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

func main() {
	author := Author{BookSales: 5, Age: 25, IsDeveloper: true}
	book := Book{Title: "Learning JSON", Author: author}

	// Encode 'struct' to 'json'
	jsonData, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonData))
}
