package routers

import (
    "github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"

	"firstAPI/authentication"
	"firstAPI/controllers"
	"firstAPI/controllers/profile_cont"
	"firstAPI/controllers/auth"
)

func NewRouter() http.Handler{
	router :=mux.NewRouter()
	// subrouter for protected routes
	protected := router.NewRoute().Subrouter()
	protected.Use(authentication.AuthMiddleware)

	
	router.HandleFunc("/auth/login", auth.LoginHandler).Methods("POST")
	router.HandleFunc("/auth/signup", auth.SignUpHandler).Methods("POST")
	protected.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	protected.HandleFunc("/profile", profile_cont.GetProfileHandler).Methods("GET")
	protected.HandleFunc("/profile/update", profile_cont.UpdateProfileHandler).Methods("PUT")
	// protected.HandleFunc("/users", users.GetUsers).Methods("GET")


	// Allow CORS for all origins, headers, and methods
    corsMiddleware := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), 
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"}), 
    )
	return corsMiddleware(router)
}