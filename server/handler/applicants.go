package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/traP-jp/h24s_10/api"
	"github.com/traP-jp/h24s_10/model"

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
			TraqID:        &traQID,
			DateOptionIDs: &dateIDs,
		})
	}
	return ctx.JSON(http.StatusOK, res)
}

// (POST /events/{eventID}/applicants)
func (h *Handler) PostEventsEventIDApplicants(ctx echo.Context, eventID api.EventID) error {
	var req api.PostEventApplicantsRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// eventIDがeventDateIDのそれぞれの値が齟齬がないか確認
	err := h.repo.ValidateEventDateIDsFromEventID(eventID, *req.DateOptionIDs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid event ID and DateOptionIDs: %v", err))
	}

	traQID := ctx.Get(traQIDKey).(string)

	// traqIDがeventDateIDのそれぞれの値が齟齬がないか確認
	err = h.repo.ValidateEventDateIDsFromTraqID(traQID, *req.DateOptionIDs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("already exists traqID and DateOptionIDs: %v", err))
	}

	dateVotes := make([]model.DateVote, 0, len(*req.DateOptionIDs))
	for _, dateOption := range *req.DateOptionIDs {
		id, err := uuid.NewV7()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		dateVotes = append(dateVotes, model.DateVote{
			ID:     id,
			TraQID: traQID,
			DateID: dateOption,
		})
	}
	err = h.repo.CreateDateVotes(dateVotes)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusCreated)
}
