// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: sign-in.maindb.sql

package maindb

import (
	"context"

	"github.com/google/uuid"
)

const SignInGetUser = `-- name: SignInGetUser :one
SELECT id, "password", "role" FROM "user" WHERE email = $1 LIMIT 1
`

type SignInGetUserRow struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Password string    `db:"password" json:"password"`
	Role     string    `db:"role" json:"role"`
}

func (q *Queries) SignInGetUser(ctx context.Context, email string) (*SignInGetUserRow, error) {
	row := q.db.QueryRowContext(ctx, SignInGetUser, email)
	var i SignInGetUserRow
	err := row.Scan(&i.ID, &i.Password, &i.Role)
	return &i, err
}
