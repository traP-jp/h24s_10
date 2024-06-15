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
	Location    string    `db:"location"`
	Description string    `db:"description"`
	IsConfirmed bool      `db:"is_confirmed"`
	CreatedAt   time.Time `db:"created_at"`
}

func (repo *Repository) GetEvents() ([]Event, error) {
	return nil, errors.New("not implemented")
}

func (repo *Repository) CreateEvent(event Event) error {
	var err error
	if event.Start.IsZero() || event.End.IsZero() {
		_, err = repo.db.Exec("INSERT INTO events (id, title, host_id, location, description, is_confirmed) VALUES (?, ?, ?, ?, ?, ?)",
			event.ID, event.Title, event.HostID, event.Location, event.Description, event.IsConfirmed)
	} else {
		_, err = repo.db.Exec("INSERT INTO events (id, title, start, end, host_id, description, is_confirmed) VALUES (?, ?, ?, ?, ?, ?, ?)",
			event.ID, event.Title, event.Start, event.End, event.HostID, event.Description, event.IsConfirmed)
	}
	return err
}
