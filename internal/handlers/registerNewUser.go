package handlers

import (
	"context"
	"house-store/internal/dto"
	"house-store/internal/entity"
	"house-store/internal/utilities/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *registerHandler) RegisterNewUser(c echo.Context) error {
	var reqDTO dto.RegisterUser_Request
	if err := c.Bind(&reqDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(reqDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := convertDTOTOEntity_RegisterUser(reqDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
	}

	resultUser, err := h.repo.User_Create(context.TODO(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, convertEntityToDTO_RegisterUser(resultUser))
}

func convertDTOTOEntity_RegisterUser(reqDTO dto.RegisterUser_Request) (entity.User, error) {
	passwordHash, err := auth.GetHash(reqDTO.Password)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		Role:         reqDTO.UserType,
		Email:        reqDTO.Email,
		PasswordHash: passwordHash,
	}, nil
}

func convertEntityToDTO_RegisterUser(user entity.User) dto.RegisterUser_Response {
	return dto.RegisterUser_Response{
		UserUUID: user.UUID,
	}
}
