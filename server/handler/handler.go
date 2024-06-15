package handler

import (
	"os"

	"github.com/traP-jp/h24s_10/model"
)

var ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")

type Handler struct {
	repo  *model.Repository
	repo2 *model.Repository2
}

func New(repo *model.Repository, repo2 *model.Repository2) *Handler {
	return &Handler{
		repo:  repo,
		repo2: repo2,
	}
}
