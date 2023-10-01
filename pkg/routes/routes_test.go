package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterBookStoreRoutes(t *testing.T) {
	// Creating a new router using Gorilla Mux.
	router := mux.NewRouter()

	// Registering the routes.
	RegisterBookStoreRoutes(router)

	// Creating a test server to handle HTTP requests.
	testServer := httptest.NewServer(router)
	defer testServer.Close()

	// Defining a set of test cases to check if the routes are registered correctly.
	testCases := []struct {
		method         string
		path           string
		expectedStatus int
	}{
		{"POST", "/book", http.StatusOK},
		{"GET", "/book", http.StatusOK},
		{"GET", "/book/14", http.StatusOK},
		{"PUT", "/book/14", http.StatusOK},
		{"DELETE", "/book/14", http.StatusOK},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest(tc.method, testServer.URL+tc.path, nil)
		if err != nil {
			t.Fatalf("Failed to create a %s request to %s: %v", tc.method, tc.path, err)
		}

		// Sending the HTTP request.
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send %s request to %s: %v", tc.method, tc.path, err)
		}
		defer resp.Body.Close()

		// Checking if the response status code matches the expected status code.
		if resp.StatusCode != tc.expectedStatus {
			t.Errorf("Expected status code %d for %s %s, but got %d", tc.expectedStatus, tc.method, tc.path, resp.StatusCode)
		}
	}
}
