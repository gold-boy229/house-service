package dto

type HouseGetById_Request struct {
	Id int `param:"id" validate:"required,min=1"`
}

type HouseGetById_Response struct {
	Flats []Flat_Response `json:"flats"`
}
