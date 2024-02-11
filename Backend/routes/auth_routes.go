// planway/backend/routes/auth_routes.go
package routes

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("secret_key") // À remplacer par une clé sécurisée dans un contexte de production

// RegisterUser enregistre un nouvel utilisateur
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// À faire: Valider les champs de l'utilisateur, vérifier s'il existe déjà, etc.
	if userExists(user.Username) {
		http.Error(w, "Utilisateur déjà enregistré", http.StatusBadRequest)
		return
	}

	// Hasher le mot de passe avant de l'enregistrer dans la base de données
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erreur lors de la création du mot de passe", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Ajouter la logique pour enregistrer l'utilisateur dans la base de données
	// db.CreateUser(user)

	// Répondre avec le nouvel utilisateur créé
	json.NewEncoder(w).Encode(user)
}

// LoginUser connecte un utilisateur existant
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// À faire: Valider les champs de connexion, vérifier l'existence de l'utilisateur, vérifier le mot de passe, etc.
	storedUser, err := getUserByUsername(user.Username)
	if err != nil || storedUser == nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusUnauthorized)
		return
	}

	// Comparer le mot de passe haché avec celui fourni
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	// Générer un token JWT
	tokenString, err := generateJWT(storedUser)
	if err != nil {
		http.Error(w, "Erreur lors de la génération du token", http.StatusInternalServerError)
		return
	}

	// Répondre avec le token JWT
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// getUserByUsername récupère un utilisateur depuis la base de données en utilisant le nom d'utilisateur
func getUserByUsername(username string) (*models.User, error) {
	// À faire: Ajouter la logique pour récupérer un utilisateur depuis la base de données par son nom d'utilisateur
	// return db.GetUserByUsername(username)
	return nil, nil
}

// userExists vérifie si un utilisateur existe déjà dans la base de données par son nom d'utilisateur
func userExists(username string) bool {
	// À faire: Ajouter la logique pour vérifier si un utilisateur existe déjà dans la base de données
	// return db.UserExists(username)
	return false
}

// generateJWT génère un token JWT pour un utilisateur
func generateJWT(user *models.User) (string, error) {
	// Créer une structure pour les revendications du token
	claims := jwt.MapClaims{
		"sub": user.ID,
		"iss": "planway-app",
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Token expire dans 1 heure
	}

	// Créer le token avec les revendications et la signature
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// À faire: Ajouter d'autres routes d'authentification au besoin
