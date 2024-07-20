package types

type Journal struct {
	Title string `json:"title"`
	Body string `json:"body"`
	ID string `json:"id"`
}

type JournalStore interface {
	GetJournals() ([]Journal, error)
}