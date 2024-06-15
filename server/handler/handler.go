package handler

import (
	"github.com/traP-jp/h24s_10/model"
)

type Handler struct {
	repo *model.Repository
}

func New(repo *model.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}
