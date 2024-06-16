package model

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

func (repo *Repository) ValidateEventDateIDsFromEventID(eventID uuid.UUID, dateIDs []uuid.UUID) error {
	if len(dateIDs) == 0 {
		return nil
	}

	var count int
	query, args, err := sqlx.In("SELECT COUNT(*) FROM event_dates WHERE event_id = ? AND id IN (?)", eventID, dateIDs)
	if err != nil {
		log.Println("Error creating SQL query")
		return err
	}
	err = repo.db.Get(&count, query, args...)
	if err != nil {
		log.Println("Error getting count from event_dates")
		return err
	}
	if count != len(dateIDs) {
		return fmt.Errorf("invalid date IDs: %v", dateIDs)
	}
	return nil
}

func (repo *Repository) ValidateEventDateIDsFromTraqID(traQID string, dateIDs []uuid.UUID) error {
	if len(dateIDs) == 0 {
		return nil
	}

	// すでにtraQIDとdateIDsの組み合わせが存在するか確認
	var count int
	query, args, err := sqlx.In("SELECT COUNT(*) FROM date_votes WHERE traq_id = ? AND event_date_id IN (?)", traQID, dateIDs)
	if err != nil {
		return err
	}
	err = repo.db.Get(&count, query, args...)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("invalid date IDs: %v", dateIDs)
	}
	return nil
}
