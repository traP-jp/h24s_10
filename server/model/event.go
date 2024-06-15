package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Start       time.Time `db:"start"`
	End         time.Time `db:"end"`
	HostID      string    `db:"host_id"`
	Description string    `db:"description"`
	IsConfirmed bool      `db:"is_confirmed"`
	CreatedAt   time.Time `db:"created_at"`
}

func (repo *Repository) GetEvents() ([]Event, error) {
	return nil, errors.New("not implemented")
}

func (repo *Repository) CreateEvent(event Event) error {
	return errors.New("not implemented")
}
