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
	_, err := repo.db.Exec("INSERT INTO comments (id, event_id, traq_id, content) VALUES (?, ?, ?, ?)",
		comment.ID, comment.EventID, comment.TraQID, comment.Content)
	return err
}
