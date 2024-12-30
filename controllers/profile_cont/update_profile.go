package profile_cont

import (
	"encoding/json"
	"net/http"
	"firstAPI/db"
	"firstAPI/utils"
)

type Profile struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Bio         string `json:"bio"`
	County      string `json:"county"`
	PhoneNumber string `json:"phone_number"`
	UserID      int64  `json:"user_id"`
}

// UpdateProfileHandler handles updating the user's profile and user details
func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get user from context (assuming the authentication middleware added the user context)
	// claims := r.Context().Value("user").(*utils.JWTClaims)
	// userID := claims.ID // You can replace it with the user ID from the JWT claims

	userID := r.Context().Value("userID").(int)

	// Parse incoming JSON body
	var profile Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		utils.Response(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}

	// Start a transaction to update both the users table and the profiles table
	tx, err := db.DB.Begin()
	if err != nil {
		utils.Response(w, map[string]string{"error": "Failed to start transaction"}, http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Update user details in the users table
	userUpdateQuery := `
		UPDATE users
		SET first_name = $1, last_name = $2
		WHERE id = $3
		RETURNING id, first_name, last_name
	`
	var updatedUser struct {
		ID        int64  `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	err = tx.QueryRow(userUpdateQuery, profile.FirstName, profile.LastName, userID).Scan(&updatedUser.ID, &updatedUser.FirstName, &updatedUser.LastName)
	if err != nil {
		utils.Response(w, map[string]string{"error": "Failed to update user details"}, http.StatusInternalServerError)
		return
	}

	// Update profile details in the profiles table
	profileUpdateQuery := `
		UPDATE profiles
		SET bio = $1, county = $2, phone_number = $3, updated_at = now()
		WHERE user_id = $4
		RETURNING id, bio, county, phone_number, user_id
	`
	var updatedProfile Profile
	err = tx.QueryRow(profileUpdateQuery, profile.Bio, profile.County, profile.PhoneNumber, userID).Scan(&updatedProfile.ID, &updatedProfile.Bio, &updatedProfile.County, &updatedProfile.PhoneNumber, &updatedProfile.UserID)
	if err != nil {
		utils.Response(w, map[string]string{"error": "Failed to update profile details"}, http.StatusInternalServerError)
		return
	}

	// Commit the transaction if both updates are successful
	err = tx.Commit()
	if err != nil {
		utils.Response(w, map[string]string{"error": "Failed to commit transaction"}, http.StatusInternalServerError)
		return
	}

	// Return the updated user and profile details as response
	utils.Response(w, map[string]interface{}{
		"id":  updatedProfile.ID,
		"first_name":  updatedUser.FirstName,
		"last_name":   updatedUser.LastName,
		"bio":         updatedProfile.Bio,
		"county":      updatedProfile.County,
		"phone_number": updatedProfile.PhoneNumber,
	}, http.StatusOK)
}
