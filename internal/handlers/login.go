package handlers

import (
	"context"
	"house-store/internal/dto"
	"house-store/internal/entity"
	"house-store/internal/utilities/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userProvider interface {
	GetUserById(ctx context.Context, userUUID string) (entity.User, error)
}

type loginHandler struct {
	repo userProvider
}

func NewLoginHandler(repo userProvider) *loginHandler {
	return &loginHandler{repo: repo}
}

func (h *loginHandler) Login(c echo.Context) error {
	var reqDTO dto.Login_Request
	if err := c.Bind(&reqDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(reqDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	userLoginData := convertDTOToEntity_Login(reqDTO)
	resultUser, err := h.repo.GetUserById(context.TODO(), userLoginData.UUID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
	}

	userExistence, err := buildUserExistence(resultUser, userLoginData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
	}
	if !userExistence.Exists {
		return c.String(http.StatusNotFound, "Invalid user_UUID")
	}
	if !userExistence.IsPasswordCorrect {
		return c.String(http.StatusBadRequest, "Invalid password")
	}
	userRole := userExistence.Role

	// want to create JWT token with claims {UserRole + UserUUID}
	token, err := auth.CreateToken(userRole)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.Login_Response{Token: token})
}

func buildUserExistence(resultUser entity.User, userLoginData entity.UserLoginData) (entity.UserExistence, error) {
	if resultUser.UUID != userLoginData.UUID {
		return entity.UserExistence{}, nil
	}

	isPasswordCorrect, err := auth.CheckPasswordCorrectness([]byte(resultUser.PasswordHash), []byte(userLoginData.Password))
	if err != nil {
		return entity.UserExistence{}, err
	}

	return entity.UserExistence{
		Exists:            true,
		IsPasswordCorrect: isPasswordCorrect,
		Role:              resultUser.Role,
	}, nil
}

func convertDTOToEntity_Login(reqDTO dto.Login_Request) entity.UserLoginData {
	return entity.UserLoginData{
		UUID:     reqDTO.UserUUID,
		Password: reqDTO.Password,
	}
}
