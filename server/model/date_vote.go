package model

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type (
	DateVote struct {
		ID        uuid.UUID `db:"id"`
		TraQID    string    `db:"traq_id"`
		DateID    uuid.UUID `db:"event_date_id"`
		CreatedAt time.Time `db:"created_at"`
	}
	DateVoteAndEventID struct {
		EventID uuid.UUID `db:"event_id"`
	}
)

func (repo *Repository) GetEventDateVotes(eventDateID uuid.UUID) ([]DateVote, error) {
	dateVotes := make([]DateVote, 0)
	err := repo.db.Select(&dateVotes, "SELECT * FROM date_votes WHERE event_date_id = ?", eventDateID)
	if err != nil {
		return nil, err
	}
	return dateVotes, nil
}

func (repo *Repository) CreateDateVotes(votes []DateVote) error {
	if len(votes) == 0 {
		return nil
	}

	tx, err := repo.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.NamedExec("INSERT INTO date_votes (id, event_date_id, traq_id) VALUES (:id, :event_date_id, :traq_id)", votes)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (repo *Repository) GetDateVotesByUser(traqID string) ([]DateVoteAndEventID, error) {
	dateVoteAndEventID := make([]DateVoteAndEventID, 0)
	err := repo.db.Select(&dateVoteAndEventID, "SELECT event_dates.event_id FROM date_votes JOIN event_dates ON date_votes.event_date_id = event_dates.id WHERE date_votes.traq_id = ?", traqID)
	if err != nil {
		return nil, err
	}
	log.Println(dateVoteAndEventID)
	return dateVoteAndEventID, nil
}
