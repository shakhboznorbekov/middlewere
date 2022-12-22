package storage

import (
	"context"

	"crud/models"
)

type StorageI interface {
	CloseDB()
	Order() OrderRepoI
	User() UserRepoI
	Book() BookRepoI
}

type OrderRepoI interface {
	Create(ctx context.Context, req *models.CreateOrder) (string, error)
	GetByPKey(ctx context.Context, req *models.OrderPrimarKey) (*models.Order, error)
	GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Update(ctx context.Context, req *models.UpdateOrder) (int64, error)
	Delete(ctx context.Context, req *models.OrderPrimarKey) error
}

type BookRepoI interface {
	Create(ctx context.Context, req *models.CreateBook) (string, error)
	GetByPKey(ctx context.Context, req *models.BookPrimarKey) (*models.Book, error)
	GetList(ctx context.Context, req *models.GetListBookRequest) (*models.GetListBookResponse, error)
	Update(ctx context.Context, req *models.UpdateBook) (int64, error)
	Delete(ctx context.Context, req *models.BookPrimarKey) error
}

type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (string, error)
	GetByPKey(ctx context.Context, req *models.UserPrimarKey) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	Delete(ctx context.Context, req *models.UserPrimarKey) error
}
