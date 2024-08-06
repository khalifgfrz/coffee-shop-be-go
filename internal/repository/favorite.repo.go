package repository

import (
	"database/sql"
	"fmt"
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoFavorite struct {
	*sqlx.DB
}

func NewFavorite(db *sqlx.DB) *RepoFavorite {
	return &RepoFavorite{db}
}

func (r *RepoFavorite) CreateFavorite(data *models.PostFavorite) error {
	query := `INSERT INTO public.favorite(
		product_id,
		user_id
		) VALUES(
		 :product_id,
		 :user_id
		)`

	_, err := r.NamedExec(query, data)
	return err
}

func (r *RepoFavorite) GetAllFavorite() (*models.Favorites, error) {
	query := `SELECT f.favorite_id, f.favorite_uuid, u.first_name, u.last_name, u.phone, u.address,
	u.email, p.product_name, p.price, p.category, p.description, f.created_at, f.updated_at FROM public.favorite f
	join public.product p on f.product_id = p.product_id
	join public.user u on f.user_id = u.user_id
	order by f.created_at DESC`
	data := models.Favorites{}

	if err := r.Select(&data, query); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoFavorite) GetDetailFavorite(id int) (*models.Favorite, error) {
	query := `SELECT f.favorite_id, f.favorite_uuid, u.first_name, u.last_name, u.phone, u.address,
	u.email, p.product_name, p.price, p.category, p.description, f.created_at, f.updated_at FROM public.favorite f
	join public.product p on f.product_id = p.product_id
	join public.user u on f.user_id = u.user_id
	WHERE f.favorite_id = :favorite_id`
	data := models.Favorite{}

	rows, err := r.DB.NamedQuery(query, map[string]interface{}{
		"favorite_id": id,
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

func (r *RepoFavorite) DeleteFavorite(id int) error {
	query := `DELETE FROM public.favorite WHERE favorite_id = :favorite_id`

	_, err := r.DB.NamedExec(query, map[string]interface{}{
		"favorite_id": id,
	})
	return err
}

func (r *RepoFavorite) UpdateFavorite(data *models.UpdateFavorite, id int) (*models.UpdateFavorite, error) {
	query := `UPDATE public.favorite SET`
	var values []interface{}
	condition := false

	if data.User_id != 0 {
		query += fmt.Sprintf(` user_id = $%d`, len(values)+1)
		values = append(values, data.User_id)
		condition = true
	}
	if data.Product_id != 0 {
		if condition {
			query += ","
		}
		query += fmt.Sprintf(` product_id = $%d`, len(values)+1)
		values = append(values, data.Product_id)
		condition = true
	}
	if !condition {
		return nil, fmt.Errorf("no fields to update")
	}

	query += fmt.Sprintf(`, updated_at = now() WHERE favorite_id = $%d RETURNING *`, len(values)+1)
	values = append(values, id)

	row := r.DB.QueryRow(query, values...)
	var favorite models.UpdateFavorite
	err := row.Scan(
		&favorite.Favorite_id,
		&favorite.Favorite_uuid,
		&favorite.User_id,
		&favorite.Product_id,
		&favorite.Created_at,
		&favorite.Updated_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(`favorite with id %d not found`, id)	
		}
		return nil, fmt.Errorf(`query execution error: %w`, err)
	}

	return &favorite, nil
}