package journal

import (
	"database/sql"

	"github.com/matteoaricci/journal-be/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetJournals() ([]types.Journal, error) {
	rows, err := s.db.Query("SELECT * FROM journals")

	if err != nil {
		return nil, err
	}

	journals := make([]types.Journal, 0)
	for rows.Next() {
		j, err := scanRowIntoJournal(rows)
		if err != nil {
			return nil, err
		}

	journals = append(journals, *j)
	}

	return journals, nil
}

func (s *Store) CreateJournal(payload types.NewJournalPayload) (*types.Journal, error) {
 	_, err := s.db.Query("INSERT INTO journals (body, title) VALUES (?, ?)", payload.Body, payload.Title)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func scanRowIntoJournal(rows *sql.Rows) (*types.Journal, error) {
	j := new(types.Journal)

	err := rows.Scan(&j.ID, &j.Body, &j.Title, &j.CreatedAt, &j.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return j, nil
}