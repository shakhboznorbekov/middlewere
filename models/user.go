package models

type UserPrimarKey struct {
	Id string `json:"user_id"`
}

type CreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Balance     string `json:"balance"`
}

type User struct {
	Id          string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Balance     string `json:"balance"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateUser struct {
	Id          string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Balance     string `json:"balance"`
}

type GetListUserRequest struct {
	Limit  int32
	Offset int32
}

type GetListUserResponse struct {
	Count int32   `json:"count"`
	Users []*User `json:"users"`
}
