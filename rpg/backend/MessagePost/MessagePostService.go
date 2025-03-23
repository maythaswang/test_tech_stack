package messagepost

import (
	"database/sql"
	"fmt"
	"log"
)

type MessagePostService struct {
	db *sql.DB
}

// Create new MessagePostService
func NewMessagePostService(db *sql.DB) *MessagePostService {
	return &MessagePostService{db: db}
}

// Add message to DB
func (s *MessagePostService) addMessage(messagePostRecord *MessagePostRecord) error {
	var query string = "INSERT INTO messages (body) VALUES ($1) RETURNING id, created_at"
	var err error = s.db.QueryRow(query, messagePostRecord.Body).Scan(&messagePostRecord.ID, &messagePostRecord.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to add message to database: %w", err)
	}
	log.Printf("New message inserted: ID = %d, CreatedAt = %s\n", messagePostRecord.ID, messagePostRecord.CreatedAt)

	return nil
}

// Get message from DB with ID
func (s *MessagePostService) GetMessage(id int) (*MessagePostRecord, error) {
	var message MessagePostRecord
	query := "SELECT id, body, created_at FROM messages WHERE id = $1"
	err := s.db.QueryRow(query, id).Scan(&message.ID, &message.Body, &message.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("message not found")
		}
		return nil, fmt.Errorf("failed to retrieve message: %w", err)
	}
	log.Printf(message.Body)
	return &message, nil
}
