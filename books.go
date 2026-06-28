package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     string
}

var Catalog = map[string]Book{
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

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title,
		book.Author, book.Copies)
}

func GetAllBooks() []Book {
	return slices.Collect(maps.Values(Catalog))
}

func GetBook(ID string) (Book, bool) {
	for _, book := range Catalog {
		if book.ID == ID {
			return book, true
		}

	}
	return Book{}, false
}

/*
When returning a value from a function we need to specify the return type as shown in
the BookToString function.
*/
