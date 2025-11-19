package repository

import (
	"context"
	"database/sql"
	"house-store/internal/entity"
	"house-store/internal/model"
)

func (repo *repository) GetUserById(ctx context.Context, userUUID string) (entity.User, error) {
	tx, err := repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return entity.User{}, err
	}
	defer tx.Rollback()

	resultUser, err := getUserById(tx, userUUID)
	if err != nil {
		return entity.User{}, err
	}

	err = tx.Commit()
	if err != nil {
		return entity.User{}, err
	}

	return convertModelToEntity_User(resultUser), nil
}

func getUserById(tx *sql.Tx, userUUID string) (model.AppUser, error) {
	query := `SELECT userId, role, email, password_hash
				FROM app_users
				WHERE userId = $1`
	var user model.AppUser
	err := tx.QueryRow(query, userUUID).Scan(
		&user.UUID,
		&user.Role,
		&user.Email,
		&user.PasswordHash,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.AppUser{}, nil
		}
		return model.AppUser{}, err
	}
	return user, nil
}

func convertModelToEntity_User(user model.AppUser) entity.User {
	return entity.User{
		UUID:         user.UUID,
		Role:         user.Role,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}
