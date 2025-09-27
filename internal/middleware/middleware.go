package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/joaquinbian/workout-api-go/internal/store"
	"github.com/joaquinbian/workout-api-go/internal/tokens"
	"github.com/joaquinbian/workout-api-go/internal/utils"
)

type UserMiddleware struct {
	UserStore store.UserStore
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
		panic("no hay usuarios en la request")
	}

	return u
}

func (um *UserMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//aqui podemos interceptar cualquier request entrante a nuestro server
		//The Vary HTTP header is used to inform caches about which request headers influence the response content
		w.Header().Set("Vary", "Authorization")

		authHeader := w.Header().Get("Authorization")

		if authHeader == "" {
			r = SetUser(r, store.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authHeader, " ")

		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "invalid auth header"})
			return
		}

		token := headerParts[1]

		user, err := um.UserStore.GetUserToken(tokens.ScopeAuth, token)

		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "invalid token"})
			return
		}

		if user == nil {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "invalid or expired token"})
			return

		}
		r = SetUser(r, user)
		next.ServeHTTP(w, r)
		return
	})
}

func (um *UserMiddleware) RequireUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := GetUser(r)

		if user.IsAnonymous() {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "you must be logged to access this route"})
			return
		}

		next.ServeHTTP(w, r)
		return
	})

}
