package models

import (
	"time"

	"github.com/evanwiseman/ppss-server/internal/database"
	"github.com/google/uuid"
)

type Wdlm struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LastSeenAt time.Time `json:"last_seen_at"`
}

func DB2Wdlm(w database.Wdlm) Wdlm {
	return Wdlm{
		ID:         w.ID,
		Name:       w.Name,
		CreatedAt:  w.CreatedAt,
		UpdatedAt:  w.UpdatedAt,
		LastSeenAt: w.LastSeenAt,
	}
}
