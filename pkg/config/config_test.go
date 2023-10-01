package config

import (
	"testing"
)

func TestConnectAndGetDB(t *testing.T) {
	// Performing a smoke test to connect to the database and retrieve the DB object.
	Connect()

	// Checking if the DB object is not nil after connecting.
	if db == nil {
		t.Errorf("Expected DB object to be initialized after connecting, but it's nil")
	}

	// Checking if the DB object can be retrieved using GetDB().
	retrievedDB := GetDB()
	if retrievedDB == nil {
		t.Errorf("Expected to retrieve a non-nil DB object using GetDB(), but it's nil")
	}

	tableName := "ebooks"

	// Checking if the 'ebooks' exists in the database.
	if !retrievedDB.HasTable(tableName) {
		t.Errorf("Expected '%s' to exist in the database, but it doesn't", tableName)
	}

	// Performing a simple query to test the database connection.
	var rowCount int
	if err := retrievedDB.Table(tableName).Count(&rowCount).Error; err != nil {
		t.Errorf("Error executing query on '%s': %v", tableName, err)
	}
}
