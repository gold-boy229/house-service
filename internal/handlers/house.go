package handlers

import (
	"context"
	"house-store/internal/entity"
)

type houseProvider interface {
	House_GetById_Client(context.Context, int) ([]entity.Flat, error)
	House_GetById_Moderator(context.Context, int) ([]entity.Flat, error)
	House_SubscribeForUpdates()
	House_Create(context.Context, entity.House) (entity.House, error)
}

type houseHandler struct {
	repo houseProvider
}

func NewHouseHandler(repo houseProvider) *houseHandler {
	return &houseHandler{repo: repo}
}
