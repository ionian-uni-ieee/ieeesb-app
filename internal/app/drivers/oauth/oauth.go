package oauth

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/handlers"
	"gopkg.in/go-oauth2/mongo.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

// DatabaseConnectionInfo contains information about a database connection
type DatabaseConnectionInfo struct {
	Origin string
	Name   string
}

// Server describes an oauth server's structure
type Server struct {
	Server *server.Server
}

// NewServer is a oauth Server factory
func NewServer(oauthServer *server.Server) Server {
	return Server{
		Server: oauthServer,
	}
}

// RunOAuthServer runs an oauth server
func RunOAuthServer(
	databaseInfo DatabaseConnectionInfo,
	authorizationHandler server.PasswordAuthorizationHandler,
	router *mux.Router,
) Server {
	manager := manage.NewDefaultManager()
	authorizeTokenConfig := &manage.Config{
		AccessTokenExp:    time.Hour * 24 * 14,
		RefreshTokenExp:   time.Hour * 24 * 15,
		IsGenerateRefresh: true,
	}
	manager.SetAuthorizeCodeTokenCfg(authorizeTokenConfig)

	// Token storage
	manager.MapTokenStorage(
		mongo.NewTokenStore(mongo.NewConfig(
			databaseInfo.Origin,
			databaseInfo.Name,
		)),
	)
	// Generate JWT access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate(
		[]byte("00000000"),
		jwt.SigningMethodHS512,
	))

	clientStore := store.NewClientStore()
	clientStore.Set("222222", &models.Client{
		ID:     "222222",
		Secret: "22222222",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	oauthServer := server.NewServer(server.NewConfig(), manager)
	oauthServer.SetAllowGetAccessRequest(true)
	oauthServer.SetClientInfoHandler(server.ClientFormHandler)

	oauthServer.SetPasswordAuthorizationHandler(authorizationHandler)

	oauthServer.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal error:", err.Error())
		return
	})

	oauthServer.SetResponseErrorHandler(func(err *errors.Response) {
		log.Println("Response error:", err.Error.Error())
		return
	})

	router.HandleFunc("/oauth2/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := oauthServer.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/oauth2/token", func(w http.ResponseWriter, r *http.Request) {
		oauthServer.HandleTokenRequest(w, r)
	}).Methods("POST")

	router.HandleFunc("/oauth2/token/{access_token}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		accessToken := vars["access_token"]
		err := manager.RemoveAccessToken(accessToken)
		if err != nil {
			handlers.SendError(
				w,
				http.StatusBadRequest,
				handlers.NewBadRequest(
					err.Error(),
				),
			)
			return
		}
		handlers.SendJSON(
			w,
			http.StatusOK,
			"Access token deleted",
		)
	}).Methods("GET")

	oauthServer.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	oauthServer.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return NewServer(oauthServer)
}
