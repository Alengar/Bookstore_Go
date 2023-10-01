package utils

import (
	"bytes"
	"net/http"
	"testing"
)

type TestStruct struct {
	Key1 string `json:"key1"`
	Key2 int    `json:"key2"`
}

func TestParseBody(t *testing.T) {
	// Creating a test JSON payload.
	payload := `{"key1": "value1", "key2": 42}`

	// Creating a new HTTP request with the JSON payload.
	req, err := http.NewRequest("POST", "/test", bytes.NewBufferString(payload))
	if err != nil {
		t.Fatalf("Failed to create a test request: %v", err)
	}

	// Calling the ParseBody function to parse the request body.
	var parsedBody TestStruct
	ParseBody(req, &parsedBody)

	// Checking if the parsed body matches the expected values.
	expected := TestStruct{Key1: "value1", Key2: 42}
	if parsedBody != expected {
		t.Errorf("Expected parsed body %+v, but got %+v", expected, parsedBody)
	}
}
