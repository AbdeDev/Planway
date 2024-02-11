package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/db"
	"backend/routes" // Utilise le chemin relatif sans le nom du module
)

func main() {
	db.InitDB()
	
	r := routes.SetupRouter()
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
