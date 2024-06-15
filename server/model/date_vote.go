package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type DateVote struct {
	ID        uuid.UUID `db:"id"`
	EventID   uuid.UUID `db:"event_id"`
	TraQID    string    `db:"traq_id"`
	DateID    uuid.UUID `db:"event_date_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (repo *Repository) GetEventDateVotes(eventDateID uuid.UUID) ([]DateVote, error) {
	dateVotes := make([]DateVote, 0)
	err := repo.db.Select(&dateVotes, "SELECT * FROM date_votes WHERE event_date_id = ?", eventDateID)
	if err != nil {
		return nil, err
	}
	return dateVotes, nil
}

func (repo *Repository) CreateDateVotes(votes []DateVote) error {
	return errors.New("not implemented")
}
