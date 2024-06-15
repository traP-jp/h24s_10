package handler

import (
	"database/sql"
	"net/http"
	"slices"
	"time"

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

	// create date options
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

	// create targets
	targets := make([]model.EventTarget, len(req.Targets))
	for i, target := range req.Targets {
		targets[i] = model.EventTarget{
			ID:      uuid.New(),
			EventID: eventID,
			TraQID:  target,
		}
	}
	err = h.repo.CreateEventTargets(targets)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := api.PostEventResponse{
		Id: eventID,
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (h *Handler) GetEventsAll(ctx echo.Context, params api.GetEventsAllParams) error {
	events, err := h.repo.GetEvents()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var omitPastEvents bool
	if params.OmitPastEvents != nil {
		omitPastEvents = *params.OmitPastEvents
	} else {
		omitPastEvents = true
	}

	res := make(api.GetAllEventsResponse, 0, len(events))
	for _, event := range events {
		if omitPastEvents && event.End.Valid && event.End.Time.Before(time.Now()) {
			continue
		}
		e := api.GetAllEventsElement{
			Id:          event.ID,
			Title:       event.Title,
			IsConfirmed: event.IsConfirmed,
		}
		if event.Start.Valid && event.End.Valid {
			e.Start = &event.Start.Time
			e.End = &event.End.Time
		}
		res = append(res, e)
	}

	return ctx.JSON(http.StatusOK, res)
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

// (PATCH /events/{eventID}/confirm)
func (h *Handler) PatchEventsEventIDConfirm(ctx echo.Context, eventID api.EventID) error {
	var req api.PatchEventConfirmRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// check if the user is the host of the event
	event, err := h.repo.GetEventByEventID(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	userID := ctx.Get(traQIDKey).(string)
	if event.HostID != userID {
		return echo.NewHTTPError(http.StatusForbidden, "you are not the host of this event")
	}

	event.IsConfirmed = req.IsConfirmed

	if req.IsConfirmed {
		if req.EventDateOptionID == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "event_date_option_id is required")
		}
		err = h.makeConfirmed(eventID, *req.EventDateOptionID, event)
	} else {
		err = h.makeUnconfirmed(event)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *Handler) makeConfirmed(eventID uuid.UUID, eventDateID uuid.UUID, event model.Event) error {
	// get all event dates of the event
	eventDates, err := h.repo.GetEventDates(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	// find the event date option which the host selected
	idx := slices.IndexFunc(eventDates, func(d model.EventDate) bool {
		return d.ID == eventDateID
	})
	if idx == -1 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid event date option")
	}
	eventDate := eventDates[idx]

	// get all votes for the event date
	votes, err := h.repo.GetEventDateVotes(eventDate.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// create participants
	participants := make([]model.Participant, len(votes))
	for i, vote := range votes {
		id, err := uuid.NewV7()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		participants[i] = model.Participant{
			ID:      id,
			EventID: eventID,
			TraQID:  vote.TraQID,
			TeamID:  1,
		}
	}
	err = h.repo.CreateParticipants(participants)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	event.Start = sql.NullTime{Time: eventDate.Start, Valid: true}
	event.End = sql.NullTime{Time: eventDate.End, Valid: true}

	return h.repo.UpdateEvent(event)
}

func (h *Handler) makeUnconfirmed(event model.Event) error {
	event.IsConfirmed = false
	event.Start = sql.NullTime{}
	event.End = sql.NullTime{}
	return h.repo.UpdateEvent(event)
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
