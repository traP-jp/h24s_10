package model

import (
	"errors"

	"github.com/google/uuid"
)

type Participant struct {
	ID      uuid.UUID `db:"id"`
	EventID uuid.UUID `db:"event_id"`
	TraQID  string    `db:"traq_id"`
	TeamID  int       `db:"team_id"`
}

func (repo *Repository) GetParticipants(eventID uuid.UUID) ([]Participant, error) {
	participants := make([]Participant, 0)
	err := repo.db.Select(&participants, "SELECT * FROM participants WHERE event_id = ?", eventID)
	if err != nil {
		return nil, err
	}
	return participants, nil
}

func (repo *Repository) CreateParticipant(participant Participant) error {
	return errors.New("not implemented")
}

func (repo *Repository) CreateParticipants(participants []Participant) error {
	if len(participants) == 0 {
		return nil
	}

	tx, err := repo.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.NamedExec("INSERT INTO `participants` (id, event_id, traq_id, team_id) VALUES (:id, :event_id, :traq_id, :team_id)", participants)
	if err != nil {
		return err
	}

	return tx.Commit()
}
