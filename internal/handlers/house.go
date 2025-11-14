package handlers

import (
	"context"
	"house-store/internal/dto"
	"house-store/internal/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type houseProvider interface {
	House_GetById()
	House_SubscribeForUpdates()
	House_Create(context.Context, entity.House) (entity.House, error)
}

type houseHandler struct {
	repo houseProvider
}

func NewHouseHandler(repo houseProvider) *houseHandler {
	return &houseHandler{repo: repo}
}

func (h *houseHandler) GetHouseById(c echo.Context) error {
	return nil
}

func (h *houseHandler) SubscribeForHouseUpdates(c echo.Context) error {
	return nil
}

func (h *houseHandler) CreateNewHouse(c echo.Context) error {
	var reqDTO dto.HouseCreateRequest
	if err := c.Bind(&reqDTO); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err := c.Validate(reqDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	houseReq := convertDTOToEntity_House(reqDTO)

	houseRes, err := h.repo.House_Create(context.Background(), houseReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, convertEntityToDTO_House(houseRes))
}

func convertDTOToEntity_House(reqDTO dto.HouseCreateRequest) entity.House {
	return entity.House{
		Address:   reqDTO.Address,
		Year:      *reqDTO.Year,
		Developer: reqDTO.Developer,
	}
}

func convertEntityToDTO_House(house entity.House) dto.HouseCreateResponse {
	return dto.HouseCreateResponse{
		Id:        house.Id,
		Address:   house.Address,
		Year:      house.Year,
		Developer: house.Developer,
		CreatedAt: house.CreatedAt,
		UpdatedAt: house.UpdatedAt,
	}
}
