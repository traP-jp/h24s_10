package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_10/api"
)

func (h *Handler) GetMe(ctx echo.Context) error {
	traQID := ctx.Get(traQIDKey).(string)
	user := api.GetMeResponse{
		TraQID: &traQID,
	}
	return ctx.JSON(http.StatusOK, user)
}