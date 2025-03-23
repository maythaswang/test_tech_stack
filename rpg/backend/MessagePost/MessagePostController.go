package messagepost

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// We are merging controller and service for this case (since its just gonna be a few functions anyway)
type MessagePostController struct {
}

func (s *MessagePostController) PostMessage(w http.ResponseWriter, r *http.Request) error {
	postBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read request body", http.StatusInternalServerError)
		return fmt.Errorf("unable to read request body")
	}

	defer r.Body.Close()

	var messagePostRecord *MessagePostRecord = &MessagePostRecord{}

	err = json.Unmarshal(postBody, &messagePostRecord)
	if err != nil {
		http.Error(w, "invalid JSON format", http.StatusInternalServerError)
		return fmt.Errorf("invalid JSON format")
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message received successfully"))

	return nil
}

func (s *MessagePostController) DeleteMessage(w http.ResponseWriter, r *http.Request) error {

}

func (s *MessagePostController) GetMessage(w http.ResponseWriter, r *http.Request) error {

}
