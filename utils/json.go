package utils

import "encoding/json"

// JSONStatus returns a message in JSON.
func JSONStatus(message string) []byte {
	m, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
	return m
}
