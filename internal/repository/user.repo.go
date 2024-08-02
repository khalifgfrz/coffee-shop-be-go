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

func (r *RepoUser) CreateUser(data *models.User) error {
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
	return err
}

func (r *RepoUser) GetAllUser() (*models.Users, error) {
	q := `SELECT * FROM public.user order by created_at DESC`
	data := models.Users{}

	if err := r.Select(&data, q); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoUser) GetDetailUser(id int) (*models.User, error) {
	q := `SELECT * FROM public.user WHERE user_id = $1`
	data := models.User{}

	if err := r.Get(&data, q, id); err != nil {
		return nil, err
	}

	if data.User_id == 0 {
		return nil, nil
	}

	return &data, nil
}

func (r *RepoUser) DeleteUser(id int) error {
	q := `DELETE FROM public.user WHERE user_id = $1`

	_, err := r.Exec(q, id)
	return err
}

func (r *RepoUser) UpdateUser(id int, user *models.User) (string, error) {
	q := `UPDATE public.user SET 
		first_name = :first_name,
		last_name = :last_name,
		phone = :phone,
		address = :address,
		birth_date = :birth_date,
		email = :email,
		password = :password,
		role = :role,
		updated_at = now()
	WHERE user_id = $1`

	_, err := r.NamedExec(q, user)
	if err != nil {
		return "", err
	}

	return "1 data user updated", nil
}