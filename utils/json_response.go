package utils

import (
	"encoding/json"
	"net/http"
)


func Response(w http.ResponseWriter, data interface{}, status int) {
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, private")
w.Header().Set("Expires", "-1")
    w.Header().Set("Content-Type", "application/json")
    // w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

    // Set the status code after headers
    w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
    // Now write the actual JSON response
    // err := 
    // if err != nil {
    //     // If encoding fails, log the error (optional)
    //     fmt.Println("Failed to encode JSON response:", err)
    // }
}