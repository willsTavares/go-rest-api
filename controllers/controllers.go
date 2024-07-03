package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willsTavares/go-rest-api/database"
	"github.com/willsTavares/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var persona []models.Personality
	database.DB.Find(&persona)
	json.NewEncoder(w).Encode(persona)
}

func PersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Personality
	database.DB.First(&persona, id)
	json.NewEncoder(w).Encode(persona)
}
