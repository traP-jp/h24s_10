package model

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID    `db:"id"`
	Title       string       `db:"title"`
	Start       sql.NullTime `db:"start"`
	End         sql.NullTime `db:"end"`
	HostID      string       `db:"host_id"`
	Location    string       `db:"location"`
	Description string       `db:"description"`
	IsConfirmed bool         `db:"is_confirmed"`
	CreatedAt   time.Time    `db:"created_at"`
}

func (repo *Repository) GetEvents() ([]Event, error) {
	var events []Event
	err := repo.db.Select(&events, "SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (repo *Repository) CreateEvent(event Event) error {
	_, err := repo.db.Exec("INSERT INTO events (id, title, start, end, host_id, location, description, is_confirmed) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		event.ID, event.Title, event.Start, event.End, event.HostID, event.Location, event.Description, event.IsConfirmed)
	return err
}

// UpdateEvent updates an event.
// This function overwrites fields, so you need the full data to update.
func (repo *Repository) UpdateEvent(event Event) error {
	_, err := repo.db.Exec("UPDATE events SET title = ?, start = ?, end = ?, location = ?, description = ?, is_confirmed = ? WHERE id = ?",
		event.Title, event.Start, event.End, event.Location, event.Description, event.IsConfirmed, event.ID)
	return err
}

func (repo *Repository) GetEventByEventID(eventID uuid.UUID) (Event, error) {
	var event Event
	err := repo.db.Get(&event, "SELECT * FROM events WHERE id = ?", eventID)
	if err != nil {
		log.Printf("failed to get event: %v", err)
		return Event{}, err
	}
	return event, nil
}
