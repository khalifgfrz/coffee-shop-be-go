package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) CreateUser(data *models.User) (string, error) {
	q := `INSERT INTO public.user(
		first_name,
		last_name,
		phone,
		address,
		birth_date,
		email,
		password,
		role
	) VALUES(
	 	:first_name,
		:last_name,
		:phone,
		:address,
		:birth_date,
		:email,
		:password,
		:role
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 data user created", nil
}

func (r *RepoUser) GetAllUser() (*models.Users, error) {
	q := `SELECT * FROM public.user order by created_at DESC`
	data := models.Users{}

	if err := r.Select(&data, q); err != nil {
		return nil, err
	}

	return &data, nil
}