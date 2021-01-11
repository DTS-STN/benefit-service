package benefits

import (
	"database/sql"

	"github.com/DTS-STN/benefit-service/models"
)

type BenefitRepo interface {
	GetAll() ([]models.Benefits, error)
	GetByID(id string) (models.Benefits, error)
}

// Store struct manages interactions with the database
type Store struct {
	db *sql.DB
}

// New returns a new Store struct
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetAll returns all benefits
func (s *Store) GetAll() ([]models.Benefits, error) {
	var results []models.Benefits

	rows, err := s.db.Query("SELECT id, title, description FROM benefits.benefit")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id          string
			title       string
			description string
		)
		err := rows.Scan(&id, &title, &description)
		if err != nil {
			return nil, err
		}
		results = append(results, models.Benefits{ID: id, Title: title, Description: description})
	}

	return results, nil
}

// GetByID returns a singe benefit based on the id
func (s *Store) GetByID(benefitID string) (models.Benefits, error) {
	var (
		id          string
		title       string
		description string
	)

	err := s.db.QueryRow("SELECT id, title, description FROM benefits.benefit WHERE id = $1", benefitID).Scan(&id, &title, &description)
	if err != nil {
		return models.Benefits{}, err
	}

	return models.Benefits{ID: id, Title: title, Description: description}, nil
}
