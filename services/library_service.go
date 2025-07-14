package services

import(
	"library-management-system/models"
)

type Library struct{
	BookCount int
	Books map[int]models.Book
	Members map[int]models.Member
}

type LibraryManager interface{
	AddBook(book models.Book)
	RemoveBook(bookID int) 
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

func (lib *Library) AddBook(book models.Book) string {
	book.ID = lib.BookCount
	lib.Books[book.ID] = book
	lib.BookCount +=  1

	return "Book created successfully!"
}

func (lib *Library) RemoveBook(bookID int) string {
	_, ok := lib.Books[bookID]
	if ok {
		delete(lib.Books, bookID)
		return "Book deleted successfully!"
	} else {
		return "The book doesn't exist!"
	}
}
func (lib *Library) BorrowBook(bookID int, memberID int) string {
	_, ok := lib.Books[bookID]
	if ok {
		book := lib.Books[bookID]
		if book.Status == "Available" {
			_, exists := lib.Members[memberID]
			if !exists{
				lib.Members[memberID] = models.Member{
					ID: memberID,
				}
			}
			member := lib.Members[memberID]
			member.BorrowedBooks = append(member.BorrowedBooks, book)
			book.Status = "Borrowed"
			lib.Books[bookID] = book
			lib.Members[memberID] = member
			return "Book borrowed Successfully!"
		}else{
			return "Book is not available!"
		}
	} else {
		return "The book doesn't exist inside the library!"
	}
}

func (lib *Library) ReturnBook(bookID int, memberID int) string {
	_, ok := lib.Books[bookID]
	if ok {
		book := lib.Books[bookID]

		_, ok := lib.Members[memberID]
		if ok {
			member := lib.Members[memberID]
			for i, b := range member.BorrowedBooks {
				if b.ID == book.ID{
					member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
					book.Status = "Available"
					lib.Books[bookID] = book
					lib.Members[memberID] = member
					return "Book returned successfully!"
				}
			}
			return "The book was not borrowed!"
		} else {
			return "Member not found!"
		}
	} else {
		return "Book not found!"
	}
}

func (lib *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book

	for _, book := range lib.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (lib *Library) ListBorrowedBooks(memberID int) []models.Book {
	_, ok := lib.Members[memberID]
	if ok {
		member := lib.Members[memberID]
		return member.BorrowedBooks
	} else {
		return nil
	}
}
