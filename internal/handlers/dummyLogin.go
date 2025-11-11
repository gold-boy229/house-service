package handlers

import (
	"errors"
	"fmt"
	"house-store/internal/dto"
	"house-store/internal/enum"
	"house-store/internal/utilities/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type dummyLoginHandler struct {
}

func NewDummyLoginHandler() *dummyLoginHandler {
	return &dummyLoginHandler{}
}

func (h *dummyLoginHandler) DummyLogin(c echo.Context) error {
	req, err := readRequestParams(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
	}

	err = validateInputData(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
	}

	res, err := createToken(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func readRequestParams(c echo.Context) (dto.DummyLoginRequest, error) {
	var req dto.DummyLoginRequest
	err := c.Bind(&req)
	if err != nil {
		return req, errors.New("bad request. Cannot parse query parameter `user_type`")
	}
	return req, nil
}

func validateInputData(req dto.DummyLoginRequest) error {
	userRole := req.UserType
	if userRole == enum.USER_ROLE_CLIENT || userRole == enum.USER_ROLE_MODERATOR {
		return nil
	}
	return fmt.Errorf("unknown UserType: %q", req.UserType)
}

func createToken(req dto.DummyLoginRequest) (dto.DummyLoginResponse, error) {
	tokenString, err := auth.CreateToken(req.UserType)
	if err != nil {
		return dto.DummyLoginResponse{}, err
	}
	return dto.DummyLoginResponse{Token: tokenString}, nil
}
