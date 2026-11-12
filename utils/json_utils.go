package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"RNTreaty/model"
)

type ModelTask model.Task

// WriteJSON writes the response for the http client
func WriteJSON(w http.ResponseWriter, status int, payload any) error {
	marshaledByte, err := json.Marshal(payload)
	if err != nil {
		return explainJSONError(err)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)

	_, writeErr := w.Write(marshaledByte)
	if writeErr != nil {
		return explainJSONError(writeErr)
	}

	return nil
}

// DecodeJSON decodes the request payload from the client http body
func DecodeJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return errors.New("misssing Body data")
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // catch typos like "emial"

	if err := dec.Decode(payload); err != nil {
		return explainJSONError(err)
	}
	// Reject trailing garbage, e.g. `{"a":1} {"b":2}`.
	if _, err := dec.Token(); !errors.Is(err, io.EOF) {
		return errors.New("unexpected data after JSON value")
	}
	return nil
}

// explainJSONError turns encoding/json errors into messages a client can act on.
func explainJSONError(err error) error {
	var (
		syntaxErr *json.SyntaxError
		typeErr   *json.UnmarshalTypeError
	)
	switch {
	case errors.As(err, &syntaxErr):
		return fmt.Errorf("malformed JSON at byte %d: %w", syntaxErr.Offset, err)
	case errors.As(err, &typeErr):
		return fmt.Errorf("field %q must be %s (got %s)", typeErr.Field, typeErr.Type, typeErr.Value)
	case errors.Is(err, io.EOF), errors.Is(err, io.ErrUnexpectedEOF):
		return errors.New("empty or truncated JSON body")
	default:
		return err // includes "unknown field" errors from DisallowUnknownFields
	}
}
