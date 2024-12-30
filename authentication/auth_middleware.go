package authentication

import (
	"context"
	"net/http"
	"fmt"

	"firstAPI/utils"
)

// AuthMiddleware is used to protect routes with JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			response := map[string]string{
				"error": "Authorization header missing",
			}
			utils.Response(w, response, http.StatusUnauthorized)
			return
		}

		// Remove the "Bearer " prefix from token string if present
		if len(tokenStr) > 6 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:]
		}

		// Parse the token and get the claims
		claims, err := ParseJWT(tokenStr)
		if err != nil {
			response := map[string]string{
				"error": "Invalid or expired token",
			}
			utils.Response(w, response, http.StatusUnauthorized)
			return
		}

		// Assuming the userID is stored in claims as "user_id"
		userID := claims.UserID // Adjust based on how you store user info in JWT

		// Add userID to context
		ctx := context.WithValue(r.Context(), "userID", userID)

		// Log authenticated user
		fmt.Printf("Authenticated user: %s\n", claims.FirstName, claims.UserID)

		// Continue to the next handler with updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
