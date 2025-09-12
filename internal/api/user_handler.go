package api

import (
	"errors"
	"log"
	"regexp"

	"github.com/joaquinbian/workout-api-go/internal/store"
)

type UserHandler struct {
	userStore *store.UserStore
	logger    *log.Logger
}

func NewUserHandler(userStore *store.UserStore, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userStore: userStore,
		logger:    logger,
	}
}

type registerUseRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	Bio          string `json:"bio"`
}

func (h *UserHandler) validateRegisterUserRequest(userTorRegister *registerUseRequest) error {
	if userTorRegister.Username == "" {
		return errors.New("username is required")
	}

	if len(userTorRegister.Username) > 50 {
		return errors.New("username must be at most 50 characters long")
	}

	if userTorRegister.Email == "" {
		return errors.New("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$ `)

	if !emailRegex.Match([]byte(userTorRegister.Email)) {
		return errors.New("email is not well formatted")
	}

	if userTorRegister.PasswordHash == "" {
		return errors.New("password is requried")
	}
	return nil
}
