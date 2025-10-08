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

func (ls *LocalServer) GetWdlmsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get all entries in wdlms table
	rows, err := ls.Queries.GetWdlms(r.Context())
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusInternalServerError,
			"wdlm table missing",
			err,
		)
		return
	}

	// Send the response of all entries
	var wdlms []models.Wdlm
	for _, row := range rows {
		wdlms = append(wdlms, models.DB2Wdlm(row))
	}
	models.RespondWithJSON(w, http.StatusOK, wdlms)
}

func (ls *LocalServer) GetWdlmsByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the wdlm id
	wdlmID, err := uuid.Parse(r.PathValue("wdlmID"))
	if err != nil {
		models.RespondWithError(w, http.StatusBadRequest, "invalid wdlm id", err)
		return
	}

	// Get the wdlm
	row, err := ls.Queries.GetWdlmByID(r.Context(), wdlmID)
	if err != nil {
		models.RespondWithError(w, http.StatusNotFound, "wdlm not found", err)
		return
	}

	// Respond with the wdlm
	resp := models.DB2Wdlm(row)
	models.RespondWithJSON(w, http.StatusOK, resp)
}
