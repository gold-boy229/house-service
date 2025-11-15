package repository

import (
	"context"
	"database/sql"
	"fmt"
	"house-store/internal/entity"
	"house-store/internal/model"
)

func (repo *repository) Flat_Create(ctx context.Context, flat entity.Flat) (entity.Flat, error) {
	tx, err := repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return entity.Flat{}, err
	}
	defer tx.Rollback()

	flatModel := convertEntityToModel(flat)

	newFlatId, err := insertRowAndGetID_Flat(tx, flatModel)
	if err != nil {
		return entity.Flat{}, err
	}

	err = updateHouseModificationTime(tx, flatModel.HouseId)
	if err != nil {
		return entity.Flat{}, err
	}

	resFlat, err := getFlatById(tx, newFlatId)
	if err != nil {
		return entity.Flat{}, err
	}

	err = tx.Commit()
	if err != nil {
		return entity.Flat{}, err
	}

	return convertModelToEntity(resFlat), nil
}

func insertRowAndGetID_Flat(tx *sql.Tx, flat model.Flat) (int64, error) {
	query := `INSERT INTO flats(houseId, price, rooms)
				VALUES ($1, $2, $3)
				RETURNING flatId`
	var id int64
	err := tx.QueryRow(query, flat.HouseId, flat.Price, flat.Rooms).Scan(&id)
	if err != nil {
		fmt.Printf("insertFlat. model = %+v; err = %v\n", flat, err)
		return 0, err
	}

	return id, nil
}

func updateHouseModificationTime(tx *sql.Tx, houseId int) error {
	query := `UPDATE houses
				SET updated_at = NOW()
				WHERE houseid = $1`
	_, err := tx.Exec(query, houseId)
	return err
}

func getFlatById(tx *sql.Tx, id int64) (model.Flat, error) {
	query := `SELECT flatId, houseId, price, rooms, status
				FROM flats
				WHERE flatId = $1`

	var flat model.Flat
	err := tx.QueryRow(query, id).Scan(
		&flat.FlatId,
		&flat.HouseId,
		&flat.Price,
		&flat.Rooms,
		&flat.Status,
	)
	if err != nil {
		return model.Flat{}, err
	}
	return flat, nil
}

func convertEntityToModel(flat entity.Flat) model.Flat {
	return model.Flat{
		HouseId: flat.HouseId,
		Price:   flat.Price,
		Rooms:   flat.Rooms,
	}
}

func convertModelToEntity(flat model.Flat) entity.Flat {
	return entity.Flat{
		FlatId:  flat.FlatId,
		HouseId: flat.HouseId,
		Price:   flat.Price,
		Rooms:   flat.Rooms,
		Status:  flat.Status,
	}
}
