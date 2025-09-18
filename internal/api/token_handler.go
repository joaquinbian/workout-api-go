package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/joaquinbian/workout-api-go/internal/store"
	"github.com/joaquinbian/workout-api-go/internal/tokens"
	"github.com/joaquinbian/workout-api-go/internal/utils"
)

type TokenHandler struct {
	tokenStore store.TokenStore
	userStore  store.UserStore
	logger     *log.Logger
}

type createTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewTokenHander(ts store.TokenStore, us store.UserStore, logger *log.Logger) *TokenHandler {
	return &TokenHandler{
		tokenStore: ts,
		userStore:  us,
		logger:     logger,
	}
}

func (th *TokenHandler) HandleCreateToken(w http.ResponseWriter, r *http.Request) {
	var req createTokenRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		th.logger.Printf("error: HandleCreateToken: decoding user: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request payload"})
		return
	}

	user, err := th.userStore.GetUserByUsername(req.Username)

	if err != nil || user == nil {
		th.logger.Printf("error: HandleCreateToken: getting userByUsername: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "iinternal server error"})
		return
	}

	matches, err := user.PasswordHash.Matches(req.Password)
	if err != nil {
		th.logger.Printf("error: HandleCreateToken: checking if password matches: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "iinternal server error"})
		return
	}

	if !matches {
		th.logger.Printf("error: HandleCreateToken: password doesnt match: %v", err)
		utils.WriteJSON(w, http.StatusForbidden, utils.Envelope{"error": "invalid user or password"})
		return
	}

	token, err := th.tokenStore.CreateNewToken(user.ID, 24*time.Hour, tokens.ScopeAuth)

	if err != nil {
		th.logger.Printf("error: HandleCreateToken: creating token: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "iinternal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"token": token})
	return
}
