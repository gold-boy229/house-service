package repository

import (
	"context"
	"database/sql"
	"fmt"
	"house-store/internal/entity"
	"house-store/internal/model"
)

func (repo *repository) User_Create(ctx context.Context, user entity.User) (entity.User, error) {
	tx, err := repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return entity.User{}, err
	}
	defer tx.Rollback()

	model := convertEntityToModel_UserCreate(user)

	newUserUUID, err := insertRowAndGetID_AppUser(tx, model)
	if err != nil {
		return entity.User{}, err
	}

	resultUser := convertModelToEntity_UserCreate(model, newUserUUID)

	err = tx.Commit()
	if err != nil {
		return entity.User{}, err
	}

	return resultUser, nil
}

func insertRowAndGetID_AppUser(tx *sql.Tx, user model.AppUser) (string, error) {
	query := `INSERT INTO app_users(role, email, password_hash) 
			VALUES ($1, $2, $3)
			RETURNING userId`
	var newRowUUID string
	err := tx.QueryRow(query, user.Role, user.Email, user.PasswordHash).Scan(&newRowUUID)
	if err != nil {
		fmt.Printf("insertAppUser. model = %+v; err = %v\n", user, err)
		return "", err
	}
	return newRowUUID, nil
}

func convertEntityToModel_UserCreate(user entity.User) model.AppUser {
	return model.AppUser{
		Role:         user.Role,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}

func convertModelToEntity_UserCreate(user model.AppUser, newUserUUID string) entity.User {
	return entity.User{
		UUID:  newUserUUID,
		Role:  user.Role,
		Email: user.Email,
	}
}
