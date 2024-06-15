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
	return nil, errors.New("not implemented")
}

func (repo *Repository) CreateParticipant(participant Participant) error {
	return errors.New("not implemented")
}
