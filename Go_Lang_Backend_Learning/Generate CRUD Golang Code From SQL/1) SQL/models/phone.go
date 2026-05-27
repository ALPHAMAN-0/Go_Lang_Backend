package models

import (
    "fmt"
    "SQL/db"
)

type Phone struct {
    ID          int
    UserID      int
    PhoneNumber string
}

// CREATE
func CreatePhone(userID int, phoneNumber string) (int, error) {
    var id int
    query := `INSERT INTO phones (user_id, phone_number) VALUES ($1, $2) RETURNING id`
    err := db.DB.QueryRow(query, userID, phoneNumber).Scan(&id)
    return id, err
}

// READ - Get phone by user
func GetPhoneByUserID(userID int) (Phone, error) {
    var p Phone
    query := `SELECT id, user_id, phone_number FROM phones WHERE user_id = $1`
    err := db.DB.QueryRow(query, userID).Scan(&p.ID, &p.UserID, &p.PhoneNumber)
    if err != nil {
        return p, fmt.Errorf("phone not found for user %d", userID)
    }
    return p, nil
}

// UPDATE phone number
func UpdatePhoneNumber(userID int, newPhone string) error {
    query := `UPDATE phones SET phone_number = $1 WHERE user_id = $2`
    _, err := db.DB.Exec(query, newPhone, userID)
    return err
}

// DELETE
func DeletePhone(userID int) error {
    _, err := db.DB.Exec(`DELETE FROM phones WHERE user_id = $1`, userID)
    return err
}