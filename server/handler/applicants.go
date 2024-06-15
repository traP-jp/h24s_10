package handler

import (
	"github.com/traP-jp/h24s_10/api"

	"github.com/labstack/echo/v4"
)

// (GET /events/{eventID}/applicants)
func (h *Handler) GetEventsEventIDApplicants(ctx echo.Context, eventID api.EventID) error {
	return nil
}

// (POST /events/{eventID}/applicants)
func (h *Handler) PostEventsEventIDApplicants(ctx echo.Context, eventID api.EventID) error {
	return nil
}
