package model

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `db:"id"`
	EventID   uuid.UUID `db:"event_id"`
	TraQID    string    `db:"traq_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

func (repo *Repository) GetComment(eventID uuid.UUID, traQID string) (Comment, error) {
	var comment Comment
	err := repo.db.Get(&comment, "SELECT * FROM comments WHERE event_id = ? AND traq_id = ?", eventID, traQID)
	return comment, err
}

func (repo *Repository) CreateComment(comment Comment) error {
	_, err := repo.db.Exec("INSERT INTO comments (id, event_id, traq_id, content) VALUES (?, ?, ?, ?)",
		comment.ID, comment.EventID, comment.TraQID, comment.Content)
	return err
}
