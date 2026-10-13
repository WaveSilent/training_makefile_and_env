package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]UserModel, error) {
	sqlQuery := `
	SELECT * FROM users
	ORDER BY id
	`

	rows, err := conn.Query(ctx, sqlQuery)

	if err != nil {
		return nil, err
	}

	users := make([]UserModel, 0)

	for rows.Next() {
		var user UserModel

		err := rows.Scan(
			&user.ID,
			&user.FullName,
			&user.PhoneNumber,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
