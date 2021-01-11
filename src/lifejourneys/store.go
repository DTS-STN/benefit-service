package lifejourneys

import (
	"database/sql"

	"github.com/DTS-STN/benefit-service/models"
)

type LifeJourneyRepo interface {
	GetAll() ([]models.LifeJourney, error)
	GetByID(id string) (models.LifeJourney, error)
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
func (s *Store) GetAll() ([]models.LifeJourney, error) {
	var results []models.LifeJourney

	rows, err := s.db.Query("SELECT * FROM benefits.life_journey")
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
		results = append(results, models.LifeJourney{ID: id, Title: title, Description: description})
	}

	return results, nil
}

// GetByID returns a singe benefit based on the id
func (s *Store) GetByID(id string) (models.LifeJourney, error) {
	var (
		lifeJourneyID string
		title         string
		description   string
	)
	err := s.db.QueryRow("SELECT * FROM benefits.life_journey WHERE ID = $1", id).Scan(&lifeJourneyID, &title, &description)

	if err != nil {
		return models.LifeJourney{}, err
	}

	return models.LifeJourney{ID: lifeJourneyID, Title: title, Description: description}, nil
}
