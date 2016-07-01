package main

import (
  "fmt"
)

type Book struct {
  title string
  author string
  subject string
  id int
}

func books() Book{
  var book1 Book;

  book1.title = "The glorious adventures of David Skrenta";
  book1.author = "David Skrenta";
  book1.subject = "Adventure";
  book1.id = 1;

  return book1;
}

func main() {
  var hello string = "Hello";
  var balance = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};

  for a := 0; a < 10; a++ {
    fmt.Println("Balance is:", balance[a]);
  }

  fmt.Println("Hello, world!");
  fmt.Println(hello);

  if(hello == "hello") {
    fmt.Println("hello is hello");
  } else {
    fmt.Println("hello is not hello");
  }

  var book Book = books();
  fmt.Println(book.title);
}
