package dto

type FlatUpdate_Request struct {
	FlatId int    `json:"id" validate:"required,min=1"`
	Status string `json:"status" validate:"required,oneof=approved declined"`
}

type FlatUpdate_Response Flat_Response
