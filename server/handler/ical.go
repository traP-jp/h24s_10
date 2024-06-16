package handler

import (
	"net/http"
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

	cal := generateIcal(event)

	return ctx.Blob(http.StatusOK, "text/calendar", []byte(cal))
}

func generateIcal(e model.Event) string {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)

	event := cal.AddEvent(uuid.NewString())
	event.SetCreatedTime(time.Now())
	event.SetStartAt(e.Start.Time)
	event.SetEndAt(e.End.Time)
	event.SetSummary(e.Title)
	event.SetDescription(e.Description)
	event.SetLocation(e.Location)

	return cal.Serialize()
}
