package auth

import (
	// "fmt"
	"encoding/json"
	"net/http"
	// "strings"

	"firstAPI/utils"
	"firstAPI/authentication"
	"firstAPI/models"
	"firstAPI/gateways/auth"
)


func LoginHandler(w http.ResponseWriter, r *http.Request){
	// Parse the request body into LoginRequest struct
	var loginReq models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)

	if err != nil {
		response:=map[string] string{
			"error": "Email and password are required",
		}
		utils.Response(w, response, http.StatusBadRequest)
		return
	}

	user, err := auth.GetUserByEmail(loginReq.Email)
	if err != nil {
		response := map[string]string{"error": "User not found"}
		utils.Response(w, response, http.StatusUnauthorized)
		return
	}

    // Check if password matches
	if !utils.CheckPasswordHash(loginReq.Password, user.Password) {
		response := map[string]string{"error": "Invalid credentials"}
		utils.Response(w, response, http.StatusUnauthorized)
		return
	}

	// Generate JWT token (assuming a util function for generating JWT)
	token, err := authentication.GenerateJWT(user)
	if err != nil {
		response := map[string]string{"error": "Failed to generate token"}
		utils.Response(w, response, http.StatusInternalServerError)
		return
	}

	// Successful login response
	response := map[string]interface{}{
		"message": "Login successful",
		"user":map[string]string{"first_name":user.FirstName, "last_name":user.LastName},
		"token":   token,
	}
	utils.Response(w, response, http.StatusOK)
}


