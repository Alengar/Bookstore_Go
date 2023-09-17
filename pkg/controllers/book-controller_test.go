package controllers

import (
	"bookstore/pkg/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetBook(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/books", GetBook)
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"ID":1,"Name":"Book1","Author":"Author1","Publication":"Publisher1"},{"ID":2,"Name":"Book2","Author":"Author2","Publication":"Publisher2"}]`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetBookById(t *testing.T) {
	// Create a new request with a book ID
	bookID := 1
	req, err := http.NewRequest("GET", "/books/"+strconv.Itoa(bookID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Create a mock router and handle the request using the GetBookById function
	router := mux.NewRouter()
	router.HandleFunc("/books/{bookId:[0-9]+}", GetBookById)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"ID":1,"Name":"Book1","Author":"Author1","Publication":"Publisher1"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateBook(t *testing.T) {
	// Create a JSON request body for creating a book
	book := models.Book{
		Name:        "New Book",
		Author:      "New Author",
		Publication: "New Publisher",
	}
	reqBody, err := json.Marshal(book)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new POST request with the JSON body
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/books", CreateBook)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"ID":3,"Name":"New Book","Author":"New Author","Publication":"New Publisher"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeleteBook(t *testing.T) {
	// Create a new request with a book ID for deletion
	bookID := 2
	req, err := http.NewRequest("DELETE", "/books/"+strconv.Itoa(bookID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/books/{bookId:[0-9]+}", DeleteBook)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"ID":2,"Name":"Book2","Author":"Author2","Publication":"Publisher2"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdateBook(t *testing.T) {
	// Create a JSON request body for updating a book
	updateData := models.Book{
		Name:        "Updated Book",
		Author:      "Updated Author",
		Publication: "Updated Publisher",
	}
	reqBody, err := json.Marshal(updateData)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new PUT request with the JSON body and book ID
	bookID := 1
	req, err := http.NewRequest("PUT", "/books/"+strconv.Itoa(bookID), bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/books/{bookId:[0-9]+}", UpdateBook)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"ID":1,"Name":"Updated Book","Author":"Updated Author","Publication":"Updated Publisher"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
