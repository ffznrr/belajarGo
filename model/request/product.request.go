package request

type ProductCreateRequest struct {
	Name  string `json:"name" validate:"required"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}