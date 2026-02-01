package services

import (
	"errors"
	"sync"
	"task4/models"
	"time"
)

type LibraryManager interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
	ReserveBook(bookID int, memberID int) error
}

type Library struct {
	Books	 map[int]models.Book
	Members	 map[int]models.Member
	mu		 sync.Mutex //? "Lock" for the library data
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
	l.mu.Lock()
	defer l.mu.Unlock()

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
	l.mu.Lock()
	defer l.mu.Unlock()

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
	l.mu.Lock()
	defer l.mu.Unlock()
	
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
	l.mu.Lock()
	defer l.mu.Unlock()

	member, memberExists := l.Members[memberID]

	// Check if member exists
	if !memberExists {
		return nil, errors.New("member not found")
	}

	return member.BorrowedBooks, nil
}

// AddBook adds a book to the library list of books
func (l *Library) AddBook(book models.Book) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, bookExists := l.Books[book.ID]

	// Check if the book already exists in the map
	if bookExists {
		return errors.New("book already exists in the Library")
	}

	// Add book to the map
	l.Books[book.ID] = book
	return nil
}

func (l *Library) ReserveBook(bookID int, memberID int) error {
	l.mu.Lock()

	book, bExists := l.Books[bookID]
	_, mExists := l.Members[memberID]

	if !bExists || !mExists {
		l.mu.Unlock()
		return errors.New("book or member not found")
	}

	if book.Status != "Available" {
		l.mu.Unlock()
		return errors.New("book is not available for reservation")
	}

	// Set status to Reserved
	book.Status = "Reserved"
	l.Books[bookID] = book
	l.mu.Unlock()

	// Asynchronous auto-cancellation
	go func() {
		time.Sleep(5*time.Second)
		l.mu.Lock()
		defer l.mu.Unlock()

		// Re-check if it's still "Reserved"
		currentBook := l.Books[bookID]
		if currentBook.Status == "Reserved" {
			currentBook.Status = "Available"
			l.Books[bookID] = currentBook
		}

	}()

	return nil
}

// RemoveBook removes a book from the library
func (l *Library) RemoveBook(bookID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	
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