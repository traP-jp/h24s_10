package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type EventDate struct {
	ID        uuid.UUID `db:"id"`
	EventID   uuid.UUID `db:"event_id"`
	Start     time.Time `db:"start"`
	End       time.Time `db:"end"`
	CreatedAt time.Time `db:"created_at"`
}

func (repo *Repository) GetEventDates(eventID uuid.UUID) ([]EventDate, error) {
	return nil, errors.New("not implemented")
}

// CreateEventDates creates event dates. This function uses a bulk insert.
func (repo *Repository) CreateEventDates(dates []EventDate) error {
	return errors.New("not implemented")
}
