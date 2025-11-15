package dto

type HouseCreate_Request struct {
	Address   string `json:"address" validate:"required"`
	Year      *int   `json:"year" validate:"required,min=0"`
	Developer string `json:"developer"`
}

type HouseCreate_Response struct {
	Id        int    `json:"id"`
	Address   string `json:"address"`
	Year      int    `json:"year"`
	Developer string `json:"developer,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"update_at"`
}
