package main

import (
	"log"

	// "fmt"
	"net/http"

	"firstAPI/routers"
	"firstAPI/db"
)


func main (){

	db.InitDB()
	
	router:=routers.NewRouter()

	log.Println("Starting the server at port 8000")
	err:=http.ListenAndServe(":8000", router)

	if err!=nil{
		log.Fatal("Server failed to start: ", err)
	}

}

// connect cockroach db to dbeaver
// use localhost django backend