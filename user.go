package main

import (
	"encoding/json"
	"go-crud/database"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID   int
	Name string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := User{}
	json.NewDecoder(r.Body).Decode(&user)

	result := database.Connector.Create(&user)

	if result.Error != nil {
		http.Error(w, "Error creating user data", http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := []User{}

	result := database.Connector.Find(&users)

	if result.Error != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result.Value)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["id"]

	user := User{}
	database.Connector.First(&user, userID)

	newUser := User{}
	json.NewDecoder(r.Body).Decode(&newUser)

	user.Name = newUser.Name

	result := database.Connector.Save(&user)
	json.NewEncoder(w).Encode(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["id"]

	user := User{
		Name: userID,
	}
	result := database.Connector.Delete(&user, userID)

	json.NewEncoder(w).Encode(result)
}
