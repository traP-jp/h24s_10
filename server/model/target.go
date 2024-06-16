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

func (repo *Repository) CreateEventTargets(targets []EventTarget) error {
	if len(targets) == 0 {
		return nil
	}

	tx, err := repo.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.NamedExec("INSERT INTO `targets` (id, event_id, traq_id) VALUES (:id, :event_id, :traq_id)", targets)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (repo *Repository) GetEventsByTargetUser(targetUser string) ([]Event, error) {
	var events []Event
	err := repo.db.Select(&events, "SELECT * FROM events WHERE id IN (SELECT event_id FROM targets WHERE traq_id = ?)", targetUser)
	if err != nil {
		return nil, err
	}
	return events, nil
}
