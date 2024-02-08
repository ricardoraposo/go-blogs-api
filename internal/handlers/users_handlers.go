package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ricardoraposo/blogs-api-go/internal/database"
	"github.com/ricardoraposo/blogs-api-go/internal/entities"
	"github.com/ricardoraposo/blogs-api-go/internal/utils"
)

type UserHandler struct {
	UserDB database.UserDBInterface
}

type CreateUserDTO struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Image       string `json:"image"`
}

func NewUserHandler(userDB database.UserDBInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var p CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.WriteToJson(w, map[string]string{"error": "invalid request"})
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user, err := entities.NewUser(p.DisplayName, p.Email, p.Password, p.Image)
	if err != nil {
		utils.WriteToJson(w, map[string]string{"error": "something went wrong"})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser, err := h.UserDB.CreateUser(user)
	if err != nil {
		utils.WriteToJson(w, map[string]string{"error": "something went wrong"})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteToJson(w, newUser)
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserDB.GetUsers()
	if err != nil {
		utils.WriteToJson(w, map[string]string{"error": "something went wrong"})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteToJson(w, users)
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.UserDB.GetByID(id)
	if err != nil {
		utils.WriteToJson(w, map[string]string{"error": "No user found with this id"})
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteToJson(w, user)
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	user, err := h.UserDB.GetByEmail(email)
	if err != nil {
		utils.WriteToJson(w, map[string]string{"error": "No user found with this email"})
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteToJson(w, user)
	w.WriteHeader(http.StatusOK)
}
