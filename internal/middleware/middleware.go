package middleware

import (
	"context"
	"net/http"

	"github.com/joaquinbian/workout-api-go/internal/store"
)

type UserMiddleware struct {
	userStore store.UserStore
}

// tenemos que usar un tipo custom para evitar colisiones de nombre en el context
// no es igual el "user" (string) que contextKey("user") (contextKey)
type contextKey string

var UserContextKey = contextKey("user")

func SetUser(r *http.Request, user *store.User) *http.Request {
	//el context se usa, entre otras cosas, para pasar valores entre las request
	ctx := context.WithValue(r.Context(), UserContextKey, user)

	return r.WithContext(ctx)
}

func GetUser(r *http.Request) *store.User {
	ctx := r.Context()
	//asegura que el type que retorna del contexto es un puntero a un user
	u, ok := ctx.Value(UserContextKey).(*store.User)
	if !ok {
		return nil
	}

	return u
}
