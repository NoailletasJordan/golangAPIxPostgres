package model

import (
	"github.com/jackc/pgx/v5"
)

func getUserFromRow(row pgx.Row) (*User, error) {
	userResult := User{}
	err := row.Scan(&userResult.Id, &userResult.Email, &userResult.Firstname,
		&userResult.Lastname, &userResult.Pass, &userResult.Active, &userResult.CreatedAt, &userResult.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &userResult, nil
}
