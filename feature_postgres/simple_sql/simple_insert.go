package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	fullName string,
	phoneNumber string,
) error {
	sqlQuery := `
	INSERT INTO users (full_name, phone_number)
	VALUES ($1, $2);
	`

	_, err := conn.Exec(ctx, sqlQuery, fullName, phoneNumber)

	return err

}
