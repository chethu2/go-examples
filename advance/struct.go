package advance

import "fmt"

type Book struct {
	BookName string
	BookId string
	Author string
	Price int
}

type Store struct {
	RackNumber int
	Books Book
}

func StructLibrary () {

	book := Book{Author:"J K Rowling",BookId:"hp-1",BookName:" Harry Potter and the Sorcerer's Stone",Price:500}
	storeEntry := Store{Books:book,RackNumber:12}

	fmt.Println("book name :", book.BookName)
	fmt.Println("book Id :",book.BookId)
	fmt.Println("book price: ", book.Price)
	fmt.Println("book author :", book.Author)
	fmt.Println("rack number ",storeEntry.RackNumber)
	fmt.Println("full book details : ",storeEntry.Books)
}
