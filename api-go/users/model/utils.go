package model

import (
	"github.com/jackc/pgx/v5"
)

func getUserFromRow(row pgx.Row) (*User, error) {
	userResult := User{}
	err := row.Scan(&userResult.Id, &userResult.Email, &userResult.Name,
		&userResult.Pass, &userResult.PermissionLevel, &userResult.CreatedAt, &userResult.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &userResult, nil
}
