package controllers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task3/models"
	"task3/services"
)

type LibraryController struct {
	Service services.LibraryManager
	Scanner *bufio.Scanner
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{
		Service: service,
		Scanner: bufio.NewScanner(os.Stdin),
	}
}

// Run starts the main loop of the application
func (c *LibraryController) Run() {
	for {
		fmt.Println("\n--- Library Management System ---")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Member's Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Select an option: ")

		c.Scanner.Scan()
		input := c.Scanner.Text()

		switch input {
		case "1":
			c.handleAddBook()
		case "2":
			c.handleRemoveBook()
		case "3":
			c.handleBorrowBook()
		case "4":
			c.handleReturnBook()
		case "5":
			c.handleListAvailable()
		case "6":
			c.handleListBorrowed()
		case "7":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

//* --- Utility to read input and trim whitespace
func (c *LibraryController) getInput() string {
	c.Scanner.Scan()
	return strings.TrimSpace(c.Scanner.Text())
}

func (c *LibraryController) getIntInput() (int, error) {
	input := c.getInput()
	intInput, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("Input should be an Integer")
	}
	return intInput, nil
}

//* --- Helper Handler Methods ---
func (c *LibraryController) handleAddBook() {
	fmt.Print("Enter Book ID (int): ")
	id, inputErr := c.getIntInput()

	if inputErr != nil {
		fmt.Println(inputErr)
		return
	}

	fmt.Print("Enter Title: ")
	title := c.getInput()

	fmt.Print("Enter Author: ")
	author := c.getInput()

	newBook := models.Book {
		ID: id,
		Title: title,
		Author: author,
		Status: "Available",
	}

	err := c.Service.AddBook(newBook)
	if err == nil  {
		fmt.Println("Book added successfully!")
		return
	} else {
		fmt.Println(err)
	}
}

func (c *LibraryController) handleRemoveBook() {
	fmt.Print("Enter Book ID to remove: ")
	id, inputErr := c.getIntInput()

	if inputErr != nil {
		fmt.Println(inputErr)
		return
	}

	err := c.Service.RemoveBook(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Book removed successfully.")
	}
}

func (c *LibraryController) handleBorrowBook() {
	fmt.Print("Enter Book ID: ")
	bookID, bookError := c.getIntInput()
	if bookError != nil {
		fmt.Println(bookError)
		return
	}
	

	fmt.Print("Enter Member ID: ")
	memberID, memberError := c.getIntInput()
	if memberError != nil {
		fmt.Println(memberError)
		return
	}
	

	err := c.Service.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
} 

func (c *LibraryController) handleReturnBook() {
	fmt.Print("Enter Book ID: ")
	bookID, bookError := c.getIntInput()
	if bookError != nil {
		fmt.Println(bookError)
		return
	}
	

	fmt.Print("Enter Member ID: ")
	memberID, memberError := c.getIntInput()
	if memberError != nil {
		fmt.Println(memberError)
		return
	}

	err := c.Service.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func (c *LibraryController) handleListAvailable() {
	books := c.Service.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No books available")
		return
	}
	fmt.Println("\nAvailable Books:")
	for _, book := range books {
		fmt.Printf("[%d] %s by %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) handleListBorrowed() {
	fmt.Print("Enter Member ID: ")
	memberID, memberError := c.getIntInput()
	if memberError != nil {
		fmt.Println(memberError)
		return
	}

	books, err := c.Service.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("This member has no borrowed books.")
		return
	}

	fmt.Printf("\nBooks borrowed by Member %d:\n", memberID)
	for _, b := range books {
		fmt.Printf("[%d] %s\n", b.ID, b.Title)
	}
}