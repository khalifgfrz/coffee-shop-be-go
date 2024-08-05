package repository

import (
	"database/sql"
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
	query := `UPDATE public.user SET`
	var values []interface{}
	condition := false

	if data.First_name != "" {
		query += fmt.Sprintf(` first_name = $%d`, len(values)+1)
		values = append(values, data.First_name)
		condition = true
	}
	if data.Last_name != "" {
		if condition {
			query += ","
		}
		query += fmt.Sprintf(` last_name = $%d`, len(values)+1)
		values = append(values, data.Last_name)
		condition = true
	}
	if data.Phone != "" {
		if condition {
			query += ","
		}
		query += fmt.Sprintf(` phone = $%d`, len(values)+1)
		values = append(values, data.Phone)
		condition = true
	}
	if data.Birth_date != "" {
		if condition {
			query += ","
		}
		query += fmt.Sprintf(` birth_date = $%d`, len(values)+1)
		values = append(values, data.Birth_date)
		condition = true
	}
	if data.Email != "" {
		if condition {
			query += ","
		}
		query += fmt.Sprintf(` email = $%d`, len(values)+1)
		values = append(values, data.Email)
		condition = true
	}
	if data.Password != "" {
		if condition {
			query += ","
		}
		query += fmt.Sprintf(` password = $%d`, len(values)+1)
		values = append(values, data.Password)
		condition = true
	}
	if data.Role != "" {
		if condition {
			query += ","
		}
		query += fmt.Sprintf(` role = $%d`, len(values)+1)
		values = append(values, data.Role)
		condition = true
	}
	if !condition {
		return nil, fmt.Errorf("no fields to update")
	}

	query += fmt.Sprintf(`, updated_at = now() WHERE user_id = $%d RETURNING *`, len(values)+1)
	values = append(values, id)

	row := r.DB.QueryRow(query, values...)
	var user models.User
	err := row.Scan(
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
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(`user with id %d not found`, id)
		}
		return nil, fmt.Errorf(`query execution error: %w`, err)
	}

	return &user, nil
}