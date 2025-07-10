package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"Todo/model"

	"github.com/golang-jwt/jwt/v5"
)

var userDB = make(map[string]string) // fake in-memory DB: username -> password

var jwtKey = []byte("my_secret_key") // 🔐 In real apps, move to env

// Register a new user
func Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Username == "" || user.Password == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if _, exists := userDB[user.Username]; exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	userDB[user.Username] = user.Password // 🔐 (store plain for now)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

// Login and generate JWT token
func Login(w http.ResponseWriter, r *http.Request) {
	var creds model.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	storedPass, exists := userDB[creds.Username]
	if !exists || creds.Password != storedPass {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Create JWT claims
	claims := &jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	// Create signed token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
