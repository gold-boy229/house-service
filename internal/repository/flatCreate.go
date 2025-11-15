package repository

import (
	"context"
	"house-store/internal/entity"
)

func (repo *repository) Flat_Create(ctx context.Context, flat entity.Flat) (entity.Flat, error) {
	return entity.Flat{}, nil
}
