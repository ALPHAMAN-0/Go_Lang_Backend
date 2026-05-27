package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://admin:secret@localhost:5432/myapp?sslmode=disable"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	if err = testDB.Ping(); err != nil {
		log.Fatal("Cannot ping database:", err)
	}

	log.Println("✅ Connected to test database")

	os.Exit(m.Run())
}
