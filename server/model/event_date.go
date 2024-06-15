package model

import (
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
	eventDates := make([]EventDate, 0)
	err := repo.db.Select(&eventDates, "SELECT * FROM event_dates WHERE event_id = ?", eventID)
	if err != nil {
		return nil, err
	}
	return eventDates, nil
}

// CreateEventDates creates event dates. This function uses a bulk insert.
func (repo *Repository) CreateEventDates(dates []EventDate) error {
	tx, err := repo.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.NamedExec("INSERT INTO `event_dates` (id, event_id, start, end) VALUES (:id, :event_id, :start, :end)", dates)
	if err != nil {
		return err
	}

	return tx.Commit()
}
