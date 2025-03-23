package messagepost

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// We are merging controller and service for this case (since its just gonna be a few functions anyway)
type MessagePostController struct {
	messagePostService *MessagePostService
}

func NewMessagePostController(messagepostservice *MessagePostService) *MessagePostController {
	return &MessagePostController{messagePostService: messagepostservice}
}

// /api/post_message
func (s *MessagePostController) PostMessage(w http.ResponseWriter, r *http.Request) error {

	// Try read request
	postBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return fmt.Errorf("unable to read request body")
	}
	defer r.Body.Close()

	// Parse JSON
	var messagePostRecord *MessagePostRecord = &MessagePostRecord{}
	err = json.Unmarshal(postBody, &messagePostRecord)
	if err != nil {
		http.Error(w, "invalid JSON format", http.StatusBadRequest)
		return fmt.Errorf("invalid JSON format")
	}

	// Add message to DB
	err = s.messagePostService.addMessage(messagePostRecord)
	if err != nil {
		http.Error(w, "unable to save message", http.StatusInternalServerError)
		return fmt.Errorf("unable to save message: %w", err)
	}

	// Successful
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message received successfully"))

	return nil
}

// /api/get_message/{message_id}
func (s *MessagePostController) GetMessage(w http.ResponseWriter, r *http.Request) error {
	// get id
	id, err := strconv.Atoi(r.PathValue("message_id"))
	if err != nil {
		http.Error(w, "message not found", http.StatusNotFound)
		return fmt.Errorf("message not found: %w", err)
	}

	// Get message from db
	messagePostRecord, err := s.messagePostService.GetMessage(id)
	if err != nil {
		http.Error(w, "message not found", http.StatusNotFound)
		return fmt.Errorf("message not found: %w", err)
	}

	// Set header
	w.Header().Set("Content-Type", "application/json")

	// Encode the messagePostRecord
	response, err := json.Marshal(messagePostRecord)
	if err != nil {
		http.Error(w, "failed to serialize message", http.StatusInternalServerError)
		return fmt.Errorf("failed to serialize message: %v", err)
	}

	// Successful
	w.WriteHeader(http.StatusOK)
	w.Write(response)

	return nil

}

func (s *MessagePostController) DeleteMessage(w http.ResponseWriter, r *http.Request) error {
	// get id
	id, err := strconv.Atoi(r.PathValue("message_id"))
	if err != nil {
		http.Error(w, "message not found", http.StatusNotFound)
		return fmt.Errorf("message not found: %w", err)
	}

	// Delete message from DB
	err = s.messagePostService.DeleteMessage(id)
	if err != nil {
		http.Error(w, "message not found", http.StatusNotFound)
		return fmt.Errorf("message not found: %w", err)
	}

	// Successful
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message deleted successfully"))

	return nil

}
