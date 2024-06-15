package model

import (
	"fmt"

	"github.com/google/uuid"
)

type (
	Target struct {
		ID      uuid.UUID `db:"id"`
		EventID uuid.UUID `db:"event_id"`
		UserID  uuid.UUID `db:"user_id"`
	}
)

func (r *Repository) GetTargetsByEventID(eventID uuid.UUID) ([]Target, error) {
	var targets []Target
	if err := r.db.Select(&targets, "SELECT * FROM targets WHERE event_id = $1", eventID); err != nil {
		return nil, fmt.Errorf("failed to get targets: %w", err)
	}
	return targets, nil
}
