package dto

type HouseGetById_Request struct {
	Id int `param:"id" validate:"required,min=1"`
}

type HouseGetById_Response struct {
	Flats []Flat_Response `json:"flats"`
}

type Flat_Response struct {
	FlatId  int    `json:"id"`
	HouseId int    `json:"house_id"`
	Price   int    `json:"price"`
	Rooms   int    `json:"rooms"`
	Status  string `json:"status"`
}
