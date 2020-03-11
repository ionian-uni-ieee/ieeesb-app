package oauth

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/handlers"
)

// ContextKey is a key for context values
type ContextKey string

// UserID is a key for the user id exported from the http oauth middleware
const UserID ContextKey = "userID"

// HTTPMiddleware authorizes a user
func (s *Server) HTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isWsProtocol := strings.HasPrefix(r.URL.String(), "/ws/")
		tokenPathRegex, err := regexp.Compile("/oauth2/token(/\\w+)?")
		if err != nil {
			log.Fatal(err)
		}
		isTokenPath := tokenPathRegex.Match([]byte(r.URL.String()))
		isNewUserRequest := r.URL.String() == "/users" && r.Method == "POST"

		if isTokenPath || isWsProtocol || isNewUserRequest {
			next.ServeHTTP(w, r)
			return
		}

		token, err := s.Server.ValidationBearerToken(r)
		if err != nil {
			handlers.SendError(
				w,
				http.StatusForbidden,
				handlers.NewHTTPError(
					"Invalid token",
					"Token was unauthorized or expired",
				),
			)
			return
		}
		ctx := context.WithValue(r.Context(), UserID, token.GetUserID())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
