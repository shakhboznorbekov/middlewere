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

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (f *userRepo) Create(ctx context.Context, user *models.CreateUser) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO users(
			user_id,
			first_name, 
			last_name,
			phone_number,
			balance,
			updated_at
		) VALUES ( $1, $2 , $3, $4, $5, now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Balance,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *userRepo) GetByPKey(ctx context.Context, pkey *models.UserPrimarKey) (*models.User, error) {

	var (
		id          sql.NullString
		firstName   sql.NullString
		lastName    sql.NullString
		phoneNumber sql.NullString
		balance     sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query := `
		SELECT
			user_id,
			first_name, 
			last_name,
			phone_number,
			balance,
			created_at,
			updated_at
		FROM
			users
		WHERE user_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&firstName,
			&lastName,
			&phoneNumber,
			&balance,
		)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:          id.String,
		FirstName:   firstName.String,
		LastName:    lastName.String,
		PhoneNumber: phoneNumber.String,
		Balance:     balance.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (f *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {

	var (
		resp   = models.GetListUserResponse{}
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
			user_id,
			first_name, 
			last_name,
			phone_number,
			balance,
			created_at,
			updated_at
		FROM
			users
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id          sql.NullString
			firstName   sql.NullString
			lastName    sql.NullString
			phoneNumber sql.NullString
			balance     sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&firstName,
			&lastName,
			&phoneNumber,
			&balance,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:          id.String,
			FirstName:   firstName.String,
			LastName:    lastName.String,
			PhoneNumber: phoneNumber.String,
			Balance:     balance.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})

	}

	return &resp, err
}

func (f *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			users
		SET
			first_name = :first_name,
			last_name = :last_name,
			phone_number = :phone_number,
			balance = balance, 
			updated_at = now()
		WHERE user_id = :user_id
	`

	params = map[string]interface{}{
		"user_id":      req.Id,
		"first_name":   req.FirstName,
		"last_name":    req.LastName,
		"phone_number": req.PhoneNumber,
		"balance":      req.Balance,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *userRepo) Delete(ctx context.Context, req *models.UserPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM users WHERE user_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
