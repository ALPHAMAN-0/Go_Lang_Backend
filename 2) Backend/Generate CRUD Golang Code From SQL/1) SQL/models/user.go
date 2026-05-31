package models

import (
	"SQL/db"
	"database/sql"
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email string
}

// CREATE
func CreateUser(name, email string) (int, error) {
	var id int
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := db.DB.QueryRow(query, name, email).Scan(&id)
	return id, err
}

// READ ALL
func GetAllUsers() ([]User, error) {
	rows, err := db.DB.Query(`SELECT id, name, email FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// READ ONE
func GetUserByID(id int) (User, error) {
	var u User
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err := db.DB.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email)
	if err == sql.ErrNoRows {
		return u, fmt.Errorf("user not found")
	}
	return u, err
}

// UPDATE name
func UpdateUserName(id int, newName string) error {
	query := `UPDATE users SET name = $1 WHERE id = $2`
	_, err := db.DB.Exec(query, newName, id)
	return err
}

// DELETE
func DeleteUser(id int) error {
	_, err := db.DB.Exec(`DELETE FROM users WHERE id = $1`, id)
	return err
}
