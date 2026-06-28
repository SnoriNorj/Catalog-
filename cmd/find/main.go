package main

import (
	"books"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: find <BOOK ID>")
		return
	}
	ID := os.Args[1]
	book, ok := books.GetBook(ID)
	if !ok {
		fmt.Println("Sorry, I couldn't find that book in the Catalog")
		return
	}
	fmt.Println(books.BookToString(book))
}

/*
 the os.Args is getting the ID when you run the program
 that is when you type go run ./cmd/find abc, os.Args is making
 a slice  os.Args = []string {"./cmd.find, xyz"}. that is we using the 2nd index for the
 ID.
*/
