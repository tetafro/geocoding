package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Respond writes a generic JSON response.
func Respond(w http.ResponseWriter, code int, msg interface{}) {
	resp := "{}"
	if msg != nil {
		b, err := json.Marshal(msg)
		if err != nil {
			code = http.StatusInternalServerError
			resp = fmt.Sprintf(`{"error": "%s"}`, http.StatusText(code))
		} else {
			resp = string(b)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintln(w, resp)
}

// RespondOK writes an empty 200 response.
func RespondOK(w http.ResponseWriter) {
	Respond(w, http.StatusOK, nil)
}

// RespondBadRequest writes a 400 error as a response.
func RespondBadRequest(w http.ResponseWriter) {
	code := http.StatusBadRequest
	err := Error(http.StatusText(code))
	Respond(w, code, err)
}

// RespondNotFound writes a 404 error as a response.
func RespondNotFound(w http.ResponseWriter) {
	code := http.StatusNotFound
	err := Error(http.StatusText(code))
	Respond(w, code, err)
}

// RespondInternalServerError writes a 500 error as a response.
func RespondInternalServerError(w http.ResponseWriter) {
	code := http.StatusInternalServerError
	err := Error(http.StatusText(code))
	Respond(w, code, err)
}
