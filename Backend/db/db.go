// planway/backend/db/db.go
package db

import (
	"backend/models"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
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
		return createdUser, fmt.Errorf("Un utilisateur avec cet e-mail existe déjà")
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
		return user, fmt.Errorf("Mot de passe incorrect")
	}

	return dbUser, nil
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