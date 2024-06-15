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

// DateTimeResponse defines model for DateTimeResponse.
type DateTimeResponse struct {
	End   time.Time `json:"end"`
	Start time.Time `json:"start"`
}

// EventMeResponse defines model for EventMeResponse.
type EventMeResponse struct {
	Description *string            `json:"description,omitempty"`
	EventId     openapi_types.UUID `json:"event_id"`
	IsAnswered  bool               `json:"isAnswered"`
	IsConfirmed bool               `json:"isConfirmed"`
	IsHost      bool               `json:"isHost"`
	Title       string             `json:"title"`
}

// EventMeResponses defines model for EventMeResponses.
type EventMeResponses = []EventMeResponse

// GetAllEventsResponse defines model for GetAllEventsResponse.
type GetAllEventsResponse = []GetAllEventsElement

// GetEventApplicantsResponse defines model for GetEventApplicantsResponse.
type GetEventApplicantsResponse = []Applicant

// GetEventParticipantsResponse defines model for GetEventParticipantsResponse.
type GetEventParticipantsResponse = []string

// GetEventResponse defines model for GetEventResponse.
type GetEventResponse struct {
	Date        *DateTimeResponse  `json:"date,omitempty"`
	DateOptions []DateOption       `json:"dateOptions"`
	Description string             `json:"description"`
	Id          openapi_types.UUID `json:"id"`
	IsConfirmed bool               `json:"isConfirmed"`
	Title       string             `json:"title"`
}

// GetEventTargetsResponse defines model for GetEventTargetsResponse.
type GetEventTargetsResponse = []string

// GetMeResponse defines model for GetMeResponse.
type GetMeResponse struct {
	TraQID *string `json:"traQID,omitempty"`
}

// GetMyParticipateEvent defines model for GetMyParticipateEvent.
type GetMyParticipateEvent struct {
	Date        DateTimeResponse   `json:"date"`
	Description string             `json:"description"`
	Id          openapi_types.UUID `json:"id"`
	Title       string             `json:"title"`
}

// GetMyParticipateEventsResponse defines model for GetMyParticipateEventsResponse.
type GetMyParticipateEventsResponse = []GetMyParticipateEvent

// GetTraQGroupsResponse defines model for GetTraQGroupsResponse.
type GetTraQGroupsResponse = []TraQGroup

// GetTraQUsersResponse defines model for GetTraQUsersResponse.
type GetTraQUsersResponse = []TraQUser

// PatchEventConfirmRequest defines model for PatchEventConfirmRequest.
type PatchEventConfirmRequest struct {
	EventDateOptionID *openapi_types.UUID `json:"eventDateOptionID,omitempty"`
	IsConfirmed       bool                `json:"isConfirmed"`
}

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

// Applicant defines model for applicant.
type Applicant struct {
	DateOptionIDs *[]openapi_types.UUID `json:"dateOptionIDs,omitempty"`
	TraqID        *string               `json:"traqID,omitempty"`
}

// DateOption defines model for dateOption.
type DateOption struct {
	End   time.Time          `json:"end"`
	Id    openapi_types.UUID `json:"id"`
	Start time.Time          `json:"start"`
}

// GetAllEventsElement defines model for getAllEventsElement.
type GetAllEventsElement struct {
	End         *time.Time         `json:"end,omitempty"`
	Id          openapi_types.UUID `json:"id"`
	IsConfirmed bool               `json:"isConfirmed"`
	Start       *time.Time         `json:"start,omitempty"`
	Title       string             `json:"title"`
}

// TraQGroup defines model for traQGroup.
type TraQGroup struct {
	Members *[]TraQUser `json:"members,omitempty"`
	Name    *string     `json:"name,omitempty"`
}

// TraQUser defines model for traQUser.
type TraQUser struct {
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
}

// EventID defines model for eventID.
type EventID = openapi_types.UUID

// GetEventsAllParams defines parameters for GetEventsAll.
type GetEventsAllParams struct {
	IncludePastEvents *bool `form:"includePastEvents,omitempty" json:"includePastEvents,omitempty"`
}

// PostEventsJSONRequestBody defines body for PostEvents for application/json ContentType.
type PostEventsJSONRequestBody = PostEventRequest

// PostEventsEventIDApplicantsJSONRequestBody defines body for PostEventsEventIDApplicants for application/json ContentType.
type PostEventsEventIDApplicantsJSONRequestBody = PostEventApplicantsRequest

// PatchEventsEventIDConfirmJSONRequestBody defines body for PatchEventsEventIDConfirm for application/json ContentType.
type PatchEventsEventIDConfirmJSONRequestBody = PatchEventConfirmRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /events)
	PostEvents(ctx echo.Context) error

	// (GET /events/all)
	GetEventsAll(ctx echo.Context, params GetEventsAllParams) error

	// (GET /events/me)
	GetEventsMe(ctx echo.Context) error

	// (GET /events/me/participate)
	GetEventsMeParticipate(ctx echo.Context) error

	// (GET /events/{eventID})
	GetEventsEventID(ctx echo.Context, eventID EventID) error

	// (GET /events/{eventID}/applicants)
	GetEventsEventIDApplicants(ctx echo.Context, eventID EventID) error

	// (POST /events/{eventID}/applicants)
	PostEventsEventIDApplicants(ctx echo.Context, eventID EventID) error

	// (PATCH /events/{eventID}/confirm)
	PatchEventsEventIDConfirm(ctx echo.Context, eventID EventID) error

	// (GET /events/{eventID}/participants)
	GetEventsEventIDParticipants(ctx echo.Context, eventID EventID) error

	// (GET /events/{eventID}/targets)
	GetEventsEventIDTargets(ctx echo.Context, eventID EventID) error

	// (GET /me)
	GetMe(ctx echo.Context) error

	// (GET /ping)
	GetPing(ctx echo.Context) error

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

// GetEventsAll converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventsAll(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetEventsAllParams
	// ------------- Optional query parameter "includePastEvents" -------------

	err = runtime.BindQueryParameter("form", true, false, "includePastEvents", ctx.QueryParams(), &params.IncludePastEvents)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter includePastEvents: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventsAll(ctx, params)
	return err
}

// GetEventsMe converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventsMe(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventsMe(ctx)
	return err
}

// GetEventsMeParticipate converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventsMeParticipate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventsMeParticipate(ctx)
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

// PatchEventsEventIDConfirm converts echo context to params.
func (w *ServerInterfaceWrapper) PatchEventsEventIDConfirm(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventID" -------------
	var eventID EventID

	err = runtime.BindStyledParameterWithOptions("simple", "eventID", ctx.Param("eventID"), &eventID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchEventsEventIDConfirm(ctx, eventID)
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

// GetMe converts echo context to params.
func (w *ServerInterfaceWrapper) GetMe(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMe(ctx)
	return err
}

// GetPing converts echo context to params.
func (w *ServerInterfaceWrapper) GetPing(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPing(ctx)
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
	router.GET(baseURL+"/events/all", wrapper.GetEventsAll)
	router.GET(baseURL+"/events/me", wrapper.GetEventsMe)
	router.GET(baseURL+"/events/me/participate", wrapper.GetEventsMeParticipate)
	router.GET(baseURL+"/events/:eventID", wrapper.GetEventsEventID)
	router.GET(baseURL+"/events/:eventID/applicants", wrapper.GetEventsEventIDApplicants)
	router.POST(baseURL+"/events/:eventID/applicants", wrapper.PostEventsEventIDApplicants)
	router.PATCH(baseURL+"/events/:eventID/confirm", wrapper.PatchEventsEventIDConfirm)
	router.GET(baseURL+"/events/:eventID/participants", wrapper.GetEventsEventIDParticipants)
	router.GET(baseURL+"/events/:eventID/targets", wrapper.GetEventsEventIDTargets)
	router.GET(baseURL+"/me", wrapper.GetMe)
	router.GET(baseURL+"/ping", wrapper.GetPing)
	router.GET(baseURL+"/traq/groups", wrapper.GetTraqGroups)
	router.GET(baseURL+"/traq/users", wrapper.GetTraqUsers)

}
