package repository

import (
	"context"
	"database/sql"
	"house-store/internal/entity"
	"house-store/internal/enum"
	"house-store/internal/model"
)

func (repo *repository) Flat_Update(ctx context.Context, flat entity.Flat) (entity.Flat, error) {
	tx, err := repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return entity.Flat{}, nil
	}
	defer tx.Rollback()

	model := convertEntityToModel_FlatUpdate(flat)

	err = updateFlatStatus(tx, model)
	if err != nil {
		return entity.Flat{}, err
	}

	resultFlat, err := getFlatById(tx, int64(flat.FlatId))
	if err != nil {
		return entity.Flat{}, err
	}

	err = tx.Commit()
	if err != nil {
		return entity.Flat{}, err
	}

	return convertModelToEntity_FlatUpdate(resultFlat), nil
}

func updateFlatStatus(tx *sql.Tx, flat model.Flat) error {
	query := `UPDATE flats
				SET status = $1			
				WHERE flatId = $2 AND status = $3`

	_, err := tx.Exec(query, flat.Status, flat.FlatId, enum.FLAT_STATUS_CREATED)
	return err
}

func convertEntityToModel_FlatUpdate(flat entity.Flat) model.Flat {
	return model.Flat{
		FlatId: flat.FlatId,
		Status: flat.Status,
	}
}

func convertModelToEntity_FlatUpdate(flat model.Flat) entity.Flat {
	return convertModelToEntity_Flat(flat)
}
