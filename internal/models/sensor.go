package models

import (
	"time"

	"github.com/evanwiseman/ppss-server/internal/database"
	"github.com/google/uuid"
)

type Sensor struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LastSeenAt time.Time `json:"last_seen_at"`
}

func DB2Sensor(s database.Sensor) Sensor {
	return Sensor{
		ID:         s.ID,
		Name:       s.Name,
		CreatedAt:  s.CreatedAt,
		UpdatedAt:  s.UpdatedAt,
		LastSeenAt: s.LastSeenAt,
	}
}
