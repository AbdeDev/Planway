// planway/backend/models/models.go
package models

import "time"

// Salon représente un salon de coiffure
type Salon struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

// TimeSlot représente un créneau d'ouverture dans un salon
type TimeSlot struct {
	ID        string    `json:"id"`
	SalonID   string    `json:"salonId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

// Reservation représente une réservation effectuée par un client
type Reservation struct {
	ID        string    `json:"id"`
	SalonID   string    `json:"salonId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	ClientID  string    `json:"clientId"`
}

// User représente un utilisateur (client ou salon)
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
