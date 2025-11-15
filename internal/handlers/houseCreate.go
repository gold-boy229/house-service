package handlers

import (
	"context"
	"house-store/internal/dto"
	"house-store/internal/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *houseHandler) CreateNewHouse(c echo.Context) error {
	var reqDTO dto.HouseCreate_Request
	if err := c.Bind(&reqDTO); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err := c.Validate(reqDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	houseReq := convertDTOToEntity_HouseCreate(reqDTO)

	houseRes, err := h.repo.House_Create(context.Background(), houseReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, convertEntityToDTO_HouseCreate(houseRes))
}

func convertDTOToEntity_HouseCreate(reqDTO dto.HouseCreate_Request) entity.House {
	return entity.House{
		Address:   reqDTO.Address,
		Year:      *reqDTO.Year,
		Developer: reqDTO.Developer,
	}
}

func convertEntityToDTO_HouseCreate(house entity.House) dto.HouseCreate_Response {
	return dto.HouseCreate_Response{
		Id:        house.Id,
		Address:   house.Address,
		Year:      house.Year,
		Developer: house.Developer,
		CreatedAt: house.CreatedAt,
		UpdatedAt: house.UpdatedAt,
	}
}
