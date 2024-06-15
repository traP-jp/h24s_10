// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// GetEventApplicantsResponse defines model for GetEventApplicantsResponse.
type GetEventApplicantsResponse = []struct {
	DateOptionIDs *[]openapi_types.UUID `json:"dateOptionIDs,omitempty"`
	TraqID        *string               `json:"traqID,omitempty"`
}

// GetEventParticipantsResponse defines model for GetEventParticipantsResponse.
type GetEventParticipantsResponse = []string

// GetEventResponse defines model for GetEventResponse.
type GetEventResponse struct {
	Date *struct {
		End   time.Time `json:"end"`
		Start time.Time `json:"start"`
	} `json:"date,omitempty"`
	DateOptions *[]DateOption       `json:"dateOptions,omitempty"`
	Description string              `json:"description"`
	Id          *openapi_types.UUID `json:"id,omitempty"`
	IsConfirmed *bool               `json:"isConfirmed,omitempty"`
	Title       string              `json:"title"`
}

// GetEventTargetsResponse defines model for GetEventTargetsResponse.
type GetEventTargetsResponse = []string

// GetTraQGroupsResponse defines model for GetTraQGroupsResponse.
type GetTraQGroupsResponse struct {
	Members *[]TraQUser `json:"members,omitempty"`
	Name    *string     `json:"name,omitempty"`
}

// GetTraQUsersResponse defines model for GetTraQUsersResponse.
type GetTraQUsersResponse = []TraQUser

// PostEventApplicantsRequest defines model for PostEventApplicantsRequest.
type PostEventApplicantsRequest struct {
	DateOptionIDs *[]openapi_types.UUID `json:"dateOptionIDs,omitempty"`
}

// PostEventRequest defines model for PostEventRequest.
type PostEventRequest struct {
	DateOptions []struct {
		End   time.Time `json:"end"`
		Start time.Time `json:"start"`
	} `json:"dateOptions"`
	Description string   `json:"description"`
	Targets     []string `json:"targets"`
	Title       string   `json:"title"`
}

// PostEventResponse defines model for PostEventResponse.
type PostEventResponse struct {
	Id openapi_types.UUID `json:"id"`
}

// DateOption defines model for dateOption.
type DateOption struct {
	End   time.Time          `json:"end"`
	Id    openapi_types.UUID `json:"id"`
	Start time.Time          `json:"start"`
}

// TraQUser defines model for traQUser.
type TraQUser struct {
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
}

// EventID defines model for eventID.
type EventID = openapi_types.UUID

// PostEventsJSONRequestBody defines body for PostEvents for application/json ContentType.
type PostEventsJSONRequestBody = PostEventRequest

// PostEventsEventIDApplicantsJSONRequestBody defines body for PostEventsEventIDApplicants for application/json ContentType.
type PostEventsEventIDApplicantsJSONRequestBody = PostEventApplicantsRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /events)
	PostEvents(ctx echo.Context) error

	// (GET /events/{eventID})
	GetEventsEventID(ctx echo.Context, eventID EventID) error

	// (GET /events/{eventID}/applicants)
	GetEventsEventIDApplicants(ctx echo.Context, eventID EventID) error

	// (POST /events/{eventID}/applicants)
	PostEventsEventIDApplicants(ctx echo.Context, eventID EventID) error

	// (GET /events/{eventID}/participants)
	GetEventsEventIDParticipants(ctx echo.Context, eventID EventID) error

	// (GET /events/{eventID}/targets)
	GetEventsEventIDTargets(ctx echo.Context, eventID EventID) error

	// (GET /traq/groups)
	GetTraqGroups(ctx echo.Context) error

	// (GET /traq/users)
	GetTraqUsers(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostEvents converts echo context to params.
func (w *ServerInterfaceWrapper) PostEvents(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostEvents(ctx)
	return err
}

// GetEventsEventID converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventsEventID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventID" -------------
	var eventID EventID

	err = runtime.BindStyledParameterWithOptions("simple", "eventID", ctx.Param("eventID"), &eventID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventsEventID(ctx, eventID)
	return err
}

// GetEventsEventIDApplicants converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventsEventIDApplicants(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventID" -------------
	var eventID EventID

	err = runtime.BindStyledParameterWithOptions("simple", "eventID", ctx.Param("eventID"), &eventID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventsEventIDApplicants(ctx, eventID)
	return err
}

// PostEventsEventIDApplicants converts echo context to params.
func (w *ServerInterfaceWrapper) PostEventsEventIDApplicants(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventID" -------------
	var eventID EventID

	err = runtime.BindStyledParameterWithOptions("simple", "eventID", ctx.Param("eventID"), &eventID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostEventsEventIDApplicants(ctx, eventID)
	return err
}

// GetEventsEventIDParticipants converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventsEventIDParticipants(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventID" -------------
	var eventID EventID

	err = runtime.BindStyledParameterWithOptions("simple", "eventID", ctx.Param("eventID"), &eventID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventsEventIDParticipants(ctx, eventID)
	return err
}

// GetEventsEventIDTargets converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventsEventIDTargets(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventID" -------------
	var eventID EventID

	err = runtime.BindStyledParameterWithOptions("simple", "eventID", ctx.Param("eventID"), &eventID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventsEventIDTargets(ctx, eventID)
	return err
}

// GetTraqGroups converts echo context to params.
func (w *ServerInterfaceWrapper) GetTraqGroups(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTraqGroups(ctx)
	return err
}

// GetTraqUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetTraqUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTraqUsers(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/events", wrapper.PostEvents)
	router.GET(baseURL+"/events/:eventID", wrapper.GetEventsEventID)
	router.GET(baseURL+"/events/:eventID/applicants", wrapper.GetEventsEventIDApplicants)
	router.POST(baseURL+"/events/:eventID/applicants", wrapper.PostEventsEventIDApplicants)
	router.GET(baseURL+"/events/:eventID/participants", wrapper.GetEventsEventIDParticipants)
	router.GET(baseURL+"/events/:eventID/targets", wrapper.GetEventsEventIDTargets)
	router.GET(baseURL+"/traq/groups", wrapper.GetTraqGroups)
	router.GET(baseURL+"/traq/users", wrapper.GetTraqUsers)

}
