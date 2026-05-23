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
//  Helper — creates one random phone for a user
// ─────────────────────────────────────────
func createRandomPhone(t *testing.T, userID int32) Phone {
	phoneNumber := util.RandomPhone()

	query := `INSERT INTO phones (user_id, phone_number)
	          VALUES ($1, $2)
	          RETURNING id, user_id, phone_number, created_at`

	var phone Phone
	err := testDB.QueryRowContext(context.Background(), query, userID, phoneNumber).
		Scan(&phone.ID, &phone.UserID, &phone.PhoneNumber, &phone.CreatedAt)

	require.NoError(t, err)
	require.NotEmpty(t, phone)
	require.Equal(t, userID, phone.UserID)
	require.Equal(t, phoneNumber, phone.PhoneNumber)
	require.NotZero(t, phone.ID)
	require.NotZero(t, phone.CreatedAt)

	return phone
}

// ─────────────────────────────────────────
//  CREATE
// ─────────────────────────────────────────
func TestCreatePhone(t *testing.T) {
	user  := createRandomUser(t)
	phone := createRandomPhone(t, user.ID)

	require.NotEmpty(t, phone)
	require.NotZero(t, phone.ID)
	require.Equal(t, user.ID, phone.UserID)
	require.NotEmpty(t, phone.PhoneNumber)
	require.WithinDuration(t, time.Now(), phone.CreatedAt, 5*time.Second)
}

func TestCreateMultiplePhones(t *testing.T) {
	// Create 3 users, each with a random phone
	for i := 0; i < 3; i++ {
		user  := createRandomUser(t)
		phone := createRandomPhone(t, user.ID)

		require.Equal(t, user.ID, phone.UserID)
		require.NotEmpty(t, phone.PhoneNumber)
	}
}

// ─────────────────────────────────────────
//  READ - Get phone by user ID
// ─────────────────────────────────────────
func TestGetPhoneByUserID(t *testing.T) {
	user    := createRandomUser(t)
	created := createRandomPhone(t, user.ID)

	query := `SELECT id, user_id, phone_number, created_at FROM phones WHERE user_id = $1`

	var fetched Phone
	err := testDB.QueryRowContext(context.Background(), query, user.ID).
		Scan(&fetched.ID, &fetched.UserID, &fetched.PhoneNumber, &fetched.CreatedAt)

	require.NoError(t, err)
	require.NotEmpty(t, fetched)
	require.Equal(t, created.ID, fetched.ID)
	require.Equal(t, created.UserID, fetched.UserID)
	require.Equal(t, created.PhoneNumber, fetched.PhoneNumber)
	require.WithinDuration(t, created.CreatedAt, fetched.CreatedAt, time.Second)
}

func TestGetPhoneNotFound(t *testing.T) {
	query := `SELECT id, user_id, phone_number, created_at FROM phones WHERE user_id = $1`

	var phone Phone
	err := testDB.QueryRowContext(context.Background(), query, int32(999999999)).
		Scan(&phone.ID, &phone.UserID, &phone.PhoneNumber, &phone.CreatedAt)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, phone)
}

// ─────────────────────────────────────────
//  UPDATE - phone number
// ─────────────────────────────────────────
func TestUpdatePhoneNumber(t *testing.T) {
	user     := createRandomUser(t)
	original := createRandomPhone(t, user.ID)

	newPhone := util.RandomPhone()
	query    := `UPDATE phones SET phone_number = $1 WHERE user_id = $2
	             RETURNING id, user_id, phone_number, created_at`

	var updated Phone
	err := testDB.QueryRowContext(context.Background(), query, newPhone, user.ID).
		Scan(&updated.ID, &updated.UserID, &updated.PhoneNumber, &updated.CreatedAt)

	require.NoError(t, err)
	require.NotEmpty(t, updated)

	// Phone number should be updated
	require.Equal(t, original.ID, updated.ID)
	require.Equal(t, original.UserID, updated.UserID)
	require.NotEqual(t, original.PhoneNumber, updated.PhoneNumber)
	require.Equal(t, newPhone, updated.PhoneNumber)
}

// ─────────────────────────────────────────
//  DELETE - phone
// ─────────────────────────────────────────
func TestDeletePhone(t *testing.T) {
	user  := createRandomUser(t)
	phone := createRandomPhone(t, user.ID)

	// Delete the phone
	deleteQuery := `DELETE FROM phones WHERE user_id = $1`
	_, err := testDB.ExecContext(context.Background(), deleteQuery, user.ID)
	require.NoError(t, err)

	// Verify it's gone
	selectQuery := `SELECT id, user_id, phone_number, created_at FROM phones WHERE user_id = $1`

	var deleted Phone
	err = testDB.QueryRowContext(context.Background(), selectQuery, phone.UserID).
		Scan(&deleted.ID, &deleted.UserID, &deleted.PhoneNumber, &deleted.CreatedAt)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

// ─────────────────────────────────────────
//  CASCADE DELETE — deleting user removes phone
// ─────────────────────────────────────────
func TestCascadeDeleteUserRemovesPhone(t *testing.T) {
	user  := createRandomUser(t)
	phone := createRandomPhone(t, user.ID)

	// Delete the user
	_, err := testDB.ExecContext(context.Background(),
		`DELETE FROM users WHERE id = $1`, user.ID)
	require.NoError(t, err)

	// Phone should also be deleted (ON DELETE CASCADE)
	var deleted Phone
	err = testDB.QueryRowContext(context.Background(),
		`SELECT id, user_id, phone_number, created_at FROM phones WHERE id = $1`,
		phone.ID,
	).Scan(&deleted.ID, &deleted.UserID, &deleted.PhoneNumber, &deleted.CreatedAt)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

// ─────────────────────────────────────────
//  FOREIGN KEY — phone without valid user fails
// ─────────────────────────────────────────
func TestCreatePhoneInvalidUserID(t *testing.T) {
	query := `INSERT INTO phones (user_id, phone_number)
	          VALUES ($1, $2)
	          RETURNING id, user_id, phone_number, created_at`

	var phone Phone
	err := testDB.QueryRowContext(context.Background(), query, int32(999999999), util.RandomPhone()).
		Scan(&phone.ID, &phone.UserID, &phone.PhoneNumber, &phone.CreatedAt)

	// Must fail — foreign key violation
	require.Error(t, err)
}
