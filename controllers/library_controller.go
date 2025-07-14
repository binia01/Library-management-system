package controllers

import (
	"bufio"
	"fmt"
	"library-management-system/models"
	"library-management-system/services"
	"os"
)

var Library = services.Library{
	BookCount: 0,
	Books:     map[int]models.Book{},
	Members:   map[int]models.Member{},
}

var reader = bufio.NewReader(os.Stdin)

func Controller(choice int) any {
	switch choice {
	case 1:
		return addBook()
	case 2:
		return removeBook()
	case 3:
		return borrowBook()
	case 4:
		return returnBook()
	case 5:
		return listAvailableBooks()
	case 6:
		return listBorrowedBooks()
	default:
		return "Invalid choice"
	}
}

func addBook() string {
	var book models.Book
	fmt.Println("Please Enter the Book Information: ")
	fmt.Println()
	fmt.Print("Author: ")
	book.Author, _ = reader.ReadString('\n')
	fmt.Print("Title: ")
	book.Title, _ = reader.ReadString('\n')
	book.Status = "Available"

	return Library.AddBook(book)
}

func removeBook() string {
	var bookId int

	fmt.Print("Enter Book Id: ")
	n, err := fmt.Scan(&bookId)
	if n != 1 || err != nil {
		fmt.Println("Invalid Input. Please Try again")
		return removeBook()
	}

	return Library.RemoveBook(bookId)
}

func borrowBook() string {
	var bookId, memberId int

	fmt.Print("Enter Book Id: ")
	n, err := fmt.Scan(&bookId)
	if n != 1 || err != nil {
		fmt.Println("Invalid Input. Please Try again")
		return borrowBook()
	}

	fmt.Print("Enter Member Id: ")
	n, err = fmt.Scan(&memberId)
	if n != 1 || err != nil {
		fmt.Println("Invalid Input. Please Try again")
		return borrowBook()
	}

	return Library.BorrowBook(bookId, memberId)
}

func returnBook() string {
	var bookId, memberId int

	fmt.Print("Enter Book Id: ")
	n, err := fmt.Scan(&bookId)
	if n != 1 || err != nil {
		fmt.Println("Invalid Input. Please Try again")
		return returnBook()
	}

	fmt.Print("Enter Member Id: ")
	n, err = fmt.Scan(&memberId)
	if n != 1 || err != nil {
		fmt.Println("Invalid Input. Please Try again")
		return returnBook()
	}

	return Library.ReturnBook(bookId, memberId)
}

func listAvailableBooks() []models.Book {
	return Library.ListAvailableBooks()
}

func listBorrowedBooks() []models.Book {
	var memberId int

	fmt.Print("Enter Member Id: ")
	n, err := fmt.Scan(&memberId)
	if n != 1 || err != nil {
		fmt.Println("Invalid Input. Please Try again")
		return listBorrowedBooks()
	}

	return Library.ListBorrowedBooks(memberId)
}
