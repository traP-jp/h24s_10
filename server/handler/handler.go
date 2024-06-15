package handler

import (
	"github.com/traP-jp/h24s_10/model"
	"github.com/traP-jp/h24s_10/traqclient"
)

type Handler struct {
	repo   *model.Repository
	client *traqclient.Client
}

func New(repo *model.Repository, client *traqclient.Client) *Handler {
	return &Handler{
		repo:   repo,
		client: client,
	}
}
