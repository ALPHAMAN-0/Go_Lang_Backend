package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    connStr := "host=localhost port=5432 user=admin password=secret dbname=myapp sslmode=disable"

    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Failed to open DB:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    fmt.Println("✅ Connected to database!")
}