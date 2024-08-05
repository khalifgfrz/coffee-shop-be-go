package repository

import (
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
	query := `INSERT INTO public.favorite(product_id) VALUES(:product_id)`

	_, err := r.NamedExec(query, data)
	return err
}

func (r *RepoFavorite) GetAllFavorite() (*models.Favorites, error) {
	query := `SELECT f.favorite_id, p.product_name, p.price, p.category, p.description, f.created_at, f.updated_at FROM public.favorite f
	join public.product p on f.product_id = p.product_id
	order by f.created_at DESC`
	data := models.Favorites{}

	if err := r.Select(&data, query); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoFavorite) GetDetailFavorite(id int) (*models.Favorite, error) {
	query := `SELECT f.favorite_id, p.product_name, p.price, p.category, p.description FROM public.favorite f
	join public.product p on f.product_id = p.product_id
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

func (r *RepoFavorite) UpdateFavorite(data *models.UpdateFavorite) error {
	query := `
		UPDATE public.favorite 
		SET product_id = :product_id, updated_at = now() WHERE favorite_id = :favorite_id
		RETURNING *
	`
	_, err := r.NamedExec(query, data)
	return err
}