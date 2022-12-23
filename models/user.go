package models

type UserPrimarKey struct {
	Id    string `json:"user_id"`
	Login string `json:"login"`
}

type CreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type User struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateUser struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type GetListUserRequest struct {
	Limit  int32
	Offset int32
}

type GetListUserResponse struct {
	Count int32   `json:"count"`
	Users []*User `json:"users"`
}
