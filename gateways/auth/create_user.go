package auth

import (
	"log"
	"errors"
	"github.com/lib/pq"
	"firstAPI/models"
	"firstAPI/db"  // Assuming this package is responsible for your DB connection
)

// CreateUser creates a new user and profile in the database
func CreateUser(user models.User) error {
	// Ensure the user data is valid
	if user.FirstName == "" || user.LastName == "" || user.Password == "" {
		return errors.New("invalid user data")
	}

	// Start a transaction
	tx, err := db.DB.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return errors.New("failed to start transaction")
	}
	defer tx.Rollback()  // Ensure the transaction is rolled back in case of an error

	// Create the SQL query to insert the new user into the users table
	userQuery := `INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id`
	var userID int64
	err = tx.QueryRow(userQuery, user.FirstName, user.LastName, user.Email, user.Password).Scan(&userID)
	if err != nil {
		// Check if the error is a unique violation for the email field
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			// Unique constraint violation error (email already exists)
			return errors.New("email already taken")
		}

		// Log other errors
		log.Println("Error inserting user into the database:", err)
		return errors.New("failed to create user")
	}

	// Create the SQL query to insert a new profile for the user
	profileQuery := `INSERT INTO profiles (user_id, bio, county, phone_number) VALUES ($1, '', 'Nairobi', '')`
	_, err = tx.Exec(profileQuery, userID)
	if err != nil {
		log.Println("Error inserting profile into the database:", err)
		return errors.New("failed to create user profile")
	}

	// Commit the transaction if both user and profile were successfully created
	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction:", err)
		return errors.New("failed to commit transaction")
	}

	return nil
}
