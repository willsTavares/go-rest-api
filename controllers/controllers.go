package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willsTavares/go-rest-api/database"
	"github.com/willsTavares/go-rest-api/models"
)

// ListAllPersonalities godoc
// @Summary List all personalities
// @Description get personalities
// @Tags Personalities
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Personality
// @Router /personalities [get]
func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var persona []models.Personality
	database.DB.Find(&persona)
	json.NewEncoder(w).Encode(persona)
}

// GetPersonalityById godoc
// @Summary Get a personality by ID
//
//	@Description get personality by ID
//	@Tags Personalities
//	@Accept  json
//	@Produce  json
//	@Param id path int true "Personality ID"
//	@Success 200 {object} models.Personality
//	@Router /personalities/{id} [get]
func PersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Personality
	database.DB.First(&persona, id)
	json.NewEncoder(w).Encode(persona)
}

// CreatePersonality godoc
// @Summary Create a personality
// @Description create personalities
// @Tags Personalities
//
//	@Accept  json
//	@Produce  json
//	@Param name body string true "Name"
//	@Param history body string true "History"
//
// @Success 200 {object} models.Personality
// @Router /personalities [post]
func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var newPersona models.Personality
	json.NewDecoder(r.Body).Decode(&newPersona)
	database.DB.Create(&newPersona)
	json.NewEncoder(w).Encode(newPersona)
}

// DeletePersonalityById godoc
// @Summary Delete a personality by ID
// @Description delete personality by ID
// @Tags Personalities
//
// @Accept  json
// @Produce  json
// @Param id path int true "Personality ID"
//
// @Success 200 {object} models.Personality
// @Router /personalities/{id} [delete]
func DeletePersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Personality
	database.DB.Delete(&persona, id)
	json.NewEncoder(w).Encode(persona)
}

// EditPersonalityById godoc
// @Summary Edit a personality by ID
// @Description edit personality by ID
// @Tags Personalities
//
// @Accept  json
// @Produce  json
// @Param id path int true "Personality ID"
//
// @Success 200 {object} models.Personality
// @Router /personalities/{id} [put]
func EditPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Personality
	database.DB.First(&persona, id)
	json.NewDecoder(r.Body).Decode(&persona)
	database.DB.Save(&persona)
	json.NewEncoder(w).Encode(persona)
}
