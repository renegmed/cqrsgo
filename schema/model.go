package schema

import "time"

// Meow struct primarily used for writing data to db
type Meow struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
