package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r *RepoProduct) CreateProduct(data *models.Product) error {
	q := `INSERT INTO public.product(
		product_name,
		price,
		category,
		description,
		stock
	) VALUES(
	 	:product_name,
		:price,
		:category,
		:description,
		:stock
	)`

	_, err := r.NamedExec(q, data)
	return err
}

func (r *RepoProduct) GetAllProduct() (*models.Products, error) {
	q := `SELECT * FROM public.product order by created_at DESC`
	data := models.Products{}

	if err := r.Select(&data, q); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoProduct) GetDetailProduct(id int) (*models.Product, error) {
	q := `SELECT * FROM public.product WHERE product_id = $1`
	data := models.Product{}

	if err := r.Get(&data, q, id); err != nil {
		return nil, err
	}

	if data.Product_id == 0 {
		return nil, nil
	}

	return &data, nil
}

func (r *RepoProduct) DeleteProduct(id int) error {
	q := `DELETE FROM public.product WHERE product_id = $1`

	_, err := r.Exec(q, id)
	return err
}