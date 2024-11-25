package structs

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func ShowBooks(book Book) {
	fmt.Printf("Title; %s \nAuthor:%s \nPages:%d\n ", book.Title, book.Author, book.Pages)
}
