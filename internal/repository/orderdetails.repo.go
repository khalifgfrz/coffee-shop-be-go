package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderDetailsRepositoryInterface interface {
	CreateOrderDetails(order_id int, products []models.ProductDetail) (string, error)
}

type RepoOrderDetails struct {
	*sqlx.DB
}

func NewOrderDetails(db *sqlx.DB) *RepoOrderDetails {
	return &RepoOrderDetails{db}
}

func (r *RepoOrderDetails) CreateOrderDetails(order_id int, products []models.ProductDetail) (string, error) {
	query := `INSERT INTO order_details (
		order_id, 
		size_id, 
		product_id, 
		qty
	) VALUES (
		:order_id, 
		:size_id, 
		:product_id, 
		:qty
	) RETURNING order_id, size_id, product_id, qty;`

	for _, product := range products {
		data := map[string]interface{}{
			"order_id":   order_id,
			"size_id":    product.Size_id,
			"product_id": product.Product_id,
			"qty":        product.Qty,
		}

		_, err := r.NamedExec(query, data)
		if err != nil {
			return "", err
		}
	}

	return "Data created", nil
}