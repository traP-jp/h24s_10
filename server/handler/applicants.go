package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/traP-jp/h24s_10/api"

	"github.com/labstack/echo/v4"
)

// (GET /events/{eventID}/applicants)
func (h *Handler) GetEventsEventIDApplicants(ctx echo.Context, eventID api.EventID) error {
	eventDates, err := h.repo.GetEventDates(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	dateOptions := make(map[string][]uuid.UUID)
	for _, date := range eventDates {
		dateVotes, err := h.repo.GetEventDateVotes(date.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		for _, vote := range dateVotes {
			dateOptions[vote.TraQID] = append(dateOptions[vote.TraQID], vote.DateID)
		}
	}

	res := make(api.GetEventApplicantsResponse, 0, len(dateOptions))
	for traQID, dateIDs := range dateOptions {
		res = append(res, api.Applicant{
			TraqID: &traQID,
			DateOptionIDs: &dateIDs,
		})
	}
	return ctx.JSON(http.StatusOK, res)
}

// (POST /events/{eventID}/applicants)
func (h *Handler) PostEventsEventIDApplicants(ctx echo.Context, eventID api.EventID) error {
	return nil
}
