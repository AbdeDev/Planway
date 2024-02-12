// planway/backend/db/db.go
package db

import (
	"backend/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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

// CreateUser crée un nouvel utilisateur dans la base de données
func CreateUser(user models.User) (models.User, error) {
	var createdUser models.User

	// Vérifier si l'utilisateur existe déjà
	existingUser, err := getUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return createdUser, fmt.Errorf("un utilisateur avec cet e-mail existe déjà")
	}

	// Hasher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return createdUser, err
	}

	// Insérer l'utilisateur dans la base de données
	err = Conn.QueryRow("INSERT INTO users(email, password, role) VALUES($1, $2, $3) RETURNING id, email, role",
		user.Email, string(hashedPassword), user.Role).
		Scan(&createdUser.ID, &createdUser.Email, &createdUser.Role)
	if err != nil {
		return createdUser, err
	}

	return createdUser, nil
}

// AuthenticateUser vérifie les informations d'identification de l'utilisateur
func AuthenticateUser(email, password string) (models.User, error) {
	var user models.User

	// Récupérer l'utilisateur depuis la base de données par e-mail
	dbUser, err := getUserByEmail(email)
	if err != nil {
		return user, err
	}

	// Vérifier le mot de passe
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password))
	if err != nil {
		return user, fmt.Errorf("mot de passe incorrect")
	}

	// Retourner la copie de l'utilisateur (pas le pointeur)
	return *dbUser, nil
}

// getUserByEmail récupère un utilisateur depuis la base de données par e-mail
func getUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := Conn.QueryRow("SELECT * FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Aucun utilisateur trouvé
		}
		return nil, err
	}

	return &user, nil
}

// CreateReservation crée une nouvelle réservation dans la base de données
func CreateReservation(reservation models.Reservation) (models.Reservation, error) {
	var createdReservation models.Reservation

	// Ajouter la logique pour créer une réservation dans la base de données
	_, err := Conn.Exec("INSERT INTO reservations(salon_id, start_time, end_time, client_id) VALUES($1, $2, $3, $4) RETURNING id, salon_id, start_time, end_time, client_id",
		reservation.SalonID, reservation.StartTime, reservation.EndTime, reservation.ClientID)
	if err != nil {
		return createdReservation, err
	}

	// Remplacer cet exemple par la réservation réellement créée dans la base de données
	return createdReservation, nil
}

// GenerateJWT génère un token JWT pour l'utilisateur
func GenerateJWT(user models.User) (string, error) {
	// Ajouter la logique pour générer un token JWT pour l'utilisateur
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
	})
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// CreateSalon crée un nouveau salon dans la base de données
func CreateSalon(newSalon models.Salon) (models.Salon, error) {
	var createdSalon models.Salon

	// Ajouter la logique pour créer un salon dans la base de données
	_, err := Conn.Exec("INSERT INTO salons(name, location) VALUES($1, $2) RETURNING id, name, location",
		newSalon.Name, newSalon.Location)
	if err != nil {
		return createdSalon, err
	}

	// Remplacer cet exemple par le salon réellement créé dans la base de données
	return createdSalon, nil
}

// UpdateSalon met à jour les détails d'un salon existant
func UpdateSalon(salonID string, updatedSalon models.Salon) (models.Salon, error) {
	err := Conn.QueryRow("UPDATE salons SET name=$1, location=$2 WHERE id=$3 RETURNING id, name, location",
		updatedSalon.Name, updatedSalon.Location, salonID).
		Scan(&updatedSalon.ID, &updatedSalon.Name, &updatedSalon.Location)
	if err != nil {
		return models.Salon{}, err
	}

	return updatedSalon, nil
}

// GetSalon récupère les détails d'un salon spécifique par son ID
func GetSalon(salonID string) (models.Salon, error) {
	var salon models.Salon
	err := Conn.QueryRow("SELECT * FROM salons WHERE id = $1", salonID).Scan(&salon.ID, &salon.Name, &salon.Location)
	if err != nil {
		return models.Salon{}, err
	}

	return salon, nil
}

// DeleteSalon supprime un salon par son ID
func DeleteSalon(salonID string) error {
	_, err := Conn.Exec("DELETE FROM salons WHERE id = $1", salonID)
	return err
}

func GetAllSalons() ([]models.Salon, error) {
	rows, err := Conn.Query("SELECT * FROM salons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var salons []models.Salon
	for rows.Next() {
		var salon models.Salon
		err := rows.Scan(&salon.ID, &salon.Name, &salon.Location)
		if err != nil {
			return nil, err
		}
		salons = append(salons, salon)
	}

	return salons, nil
}