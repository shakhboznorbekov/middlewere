package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"crud/models"
	"crud/pkg/helper"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (f *orderRepo) Create(ctx context.Context, order *models.CreateOrder) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO orders(
			order_id,
			book_id, 
			user_id, 
			updated_at
		) VALUES ( $1, $2 , $3, now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		order.BookId,
		order.UserId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *orderRepo) GetByPKey(ctx context.Context, pkey *models.OrderPrimarKey) (*models.Order, error) {

	var (
		id        sql.NullString
		bookId    sql.NullString
		userId    sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query := `
		SELECT
			order_id,
			book_id,
			user_id, 
			created_at,
			updated_at
		FROM
			orders
		WHERE order_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&bookId,
			&userId,
		)

	if err != nil {
		return nil, err
	}

	return &models.Order{
		Id:        id.String,
		BookId:    bookId.String,
		UserId:    userId.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (f *orderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {

	var (
		resp   = models.GetListOrderResponse{}
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			order_id,
			book_id,
			user_id, 
			created_at,
			updated_at
		FROM
			orders
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id        sql.NullString
			bookId    sql.NullString
			userId    sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&bookId,
			&userId,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Orders = append(resp.Orders, &models.Order{
			Id:        id.String,
			BookId:    bookId.String,
			UserId:    userId.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})

	}

	return &resp, err
}

func (f *orderRepo) Update(ctx context.Context, req *models.UpdateOrder) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			orders
		SET
			book_id = :book_id,
			user_id = :user_id, 
			updated_at = now()
		WHERE order_id = :order_id
	`

	params = map[string]interface{}{
		"order_id": req.Id,
		"book_id":  req.BookId,
		"user_id":  req.UserId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *orderRepo) Delete(ctx context.Context, req *models.OrderPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM orders WHERE order_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
