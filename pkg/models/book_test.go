package models

import (
	"testing"
)

func TestCreateBook(t *testing.T) {
	// Create a new book
	book := &Book{
		Name:        "Test Book",
		Author:      "Test Author",
		Publication: "Test Publisher",
	}

	// Create the book
	createdBook := book.CreateBook()

	// Check if the created book has a valid ID (greater than 0)
	if createdBook.ID <= 0 {
		t.Errorf("CreateBook() did not return a valid ID")
	}

	db.Delete(&createdBook)
}

func TestGetAllBooks(t *testing.T) {
	// Insert some test data into the database
	db.Create(&Book{Name: "Book1", Author: "Author1", Publication: "Publisher1"})
	db.Create(&Book{Name: "Book2", Author: "Author2", Publication: "Publisher2"})

	// Get all books
	books := GetAllBooks()

	// Check if there are at least two books
	if len(books) < 2 {
		t.Errorf("GetAllBooks() did not return expected number of books")
	}

	db.Where("name LIKE ?", "Book%").Delete(&Book{})
}

func TestGetBookById(t *testing.T) {
	// Insert a test book into the database
	createdBook := &Book{Name: "Test Book", Author: "Test Author", Publication: "Test Publisher"}
	db.Create(createdBook)

	// Get the book by its ID (convert uint to int64)
	book, _ := GetBookById(int64(createdBook.ID)) // Convert uint to int64

	// Check if the retrieved book matches the expected data
	if book.Name != createdBook.Name || book.Author != createdBook.Author || book.Publication != createdBook.Publication {
		t.Errorf("GetBookById() did not return the expected book")
	}

	db.Delete(&createdBook)
}
