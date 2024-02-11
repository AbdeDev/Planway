package routes

import (
	"encoding/json"
	"net/http"

	"backend/db"
	"backend/models"
	"github.com/gorilla/mux"
)

// GetAllSalons renvoie la liste de tous les salons
func GetAllSalons(w http.ResponseWriter, r *http.Request) {
	salons, err := db.GetAllSalons()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(salons)
}

// GetSalon renvoie les détails d'un salon spécifique
func GetSalon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	salonID := vars["id"]

	salon, err := db.GetSalon(salonID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(salon)
}

// CreateSalon crée un nouveau salon
func CreateSalon(w http.ResponseWriter, r *http.Request) {
	var newSalon models.Salon

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newSalon); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdSalon, err := db.CreateSalon(newSalon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdSalon)
}

// UpdateSalon met à jour les détails d'un salon existant
func UpdateSalon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	salonID := vars["id"]

	var updatedSalon *models.Salon // Déclaration comme un pointeur

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedSalon); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedSalon, err := db.UpdateSalon(salonID, *updatedSalon) // Passer la valeur pointée
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedSalon)
}

// DeleteSalon supprime un salon
func DeleteSalon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	salonID := vars["id"]

	err := db.DeleteSalon(salonID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Salon supprimé avec succès"))
}
