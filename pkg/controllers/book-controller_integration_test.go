package controllers_test

import (
	"bookstore/pkg/controllers"
	"bookstore/pkg/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *mux.Router {
	router := mux.NewRouter()

	// Attach the controller functions to the router
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PATCH")

	return router
}

func TestGetBookById(t *testing.T) {
	// Create a test HTTP server
	router := setupTestRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	bookID := "1"

	// Send a GET request to /books/{bookId}
	response, err := http.Get(server.URL + "/books/" + bookID)
	assert.NoError(t, err)
	defer response.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Decode the response body
	var book models.Book
	decoder := json.NewDecoder(response.Body)
	assert.NoError(t, decoder.Decode(&book))

}

func TestCreateBook(t *testing.T) {
	router := setupTestRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	// Define the book data for creation
	bookData := models.Book{
		Name:        "Test Book",
		Author:      "Test Author",
		Publication: "Test Publisher",
	}

	// Convert bookData to JSON
	bookJSON, err := json.Marshal(bookData)
	assert.NoError(t, err)

	// Send a POST request to /books to create a new book
	response, err := http.Post(server.URL+"/books", "application/json", bytes.NewReader(bookJSON))
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var createdBook models.Book
	decoder := json.NewDecoder(response.Body)
	assert.NoError(t, decoder.Decode(&createdBook))

}

func TestDeleteBook(t *testing.T) {
	router := setupTestRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	bookID := "1"

	// Send a DELETE request to /books/{bookId}
	request, err := http.NewRequest("DELETE", server.URL+"/books/"+bookID, nil)
	assert.NoError(t, err)

	client := &http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)
	defer response.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

}

func TestUpdateBook(t *testing.T) {
	router := setupTestRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	bookID := "1"

	updateData := models.Book{
		Name:        "Updated Book Name",
		Author:      "Updated Author",
		Publication: "Updated Publisher",
	}

	updateJSON, err := json.Marshal(updateData)
	assert.NoError(t, err)

	// Send a PATCH request to /books/{bookId} to update the book
	request, err := http.NewRequest("PATCH", server.URL+"/books/"+bookID, bytes.NewReader(updateJSON))
	assert.NoError(t, err)

	client := &http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)
	defer response.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Decode the response body
	var updatedBook models.Book
	decoder := json.NewDecoder(response.Body)
	assert.NoError(t, decoder.Decode(&updatedBook))

}
