// planway/backend/main.go
package main

import (
	"backend/db"
	"backend/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialiser la base de données
	db.InitDB()

	// Créer un routeur Gorilla Mux
	router := mux.NewRouter()

	// Routes pour l'authentification
	router.HandleFunc("/register", routes.RegisterUser).Methods("POST")
	router.HandleFunc("/login", routes.LoginUser).Methods("POST")

	// Routes pour les salons
	router.HandleFunc("/salons", routes.CreateSalon).Methods("POST")
	router.HandleFunc("/salons/{id}", routes.GetSalonByID).Methods("GET")
	router.HandleFunc("/salons/{id}", routes.UpdateSalon).Methods("PUT")

	// Route pour les réservations
	router.HandleFunc("/reservations", routes.CreateReservation).Methods("POST")

	// À faire: Ajouter d'autres routes au besoin

	// Démarrer le serveur
	port := 8080
	fmt.Printf("Serveur en cours d'exécution sur le port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
