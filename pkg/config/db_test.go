package config

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	// This test checks if the database connection can be established.
	Connect()

	if db == nil {
		t.Error("Expected a valid database connection, but got nil")
	}

	if err := db.DB().Ping(); err != nil {
		t.Errorf("Failed to ping the database: %v", err)
	}

	db.Close()
}

func TestGetDB(t *testing.T) {
	// This test checks if the GetDB function returns a non-nil database connection.
	Connect()

	retrievedDB := GetDB()
	if retrievedDB == nil {
		t.Error("Expected a valid database connection from GetDB, but got nil")
	}

	db.Close()
}

func TestMain(m *testing.M) {
	// The TestMain function is used to set up and tear down the database connection for the tests.

	Connect()
	defer db.Close()

	exitCode := m.Run()
	os.Exit(exitCode)
}
