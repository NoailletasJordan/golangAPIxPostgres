package model

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Pass      string    `json:"pass"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Model struct {
	DBConnection *pgx.Conn
}

func (model *Model) GetOne(query string) (*User, error) {
	row := model.DBConnection.QueryRow(context.Background(), query)
	userResult, err := getUserFromRow(row)
	return userResult, err
}

func (model Model) GetAll(query string) ([]*User, error) {
	var users = []*User{}
	rows, err := model.DBConnection.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user, err := getUserFromRow(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (model Model) Execute(query string) error {
	_, err := model.DBConnection.Exec(context.Background(), query)
	return err
}
