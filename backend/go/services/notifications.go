package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func QueueEmailNotification(userID, eventType string, payload interface{}) {
	body, err := json.Marshal(map[string]interface{}{
		"user_id":    userID,
		"event_type": eventType,
		"payload":    payload,
	})
	if err != nil {
		return
	}

	http.Post("http://localhost:8000/internal/notifications/queue", "application/json", bytes.NewReader(body))
}