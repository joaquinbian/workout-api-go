package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Envelope: map donde el key es un string y el value es un interface vacio, que quiere decir que puede ser cualquier tipo de dato
type Envelope map[string]interface{}

func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	js = append(js, '\n') // Agrega un salto de línea al final del JSON

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}

func ReadIdParam(w http.ResponseWriter, r *http.Request) (int64, error) {
	paramID := chi.URLParam(r, "id")

	if paramID == "" {
		http.NotFound(w, r)
		return 0, errors.New("no param was provided")
	}

	id, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		http.NotFound(w, r)
		return 0, err
	}

	return id, nil
}
