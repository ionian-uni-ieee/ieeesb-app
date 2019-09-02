package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database"
	httpHandler "github.com/ionian-uni-ieee/ieee-webapp/internal/app/handlers/http"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
	"github.com/joho/godotenv"
)

// Application holds and initiates the software structure
type Application struct {
	Host             string
	Port             string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUsername string
	DatabasePass     string
}

// Default values
const defaultHost string = "localhost"
const defaultPort string = "8081"
const defaultDatabaseHost string = "mongodb://localhost"
const defaultDatabasePort string = "27017"
const defaultDatabaseName string = "test"

// Initialize initiates the application's server and communication channels
// A webserver and a database server are configured and set in cooperation
func (a *Application) Initialize(databaseDriver database.Driver) {
	if *a != (Application{}) {
		log.Println("⚠️ ENVIRONMENT VARIABLES WILL BE OVERRIDEN BY Application{} VARIABLES IN /main.go")
	}

	a.initializeVariables()

	// Get database starting
	err := databaseDriver.Connect(a.DatabaseHost, a.DatabasePort, a.DatabaseUsername, a.DatabasePass, a.DatabaseName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("✔ Database server running at " + a.DatabaseHost + ":" + a.DatabasePort + "/" + a.DatabaseName)

	// HTTP Server IO
	r := mux.NewRouter()
	// Trims trailing slash
	r.StrictSlash(true)

	rep := repositories.MakeRepositories(databaseDriver)
	h := httpHandler.MakeHandler(rep)

	r.HandleFunc("/ping", h.GetPing).Methods("GET")

	// USER AUTH
	r.HandleFunc("/v1/register", h.PostRegister).Methods("POST")
	r.HandleFunc("/v1/login", h.PostLogin).Methods("POST")
	r.HandleFunc("/v1/logout", h.GetLogout).Methods("GET")
	r.HandleFunc("/v1/profile", h.GetProfile).Methods("GET")
	// USERS
	r.HandleFunc("/v1/users/{userID}", h.DeleteUser).Methods("DELETE")
	// r.HandleFunc("/v1/users/{userID}", h.PutUser).Methods("PUT")
	// EVENTS
	// r.HandleFunc("/v1/events", h.PostEvent).Methods("POST")
	// r.HandleFunc("/v1/events/{eventID}", h.PutEvent).Methods("PUT")
	// r.HandleFunc("/v1/events/{eventID}", h.DeleteEvent).Methods("DELETE")
	// SPONSORS
	// r.HandleFunc("/v1/sponsors", h.PostSponsor).Methods("POST")
	// r.HandleFunc("/v1/sponsors/{sponsorID}", h.PutSponsor).Methods("PUT")
	// r.HandleFunc("/v1/sponsors/{sponsorID}", h.DeleteSponsor).Methods("DELETE")
	// TICKETS
	// r.HandleFunc("/v1/tickets", h.PostTicket).Methods("POST")
	// r.HandleFunc("/v1/tickets/{ticketID}/respond", h.PostTicketResponse).Methods("POST")
	// MEDIA
	// r.HandleFunc("/v1/images", h.PostImage).Methods("POST")
	// r.HandleFunc("/v1/video", h.PostVideo).Methods("POST")

	// Logging middleware
	handler := handlers.LoggingHandler(os.Stdout, r)

	log.Println("✔ API server running at " + a.Host + ":" + a.Port)
	log.Fatal(http.ListenAndServe(a.Host+":"+a.Port, handler))
}

// initializeVariables initiates all the application variables
// they're taken from the environment variables,
// which if non-existant are replaced by default values instead
func (a *Application) initializeVariables() {
	// Load .env file(s) if existant
	godotenv.Load(".env")
	// If variable is empty assign the according environment variable to it
	if len(a.Host) == 0 {
		a.Host = os.Getenv("API_HOST")
	}
	if len(a.Port) == 0 {
		a.Port = os.Getenv("API_PORT")
	}
	if len(a.DatabaseHost) == 0 {
		a.DatabaseHost = os.Getenv("API_DATABASE_HOST")
	}
	if len(a.DatabasePort) == 0 {
		a.DatabasePort = os.Getenv("API_DATABASE_PORT")
	}
	if len(a.DatabaseName) == 0 {
		a.DatabaseName = os.Getenv("API_DATABASE_NAME")
	}
	if len(a.DatabaseUsername) == 0 {
		a.DatabaseUsername = os.Getenv("API_DATABASE_USERNAME")
	}
	if len(a.DatabasePass) == 0 {
		a.DatabasePass = os.Getenv("API_DATABASE_PASS")
	}
	// Set to default values if empty
	if len(a.Host) == 0 {
		a.Host = defaultHost
	}
	if len(a.Port) == 0 {
		a.Port = defaultPort
	}
	if len(a.DatabaseHost) == 0 {
		a.DatabaseHost = defaultDatabaseHost
	}
	if len(a.DatabasePort) == 0 {
		a.DatabasePort = defaultDatabasePort
	}
	if len(a.DatabaseName) == 0 {
		a.DatabaseName = defaultDatabaseName
	}
}
