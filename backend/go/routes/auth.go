package routes

import (
	"net/http"

	"selleasy/models"
	"encoding/json"
	"golang.org/x/crypto/bcrypt" 
)
type registerRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
}

type loginRequest struct{
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterAuthRoutes(mux *http.ServeMux, store *models.Store) {
	mux.HandleFunc("POST /api/v1/register", func(w http.ResponseWriter, r *http.Request) {
		handleRegister(w, r, store)
	})
	mux.HandleFunc("POST /api/v1/login", func(w http.ResponseWriter, r *http.Request){
		handleLogin(w, r, store)
	})
}

func handleRegister(w http.ResponseWriter, r *http.Request, store *models.Store) {
	var req registerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if req.Email == "" || req.Password == "" || req.FullName == ""{
		http.Error(w, "password, email, and full_name are required, cannot be empty", http.StatusBadRequest)
		return
	}
	_, exists := store.GetUserByEmail(req.Email)
	if exists{
		http.Error(w, "an account already exists with this email", http.StatusConflict)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "could not hash password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		FullName:     req.FullName,
		Email:        req.Email,
		Phone:        req.Phone,
		City:         req.City,
		State:        req.State,
		Country:      req.Country,
		PasswordHash: string(hash),
		Role:         models.RoleBoth,
	}

	created := store.CreateUser(user)

	w.Write([]byte("created user: " + created.ID))
}
func handleLogin(w http.ResponseWriter, r *http.Request, store *models.Store) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, "json invalid", http.StatusBadRequest)
		return
	}
	user, exists := store.GetUserByEmail(req.Email)
	if !exists{
		http.Error(w, "invalid email and password", http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil{
		http.Error(w, "invalid email and password", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("logged in: " + user.ID))
}