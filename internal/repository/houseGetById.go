package repository

import (
	"context"
	"house-store/internal/entity"
)

// Для модераторов возвращаются квартиры в любом статусе [created, approved, declined, on moderation]
func (repo *repository) House_GetById_Moderator(ctx context.Context, houseId int) ([]entity.Flat, error) {
	return []entity.Flat{}, nil
}

// Для обычных пользователей возвращаются только квартиры в статусе approved
func (repo *repository) House_GetById_Client(ctx context.Context, houseId int) ([]entity.Flat, error) {
	return []entity.Flat{}, nil
}
