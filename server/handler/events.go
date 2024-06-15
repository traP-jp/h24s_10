package handler

import (
	"net/http"

	"github.com/traP-jp/h24s_10/api"

	"github.com/labstack/echo/v4"
)

// (POST /events)
func (h *Handler) PostEvents(ctx echo.Context) error {
	return nil
}

// (GET /events/{eventID})
func (h *Handler) GetEventsEventID(ctx echo.Context, eventID api.EventID) error {
	return nil
}

// (GET /events/{eventID}/participants)
func (h *Handler) GetEventsEventIDParticipants(ctx echo.Context, eventID api.EventID) error {
	return nil
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
