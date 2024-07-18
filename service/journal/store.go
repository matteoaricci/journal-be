package journal

import (
	"database/sql"
	"fmt"

	"github.com/matteoaricci/journal-be/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllJournals() (*types.Journal, error) {
	rows, err := s.db.Query("SELECT * FROM journals")

	if err != nil {
		return nil, err
	}

	j := new(types.Journal)
	for rows.Next() {
		j, err = scanRowIntoJournal(rows)
		if err != nil {
			return nil, err
		}
	}

	if j == nil {
		return nil, fmt.Errorf("unable to find any journals")
	}

	return j, nil

}

func scanRowIntoJournal(rows *sql.Rows) (*types.Journal, error) {
	j := new(types.Journal)

	err := rows.Scan(&j.ID, &j.Body, &j.Title)

	if err != nil {
		return nil, err
	}

	return j, nil
}