package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_10/api"
	"github.com/traPtitech/go-traq"
)

// (GET /traq/groups)
func (h *Handler) GetTraqGroups(ctx echo.Context) error {
	if ACCESS_TOKEN == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("access token is missing"))
	}
	ctxWithToken := context.WithValue(ctx.Request().Context(), traq.ContextAccessToken, ACCESS_TOKEN)
	groupList, err := h.repo2.GetUserGroups(ctxWithToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	userMap, err := h.repo2.GetUsersMap(ctxWithToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	groupListResponse := api.GetTraQGroupsResponse{}
	for _, group := range groupList {
		var groupMember []api.TraQUser
		for _, member := range group.Members {
			if _, ok := userMap[member.Id]; !ok {
				continue
			}
			if userMap[member.Id].Bot {
				continue
			}
			groupMember = append(groupMember, api.TraQUser{
				DisplayName: userMap[member.Id].DisplayName,
				Name:        userMap[member.Id].Name,
			})
		}
		groupListResponse = append(groupListResponse, api.TraQGroup{
			Members: &groupMember,
			Name:    &group.Name,
		})
	}
	return ctx.JSON(http.StatusOK, groupListResponse)
}

// (GET /traq/users)
func (h *Handler) GetTraqUsers(ctx echo.Context) error {
	if ACCESS_TOKEN == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("access token is missing"))
	}
	userList, err := h.repo2.GetUsers(context.WithValue(ctx.Request().Context(), traq.ContextAccessToken, ACCESS_TOKEN))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userListResponse := api.GetTraQUsersResponse{}
	for _, user := range userList {
		if user.Bot {
			continue
		}
		userListResponse = append(userListResponse, api.TraQUser{
			DisplayName: user.DisplayName,
			Name:        user.Name,
		})
	}
	return ctx.JSON(http.StatusOK, userListResponse)
}
