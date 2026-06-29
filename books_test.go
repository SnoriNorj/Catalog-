package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func getTestCatalog() map[string]books.Book {
	return map[string]books.Book{
		"abc": {
			Title:  " In Company of Cheeful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
			ID:     "abc",
		},
		"xyz": {
			Title:  "White Heat",
			Author: "Dominic Sandbook",
			Copies: 2,
			ID:     "xyz",
		},
	}
}

/*
	 func TestBookToString_FormatsBookInAsString(t *testing.T) {
		t.Parallel()
		catalog := getTestCatalog()
		input := books.Book{
			Title:  "Sea Room",
			Author: "Adam Nicolson",
			Copies: 2,
		}
		want := "Sea Room by Adam Nicolson (copies: 2)"
		got := books.BookToString(input)
		if want != got {
			t.Fatalf("want %q, got %q", want, got)
		}
	}
*/
func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := []books.Book{
		{
			Title:  "In The Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
			ID:     "abc",
		},
		{
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
			ID:     "xyz",
		},
	}
	got := books.GetAllBooks(catalog)
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})

	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_FindsBookInCatalogByID(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := books.Book{
		Title:  "In The Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
		ID:     "abc",
	}
	got, ok := books.GetBook(catalog, "abc")
	if !ok {
		t.Fatalf("book not found")
	}
	if want != got {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := books.GetBook(catalog, "nonexistent ID")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AddsBookintoCatalogByAppending(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := books.GetBook(catalog, "123")
	if ok {
		t.Fatal("book already present")
	}
	books.AddBook(catalog, books.Book{
		Title:  "The Prize of all the Oceans",
		Author: "Glyn Williams",
		Copies: 2,
		ID:     "123",
	})
	_, ok = books.GetBook(catalog, "123")
	if !ok {
		t.Fatal("added book not found")
	}
}

/*
 when using t.Fataf() we are not returning the values to the user, we
 are running a test and it produces the test after running it.
*/
