package model

import "github.com/google/uuid"

type Comment struct {
	ID      uuid.UUID `db:"id"`
	EventID uuid.UUID `db:"event_id"`
	TraQID  string    `db:"traq_id"`
	Content string    `db:"comment"`
}

func (repo *Repository) GetComments(eventID uuid.UUID) ([]Comment, error) {
	return nil, nil
}

func (repo *Repository) CreateComment(comment Comment) error {
	return nil
}
