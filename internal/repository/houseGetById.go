package repository

import (
	"context"
	"database/sql"
	"house-store/internal/entity"
	"house-store/internal/enum"
	"house-store/internal/model"
)

// Для модераторов возвращаются квартиры в любом статусе [created, approved, declined, on moderation]
func (repo *repository) House_GetById_Moderator(ctx context.Context, houseId int) ([]entity.Flat, error) {
	tx, err := repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return []entity.Flat{}, err
	}
	defer tx.Rollback()

	resultFlats, err := getAllFlatsByHouseId(tx, houseId)
	if err != nil {
		return []entity.Flat{}, err
	}

	err = tx.Commit()
	if err != nil {
		return []entity.Flat{}, nil
	}

	return convertModelToEntity_Flats(resultFlats), nil
}

func getAllFlatsByHouseId(tx *sql.Tx, houseId int) ([]model.Flat, error) {
	query := `SELECT flatId, houseId, price, rooms, status
				FROM flats
				WHERE houseid = $1`
	return getFlats_WithQueryAndParams(tx, query, houseId)
}

// Для обычных пользователей возвращаются только квартиры в статусе approved
func (repo *repository) House_GetById_Client(ctx context.Context, houseId int) ([]entity.Flat, error) {
	tx, err := repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return []entity.Flat{}, err
	}
	defer tx.Rollback()

	resultFlats, err := getAllApprovedFlatsByHouseId(tx, houseId)
	if err != nil {
		return []entity.Flat{}, err
	}

	err = tx.Commit()
	if err != nil {
		return []entity.Flat{}, nil
	}

	return convertModelToEntity_Flats(resultFlats), nil
}

func getAllApprovedFlatsByHouseId(tx *sql.Tx, houseId int) ([]model.Flat, error) {
	query := `SELECT flatId, houseId, price, rooms, status
				FROM flats
				WHERE houseid = $1 AND status = $2`
	return getFlats_WithQueryAndParams(tx, query, houseId, enum.FLAT_STATUS_APPROVED)
}

func getFlats_WithQueryAndParams(tx *sql.Tx, query string, params ...interface{}) ([]model.Flat, error) {
	rows, err := tx.Query(query, params...)
	if err != nil {
		return []model.Flat{}, err
	}
	defer rows.Close()

	var (
		resultFlats = make([]model.Flat, 0, 10)
		currentFlat = model.Flat{}
	)

	for rows.Next() {
		err := rows.Scan(
			&currentFlat.FlatId,
			&currentFlat.HouseId,
			&currentFlat.Price,
			&currentFlat.Rooms,
			&currentFlat.Status,
		)
		if err != nil {
			return []model.Flat{}, err
		}
		resultFlats = append(resultFlats, currentFlat)
	}

	err = rows.Err()
	if err != nil {
		return []model.Flat{}, err
	}

	return resultFlats, nil
}

func convertModelToEntity_Flats(flats []model.Flat) []entity.Flat {
	result := make([]entity.Flat, 0, len(flats))
	for _, flat := range flats {
		result = append(result, convertModelToEntity_Flat(flat))
	}
	return result
}

func convertModelToEntity_Flat(from model.Flat) entity.Flat {
	return entity.Flat{
		FlatId:  from.FlatId,
		HouseId: from.HouseId,
		Price:   from.Price,
		Rooms:   from.Rooms,
		Status:  from.Status,
	}
}
