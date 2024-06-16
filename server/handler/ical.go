package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_10/api"
	"github.com/traP-jp/h24s_10/model"
)

func (h *Handler) GetEventsEventIDCalendar(ctx echo.Context, eventID api.EventID) error {
	event, err := h.repo.GetEventByEventID(eventID)
	if err != nil {
		return err
	}

	if !event.IsConfirmed {
		return echo.NewHTTPError(http.StatusBadRequest, "Event is not confirmed")
	}

	participants, err := h.repo.GetParticipants(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get participants")
	}

	cal := generateIcal(event, participants)

	return ctx.Blob(http.StatusOK, "text/calendar", []byte(cal))
}

func generateIcal(e model.Event, participants []model.Participant) string {
	attendees := make([]string, len(participants))
	for i, p := range participants {
		attendees[i] = p.TraQID
	}
	description := fmt.Sprintf(
		"Host: %s\nAttendees: %s\nDescription: %s",
		e.HostID,
		strings.Join(attendees, ", "),
		e.Description,
	)

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)

	event := cal.AddEvent(uuid.NewString())
	event.SetCreatedTime(time.Now())
	event.SetStartAt(e.Start.Time)
	event.SetEndAt(e.End.Time)
	event.SetSummary(e.Title)
	event.SetDescription(description)
	event.SetLocation(e.Location)

	return cal.Serialize()
}
