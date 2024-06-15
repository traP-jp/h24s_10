package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/traP-jp/h24s_10/api"
	"github.com/traP-jp/h24s_10/model"

	"github.com/labstack/echo/v4"
)

// (GET /events/me/participate)
func (h *Handler) GetEventsMeParticipate(ctx echo.Context) error {
	traQID := ctx.Get(traQIDKey).(string)

	participateEventsDetail, err := h.repo.GetParticipateEventsDetailByTraQID(traQID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	res := make(api.GetMyParticipateEventsResponse, 0, len(participateEventsDetail))
	for _, event := range participateEventsDetail {
		date := api.DateTimeResponse{}
		if event.Start.Valid && event.End.Valid {
			date.Start = event.Start.Time
			date.End = event.End.Time
		}
		res = append(res, api.GetMyParticipateEvent{
			Date:        date,
			Description: event.Description,
			Id:          event.ID,
			Title:       event.Title,
		})
	}
	return ctx.JSON(http.StatusOK, res)
}

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

	dateVotes := []model.DateVote{}
	for _, dateOption := range *req.DateOptionIDs {
		id, err := uuid.NewV7()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		traQID := ctx.Get(traQIDKey).(string)
		dateVotes = append(dateVotes, model.DateVote{
			ID:      id,
			EventID: eventID,
			TraQID:  traQID,
			DateID:  dateOption,
		})
	}

	err := h.repo.CreateDateVotes(dateVotes)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusCreated)
}
