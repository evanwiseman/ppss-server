package models

import (
	"time"

	"github.com/evanwiseman/ppss-server/internal/database/local"
)

type Device struct {
	SerialNumber string     `json:"serial_number"`
	Name         string     `json:"name"`
	IpAddress    string     `json:"ip_address"`
	DeviceType   string     `json:"device_type"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	LastSeen     *time.Time `json:"last_seen,omitempty"`
}

func DB2Device(d local.Device) Device {
	var lastSeen *time.Time
	if d.LastSeen.Valid {
		lastSeen = &d.LastSeen.Time
	}

	return Device{
		SerialNumber: d.SerialNumber,
		Name:         d.Name,
		IpAddress:    d.IpAddress,
		DeviceType:   d.DeviceType,
		CreatedAt:    &d.CreatedAt,
		UpdatedAt:    &d.UpdatedAt,
		LastSeen:     lastSeen,
	}
}
