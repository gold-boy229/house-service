package repository

import (
	"context"
	"database/sql"
	"fmt"
	"house-store/internal/entity"
	"house-store/internal/model"
)

func (repo *repository) House_Create(ctx context.Context, house entity.House) (entity.House, error) {
	tx, err := repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return entity.House{}, err
	}
	defer tx.Rollback()

	model := convertEntiryToModel_House(house)

	newHouseId, err := insertRowAndGetID_House(tx, model)
	if err != nil {
		return entity.House{}, err
	}

	resultModel, err := getHouseById(tx, newHouseId)
	if err != nil {
		return entity.House{}, err
	}

	err = tx.Commit()
	if err != nil {
		return entity.House{}, err
	}

	return convertModelToEntity_House(resultModel), nil
}

func insertRowAndGetID_House(tx *sql.Tx, house model.House) (int64, error) {
	query := `INSERT INTO houses(address, year, developer) 
				VALUES ($1, $2, $3)
				RETURNING houseId`
	var id int64
	err := tx.QueryRow(query, house.Address, house.Year, house.Developer).Scan(&id)
	if err != nil {
		fmt.Printf("insertHouse. model = %+v; err = %v\n", house, err)
		return 0, err
	}

	return id, nil
}

func getHouseById(tx *sql.Tx, id int64) (model.House, error) {
	query := `SELECT houseId, address, year, developer, created_at, updated_at
			FROM houses
			WHERE houseId = $1`

	var h model.House
	err := tx.QueryRow(query, id).Scan(
		&h.Id,
		&h.Address,
		&h.Year,
		&h.Developer,
		&h.CreatedAt,
		&h.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.House{}, fmt.Errorf("getHouseById. There is no row in houses table with houseId = %v", id)
		}
		return model.House{}, fmt.Errorf("getHouseById. err = %v", err)
	}
	return h, nil
}

func convertEntiryToModel_House(from entity.House) model.House {
	return model.House{
		Address:   from.Address,
		Year:      from.Year,
		Developer: sql.NullString{String: from.Developer, Valid: from.Developer != ""},
	}
}

func convertModelToEntity_House(from model.House) entity.House {
	return entity.House{
		Id:        from.Id,
		Address:   from.Address,
		Year:      from.Year,
		Developer: from.Developer.String,
		CreatedAt: from.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: from.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
