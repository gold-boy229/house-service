package handlers

import (
	"context"
	"house-store/internal/dto"
	"house-store/internal/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *flatHandler) Update(c echo.Context) error {
	var reqDTO dto.FlatUpdate_Request
	if err := c.Bind(&reqDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(reqDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	flat := convertDTOToEntity_FlatUpdate(reqDTO)

	resultFlat, err := h.repo.Flat_Update(context.TODO(), flat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, convertEntityToDTO_FlatUpdate(resultFlat))
}

func convertDTOToEntity_FlatUpdate(reqDTO dto.FlatUpdate_Request) entity.Flat {
	return entity.Flat{
		FlatId: reqDTO.FlatId,
		Status: reqDTO.Status,
	}
}

func convertEntityToDTO_FlatUpdate(flat entity.Flat) dto.FlatUpdate_Response {
	return dto.FlatUpdate_Response{
		FlatId:  flat.FlatId,
		HouseId: flat.HouseId,
		Price:   flat.Price,
		Rooms:   flat.Rooms,
		Status:  flat.Status,
	}
}
