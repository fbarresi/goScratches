package main

import "fmt"

type Book struct {
   title string
   author string
   subject string
   book_id int
}

func (b Book)String() string {
	return b.title + " - " + b.author
}

func main() {
	var Book1 Book
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	fmt.Println(Book1)
}


