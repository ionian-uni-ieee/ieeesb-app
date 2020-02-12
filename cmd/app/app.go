package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database"
	httpHandler "github.com/ionian-uni-ieee/ieeesb-app/internal/app/handlers/rest"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories"
	"github.com/joho/godotenv"
)

// Application holds and initiates the software structure
type Application struct {
	host             string
	port             string
	databaseHost     string
	databasePort     string
	databaseName     string
	databaseUsername string
	databasePass     string
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
	a.initializeVariables()

	// Get database starting
	err := databaseDriver.Connect(a.databaseHost, a.databasePort, a.databaseUsername, a.databasePass, a.databaseName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("✔ Database server running at " + a.databaseHost + ":" + a.databasePort + "/" + a.databaseName)

	// HTTP Server IO
	r := mux.NewRouter()
	// Trims trailing slash
	r.StrictSlash(true)

	reps := repositories.MakeRepositories(databaseDriver)
	h := httpHandler.MakeHandler(reps)

	// UTILITIES
	r.HandleFunc("/ping", h.GetPing).Methods("GET")
	// AUTH
	r.HandleFunc("/v1/register", h.PostRegister).Methods("POST")
	r.HandleFunc("/v1/login", h.PostLogin).Methods("POST")
	r.HandleFunc("/v1/logout", h.PostLogout).Methods("POST")
	r.HandleFunc("/v1/profile", h.GetProfile).Methods("GET")
	// USERS
	r.HandleFunc("/v1/users/{userID}", h.DeleteUser).Methods("DELETE")
	r.HandleFunc("/v1/users/{userID}", h.PutUser).Methods("PUT")
	// SPONSORS
	r.HandleFunc("/v1/sponsors", h.PostSponsor).Methods("POST")
	r.HandleFunc("/v1/sponsors/{sponsorID}", h.PutSponsor).Methods("PUT")
	r.HandleFunc("/v1/sponsors/{sponsorID}", h.DeleteSponsor).Methods("DELETE")
	// TICKETS
	r.HandleFunc("/v1/tickets", h.GetTickets).Methods("GET")
	r.HandleFunc("/v1/tickets/contact", h.PostTicketContact).Methods("POST")

	// Logging middleware
	handler := handlers.LoggingHandler(os.Stdout, r)

	log.Println("✔ API server running at " + a.host + ":" + a.port)
	log.Fatal(http.ListenAndServe(a.host+":"+a.port, handler))
}

// initializeVariables initiates all the application variables
// they're taken from the environment variables,
// which if non-existant are replaced by default values instead
func (a *Application) initializeVariables() {
	// Load .env file(s) if existant
	godotenv.Load(".env")
	initializeVariable(&a.host, "API_HOST", defaultHost)
	initializeVariable(&a.port, "API_PORT", defaultPort)
	initializeVariable(&a.databaseHost, "API_DATABASE_HOST", defaultDatabaseHost)
	initializeVariable(&a.databasePort, "API_DATABASE_PORT", defaultDatabasePort)
	initializeVariable(&a.databaseName, "API_DATABASE_NAME", defaultDatabaseName)
	initializeVariable(&a.databaseUsername, "API_DATABASE_NAME", "")
	initializeVariable(&a.databasePass, "API_DATABASE_PASS", "")
}

func initializeVariable(ptr *string, envName string, defaultValue string) {
	if len(*ptr) == 0 {
		*ptr = os.Getenv("API_HOST")
		if len(*ptr) == 0 {
			*ptr = defaultValue
		}
	}
}
