package repository

import (
	"context"
	"database/sql"
	"fmt"
	"house-store/internal/entity"
	"house-store/internal/model"
)

func (r *repository) House_GetById() {

}

func (r *repository) House_SubscribeForUpdates() {

}

func (repo *repository) House_Create(ctx context.Context, house entity.House) (entity.House, error) {
	model := convertEntiryToModel_House(house)

	newHouseId, err := repo.insertRowAndGetID_House(model)
	if err != nil {
		return entity.House{}, err
	}

	resultModel, err := repo.getHouseById(newHouseId)
	if err != nil {
		return entity.House{}, fmt.Errorf("House_Create. err = %v", err)
	}

	return convertModelToEntity_House(resultModel), nil
}

func (repo *repository) insertRowAndGetID_House(house model.House) (int64, error) {
	query := `INSERT INTO houses(address, year, developer) 
				VALUES ($1, $2, $3)
				RETURNING houseId`
	var id int64
	err := repo.Db.QueryRow(query, house.Address, house.Year, house.Developer).Scan(&id)
	if err != nil {
		fmt.Printf("insertHouse. model = %+v; err = %v\n", house, err)
		return 0, err
	}

	return id, nil
}

func (repo *repository) getHouseById(id int64) (model.House, error) {
	query := `SELECT houseId, address, year, developer, created_at, updated_at
				FROM houses
				WHERE houseId = $1`

	var h model.House
	err := repo.Db.QueryRow(query, id).Scan(
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
