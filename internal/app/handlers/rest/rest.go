package rest

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/services/events"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/services/sponsors"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/services/tickets"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/services/users"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/reflections"
)

// Handler contains required services
type Handler struct {
	events   *events.Service
	users    *users.Service
	sponsors *sponsors.Service
	tickets  *tickets.Service
}

// MakeHandler is a Handler factory
func MakeHandler(router *mux.Router, rep *repositories.Repositories) *Handler {
	handler := &Handler{
		events:   events.MakeService(rep),
		users:    users.MakeService(rep),
		sponsors: sponsors.MakeService(rep),
		tickets:  tickets.MakeService(rep),
	}

	addRESTHandlers(
		router,
		apiResources{
			Event: resource{
				Path:     "/event",
				Handlers: resourceHandlers{},
			},
			Events: resource{
				Path:     "/events",
				Handlers: resourceHandlers{},
			},
			User: resource{
				Path:     "/user",
				Handlers: resourceHandlers{},
			},
			Users: resource{
				Path:     "/users",
				Handlers: resourceHandlers{},
			},
		},
	)

	return handler
}

// resourceHandlers contains all resource's handlers
type resourceHandlers struct {
	Get    func(http.ResponseWriter, *http.Request)
	Post   func(http.ResponseWriter, *http.Request)
	Put    func(http.ResponseWriter, *http.Request)
	Patch  func(http.ResponseWriter, *http.Request)
	Delete func(http.ResponseWriter, *http.Request)
}

// resource describes a resource's path and its handlers
type resource struct {
	Path     string
	Handlers resourceHandlers
}

// apiresources describes all REST API's resources
type apiResources struct {
	User     resource
	Users    resource
	Event    resource
	Events   resource
	Ticket   resource
	Tickets  resource
	Sponsor  resource
	Sponsors resource
}

// addRESTHandlers sets all the HTTP REST API's handlers onto the mux router
func addRESTHandlers(router *mux.Router, api apiResources) {
	v := reflect.ValueOf(api)
	typeOfV := v.Type()
	for fieldIndex := 0; fieldIndex < typeOfV.NumField(); fieldIndex++ {
		resource := v.Field(fieldIndex).Interface().(resource)
		addRESTHandler(router, resource)
	}
}

// addRESTHandler adds a single handler
func addRESTHandler(router *mux.Router, resource resource) {
	handlerNames, err := reflections.GetFieldNames(
		reflect.TypeOf(resource.Handlers),
	)
	if err != nil {
		panic(err)
	}
	for _, handlerName := range handlerNames {
		handlerInterface, err := reflections.GetFieldValue(
			resource.Handlers,
			handlerName,
		)
		if err != nil {
			panic(err)
		}
		handler := handlerInterface.(func(http.ResponseWriter, *http.Request))
		if handler != nil {
			handlerNameUppercase := strings.ToUpper(handlerName)
			router.HandleFunc(resource.Path, handler).Methods(handlerNameUppercase)
		}
	}
}
