package repository

import (
	"fmt"
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryInterface interface {
	CreateUser(data *models.User) (string, error)
	GetAllUser(que *models.UserQuery) (*models.GetUsers, error)
	GetDetailUser(id string) (*models.GetUser, error)
	DeleteUser(id string) (string, error)
	UpdateUser(data *models.User, id string) (string, error)
}

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) CreateUser(data *models.User) (string, error) {
	query := `INSERT INTO public.user(
		first_name,
		last_name,
		phone,
		address,
		birth_date,
		email,
		password,
		role,
	) VALUES(
	 	:first_name,
		:last_name,
		:phone,
		:address,
		:birth_date,
		:email,
		:password,
		:role,
	)`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "Data created", nil
}

func (r *RepoUser) GetAllUser(que *models.UserQuery) (*models.GetUsers, error) {
	query := `SELECT * FROM public.user`
	var values []interface{}

	if que.Page > 0 {
		limit := 5
		offset := (que.Page - 1) * limit
		query += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, len(values)+1, len(values)+2)
		values = append(values, limit, offset)
	}
	
	data := models.GetUsers{}

	err := r.Select(&data, query, values...)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoUser) GetDetailUser(id string) (*models.GetUser, error) {
	query := `SELECT * FROM public.user WHERE user_id=$1`
	data := models.GetUser{}
	err := r.Get(&data, query, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *RepoUser) DeleteUser(id string) (string, error) {
	query := `DELETE FROM public.user WHERE user_id=$1`
	_, err := r.Exec(query, id)
	if err != nil {
		return "", err
	}
	return "Data deleted", nil
}

func (r *RepoUser) UpdateUser(data *models.User, id string) (string, error) {
	query := `UPDATE public.user SET
		first_name = COALESCE(NULLIF($1, ''), first_name),
		last_name = COALESCE(NULLIF($2, ''), last_name),
		phone = COALESCE(NULLIF($3, ''), phone),
		address = COALESCE(NULLIF($4, ''), address),
		birth_date = COALESCE(NULLIF($5, '')::date, birth_date),
		email = COALESCE(NULLIF($6, ''), email),
		password = COALESCE(NULLIF($7, ''), password),
		role = COALESCE(NULLIF($8, ''), role),
		user_image = COALESCE(NULLIF($9, ''), user_image),
		updated_at = now()
	WHERE user_id = $10`

	params := []interface{}{
		data.First_name,
		data.Last_name,
		data.Phone,
		data.Address,
		data.Birth_date,
		data.Email,
		data.Password,
		data.Role,
		data.User_image,
		id,
	}

	_, err := r.Exec(query, params...)
	if err != nil {
		return "", err
	}
	return "Data updated", nil
}