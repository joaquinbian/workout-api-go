package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/joaquinbian/workout-api-go/internal/store"
	"github.com/joaquinbian/workout-api-go/internal/utils"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
	logger       *log.Logger
}

func NewWorkoutHandler(workoutStore store.WorkoutStore, logger *log.Logger) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
		logger:       logger,
	}
}

func (wh *WorkoutHandler) GetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	workoutID, err := utils.ReadIdParam(w, r)

	if err != nil {
		wh.logger.Printf("error: ReadIdParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"message": "error leyendo id de los params"})
		return
	}

	workout, err := wh.workoutStore.GetWorkoutByID(workoutID)

	if err != nil {
		wh.logger.Printf("error: GetWorkoutByID: %v", err)
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"message": "workout no encontrado"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workout)
}

func (wh *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout

	//decodea el body (data) en el struct de workout
	err := json.NewDecoder(r.Body).Decode(&workout)

	if err != nil {
		wh.logger.Printf("error: decoding workout: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"message": "formato incorrecto"})
		return
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)

	if err != nil {
		wh.logger.Printf("error: creating workout: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"message": "no pudimos procesar la solicitud"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWorkout)
}

func (wh *WorkoutHandler) GetWorkouts(w http.ResponseWriter, r *http.Request) {

	workouts, err := wh.workoutStore.GetWorkouts()

	if err != nil {
		wh.logger.Printf("error: GetWorkouts: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"message": "error obteniendo los workouts"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workouts)
}

func (wh *WorkoutHandler) UpdateWorkout(w http.ResponseWriter, r *http.Request) {

	workoutID, err := utils.ReadIdParam(w, r)

	if err != nil {
		wh.logger.Printf("error: ReadIdParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"message": "error leyendo id de los params"})
		return
	}

	existingWorkout, err := wh.workoutStore.GetWorkoutByID(workoutID)

	if err != nil {
		wh.logger.Printf("error: GetWorkoutByID: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"message": "error al obtener el workout"})
		return
	}

	if existingWorkout == nil {
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"message": "workout no encontrado"})
		return
	}

	var updateWorkoutRequest struct {
		Title           *string              `json:"title"`
		Description     *string              `json:"description"`
		DurationMinutes *int                 `json:"duration_minutes"`
		CaloriesBurned  *int                 `json:"calories_burned"`
		Entries         []store.WorkoutEntry `json:"entries"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateWorkoutRequest)

	if updateWorkoutRequest.Title != nil {
		existingWorkout.Title = *updateWorkoutRequest.Title
	}
	if updateWorkoutRequest.Description != nil {
		existingWorkout.Description = *updateWorkoutRequest.Description
	}
	if updateWorkoutRequest.DurationMinutes != nil {
		existingWorkout.DurationMinutes = *updateWorkoutRequest.DurationMinutes
	}
	if updateWorkoutRequest.CaloriesBurned != nil {
		existingWorkout.CaloriesBurned = *updateWorkoutRequest.CaloriesBurned
	}
	if updateWorkoutRequest.Entries != nil {
		existingWorkout.Entries = updateWorkoutRequest.Entries
	}

	err = wh.workoutStore.UpdateWorkout(existingWorkout)
	if err != nil {
		wh.logger.Printf("error: UpdateWorkout: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"message": "error al modificar el workout"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingWorkout)

}

func (wh *WorkoutHandler) DeleteWorkout(w http.ResponseWriter, r *http.Request) {

	workoutID, err := utils.ReadIdParam(w, r)

	if err != nil {
		wh.logger.Printf("error: ReadIdParam: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"message": "error leyendo el id de los params"})
		return
	}

	err = wh.workoutStore.DeleteWorkout(workoutID)

	if err == sql.ErrNoRows {
		wh.logger.Printf("error: DeleteWorkout: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"message": "ningun workout fue eliminado"})
		return
	}

	if err != nil {
		wh.logger.Printf("error: DeleteWorkout: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"message": "error eliminando el workout"})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Workout eliminado")

}
