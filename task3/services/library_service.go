package services

import (
	"errors"
	"task3/models"
)

type LibraryManager interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}

type Library struct {
	Books map[int]models.Book
	Members map[int]models.Member
}

// NewLibrary initializes the Library object and 
// return pointer to the object
func NewLibrary() *Library {
	return &Library{
		Books: make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

// BorrowBook updates book stutus and member info if valid
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID]
	member, memberExists := l.Members[memberID]

	if !bookExists {
		return errors.New("book not found")
	}
	if !memberExists {
		return errors.New("member not found") 
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	// Update Book status
	book.Status = "Borrowed"
	l.Books[bookID] = book

	// Add to Members borrowed list
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member // Update user in library member

	return nil
}


func (l *Library) ReturnBook(bookID int, memberID int) error {
	
	book, bookExists := l.Books[bookID]
	member, memberExists := l.Members[memberID]

	if !bookExists {
		return errors.New("book not found")
	}
	if !memberExists {
		return errors.New("member not found") 
	}
	if book.Status == "Available" {
		return errors.New("book was not borrowed")
	}

	// Update book status
	book.Status = "Available"
	l.Books[bookID] = book

	// Update member information
	newBorrowedList := []models.Book{}
	for _, book := range member.BorrowedBooks {
		if book.ID == bookID {
			continue
		}
		newBorrowedList = append(newBorrowedList, book)
	}

	if len(newBorrowedList) == len(member.BorrowedBooks) {
		return  errors.New("this member did not borrow this book")
	}

	member.BorrowedBooks = newBorrowedList

	return nil
}

// ListAvailableBooks list all books with status == 'Available'
func (l*Library) ListAvailableBooks() []models.Book {
	AvailableBooks := []models.Book{}
	for _, book := range l.Books {
		if book.Status == "Available" {
			AvailableBooks = append(AvailableBooks, book)
		} 
	}

	return AvailableBooks
}

// ListBorrowedBooks list all books borrowed by a member
func (l*Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, memberExists := l.Members[memberID]

	// Check if member exists
	if !memberExists {
		return nil, errors.New("member not found")
	}

	return member.BorrowedBooks, nil
}

// AddBook adds a book to the library list of books
func (l *Library) AddBook(book models.Book) error {
	_, bookExists := l.Books[book.ID]

	// Check if the book already exists in the map
	if !bookExists {
		return errors.New("book already exists in the Library")
	}

	// Add book to the map
	l.Books[book.ID] = book
	return nil
}

// RemoveBook removes a book from the library
func (l *Library) RemoveBook(bookID int) error {
	book, bookExists := l.Books[bookID]

	// Check if the book exists
	if !bookExists {
		return errors.New("book not found")
	}
	// Check if the book is not borrowed
	if book.Status == "Borrowed" {
		return errors.New("cannot remove a book that is currently borrowed")
	}

	delete(l.Books, bookID)

	return nil
}