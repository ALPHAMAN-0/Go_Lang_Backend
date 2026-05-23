package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

const connStr = "host=localhost port=5432 user=admin password=secret dbname=mydb sslmode=disable"

func main() {
    // Open connection
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error opening DB:", err)
    }
    defer db.Close()

    // Test the connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Cannot connect to DB:", err)
    }

    fmt.Println("✅ Connected to PostgreSQL successfully!")

    // Create a table
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id   SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        age  INT
    )`)
    if err != nil {
        log.Fatal("Error creating table:", err)
    }
    fmt.Println("✅ Table created!")

    // Insert a row
    _, err = db.Exec(`INSERT INTO users (name, age) VALUES ($1, $2)`, "Alice", 30)
    if err != nil {
        log.Fatal("Error inserting:", err)
    }
    fmt.Println("✅ Row inserted!")

    // Query rows
    rows, err := db.Query(`SELECT id, name, age FROM users`)
    if err != nil {
        log.Fatal("Error querying:", err)
    }
    defer rows.Close()

    fmt.Println("\n--- Users ---")
    for rows.Next() {
        var id, age int
        var name string
        rows.Scan(&id, &name, &age)
        fmt.Printf("ID: %d | Name: %s | Age: %d\n", id, name, age)
    }
}