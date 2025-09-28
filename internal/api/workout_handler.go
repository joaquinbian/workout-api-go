package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/joaquinbian/workout-api-go/internal/middleware"
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

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout": workout})
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

	currentUser := middleware.GetUser(r)
	if currentUser == nil || currentUser == store.AnonymousUser {
		wh.logger.Printf("error: Create Workout handler: invalid user: %v", err)
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"message": "debes estar logueado para ejecutar esta operacion"})
		return
	}
	workout.UserID = currentUser.ID

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)

	if err != nil {
		wh.logger.Printf("error: creating workout: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"message": "no pudimos procesar la solicitud"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout": createdWorkout})
}

func (wh *WorkoutHandler) GetWorkouts(w http.ResponseWriter, r *http.Request) {

	workouts, err := wh.workoutStore.GetWorkouts()

	if err != nil {
		wh.logger.Printf("error: GetWorkouts: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"message": "error obteniendo los workouts"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workouts": workouts})
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

	userReq := middleware.GetUser(r)

	workoutOwner, err := wh.workoutStore.GetWorkoutOwner(workoutID)

	if err != nil {
		wh.logger.Printf("error: update workout: get workout owner: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"message": "workout inexistente"})
			return
		}
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"message": "internal server error"})
		return

	}
	if userReq == nil || userReq == store.AnonymousUser {
		wh.logger.Printf("error: Create Workout handler: invalid user: %v", err)
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"message": "debes estar logueado para ejectuar esta operacion"})
		return
	}

	if userReq.ID != workoutOwner {
		wh.logger.Printf("error: Create Workout handler: unauthorized user: %v", err)
		utils.WriteJSON(w, http.StatusForbidden, utils.Envelope{"message": "no puedes modificar este workout"})
		return
	}

	err = wh.workoutStore.UpdateWorkout(existingWorkout)
	if err != nil {
		wh.logger.Printf("error: UpdateWorkout: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"message": "error al modificar el workout"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout": existingWorkout})

}

func (wh *WorkoutHandler) DeleteWorkout(w http.ResponseWriter, r *http.Request) {

	workoutID, err := utils.ReadIdParam(w, r)

	if err != nil {
		wh.logger.Printf("error: ReadIdParam: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"message": "error leyendo el id de los params"})
		return
	}

	userReq := middleware.GetUser(r)

	workoutOwner, err := wh.workoutStore.GetWorkoutOwner(workoutID)

	if err != nil {
		wh.logger.Printf("error: update workout: get workout owner: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"message": "workout inexistente"})
			return
		}
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"message": "internal server error"})
		return

	}
	if userReq == nil || userReq == store.AnonymousUser {
		wh.logger.Printf("error: Create Workout handler: invalid user: %v", err)
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"message": "debes estar logueado para ejectuar esta operacion"})
		return
	}

	if userReq.ID != workoutOwner {
		wh.logger.Printf("error: Create Workout handler: unauthorized user: %v", err)
		utils.WriteJSON(w, http.StatusForbidden, utils.Envelope{"message": "no puedes eliminar este workout"})
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

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "workout eliminado"})

}
