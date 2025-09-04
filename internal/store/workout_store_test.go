package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")

	if err != nil {
		t.Fatalf("Connect to DB: %v", err)
	}

	err = Migrate(db, "../../migrations")

	if err != nil {
		t.Fatalf("Running migrations: %v", err)
	}

	_, err = db.Exec("TRUNCATE workouts, workout_entries")

	if err != nil {
		t.Fatalf("Truncating tables: %v", err)
	}

	return db
}

func TestCreateWorkout(t *testing.T) {
	db := setupDB(t)

	defer db.Close()

	store := NewPostgresWorkoutStore(db)

	tests := []struct {
		name    string
		workout *Workout
		wantErr bool
	}{
		{
			name: "valid workout",
			workout: &Workout{
				Title:           "push day",
				Description:     "upper push day",
				DurationMinutes: 60,
				CaloriesBurned:  200,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Bench press",
						Sets:         3,
						Reps:         IntPtr(10),
						Weight:       FloatPtr(135.5),
						Notes:        "warm up properly",
						OrderIndex:   1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "workout with invalid entries",
			workout: &Workout{
				Title:           "full body",
				Description:     "complete workout",
				DurationMinutes: 90,
				CaloriesBurned:  500,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Plank",
						Sets:         3,
						Reps:         IntPtr(60),
						Notes:        "keeps form",
						OrderIndex:   1,
					},
					{
						ExerciseName:    "squats",
						Sets:            4,
						Reps:            IntPtr(12),
						DurationSeconds: IntPtr(60),
						Weight:          FloatPtr(185.5),
						Notes:           "full depth",
						OrderIndex:      2,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			createdWorkout, err := store.CreateWorkout(c.workout)
			if c.wantErr {
				assert.Error(t, err)
				return
			}
			//if there is no error, we require error to be nil
			require.NoError(t, err)
			assert.Equal(t, c.workout.Title, createdWorkout.Title)
			assert.Equal(t, c.workout.Description, createdWorkout.Description)
			assert.Equal(t, c.workout.CaloriesBurned, createdWorkout.CaloriesBurned)
			assert.Equal(t, c.workout.DurationMinutes, createdWorkout.DurationMinutes)

			retrieved, err := store.GetWorkoutByID(int64(createdWorkout.ID))

			require.NoError(t, err)

			assert.Equal(t, retrieved.ID, createdWorkout.ID)
			assert.Equal(t, len(retrieved.Entries), len(createdWorkout.Entries))

			for i, entry := range retrieved.Entries {
				assert.Equal(t, c.workout.Entries[i].ExerciseName, entry.ExerciseName)
				assert.Equal(t, c.workout.Entries[i].DurationSeconds, entry.DurationSeconds)
				assert.Equal(t, c.workout.Entries[i].Notes, entry.Notes)
			}
		})
	}
}

// helper funcition para obtener el puntero de una variable int rapido
func IntPtr(i int) *int {
	return &i
}

func FloatPtr(f float64) *float64 {
	return &f
}
