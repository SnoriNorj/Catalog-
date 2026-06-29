package main

import (
	"books"
	"fmt"
)

func main() {
	catalog := books.GetCatalog()
	for _, book := range books.GetAllBooks(catalog) {
		fmt.Println(books.BookToString(book))
	}
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

the for _, book := range books.GetAllBook(){
fmt.Println(books.BooktToString(book))
}
-, is a blank identifier that tells the go complier to ignore the index value of the slice of book, since the range operator gives two values.
since we are not asking for index , instead only need the values of book[0], otherwise I would need to use some other variable i , in order to store the index
and return in fmt.Println() .moreover the index start with zero just like python.
*/
