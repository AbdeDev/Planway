package models

// Salon représente un salon de coiffure
type Salon struct {
  ID      int    `json:"id"`
  Name    string `json:"name"`
  Address string `json:"address"`
  // Ajoute d'autres champs en fonction des besoins
}
