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
		w.Write([]byte("Bienvenue sur l'API de réservation de salon de coiffure"))
	})

	return r
}
