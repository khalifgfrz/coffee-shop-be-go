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
	q := `INSERT INTO public.favorite(product_id) VALUES(:product_id)`

	_, err := r.NamedExec(q, data)
	return err
}

func (r *RepoFavorite) GetAllFavorite() (*models.Favorites, error) {
	q := `SELECT f.favorite_id, p.product_name, p.price, p.category, p.description, f.created_at, f.updated_at FROM public.favorite f
	join public.product p on f.product_id = p.product_id
	order by f.created_at DESC`
	data := models.Favorites{}

	if err := r.Select(&data, q); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoFavorite) GetDetailFavorite(id int) (*models.Favorite, error) {
	q := `SELECT f.favorite_id, p.product_name, p.price, p.category, p.description FROM public.favorite f
	join public.product p on f.product_id = p.product_id
	WHERE f.favorite_id = $1`
	data := models.Favorite{}

	if err := r.Get(&data, q, id); err != nil {
		return nil, err
	}

	if data.Favorite_id == 0 {
		return nil, nil
	}

	return &data, nil
}

func (r *RepoFavorite) DeleteFavorite(id int) error {
	q := `DELETE FROM public.favorite WHERE favorite_id = $1`

	_, err := r.Exec(q, id)
	return err
}