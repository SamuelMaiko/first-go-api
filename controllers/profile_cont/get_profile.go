package profile_cont

import (
    "net/http"
    "firstAPI/db"
    "encoding/json"
    "log"
)

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(int) // Assuming you have middleware to get the user ID

    var profile struct {
        ID         int     `json:"id"`
        Image      string  `json:"image"`
        Bio        string  `json:"bio"`
        County     string  `json:"county"`
        PhoneNumber string `json:"phone_number"`
        Email string `json:"email"`
        FirstName  string  `json:"first_name"`
        LastName   string  `json:"last_name"`
    }

    query := `
        SELECT p.id, p.image, p.bio, p.county, p.phone_number, u.email, u.first_name, u.last_name
        FROM profiles p
        JOIN users u ON u.id = p.user_id
        WHERE p.user_id = $1
    `

    err := db.DB.QueryRow(query, userID).Scan(&profile.ID, &profile.Image, &profile.Bio, &profile.County, &profile.PhoneNumber,&profile.Email, &profile.FirstName, &profile.LastName)
    if err != nil {
        log.Println("Error retrieving profile:", err)
        http.Error(w, "Profile not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(profile)
}
