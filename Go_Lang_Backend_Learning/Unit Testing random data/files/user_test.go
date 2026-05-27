package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"SQL/util"

	"github.com/stretchr/testify/require"
)

// ─────────────────────────────────────────
//  Helper — creates one random user in DB
// ─────────────────────────────────────────
func createRandomUser(t *testing.T) User {
	name  := util.RandomName()
	email := util.RandomEmail()

	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at`

	var user User
	err := testDB.QueryRowContext(context.Background(), query, name, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, name, user.Name)
	require.Equal(t, email, user.Email)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

// ─────────────────────────────────────────
//  CREATE
// ─────────────────────────────────────────
func TestCreateUser(t *testing.T) {
	user := createRandomUser(t)

	require.NotEmpty(t, user)
	require.NotZero(t, user.ID)
	require.NotEmpty(t, user.Name)
	require.NotEmpty(t, user.Email)
	require.WithinDuration(t, time.Now(), user.CreatedAt, 5*time.Second)
}

func TestCreateMultipleUsers(t *testing.T) {
	// Create 5 users with random data
	users := make([]User, 5)
	for i := 0; i < 5; i++ {
		users[i] = createRandomUser(t)
	}

	// Make sure each user has a unique ID and email
	ids    := make(map[int32]bool)
	emails := make(map[string]bool)

	for _, u := range users {
		require.False(t, ids[u.ID], "duplicate user ID found")
		require.False(t, emails[u.Email], "duplicate email found")
		ids[u.ID]       = true
		emails[u.Email] = true
	}
}

// ─────────────────────────────────────────
//  READ - Get by ID
// ─────────────────────────────────────────
func TestGetUser(t *testing.T) {
	// Create a user first
	created := createRandomUser(t)

	// Read it back
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`

	var fetched User
	err := testDB.QueryRowContext(context.Background(), query, created.ID).
		Scan(&fetched.ID, &fetched.Name, &fetched.Email, &fetched.CreatedAt)

	require.NoError(t, err)
	require.NotEmpty(t, fetched)

	// Check all fields match exactly
	require.Equal(t, created.ID, fetched.ID)
	require.Equal(t, created.Name, fetched.Name)
	require.Equal(t, created.Email, fetched.Email)
	require.WithinDuration(t, created.CreatedAt, fetched.CreatedAt, time.Second)
}

func TestGetUserNotFound(t *testing.T) {
	// Use a very large ID that doesn't exist
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`

	var user User
	err := testDB.QueryRowContext(context.Background(), query, int32(999999999)).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user)
}

// ─────────────────────────────────────────
//  READ - List all
// ─────────────────────────────────────────
func TestListUsers(t *testing.T) {
	// Create 3 users
	for i := 0; i < 3; i++ {
		createRandomUser(t)
	}

	query := `SELECT id, name, email, created_at FROM users ORDER BY id`

	rows, err := testDB.QueryContext(context.Background(), query)
	require.NoError(t, err)
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
		require.NoError(t, err)
		require.NotEmpty(t, u)
		users = append(users, u)
	}

	require.NoError(t, rows.Err())
	require.GreaterOrEqual(t, len(users), 3)
}

// ─────────────────────────────────────────
//  UPDATE - name
// ─────────────────────────────────────────
func TestUpdateUserName(t *testing.T) {
	// Create user with random name
	original := createRandomUser(t)

	// Update to a new random name
	newName := util.RandomName()
	query   := `UPDATE users SET name = $1 WHERE id = $2 RETURNING id, name, email, created_at`

	var updated User
	err := testDB.QueryRowContext(context.Background(), query, newName, original.ID).
		Scan(&updated.ID, &updated.Name, &updated.Email, &updated.CreatedAt)

	require.NoError(t, err)
	require.NotEmpty(t, updated)

	// Name should be changed
	require.Equal(t, original.ID, updated.ID)
	require.NotEqual(t, original.Name, updated.Name)
	require.Equal(t, newName, updated.Name)

	// Email should stay the same
	require.Equal(t, original.Email, updated.Email)
}

// ─────────────────────────────────────────
//  DELETE
// ─────────────────────────────────────────
func TestDeleteUser(t *testing.T) {
	// Create a user
	user := createRandomUser(t)

	// Delete it
	deleteQuery := `DELETE FROM users WHERE id = $1`
	_, err := testDB.ExecContext(context.Background(), deleteQuery, user.ID)
	require.NoError(t, err)

	// Try to read it back — should be gone
	selectQuery := `SELECT id, name, email, created_at FROM users WHERE id = $1`

	var deleted User
	err = testDB.QueryRowContext(context.Background(), selectQuery, user.ID).
		Scan(&deleted.ID, &deleted.Name, &deleted.Email, &deleted.CreatedAt)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

// ─────────────────────────────────────────
//  DUPLICATE EMAIL (constraint test)
// ─────────────────────────────────────────
func TestCreateDuplicateEmail(t *testing.T) {
	user := createRandomUser(t)

	// Try inserting same email again
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at`

	var duplicate User
	err := testDB.QueryRowContext(context.Background(), query, util.RandomName(), user.Email).
		Scan(&duplicate.ID, &duplicate.Name, &duplicate.Email, &duplicate.CreatedAt)

	// Must fail with unique constraint error
	require.Error(t, err)
}
