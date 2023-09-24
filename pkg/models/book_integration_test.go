package models_test

import (
	"bookstore/pkg/config"
	"bookstore/pkg/models"

	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

var testDB *gorm.DB

func setupTestDatabase() *gorm.DB {
	config.TestConnect()
	db := config.GetDB()
	db.AutoMigrate(&models.Book{})
	return db
}

func cleanupTestDatabase(db *gorm.DB) {
	db.DropTableIfExists(&models.Book{})
}

func TestMain(m *testing.M) {
	testDB = setupTestDatabase()
	defer testDB.Close()

	code := m.Run()

	cleanupTestDatabase(testDB)
	testDB.Close()

	os.Exit(code)
}

func TestIntegrationCreateBook(t *testing.T) {
	book := &models.Book{
		Name:        "Integration Test Book",
		Author:      "Integration Test Author",
		Publication: "Integration Test Publisher",
	}

	createdBook := book.CreateBook()

	id := int64(createdBook.ID)

	// Query the database to retrieve the created book
	retrievedBook, _ := models.GetBookById(id)

	// Performing assertions to validate that the book was created and retrieved correctly
	assert.NotNil(t, retrievedBook)
	assert.Equal(t, book.Name, retrievedBook.Name)
	assert.Equal(t, book.Author, retrievedBook.Author)
	assert.Equal(t, book.Publication, retrievedBook.Publication)
}

func TestIntegrationGetAllBooks(t *testing.T) {
	testBooks := []models.Book{
		{
			Name:        "Book 1",
			Author:      "Author 1",
			Publication: "Publisher 1",
		},
		{
			Name:        "Book 2",
			Author:      "Author 2",
			Publication: "Publisher 2",
		},
	}

	for _, book := range testBooks {
		db := config.GetDB()
		db.Create(&book)
		db.Close()
	}

	books := models.GetAllBooks()

	// Performing assertions to validate that the retrieved books match the expected ones
	assert.Len(t, books, len(testBooks))

	for i, expectedBook := range testBooks {
		assert.Equal(t, expectedBook.Name, books[i].Name)
		assert.Equal(t, expectedBook.Author, books[i].Author)
		assert.Equal(t, expectedBook.Publication, books[i].Publication)
	}
}

func TestIntegrationGetBookById(t *testing.T) {
	book := &models.Book{
		Name:        "Integration Test Book",
		Author:      "Integration Test Author",
		Publication: "Integration Test Publisher",
	}

	createdBook := book.CreateBook()

	id := int64(createdBook.ID)

	// Query the database to retrieve the created book
	retrievedBook, _ := models.GetBookById(id)

	// Performing assertions to validate that the book was retrieved correctly
	assert.NotNil(t, retrievedBook)
	assert.Equal(t, book.Name, retrievedBook.Name)
	assert.Equal(t, book.Author, retrievedBook.Author)
	assert.Equal(t, book.Publication, retrievedBook.Publication)
}

func TestIntegrationDeleteBook(t *testing.T) {
	book := &models.Book{
		Name:        "Integration Test Book",
		Author:      "Integration Test Author",
		Publication: "Integration Test Publisher",
	}

	createdBook := book.CreateBook()

	id := int64(createdBook.ID)

	deletedBook := models.DeleteBook(id)

	// Verifying that the book no longer exists in the database
	_, db := models.GetBookById(id)
	defer db.Close()

	assert.Error(t, db.Error) // Expect an error indicating that the record was not found
	assert.Empty(t, deletedBook.Name)
	assert.Empty(t, deletedBook.Author)
	assert.Empty(t, deletedBook.Publication)
}

func TestIntegrationUpdateBookPartial(t *testing.T) {
	book := &models.Book{
		Name:        "Test Book",
		Author:      "Test Author",
		Publication: "Test Publisher",
	}

	createdBook := book.CreateBook()

	updates := map[string]interface{}{
		"Name": "Updated Book Name",
	}

	db := config.GetDB()
	db.Model(&createdBook).Updates(updates)
	defer db.Close()

	retrievedBook, _ := models.GetBookById(createdBook.ID)

	// Perform assertions to validate that the book was partially updated correctly
	assert.NotNil(t, retrievedBook)
	assert.Equal(t, "Updated Book Name", retrievedBook.Name)
	assert.Equal(t, "Test Author", retrievedBook.Author)         // Author remain unchanged
	assert.Equal(t, "Test Publisher", retrievedBook.Publication) // Publication remain unchanged
}

func TestIntegrationUpdateBookNonExistent(t *testing.T) {
	nonExistentID := int64(-1)

	updatedBook := &models.Book{
		ID:          nonExistentID,
		Name:        "Updated Book Name",
		Author:      "Updated Book Author",
		Publication: "Updated Book Publisher",
	}

	updatedBook = updatedBook.UpdateBook()

	assert.Nil(t, updatedBook)
}
