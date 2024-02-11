package db

import (
	"backend/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

func InitDB() {
	var err error

	// Construire la chaîne de connexion à PostgreSQL
	connStr := "user=ton_utilisateur dbname=ton_nom_de_base password=ton_mot_de_passe sslmode=disable"

	// Connexion à la base de données
	Conn, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Vérification de la connexion à la base de données
	err = Conn.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to the database")

	// À faire: Initialisation de la base de données, création de tables, etc.
}

// GetAllSalons renvoie la liste de tous les salons
func GetAllSalons() ([]models.Salon, error) {
	rows, err := Conn.Query("SELECT * FROM salons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var salons []models.Salon
	for rows.Next() {
		var salon models.Salon
		err := rows.Scan(&salon.ID, &salon.Name, &salon.Address)
		if err != nil {
			return nil, err
		}
		salons = append(salons, salon)
	}

	return salons, nil
}

// GetSalon renvoie les détails d'un salon spécifique
func GetSalon(salonID string) (*models.Salon, error) {
	var salon models.Salon
	err := Conn.QueryRow("SELECT * FROM salons WHERE id = $1", salonID).Scan(&salon.ID, &salon.Name, &salon.Address)
	if err != nil {
		return nil, err
	}

	return &salon, nil
}

// CreateSalon crée un nouveau salon
func CreateSalon(newSalon models.Salon) (*models.Salon, error) {
	err := Conn.QueryRow("INSERT INTO salons(name, address) VALUES($1, $2) RETURNING id, name, address", newSalon.Name, newSalon.Address).
		Scan(&newSalon.ID, &newSalon.Name, &newSalon.Address)
	if err != nil {
		return nil, err
	}

	return &newSalon, nil
}

// UpdateSalon met à jour les détails d'un salon existant
func UpdateSalon(salonID string, updatedSalon models.Salon) (*models.Salon, error) {
	err := Conn.QueryRow("UPDATE salons SET name=$1, address=$2 WHERE id=$3 RETURNING id, name, address",
		updatedSalon.Name, updatedSalon.Address, salonID).
		Scan(&updatedSalon.ID, &updatedSalon.Name, &updatedSalon.Address)
	if err != nil {
		return nil, err
	}

	return &updatedSalon, nil
}

// DeleteSalon supprime un salon
func DeleteSalon(salonID string) error {
	_, err := Conn.Exec("DELETE FROM salons WHERE id=$1", salonID)
	if err != nil {
		return err
	}

	return nil
}
