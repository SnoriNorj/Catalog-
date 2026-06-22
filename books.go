package books

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title,
		book.Author, book.Copies)
}

/*
When returning a value from a function we need to specify the return type as shown in
the BookToString function.
*/
