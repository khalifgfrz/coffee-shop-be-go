package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderDetailsRepositoryInterface interface {
	CreateOrderDetails(order_id int, products []models.ProductDetail) (string, error)
	GetOrderDetails(order_id int) (*[]models.GetOrderDetail, error)
}

type RepoOrderDetails struct {
	*sqlx.DB
}

func NewOrderDetails(db *sqlx.DB) *RepoOrderDetails {
	return &RepoOrderDetails{db}
}

func (r *RepoOrderDetails) CreateOrderDetails(order_id int, products []models.ProductDetail) (string, error) {
	query := `INSERT INTO public.order_details (
		order_id, 
		size_id, 
		product_id, 
		qty
	) VALUES (
		:order_id, 
		:size_id, 
		:product_id, 
		:qty
	)`

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

func (r *RepoOrderDetails) GetOrderDetails(order_id int) (*[]models.GetOrderDetail, error) {
	query := `select od.order_id, p.product_image, p.product_name, p.price, s.size, s.added_cost, od.qty from order_details od
    join public.product p on od.product_id = p.product_id
    join public.sizes s on od.size_id = s.size_id
    where od.order_id=$1`
	var data []models.GetOrderDetail
	err := r.Select(&data, query, order_id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

