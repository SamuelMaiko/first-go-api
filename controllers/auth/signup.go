package auth

import (
	"encoding/json"
	"net/http"
	"firstAPI/gateways/auth"   // For the gateway to interact with the database
	"firstAPI/models"          // Assuming user-related models are in this package
	"firstAPI/utils"           // For sending responses
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into RegisterRequest struct
	var registerReq models.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&registerReq)
	if err != nil {
		response := map[string]string{
			"error": "Invalid request body",
		}
		utils.Response(w, response, http.StatusBadRequest)
		return
	}

	// Validate that all required fields are provided
	if registerReq.FirstName == "" || registerReq.FirstName == "" || registerReq.Password == "" || registerReq.Email == "" {
		response := map[string]string{
			"error": "First name, lastname , email and password are required",
		}
		utils.Response(w, response, http.StatusBadRequest)
		return
	}

	// Check if the user already exists
	existingUser, err := auth.GetUserByEmail(registerReq.Email)
	if err == nil && existingUser.Email != "" {
		// If user already exists, return error
		response := map[string]string{
			"error": "Email already taken",
		}
		utils.Response(w, response, http.StatusBadRequest)
		return
	}

	
	// Hash the password using the existing utility function
	hashedPassword, err := utils.HashPassword(registerReq.Password)
	if err != nil {
		response := map[string]string{
			"error": "Failed to hash password",
		}
		utils.Response(w, response, http.StatusInternalServerError)
		return
	}

	// Create the user in the database
	user := models.User{
		FirstName: registerReq.FirstName,
		LastName: registerReq.LastName,
		Email: registerReq.Email,
		Password: string(hashedPassword),
	}

	err = auth.CreateUser(user)
	if err != nil {
		// If the error is "email already taken"
		if err.Error() == "email already taken" {
			response := map[string]string{
				"error": "Email is already taken",
			}
			utils.Response(w, response, http.StatusBadRequest)
			return
		}
	
		// Generic failure message
		response := map[string]string{
			"error": "Failed to create user",
		}
		utils.Response(w, response, http.StatusInternalServerError)
		return
	}

	// Respond with success message
	response := map[string]string{
		"message": "User registered successfully",
	}
	utils.Response(w, response, http.StatusOK)
}
