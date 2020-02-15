package app

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/handlers/rest"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories"
	"github.com/joho/godotenv"
)

// Default values
const defaultHost string = "localhost"
const defaultPort string = "8081"
const defaultDatabaseHost string = "mongodb://localhost"
const defaultDatabasePort string = "27017"
const defaultDatabaseName string = "test"

// Initialize initiates the application's server and communication channels
// A webserver and a database server are configured and set in cooperation
func Initialize(databaseDriver database.Driver) {
	initializeVariables()

	// Get database starting
	err := databaseDriver.Connect(
		os.Getenv("API_DATABASE_HOST"),
		os.Getenv("API_DATABASE_UNAME"),
		os.Getenv("API_DATABASE_PASS"),
		os.Getenv("API_DATABASE_NAME"),
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(
		"✔ " +
			strings.Title(databaseDriver.GetDatabaseType()) +
			" Database server running at " +
			os.Getenv("DATABASE_ORIGIN") +
			"/" +
			os.Getenv("DATABASE_NAME"),
	)

	/*
	 * HTTP ROUTER
	 *
	 */
	r := mux.NewRouter()
	// Trims trailing slash
	r.StrictSlash(true)

	rep := repositories.MakeRepositories(databaseDriver)
	rest.MakeHandler(r, rep)

	// Run HTTP Server
	httpOrigin := os.Getenv("API_ORIGIN")
	log.Println("✔ API server running at " + httpOrigin)
	log.Fatal(http.ListenAndServe(
		httpOrigin,
		r,
	))
}

// initializeVariables initiates all the application variables
// they're taken from the environment variables,
// which if non-existant are replaced by default values instead
func initializeVariables() {
	// Load .env file(s) if existant
	godotenv.Load(".env")
	initializeVariable("API_HOST", defaultHost)
	initializeVariable("API_PORT", defaultPort)
	initializeVariable("API_DATABASE_HOST", defaultDatabaseHost)
	initializeVariable("API_DATABASE_PORT", defaultDatabasePort)
	initializeVariable("API_DATABASE_NAME", defaultDatabaseName)
	initializeVariable("API_DATABASE_UNAME", "")
	initializeVariable("API_DATABASE_PASS", "")
}

func initializeVariable(envName string, defaultValue string) {
	if os.Getenv(envName) != "" {
		return
	}
	// Load .env file(s) if existant
	var envVars map[string]string
	envVars, err := godotenv.Read(".env")
	if err != nil {
		if envVars[envName] != "" {
			os.Setenv(envName, envVars[envName])
			return
		}
	}
	os.Setenv(envName, defaultValue)
}
