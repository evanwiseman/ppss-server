package server

import (
	"encoding/json"
	"net/http"

	"github.com/evanwiseman/ppss-server/internal/database"
	"github.com/evanwiseman/ppss-server/internal/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (ls *LocalServer) PostWdlmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse the JSON
	var params struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		models.RespondWithError(w, http.StatusBadRequest, "invalid JSON", err)
		return
	}

	// Create the sensor in the database
	created, err := ls.Queries.CreateWdlm(r.Context(), params.Name)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// unique_violation
			if pqErr.Code == "23505" {
				models.RespondWithError(
					w,
					http.StatusConflict,
					"device already exists",
					err,
				)
				return
			}
		}
		models.RespondWithError(
			w,
			http.StatusInternalServerError,
			"unable to create device",
			err,
		)
		return
	}

	// Respond with created entry
	resp := models.DB2Wdlm(created)
	models.RespondWithJSON(w, http.StatusCreated, resp)
}

func (ls *LocalServer) PutWdlmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse the JSON
	var params struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		models.RespondWithError(w, http.StatusBadRequest, "invalid JSON", err)
		return
	}

	// Update the wdlm
	updated, err := ls.Queries.UpdateWdlm(r.Context(), database.UpdateWdlmParams{
		ID:   params.ID,
		Name: params.Name,
	})
	if err != nil {
		models.RespondWithError(w, http.StatusInternalServerError, "unable to update wdlm", err)
		return
	}

	// Respond with updated entry
	resp := models.DB2Wdlm(updated)
	models.RespondWithJSON(w, http.StatusOK, resp)
}
