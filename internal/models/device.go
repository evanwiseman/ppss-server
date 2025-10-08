package models

import (
	"time"

	"github.com/evanwiseman/ppss-server/internal/database"
	"github.com/google/uuid"
)

type Device struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"device_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LastSeenAt time.Time `json:"last_seen"`
}

func DB2Device(d database.Device) Device {
	return Device{
		ID:         d.ID,
		Name:       d.Name,
		Type:       d.Type,
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
		LastSeenAt: d.LastSeenAt,
	}
}
