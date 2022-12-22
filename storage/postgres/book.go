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

type bookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) *bookRepo {
	return &bookRepo{
		db: db,
	}
}

func (f *bookRepo) Create(ctx context.Context, book *models.CreateBook) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO book(
			book_id,
			name, 
			author_name,
			price,
			date,
			created_at,
			updated_at
		) VALUES ( $1, $2 , $3, $4, $5, now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		book.Name,
		book.AuthorName,
		book.Price,
		book.Date,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *bookRepo) GetByPKey(ctx context.Context, pkey *models.BookPrimarKey) (*models.Book, error) {

	var (
		id         sql.NullString
		name       sql.NullString
		authorName sql.NullString
		price      sql.NullString
		date       sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	query := `
		SELECT
			book_id,
			name, 
			author_name,
			price,
			date,
			created_at,
			updated_at
		FROM
			book
		WHERE book_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&name,
			&authorName,
			&price,
			&date,
		)
	if err != nil {
		return nil, err
	}

	return &models.Book{
		Id:         id.String,
		Name:       name.String,
		AuthorName: authorName.String,
		Price:      price.String,
		Date:       date.String,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}, nil
}

func (f *bookRepo) GetList(ctx context.Context, req *models.GetListBookRequest) (*models.GetListBookResponse, error) {

	var (
		resp   = models.GetListBookResponse{}
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
			book_id,
			name, 
			author_name,
			price,
			date,
			created_at,
			updated_at
		FROM
			book
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id         sql.NullString
			name       sql.NullString
			authorName sql.NullString
			price      sql.NullString
			date       sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&authorName,
			&price,
			&date,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Books = append(resp.Books, &models.Book{
			Id:         id.String,
			Name:       name.String,
			AuthorName: authorName.String,
			Price:      price.String,
			Date:       date.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})

	}

	return &resp, err
}

func (f *bookRepo) Update(ctx context.Context, req *models.UpdateBook) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
			UPDATE
				book
			SET
				name = :name,
				author_name = :author_name,
				price = :price,
				date = date, 
				updated_at = now()
			WHERE book_id = :book_id
		`

	params = map[string]interface{}{
		"book_id":     req.Id,
		"name":        req.Name,
		"author_name": req.AuthorName,
		"price":       req.Price,
		"date":        req.Date,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *bookRepo) Delete(ctx context.Context, req *models.BookPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM book WHERE book_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
