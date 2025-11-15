package handlers

import (
	"context"
	"house-store/internal/dto"
	"house-store/internal/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *flatHandler) Create(c echo.Context) error {
	var reqDTO dto.FlatCreate_Request
	if err := c.Bind(&reqDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(reqDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	flat := convertDTOToEntity_FlatCreate(reqDTO)

	resFlat, err := h.repo.Flat_Create(context.TODO(), flat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, convertEntityToDTO_FlatCreate(resFlat))
}

func convertDTOToEntity_FlatCreate(reqDTO dto.FlatCreate_Request) entity.Flat {
	return entity.Flat{
		HouseId: reqDTO.HouseId,
		Price:   *reqDTO.Price,
		Rooms:   reqDTO.Rooms,
	}
}

func convertEntityToDTO_FlatCreate(flat entity.Flat) dto.FlatCreate_Response {
	return dto.FlatCreate_Response{
		FlatId:  flat.FlatId,
		HouseId: flat.HouseId,
		Price:   flat.Price,
		Rooms:   flat.Rooms,
		Status:  flat.Status,
	}
}
