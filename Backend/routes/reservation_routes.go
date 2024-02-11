// planway/backend/routes/reservation_routes.go
package routes

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
)

// CreateReservation crée une nouvelle réservation
func CreateReservation(w http.ResponseWriter, r *http.Request) {
	var reservation models.Reservation
	_ = json.NewDecoder(r.Body).Decode(&reservation)

	// Valider les données de réservation
	if reservation.SalonID == "" || reservation.StartTime.IsZero() || reservation.EndTime.IsZero() || reservation.ClientID == "" {
		http.Error(w, "Données de réservation invalides", http.StatusBadRequest)
		return
	}

	// Ajouter la logique pour créer une réservation dans la base de données
	createdReservation, err := db.CreateReservation(reservation)
	if err != nil {
		http.Error(w, "Erreur lors de la création de la réservation", http.StatusInternalServerError)
		return
	}

	// Répondre avec la nouvelle réservation créée
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdReservation)
}
