// planway/backend/routes/salon_routes.go
package routes

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateSalon crée un nouveau salon
func CreateSalon(w http.ResponseWriter, r *http.Request) {
	var salon models.Salon
	_ = json.NewDecoder(r.Body).Decode(&salon)

	// Valider les données du salon
	if salon.Name == "" || salon.Location == "" {
		http.Error(w, "Nom et emplacement du salon requis", http.StatusBadRequest)
		return
	}

	// Ajouter la logique pour créer un salon dans la base de données
	createdSalon, err := db.CreateSalon(salon)
	if err != nil {
		http.Error(w, "Erreur lors de la création du salon", http.StatusInternalServerError)
		return
	}

	// Répondre avec le nouveau salon créé
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdSalon)
}

// GetSalonByID récupère les détails d'un salon spécifique par son ID
func GetSalonByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	salonID := params["id"]

	// Ajouter la logique pour récupérer un salon depuis la base de données par son ID
	salon, err := db.GetSalon(salonID)
	if err != nil {
		http.Error(w, "Salon non trouvé", http.StatusNotFound)
		return
	}

	// Répondre avec les détails du salon
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(salon)
}

// UpdateSalon met à jour les détails d'un salon existant par son ID
func UpdateSalon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	salonID := params["id"]

	var updatedSalon models.Salon
	_ = json.NewDecoder(r.Body).Decode(&updatedSalon)

	// Ajouter la logique pour mettre à jour les détails d'un salon dans la base de données
	updatedSalon, err := db.UpdateSalon(salonID, updatedSalon)
	if err != nil {
		http.Error(w, "Erreur lors de la mise à jour du salon", http.StatusInternalServerError)
		return
	}

	// Répondre avec le salon mis à jour
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedSalon)
}

// DeleteSalon supprime un salon par son ID
func DeleteSalon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	salonID := params["id"]

	// Ajouter la logique pour supprimer un salon dans la base de données
	err := db.DeleteSalon(salonID)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression du salon", http.StatusInternalServerError)
		return
	}

	// Répondre avec un message de succès
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Salon supprimé avec succès"})
}

// GetAllSalons récupère la liste de tous les salons
func GetAllSalons(w http.ResponseWriter, r *http.Request) {
	// Ajouter la logique pour récupérer tous les salons depuis la base de données
	salons, err := db.GetAllSalons()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des salons", http.StatusInternalServerError)
		return
	}

	// Répondre avec la liste des salons
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(salons)
}
