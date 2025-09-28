package store

import (
	"database/sql"
)

type Workout struct {
	ID              int            `json:"id"`
	UserID          int            `json:"user_id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	DurationMinutes int            `json:"duration_minutes"`
	CaloriesBurned  int            `json:"calories_burned"`
	Entries         []WorkoutEntry `json:"entries"`
}

type WorkoutEntry struct {
	ID              int      `json:"id"`
	ExerciseName    string   `json:"exercise_name"`
	Sets            int      `json:"sets"`
	Reps            *int     `json:"reps"`
	DurationSeconds *int     `json:"duration_seconds"`
	Weight          *float64 `json:"weight"`
	Notes           string   `json:"notes"`
	OrderIndex      int      `json:"order_index"`
}

type PostgresWorkoutStore struct {
	db *sql.DB
}

func NewPostgresWorkoutStore(db *sql.DB) *PostgresWorkoutStore {
	return &PostgresWorkoutStore{db: db}
}

// asi es como tiene que verse nuestro store
// si en lugar de usar postgres usamos otra db, solo tiene que respetar esta interface
// la app trabajara con esta interface
type WorkoutStore interface {
	CreateWorkout(*Workout) (*Workout, error)
	GetWorkoutByID(id int64) (*Workout, error)
	GetWorkouts() ([]*Workout, error)
	UpdateWorkout(*Workout) error
	GetWorkoutOwner(id int64) (int, error)
	DeleteWorkout(id int64) error
}

func (pg *PostgresWorkoutStore) CreateWorkout(w *Workout) (*Workout, error) {
	tx, err := pg.db.Begin()

	if err != nil {
		return nil, err
	}

	//siempre llamamos a rollback, pq tx en su estado interno colecciona
	//el estado de las trasnacciones activas, en caso de que detecte un error
	//hace rollback
	defer tx.Rollback()

	query := `INSERT INTO workouts (title, user_id, description, duration_minutes, calories_burned)
	VALUES($1, $2, $3, $4, $5)
	RETURNING id
	`
	//Scan es el mecanismo que copia y convierte las columnas de la query en tus variables Go.
	//En .Scan(&w.ID) cada argumento debe ser un puntero a la variable donde querés guardar la columna.
	err = tx.QueryRow(query, w.Title, w.UserID, w.Description, w.DurationMinutes, w.CaloriesBurned).Scan(&w.ID)
	if err != nil {
		return nil, err
	}

	for _, entry := range w.Entries {
		query := `INSERT INTO workout_entries (workout_id, exercise_name, sets, reps, duration_seconds, weight, notes, order_index)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id
    `
		err = tx.QueryRow(query, w.ID, entry.ExerciseName, entry.Sets, entry.Reps, entry.DurationSeconds, entry.Weight, entry.Notes, entry.OrderIndex).Scan(&entry.ID)

		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return w, nil

}

func (pg *PostgresWorkoutStore) GetWorkoutByID(id int64) (*Workout, error) {
	w := &Workout{}
	query := `
  SELECT id, title, description, duration_minutes, calories_burned
  FROM workouts
  WHERE id = $1
  `

	err := pg.db.QueryRow(query, id).Scan(&w.ID, &w.Title, &w.Description, &w.DurationMinutes, &w.CaloriesBurned)

	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}

	if err != nil {
		return nil, err
	}

	workoutEntries, err := getWorkoutEntriesOfWorkout(pg.db, id)

	if err != nil {
		return nil, err
	}

	w.Entries = workoutEntries

	return w, nil
}

func (pg *PostgresWorkoutStore) GetWorkouts() ([]*Workout, error) {
	var workouts []*Workout

	query := `
  SELECT id, title, description, duration_minutes, calories_burned
  FROM workouts;

  `
	rows, err := pg.db.Query(query)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		workout := &Workout{}
		err := rows.Scan(&workout.ID, &workout.Title, &workout.Description, &workout.DurationMinutes, &workout.CaloriesBurned)

		if err != nil {
			return nil, err
		}

		workouts = append(workouts, workout)
	}

	for _, w := range workouts {
		workoutsEntries, err := getWorkoutEntriesOfWorkout(pg.db, int64(w.ID))
		if err != nil {
			return nil, err
		}
		w.Entries = workoutsEntries

	}

	return workouts, nil
}

func (pg *PostgresWorkoutStore) UpdateWorkout(w *Workout) error {

	//como es algo que cambia la bd, hacemos una transaccion

	tx, err := pg.db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `UPDATE workouts
  SET title = $1, description = $2, duration_minutes = $3, calories_burned = $4
  WHERE id = $5
  `

	result, err := pg.db.Exec(query, w.Title, w.Description, w.DurationMinutes, w.CaloriesBurned, w.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	//para actualizar los workout entries hacemos:
	//borramos todos los workout entries del workout que acabamos de actualizar
	_, err = pg.db.Exec(`DELETE FROM workout_entries WHERE workout_id = $1`, w.ID)

	if err != nil {
		return err
	}

	//actualizamos los workout entries del workout con los que vienen en la  llamada
	for _, entry := range w.Entries {
		query := `INSERT INTO workout_entries (workout_id, exercise_name, sets, reps, duration_seconds, weight, notes, order_index)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id
    `
		_, err := tx.Exec(query, w.ID, entry.ExerciseName, entry.Sets, entry.Reps, entry.DurationSeconds, entry.Weight, entry.Notes, entry.OrderIndex)

		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (pg *PostgresWorkoutStore) DeleteWorkout(id int64) error {

	query := `DELETE FROM workouts
  WHERE id = $1
  `
	result, err := pg.db.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func (pg *PostgresWorkoutStore) GetWorkoutOwner(workoutID int64) (int, error) {

	query := `SELECT user_id FROM workouts WHERE id = $1`

	var id int

	err := pg.db.QueryRow(query, workoutID).Scan(&id)

	if err != nil {
		return -1, err
	}

	if err == sql.ErrNoRows {
		return -1, err
	}

	return id, nil
}

// ver de poner esta en otro archivo
func getWorkoutEntriesOfWorkout(db *sql.DB, id int64) ([]WorkoutEntry, error) {
	var workoutEntries []WorkoutEntry

	entryQuery := `
  SELECT id, exercise_name, sets, reps, duration_seconds, weight, notes, order_index
  FROM workout_entries
  WHERE workout_id = $1
  ORDER BY order_index
  `

	rows, err := db.Query(entryQuery, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		wEntry := &WorkoutEntry{}

		err := rows.Scan(&wEntry.ID, &wEntry.ExerciseName, &wEntry.Sets, &wEntry.Reps, &wEntry.DurationSeconds, &wEntry.Weight, &wEntry.Notes, &wEntry.OrderIndex)

		if err != nil {
			return nil, err
		}

		workoutEntries = append(workoutEntries, *wEntry)
	}

	return workoutEntries, nil
}
