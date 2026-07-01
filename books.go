package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     string
}

func (book Book) String() string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title,
		book.Author, book.Copies)
}

type Catalog map[string]Book

func (catalog Catalog) GetAllBooks() []Book {
	return slices.Collect(maps.Values(catalog))
}

func (catalog Catalog) GetBook(ID string) (Book, bool) {
	book, ok := catalog[ID]
	return book, ok
}

func (catalog Catalog) AddBook(book Book) {
	catalog[book.ID] = book
}

func (book *Book) SetCopies(copies int) error {
	if copies < 0 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}
	book.Copies = copies
	return nil
}
func GetCatalog() Catalog {
	return Catalog{
		"abc": {
			Title:  "In The Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
			ID:     "abc",
		},
		"xyz": {
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
			ID:     "xyz",
		},
	}
}

func OpenCatalog(path string) (Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	catalog := Catalog{}
	err = json.NewDecoder(file).Decode(&catalog)
	if err != nil {
		return nil, err
	}
	return catalog, nil
}

/*
When returning a value from a function we need to specify the return type as shown in
the BookToString function.
*/
