package repository

import (
	"fmt"
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type FavoriteRepositoryInterface interface {
	CreateFavorite(data *models.Favorite) (string, error)
	GetAllFavorite(que *models.FavoriteQuery) (*models.GetFavorites, error)
	GetDetailFavorite(id string) (*models.GetFavorite, error)
	DeleteFavorite(id string) (string, error)
	// UpdateFavorite(data *models.Favorite, id string) (string, error)
}

type RepoFavorite struct {
	*sqlx.DB
}

func NewFavorite(db *sqlx.DB) *RepoFavorite {
	return &RepoFavorite{db}
}

func (r *RepoFavorite) CreateFavorite(data *models.Favorite) (string, error) {
	query := `INSERT INTO public.favorite(
		product_id,
		user_id
		) VALUES(
		 :product_id,
		 :user_id
		)`

		_, err := r.NamedExec(query, data)
		if err != nil {
			return "", err
		}
		return "data created", nil
}

func (r *RepoFavorite) GetAllFavorite(que *models.FavoriteQuery) (*models.GetFavorites, error) {
	query := `SELECT f.favorite_id, f.favorite_uuid, u.first_name, u.last_name, u.phone, u.address,
	u.email, u.user_image, p.product_name, p.price, p.category, p.description, p.product_image, f.created_at, f.updated_at FROM public.favorite f
	join public.product p on f.product_id = p.product_id
	join public.user u on f.user_id = u.user_id
	order by f.created_at DESC`
	var values []interface{}

	if que.Page > 0 {
		limit := 5
		offset := (que.Page - 1) * limit
		query += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, len(values)+1, len(values)+2)
		values = append(values, limit, offset)
	}
	
	data := models.GetFavorites{}

	err := r.DB.Select(&data, query, values...)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoFavorite) GetDetailFavorite(id string) (*models.GetFavorite, error) {
	query := `SELECT f.favorite_id, f.favorite_uuid, u.first_name, u.last_name, u.phone, u.address,
	u.email, u.user_image, p.product_name, p.price, p.category, p.description, p.product_image,
	f.created_at, f.updated_at FROM public.favorite f
	join public.product p on f.product_id = p.product_id
	join public.user u on f.user_id = u.user_id
	WHERE f.favorite_id=$1`
	data := models.GetFavorite{}
	err := r.Get(&data, query, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *RepoFavorite) DeleteFavorite(id string) (string, error) {
	query := `DELETE FROM public.favorite WHERE favorite_id=$1`
	_, err := r.Exec(query, id)
	if err != nil {
		return "", err
	}
	return "data deleted", nil
}

// func (r *RepoFavorite) UpdateFavorite(data *models.Favorite, id string) (string, error) {
// 	query := `UPDATE public.favorite SET
// 		user_id = COALESCE(NULLIF(:user_id, 0), user_id),
// 		product_id = COALESCE(NULLIF(:product_id, 0), product_id),
// 		updated_at = now()
// 	WHERE favorite_id = :id`

// 	params := map[string]interface{}{
// 		"user_id": 		data.User_id,
// 		"product_id":   data.Product_id,
// 		"id":           id,
// 	}

// 	_, err := r.NamedExec(query, params)
// 	if err != nil {
// 		return "", err
// 	}
// 	return "data updated", nil
// }