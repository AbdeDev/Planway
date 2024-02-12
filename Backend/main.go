// main.go
package main

import (
	"backend/db"
	"backend/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialiser la base de données
	db.InitDB()

	// Créer un routeur Gorilla mux
	router := mux.NewRouter()

	// Routes pour les salons
	router.HandleFunc("/api/salons", routes.CreateSalon).Methods("POST")
	router.HandleFunc("/api/salons/{id}", routes.GetSalonByID).Methods("GET")
	router.HandleFunc("/api/salons/{id}", routes.UpdateSalon).Methods("PUT")
	router.HandleFunc("/api/salons/{id}", routes.DeleteSalon).Methods("DELETE")
	router.HandleFunc("/api/salons", routes.GetAllSalons).Methods("GET")

	// ... autres routes pour d'autres fonctionnalités si nécessaire

	// Définir le routeur comme route principale
	http.Handle("/", router)

	// Démarrer le serveur
	http.ListenAndServe(":8080", nil)
}

