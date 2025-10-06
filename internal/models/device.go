package models

import "time"

type Device struct {
	SerialNumber string     `json:"serial_number"`
	Name         string     `json:"name"`
	IPAddress    string     `json:"ip_address"`
	DeviceType   string     `json:"device_type"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	LastSeen     *time.Time `json:"last_seen,omitempty"`
}
