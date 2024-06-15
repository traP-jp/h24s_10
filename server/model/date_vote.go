package model

import (
	"errors"

	"github.com/google/uuid"
)

type DateVote struct {
	ID      uuid.UUID `db:"id"`
	EventID uuid.UUID `db:"event_id"`
	TraQID  string    `db:"traq_id"`
	DateID  uuid.UUID `db:"event_date_id"`
	Comment string    `db:"comment"`
}

func (repo *Repository) GetEventDateVotes(eventID uuid.UUID) ([]DateVote, error) {
	return nil, errors.New("not implemented")
}

func (repo *Repository) CreateDateVotes(votes []DateVote) error {
	return errors.New("not implemented")
}
