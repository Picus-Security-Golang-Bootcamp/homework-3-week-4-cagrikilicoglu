package bookList

import (
	"bookApp/pkg/Base/book"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Catalog struct {
	BookList []book.Book
}

func CreateList() (*Catalog, error) {
	BookList := Catalog{}

	err := BookList.AddBooks()
	if err != nil {
		return nil, err
	}
	fmt.Println(BookList)
	return &BookList, nil

}

func (c *Catalog) AddBooks() error {

	var books []book.Book
	// open the books json file
	jsonFile, err := os.Open("books.json")
	if err != nil {
		fmt.Println(err)
		return err

	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	json.Unmarshal(byteValue, &books)
	c.BookList = books
	return nil
}
