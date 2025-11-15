package handlers

import (
	"context"
	"house-store/internal/consts"
	"house-store/internal/dto"
	"house-store/internal/entity"
	"house-store/internal/enum"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *houseHandler) GetHouseById(c echo.Context) error {
	var reqDTO dto.HouseGetById_Request
	if err := c.Bind(&reqDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(reqDTO); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	userRole := c.Get(consts.ECHO_CONTEXT_USER_ROLE_KEY).(string)

	var (
		resultFlats []entity.Flat
		err         error
	)
	if userRole == enum.USER_ROLE_MODERATOR {
		resultFlats, err = h.repo.House_GetById_Moderator(context.TODO(), reqDTO.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
		}
	} else {
		resultFlats, err = h.repo.House_GetById_Client(context.TODO(), reqDTO.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse_5xx{Message: err.Error()})
		}
	}

	return c.JSON(http.StatusOK, createDTOResponse_HouseGetById(resultFlats))
}

func createDTOResponse_HouseGetById(flats []entity.Flat) dto.HouseGetById_Response {
	return dto.HouseGetById_Response{
		Flats: convertEntityToDTO_Flats(flats),
	}
}

func convertEntityToDTO_Flats(flats []entity.Flat) []dto.Flat_Response {
	result := make([]dto.Flat_Response, 0, len(flats))
	for _, flat := range flats {
		result = append(result, convertEntityToDTO_Flat(flat))
	}
	return result
}

func convertEntityToDTO_Flat(from entity.Flat) dto.Flat_Response {
	return dto.Flat_Response{
		FlatId:  from.FlatId,
		HouseId: from.HouseId,
		Price:   from.Price,
		Rooms:   from.Rooms,
		Status:  from.Status,
	}
}
