package auth

import (
	"database/sql"
	"errors"
	"firstAPI/models"
	"firstAPI/db"
)


func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, first_name, last_name FROM users WHERE email = $1 LIMIT 1`

	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
