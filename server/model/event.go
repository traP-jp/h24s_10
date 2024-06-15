package model

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID  `db:"id"`
	Title       string     `db:"title"`
	Start       *time.Time `db:"start"`
	End         *time.Time `db:"end"`
	HostID      string     `db:"host_id"`
	Location    string     `db:"location"`
	Description string     `db:"description"`
	IsConfirmed bool       `db:"is_confirmed"`
	CreatedAt   time.Time  `db:"created_at"`
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
	var err error
	if event.Start == nil || event.End == nil {
		_, err = repo.db.Exec("INSERT INTO events (id, title, host_id, location, description, is_confirmed) VALUES (?, ?, ?, ?, ?, ?)",
			event.ID, event.Title, event.HostID, event.Location, event.Description, event.IsConfirmed)
	} else {
		_, err = repo.db.Exec("INSERT INTO events (id, title, start, end, host_id, description, is_confirmed) VALUES (?, ?, ?, ?, ?, ?, ?)",
			event.ID, event.Title, event.Start, event.End, event.HostID, event.Description, event.IsConfirmed)
	}
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
