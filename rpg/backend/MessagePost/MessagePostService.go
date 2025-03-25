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
func (s *MessagePostService) addMessage(message *MessagePostRecord) error {
	var query string = "INSERT INTO messages (body) VALUES ($1) RETURNING id, created_at"
	var err error = s.db.QueryRow(query, message.Body).Scan(&message.ID, &message.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to add message to database: %w", err)
	}
	log.Printf("New message inserted: ID = %d, CreatedAt = %s\n", message.ID, message.CreatedAt)

	return nil
}

// Get message from DB with ID
func (s *MessagePostService) GetMessage(id int) (*MessagePostRecord, error) {
	var message MessagePostRecord
	var query string = "SELECT id, body, created_at FROM messages WHERE id = $1"
	var err error = s.db.QueryRow(query, id).Scan(&message.ID, &message.Body, &message.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("message not found")
		}
		return nil, fmt.Errorf("failed to retrieve message: %w", err)
	}
	return &message, nil
}

// Get all messages from DB
func (s *MessagePostService) GetAllMessages() ([]*MessagePostRecord, error) {
	var allMessage []*MessagePostRecord
	var query string = "SELECT * FROM messages"
	rows, err := s.db.Query(query)

	// Get records from db
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("message not found")
		}
		return nil, fmt.Errorf("failed to retrieve message: %w", err)
	}

	// Go through each record
	for rows.Next() {
		var message MessagePostRecord
		var err error = rows.Scan(&message.ID, &message.Body, &message.CreatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("failed to read message")
			}
			return nil, fmt.Errorf("failed to read message: %w", err)
		}
		allMessage = append(allMessage, &message)

	}

	return allMessage, nil
}

// Delete message from DB with ID
func (s *MessagePostService) DeleteMessage(id int) error {
	var query string = "DELETE FROM messages WHERE id = $1"
	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no message found with ID %d", id)
	}
	return nil
}
