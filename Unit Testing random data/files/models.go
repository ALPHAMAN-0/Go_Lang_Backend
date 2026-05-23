package db

import "time"

// User represents a row in the users table
type User struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Phone represents a row in the phones table
type Phone struct {
	ID          int32     `json:"id"`
	UserID      int32     `json:"user_id"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}
