package handlers

import (
	"context"
	"house-store/internal/entity"
)

type userCreator interface {
	User_Create(context.Context, entity.User) (entity.User, error)
}

type registerHandler struct {
	repo userCreator
}

func NewRegisterHandler(repo userCreator) *registerHandler {
	return &registerHandler{repo: repo}
}
