package model

import (
	"time"

	"github.com/google/uuid"
)

type Applicant struct {
	ID        uuid.UUID `db:"id"`
	EventID   uuid.UUID `db:"event_id"`
	TraQID    string    `db:"traq_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (repo *Repository) CreateApplicant(applicant Applicant) error {
	_, err := repo.db.Exec("INSERT INTO applicants (id, event_id, traq_id) VALUES (?, ?, ?)",
		applicant.ID, applicant.EventID, applicant.TraQID)
	return err
}

func (repo *Repository) GetApplicants(eventID uuid.UUID) ([]Applicant, error) {
	applicants := make([]Applicant, 0)
	err := repo.db.Select(&applicants, "SELECT * FROM applicants WHERE event_id = ?", eventID)
	if err != nil {
		return nil, err
	}
	return applicants, nil
}
