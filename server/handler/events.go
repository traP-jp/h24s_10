package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/traP-jp/h24s_10/api"
	"github.com/traP-jp/h24s_10/model"

	"github.com/labstack/echo/v4"
)

// (POST /events)
func (h *Handler) PostEvents(ctx echo.Context) error {
	var req api.PostEventRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	hostID := ctx.Get(traQIDKey).(string)
	eventID, err := uuid.NewV7()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	event := model.Event{
		ID:          eventID,
		Title:       req.Title,
		HostID:      hostID,
		Description: req.Description,
		IsConfirmed: false,
	}
	err = h.repo.CreateEvent(event)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	options := make([]model.EventDate, len(req.DateOptions))
	for i, option := range req.DateOptions {
		id, err := uuid.NewV7()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		options[i] = model.EventDate{
			ID:      id,
			EventID: eventID,
			Start:   option.Start,
			End:     option.End,
		}
	}
	err = h.repo.CreateEventDates(options)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := api.PostEventResponse{
		Id: eventID,
	}

	return ctx.JSON(http.StatusCreated, res)
}

// (GET /events/{eventID})
func (h *Handler) GetEventsEventID(ctx echo.Context, eventID api.EventID) error {
	event, err := h.repo.GetEventByEventID(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	dateOptionResponse, err := h.repo.GetEventDates(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	var dateOptionsResponse []api.DateOption
	for _, dateOption := range dateOptionResponse {
		dateOptionsResponse = append(dateOptionsResponse, api.DateOption{
			End:   dateOption.End,
			Id:    dateOption.ID,
			Start: dateOption.Start,
		})
	}
	getEventsByEventIDResponse := api.GetEventResponse{
		DateOptions: &dateOptionsResponse,
		Description: event.Description,
		Id:          &event.ID,
		IsConfirmed: &event.IsConfirmed,
		Title:       event.Title,
	}
	if event.Start.Valid && event.End.Valid {
		getEventsByEventIDResponse.Date = &api.DateTimeResponse{
			End:   event.End.Time,
			Start: event.Start.Time,
		}
	}
	return ctx.JSON(http.StatusOK, getEventsByEventIDResponse)
}

// (GET /events/{eventID}/participants)
func (h *Handler) GetEventsEventIDParticipants(ctx echo.Context, eventID api.EventID) error {
	participants, err := h.repo.GetParticipants(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	res := make(api.GetEventParticipantsResponse, len(participants))
	for i, participant := range participants {
		res[i] = participant.TraQID
	}
	return ctx.JSON(http.StatusOK, res)
}

// (GET /events/{eventID}/targets)
func (h *Handler) GetEventsEventIDTargets(ctx echo.Context, eventID api.EventID) error {
	targets, err := h.repo.GetEventTargets(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	res := make(api.GetEventTargetsResponse, len(targets))
	for i, target := range targets {
		res[i] = target.TraQID
	}
	return ctx.JSON(http.StatusOK, res)
}
