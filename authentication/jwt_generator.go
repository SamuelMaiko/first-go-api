package authentication

import (
	"time"
	"github.com/golang-jwt/jwt/v5"


	"firstAPI/models"
)

// Secret key for signing tokens
var jwtKey = []byte("iyu63867hjhji783")

// Claims represents the JWT payload
type Claims struct {
	UserID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT for the given user
func GenerateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
