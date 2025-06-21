package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) GetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")

	if paramID == "" {
		http.NotFound(w, r)
		return
	}

	workoutID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "The workout id is %d\n", workoutID)
}

func (wh *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created workout\n")
}
