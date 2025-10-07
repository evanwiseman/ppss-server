package models

import (
	"time"

	"github.com/evanwiseman/ppss-server/internal/database"
	"github.com/google/uuid"
)

type Device struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	IpAddress  string    `json:"ip_address"`
	DeviceType string    `json:"device_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LastSeen   time.Time `json:"last_seen"`
}

func DB2Device(d database.Device) Device {
	return Device{
		ID:         d.ID,
		Name:       d.Name,
		IpAddress:  d.IpAddress,
		DeviceType: d.DeviceType,
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
		LastSeen:   d.LastSeen,
	}
}
