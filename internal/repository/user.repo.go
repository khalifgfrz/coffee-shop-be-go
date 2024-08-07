package repository

import (
	"fmt"
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
	query := `INSERT INTO public.user(
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

	_, err := r.NamedExec(query, data)
	return err
}

func (r *RepoUser) GetAllUser(que *models.UserQuery) (*models.Users, error) {
	query := `SELECT * FROM public.user order by created_at DESC`
	var values []interface{}

	if que.Page > 0 {
		limit := 5
		offset := (que.Page - 1) * limit
		query += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, len(values)+1, len(values)+2)
		values = append(values, limit, offset)
	}
	
	rows, err := r.DB.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data models.Users
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.User_id,
			&user.User_uuid,
			&user.First_name,
			&user.Last_name,
			&user.Phone,
			&user.Address,
			&user.Birth_date,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.Created_at,
			&user.Updated_at,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoUser) GetDetailUser(id int) (*models.User, error) {
	query := `SELECT * FROM public.user WHERE user_id = :user_id`
	data := models.User{}

	rows, err := r.DB.NamedQuery(query, map[string]interface{}{
		"user_id": id,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.StructScan(&data)
		if err != nil {
			return nil, err
		}
		return &data, nil
	}

	return nil, nil
}

func (r *RepoUser) DeleteUser(id int) error {
	query := `DELETE FROM public.user WHERE user_id = :user_id`

	_, err := r.DB.NamedExec(query, map[string]interface{}{
		"user_id": id,
	})
	return err
}

func (r *RepoUser) UpdateUser(data *models.User, id int) (*models.User, error) {
	query := `UPDATE public.user SET
		first_name = COALESCE(NULLIF(:first_name, ''), first_name),
		last_name = COALESCE(NULLIF(:last_name, 0), last_name),
		phone = COALESCE(NULLIF(:phone, ''), phone),
		address = COALESCE(NULLIF(:address, ''), address),
		birth_date = COALESCE(NULLIF(:birth_date, ''), birth_date),
		email = COALESCE(NULLIF(:email, 0), email),
		password = COALESCE(NULLIF(:password, 0), password),
		role = COALESCE(NULLIF(:role, 0), role),
		updated_at = now()
	WHERE user_id = :id RETURNING *`

	params := map[string]interface{}{
		"first_name": 		data.First_name,
		"last_name":        data.Last_name,
		"phone":     		data.Phone,
		"address":     		data.Address,
		"birth_date":  		data.Birth_date,
		"email":        	data.Email,
		"password":        	data.Password,
		"role":        		data.Role,
		"id":           	id,
	}

	rows, err := r.DB.NamedQuery(query, params)
	if err != nil {
		return nil, fmt.Errorf("query execution error: %w", err)
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err := rows.Scan(
			&user.User_id,
			&user.User_uuid,
			&user.First_name,
			&user.Last_name,
			&user.Phone,
			&user.Address,
			&user.Birth_date,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.Created_at,
			&user.Updated_at,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
	} else {
		return nil, fmt.Errorf("user with id %d not found", id)
	}

	return &user, nil
}