package handler

import (
	"fmt"
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

	applicants, err := h.repo.GetApplicants(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	dateOptions := make(map[string][]uuid.UUID, len(applicants))
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
	for _, applicant := range applicants {
		comment, err := h.repo.GetComment(eventID, applicant.TraQID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		dateOptionIDs, ok := dateOptions[applicant.TraQID]
		if !ok {
			dateOptionIDs = make([]uuid.UUID, 0)
		}
		res = append(res, api.Applicant{
			TraqID:        applicant.TraQID,
			DateOptionIDs: dateOptionIDs,
			Comment:       comment.Content,
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
	err := h.repo.ValidateEventDateIDsFromEventID(eventID, req.DateOptionIDs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid event ID and DateOptionIDs: %v", err))
	}

	traQID := ctx.Get(traQIDKey).(string)

	// traqIDがeventDateIDのそれぞれの値が齟齬がないか確認
	err = h.repo.ValidateEventDateIDsFromTraqID(traQID, req.DateOptionIDs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("already exists traqID and DateOptionIDs: %v", err))
	}

	applicantID, err := uuid.NewV7()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	err = h.repo.CreateApplicant(model.Applicant{
		ID:      applicantID,
		EventID: eventID,
		TraQID:  traQID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	dateVotes := make([]model.DateVote, 0, len(req.DateOptionIDs))
	for _, dateOption := range req.DateOptionIDs {
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

	commentID, err := uuid.NewV7()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	h.repo.CreateComment(model.Comment{
		ID:      commentID,
		EventID: eventID,
		TraQID:  traQID,
		Content: req.Comment,
	})

	return ctx.NoContent(http.StatusCreated)
}
