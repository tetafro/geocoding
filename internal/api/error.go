package api

import (
	"fmt"
)

// Error is a custom error type which can be represented
// as a JSON error message {"error": "message"}.
type Error string

func (err Error) Error() string {
	return string(err)
}

// MarshalJSON implements json.Marshaler.
func (err Error) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())), nil
}
