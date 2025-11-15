package handlers

import (
	"context"
	"house-store/internal/entity"
)

type flatProvider interface {
	Flat_Create(context.Context, entity.Flat) (entity.Flat, error)
	Flat_Update()
}

type flatHandler struct {
	repo flatProvider
}

func NewFlatHandler(repo flatProvider) *flatHandler {
	return &flatHandler{repo: repo}
}
