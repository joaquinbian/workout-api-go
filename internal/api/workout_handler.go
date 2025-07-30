package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joaquinbian/workout-api-go/internal/store"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
}

func NewWorkoutHandler(workoutStore store.WorkoutStore) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
	}
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
	var workout store.Workout

	//decodea el body (data) en el struct de workout
	err := json.NewDecoder(r.Body).Decode(&workout)

	if err != nil {
		fmt.Println(err) //luego lo mejoramos
		http.Error(w, "No pudimos procesar su solicitud. Formato incorrecto", http.StatusInternalServerError)
		return
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)

	if err != nil {
		fmt.Println(err) //luego lo mejoramos
		http.Error(w, "No pudimos procesar su solicitud", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWorkout)
}
