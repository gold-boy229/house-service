package dto

type FlatCreate_Request struct {
	HouseId int  `json:"house_id" validate:"required,min=1"`
	Price   *int `json:"price" validate:"required,min=0"`
	Rooms   int  `json:"rooms" validate:"required,min=1"`
}

type FlatCreate_Response Flat_Response
