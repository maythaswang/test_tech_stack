package messagepost

import (
	"time"
)

type MessagePostRecord struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
