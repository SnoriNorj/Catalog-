package main

import (
	"books"
	"fmt"
)

func main() {
	book := books.Book{
		Title:  "Engineering in Plain Sight",
		Author: "Grady Hillhouse",
		Copies: 2,
	}
	fmt.Println(books.BookToString(book))
}

/*
Constructing a struct , we can access the values that stored by usign the dot operator ,
remembering that the stuct is a key-value pair. so the dot will access the value of TItle
and so on.

printf will return a formatted string immediately to standardout , while Sprintf will return the string value to used
or to be modified.

we need the import "books" in order to use the functions in books.go and morover, we create
the package in books.go by using the syntax package books.

Moreover, since the struct likes in a different folder , we need to call it using the books.Book
when declaring our book variable.
*/
