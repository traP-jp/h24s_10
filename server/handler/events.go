package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"slices"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/traP-jp/h24s_10/api"
	"github.com/traP-jp/h24s_10/model"
	"github.com/traP-jp/h24s_10/traqclient"
	"golang.org/x/sync/errgroup"

	"github.com/labstack/echo/v4"
)

// (POST /events)
func (h *Handler) PostEvents(ctx echo.Context) error {
	var req api.PostEventRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// check if exist same pair of date start and date end
	for i, dateOption := range req.DateOptions {
		for j := i + 1; j < len(req.DateOptions); j++ {
			if dateOption.Start == req.DateOptions[j].Start && dateOption.End == req.DateOptions[j].End {
				return echo.NewHTTPError(http.StatusBadRequest, "date options must be unique")
			}
		}
	}

	// check if the targets are valid
	eg, c := errgroup.WithContext(ctx.Request().Context())
	eg.SetLimit(10)
	for _, target := range req.Targets {
		eg.Go(func() error {
			_, err := h.client.GetUser(c, target)
			if err == traqclient.ErrUserNotFound {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user %s not found", target))
			} else if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	hostID := ctx.Get(traQIDKey).(string)
	eventID, err := uuid.NewV7()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	var location string
	if req.Location != nil {
		location = *req.Location
	}

	event := model.Event{
		ID:          eventID,
		Title:       req.Title,
		HostID:      hostID,
		Description: req.Description,
		Location:    location,
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

	var includePathEvents bool
	if params.IncludePastEvents != nil {
		includePathEvents = *params.IncludePastEvents
	} else {
		includePathEvents = true
	}

	res := make(api.GetAllEventsResponse, 0, len(events))
	for _, event := range events {
		if !includePathEvents && event.End.Valid && event.End.Time.Before(time.Now()) {
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

// (GET /events/me)
func (h *Handler) GetEventsMe(ctx echo.Context) error {
	userID := ctx.Get(traQIDKey).(string)
	// ホストとなっているイベントの取得
	eventsByHost, err := h.repo.GetEventsByHost(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// 回答候補となっているイベントの取得
	eventsByTarget, err := h.repo.GetEventsByTargetUser(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// ホストとなっていないイベントのみに限定
	removedHostEvent := removeEventHostFromTarget(eventsByHost, eventsByTarget)

	// 回答済みのイベントの取得
	eventAnsweredStruct, err := h.repo.GetDateVotesByUser(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	eventAnsweredMap := make(map[uuid.UUID]struct{})
	for _, e := range eventAnsweredStruct {
		eventAnsweredMap[e.EventID] = struct{}{}
		log.Printf("eventAnsweredMap: %v", e.EventID)
	}

	res := make([]api.EventMeResponse, 0, len(eventsByHost)+len(removedHostEvent))
	for _, event := range eventsByHost {
		_, isAnswered := eventAnsweredMap[event.ID]
		res = append(res, api.EventMeResponse{
			Description: &event.Description,
			EventId:     event.ID,
			IsAnswered:  isAnswered,
			IsConfirmed: event.IsConfirmed,
			IsHost:      true,
			Title:       event.Title,
		})
	}
	for _, event := range removedHostEvent {
		_, isAnswered := eventAnsweredMap[event.ID]
		res = append(res, api.EventMeResponse{
			Description: &event.Description,
			EventId:     event.ID,
			IsAnswered:  isAnswered,
			IsConfirmed: event.IsConfirmed,
			IsHost:      false,
			Title:       event.Title,
		})
	}
	return ctx.JSON(http.StatusOK, res)
}

func removeEventHostFromTarget(hostEvent, targetEvent []model.Event) []model.Event {
	removedEvent := []model.Event{}
	for i := 0; i < len(targetEvent); i++ {
		check := false
		for j := 0; j < len(hostEvent); j++ {
			if targetEvent[i].ID == hostEvent[j].ID {
				check = true
				break
			}
		}
		if !check {
			removedEvent = append(removedEvent, targetEvent[i])
		}
	}
	return removedEvent
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
	var location *string
	if event.Location != "" {
		location = &event.Location
	}
	getEventsByEventIDResponse := api.GetEventResponse{
		DateOptions: dateOptionsResponse,
		Description: event.Description,
		Location:    location,
		Id:          event.ID,
		IsConfirmed: event.IsConfirmed,
		Title:       event.Title,
		HostID:      event.HostID,
	}
	if event.Start.Valid && event.End.Valid {
		getEventsByEventIDResponse.Date = &api.DateTimeResponse{
			End:   event.End.Time,
			Start: event.Start.Time,
		}
	}
	return ctx.JSON(http.StatusOK, getEventsByEventIDResponse)
}

// (POST /events/{eventID}/confirm)
func (h *Handler) PostEventsEventIDConfirm(ctx echo.Context, eventID api.EventID) error {
	var req api.PostEventConfirmRequest
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

	// participantsを取得
	participants, err := h.repo.GetParticipants(eventID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// グループ名一覧を取得
	groupsMap, err := h.client.GetUserGroupsMap(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	// グループ名が存在しないものを使う
	groupName, err := editGroupName(event.Title, groupsMap)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	eventDescription := []rune(event.Description)

	// eventDescriptionが100以上のときは100文字までにする
	if utf8.RuneCountInString(string(eventDescription)) > 100 {
		eventDescription = eventDescription[:100]
	}

	group, err := h.client.CreateUserGroup(ctx.Request().Context(), groupName, string(eventDescription), "あのにます", event.HostID, participants)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	groupMembers := make([]api.TraQUser, 0, len(group.Members))
	for _, member := range group.Members {
		groupMembers = append(groupMembers, api.TraQUser{
			DisplayName: member.DisplayName,
			Name:        member.Name,
		})
	}

	groupResponse := api.TraQGroup{
		Name:    &group.Name,
		Members: &groupMembers,
	}

	response := &api.PostEventConfirmResponse{
		Group: groupResponse,
	}

	return ctx.JSON(http.StatusOK, response)
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

func editGroupName(name string, groupsMap map[string]traqclient.Group) (string, error) {
	runeName := []rune(name)
	if utf8.RuneCountInString(string(runeName)) > 30 {
		runeName = runeName[:30]
	}
	originalName := string(runeName)
	// 既存のグループ名でない場合、元の名前を返す
	if _, ok := groupsMap[originalName]; !ok {
		return originalName, nil
	}
	// 新しいグループ名を生成する
	for i := 0; i < 10; i++ {
		if len(runeName) < 30 {
			// runeNameが30文字未満の場合は末尾に追加
			runeName = append(runeName, rune('0'+i%10))
		} else {
			// それ以外の場合は最後の文字を変更
			runeName[29] = rune('0' + i%10)
		}
		newName := string(runeName)
		if _, ok := groupsMap[newName]; !ok {
			return newName, nil
		}
	}
	return "", fmt.Errorf("group name error")
}
