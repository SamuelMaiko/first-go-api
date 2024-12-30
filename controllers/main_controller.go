package controllers

import (
	"net/http"

	"firstAPI/utils"
)

func HomeHandler (w http.ResponseWriter, r *http.Request){
	response:=map[string]string{
		"message":"Hoorah Maiko! Your first Go API",
	}

	utils.Response(w, response, http.StatusOK)
}