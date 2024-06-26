// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package product

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

// START OMIT

const create = `-- name: Create :one
INSERT INTO product (
  name, price
) VALUES ($1, $2)
RETURNING id, name, price
`

type CreateParams struct {
	Name  string
	Price pgtype.Numeric
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (Product, error) {
	row := q.db.QueryRow(ctx, create, arg.Name, arg.Price)
	var i Product
	err := row.Scan(&i.ID, &i.Name, &i.Price)
	return i, err
}

// END OMIT

const delete = `-- name: Delete :exec
DELETE FROM product
WHERE id = $1
`

func (q *Queries) Delete(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, delete, id)
	return err
}

const getAll = `-- name: GetAll :many
SELECT id, name, price FROM product
ORDER BY name
`

func (q *Queries) GetAll(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(&i.ID, &i.Name, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getByID = `-- name: GetByID :one
SELECT id, name, price FROM product
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetByID(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRow(ctx, getByID, id)
	var i Product
	err := row.Scan(&i.ID, &i.Name, &i.Price)
	return i, err
}

const update = `-- name: Update :exec
UPDATE product
  set name = $2,
  price = $3
WHERE id = $1
`

type UpdateParams struct {
	ID    int64
	Name  string
	Price pgtype.Numeric
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) error {
	_, err := q.db.Exec(ctx, update, arg.ID, arg.Name, arg.Price)
	return err
}
