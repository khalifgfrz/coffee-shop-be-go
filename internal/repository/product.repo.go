package repository

import (
	"fmt"
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type ProductRepositoryInterface interface {
	CreateProduct(data *models.Product) (string, error)
	GetAllProduct(que *models.ProductQuery) (*models.GetProducts, error)
	GetDetailProduct(id string) (*models.GetProduct, error)
	DeleteProduct(id string) (string, error)
	UpdateProduct(data *models.Product, id string) (string, error)
}

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r *RepoProduct) CreateProduct(data *models.Product) (string, error) {
	query := `INSERT INTO public.product(
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

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "Data created", nil
}

func (r *RepoProduct) GetAllProduct(que *models.ProductQuery) (*models.GetProducts, error) {
	query := `SELECT * FROM public.product`
	var values []interface{}
	condition := false

	if que.Product_name != "" {
		query += fmt.Sprintf(` WHERE product_name ILIKE $%d`, len(values)+1)
		values = append(values, "%"+que.Product_name+"%")
		condition = true
	}
	if que.MinPrice != 0 {
		if condition {
			query += " AND "
		} else {
			query += " WHERE "
		}
		query += fmt.Sprintf(` price > $%d`, len(values)+1)
		values = append(values, que.MinPrice)
		condition = true
	}
	if que.MaxPrice != 0 {
		if condition {
			query += " AND "
		} else {
			query += " WHERE "
		}
		query += fmt.Sprintf(` price < $%d`, len(values)+1)
		values = append(values, que.MaxPrice)
		condition = true
	}
	if que.Category != "" {
		if condition {
			query += " AND "
		} else {
			query += " WHERE "
		}
		query += fmt.Sprintf(` category = $%d`, len(values)+1)
		values = append(values, que.Category)
		condition = true
	}

	switch que.SortBy {
	case "alphabet":
		query += " ORDER BY product_name ASC"
	case "price":
		query += " ORDER BY price ASC"
	case "latest":
		query += " ORDER BY created_at DESC"
	case "oldest":
		query += " ORDER BY created_at ASC"
	}

	if que.Page > 0 {
		limit := 5
		offset := (que.Page - 1) * limit
		query += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, len(values)+1, len(values)+2)
		values = append(values, limit, offset)
	}
	
	data := models.GetProducts{}

	err := r.DB.Select(&data, query, values...)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoProduct) GetDetailProduct(id string) (*models.GetProduct, error) {
	query := `SELECT * FROM public.product WHERE product_id=$1`
	data := models.GetProduct{}
	err := r.Get(&data, query, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (r *RepoProduct) DeleteProduct(id string) (string, error) {
	query := `DELETE FROM public.product WHERE product_id=$1`
	_, err := r.Exec(query, id)
	if err != nil {
		return "", err
	}
	return "Data deleted", nil
}

func (r *RepoProduct) UpdateProduct(data *models.Product, id string) (string, error) {
	query := `UPDATE public.product SET
		product_name = COALESCE(NULLIF(:product_name, ''), product_name),
		price = COALESCE(NULLIF(:price, 0), price),
		category = COALESCE(NULLIF(:category, ''), category),
		description = COALESCE(NULLIF(:description, ''), description),
		stock = COALESCE(NULLIF(:stock, 0), stock),
		product_image = COALESCE(NULLIF(:product_image, ''), product_image),
		updated_at = now()
	WHERE product_id = :id
	`

	params := map[string]interface{}{
		"product_name":  data.Product_name,
		"price":         data.Price,
		"category":      data.Category,
		"description":   data.Description,
		"stock":         data.Stock,
		"product_image": data.Product_image,
		"id":            id,
	}

	_, err := r.NamedExec(query, params)
	if err != nil {
		return "", err
	}
	return "Data updated", nil
}