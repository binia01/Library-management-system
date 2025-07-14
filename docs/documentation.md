# Library Management System Documentation

## Overview

This Library Management System is a simple command-line application written in Go. It allows users to manage books and members in a library, including adding/removing books, borrowing/returning books, and listing available or borrowed books.

## Features

- Add a new book
- Remove an existing book
- Borrow a book
- Return a book
- List all available books
- List all borrowed books by a member

## Project Structure

- `main.go`: Entry point of the application. Handles user interaction and operation selection.
- `controllers/library_controller.go`: Contains controller functions for each operation, handling input/output and calling service methods.
- `services/library_service.go`: Implements the core library logic and operations.
- `models/book.go`: Defines the `Book` struct.
- `models/member.go`: Defines the `Member` struct.
- `docs/documentation.md`: Project documentation.

## Data Models

### Book

```go
type Book struct {
	ID     int
	Title  string
	Author string
	Status string // "Available" or "Borrowed"
}
```

### Member

```go
type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}
```

## Usage

1. Run the application using `go run main.go`.
2. Follow the prompts to select operations and enter required information.
3. The system will display results and allow you to perform further operations or exit.

## How Operations Work

- **Add Book**: Enter book details. The book is added with status "Available".
- **Remove Book**: Enter the book ID. The book is removed if it exists.
- **Borrow Book**: Enter book and member IDs. If the book is available, it is marked as borrowed by the member.
- **Return Book**: Enter book and member IDs. If the member has borrowed the book, it is returned and marked as available.
- **List Available Books**: Displays all books with status "Available".
- **List Borrowed Books**: Enter member ID to see all books borrowed by that member.

## Notes

- Member names are not prompted; only IDs are used.
- All data is stored in memory and will be lost when the program exits.
- Input validation is performed for all operations.

## Example Session

```
Welcome to the Library Management System!
Here are the operations you can perform:
1. Add a new book
2. Remove an existing book
3. Borrow a book
4. Return a book
5. List all available books
6. List all borrowed books by a member
Please select the number of the operation you want to perform: 1
Please Enter the Book Information:
Author: J.K. Rowling
Title: Harry Potter
Book created successfully!
```

