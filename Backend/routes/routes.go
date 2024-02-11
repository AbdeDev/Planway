package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouter configure les routes pour le routeur Gorilla Mux
func SetupRouter() *mux.Router {
  r := mux.NewRouter()

  // Routes de test
  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Bienvenue sur l'API de Planway"))
  })

  // Routes pour les salons de coiffure
  r.HandleFunc("/salons", GetAllSalons).Methods("GET")
  r.HandleFunc("/salons/{id}", GetSalon).Methods("GET")
  r.HandleFunc("/salons", CreateSalon).Methods("POST")
  r.HandleFunc("/salons/{id}", UpdateSalon).Methods("PUT")
  r.HandleFunc("/salons/{id}", DeleteSalon).Methods("DELETE")

  return r
}
