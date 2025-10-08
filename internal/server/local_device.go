package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/evanwiseman/ppss-server/internal/database"
	"github.com/evanwiseman/ppss-server/internal/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (ls *LocalServer) PostDeviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse JSON request
	var params struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusBadRequest,
			"invalid JSON",
			err,
		)
		return
	}

	// Create entry in database
	created, err := ls.Queries.CreateDevice(r.Context(), database.CreateDeviceParams{
		Name: params.Name,
		Type: params.Type,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // unique_violation
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

	// Respond with created
	resp := models.DB2Device(created)
	models.RespondWithJSON(w, http.StatusCreated, resp)
}

func (ls *LocalServer) PutDevicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse JSON request
	var params struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Type string    `json:"type"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		models.RespondWithError(w, http.StatusBadRequest, "invalid JSON", err)
		return
	}

	// Update the device
	updated, err := ls.Queries.UpdateDevice(r.Context(), database.UpdateDeviceParams{
		ID:   params.ID,
		Name: params.Name,
		Type: params.Type,
	})
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusInternalServerError,
			"unable to update device",
			err,
		)
		return
	}

	// Send the response
	resp := models.DB2Device(updated)
	models.RespondWithJSON(w, http.StatusOK, resp)
}

func (ls *LocalServer) GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get all devices
	rows, err := ls.Queries.GetDevices(r.Context())
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusNotFound,
			"no devices found",
			err,
		)
		return
	}

	// Send the response
	var resp []models.Device
	for _, row := range rows {
		resp = append(resp, models.DB2Device(row))
	}
	models.RespondWithJSON(w, http.StatusOK, resp)
}

func (ls *LocalServer) GetDeviceByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the device ID
	deviceID, err := uuid.Parse(r.PathValue("deviceID"))
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusBadRequest,
			"invalid id",
			err,
		)
		return
	}

	// Get the device
	row, err := ls.Queries.GetDeviceByID(r.Context(), deviceID)
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusNotFound,
			"device not found",
			err,
		)
		return
	}

	// Send the response
	resp := models.DB2Device(row)
	models.RespondWithJSON(w, http.StatusOK, resp)
}

func (ls *LocalServer) DeleteDeviceByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the device ID
	deviceID, err := uuid.Parse(r.PathValue("deviceID"))
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusBadRequest,
			"invalid id",
			err,
		)
		return
	}

	// Check the device exists
	_, err = ls.Queries.GetDeviceByID(r.Context(), deviceID)
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusNotFound,
			"device not found",
			err,
		)
		return
	}

	// Delete the device with the matching ID
	err = ls.Queries.DeleteDeviceByID(r.Context(), deviceID)
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusInternalServerError,
			"unable to delete device",
			err,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ls *LocalServer) ResetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check that platform is in dev
	// TODO: Check if user is authenticated to reset
	if ls.Cfg.Platform != "dev" {
		models.RespondWithError(
			w,
			http.StatusUnauthorized,
			"not authorized to reset devices",
			errors.New("database not on dev platform"),
		)
		return
	}

	// Reset devices in the database
	err := ls.Queries.ResetDevices(r.Context())
	if err != nil {
		models.RespondWithError(
			w,
			http.StatusInternalServerError,
			"unable to reset devices",
			err,
		)
	}

	w.WriteHeader(http.StatusNoContent)
}
