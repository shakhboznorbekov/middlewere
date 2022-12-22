package models

type OrderPrimarKey struct {
	Id string `json:"order_id"`
}

type CreateOrder struct {
	BookId string `json:"book_id"`
	UserId string `json:"user_id"`
}

type Order struct {
	Id        string `json:"order_id"`
	BookId    string `json:"book_id"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateOrder struct {
	Id     string `json:"order_id"`
	BookId string `json:"book_id"`
	UserId string `json:"user_id"`
}

type GetListOrderRequest struct {
	Limit  int32
	Offset int32
}

type GetListOrderResponse struct {
	Count  int32    `json:"count"`
	Orders []*Order `json:"orders"`
}
