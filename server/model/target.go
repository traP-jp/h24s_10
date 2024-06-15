package model

import (
	"time"

	"github.com/google/uuid"
)

type EventTarget struct {
	ID        uuid.UUID `db:"id"`
	EventID   uuid.UUID `db:"event_id"`
	TraQID    string    `db:"traq_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (repo *Repository) GetEventTargets(eventID uuid.UUID) ([]EventTarget, error) {
	targets := make([]EventTarget, 0)
	err := repo.db.Select(&targets, "SELECT * FROM targets WHERE event_id = ?", eventID)
	if err != nil {
		return nil, err
	}
	return targets, nil
}
