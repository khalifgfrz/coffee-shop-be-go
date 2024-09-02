package repository

import (
	"fmt"
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderRepositoryInterface interface {
	CreateOrder(order *models.Order) (int, error)
	GetAllOrder(que *models.OrderQuery) (*models.GetOrders, error)
	GetDetailOrder(uuid string) (*models.GetOrder, error)
	GetHistoryOrder(id string) (*models.GetOrder, error)
	UpdateOrder(data *models.Order, uuid string) (string, error)
}

type RepoOrder struct {
	*sqlx.DB
}

func NewOrder(db *sqlx.DB) *RepoOrder {
	return &RepoOrder{db}
}

func (r *RepoOrder) CreateOrder(order *models.Order) (int, error) {
	query := `
		INSERT INTO public.order_list (
			user_id,
			subtotal,
			tax,
			payment_id,
			delivery_id,
			status,
			grand_total
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		) RETURNING orderlist_id
	`

	params := []interface{}{
		order.User_id,
		order.Subtotal,
		order.Tax,
		order.Payment_id,
		order.Delivery_id,
		order.Status,
		order.Grand_total, 
	}

	var id int
	err := r.Get(&id, query, params...)
	if err != nil {
		return 0, err
	}

	return id, nil
}


func (r *RepoOrder) GetAllOrder(que *models.OrderQuery) (*models.GetOrders, error) {
	query := `SELECT ol.orderlist_id, ol.orderlist_uuid, u.first_name, u.last_name, u.phone, u.address, u.email, ol.subtotal, ol.tax, p.payment_method, d.delivery_method,
  	d.added_cost, ol.status, ol.grand_total from order_list ol
    join public.user u on ol.user_id = u.user_id
    join public.payments p on ol.payment_id = p.payment_id
    join public.deliveries d on ol.delivery_id = d.delivery_id`
	var values []interface{}

	if que.Page > 0 {
		limit := 5
		offset := (que.Page - 1) * limit
		query += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, len(values)+1, len(values)+2)
		values = append(values, limit, offset)
	}
	
	data := models.GetOrders{}

	err := r.Select(&data, query, values...)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoOrder) GetDetailOrder(uuid string) (*models.GetOrder, error) {
	query := `SELECT ol.orderlist_id, ol.orderlist_uuid, u.first_name, u.last_name, u.phone, u.address, u.email, ol.subtotal, ol.tax, p.payment_method, d.delivery_method,
  	d.added_cost, ol.status, ol.grand_total from order_list ol
    join public.user u on ol.user_id = u.user_id
    join public.payments p on ol.payment_id = p.payment_id
    join public.deliveries d on ol.delivery_id = d.delivery_id
	WHERE ol.orderlist_uuid=$1`
	data := models.GetOrder{}
	err := r.Get(&data, query, uuid)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *RepoOrder) GetHistoryOrder(id string) (*models.GetOrder, error) {
	query := `SELECT ol.orderlist_id, ol.orderlist_uuid, u.first_name, u.last_name, u.phone, u.address, u.email, ol.subtotal, ol.tax, p.payment_method, d.delivery_method,
  	d.added_cost, ol.status, ol.grand_total from order_list ol
    join public.user u on ol.user_id = u.user_id
    join public.payments p on ol.payment_id = p.payment_id
    join public.deliveries d on ol.delivery_id = d.delivery_id
	WHERE u.user_id=$1`
	data := models.GetOrder{}
	err := r.Get(&data, query, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *RepoOrder) UpdateOrder(data *models.Order, uuid string) (string, error) {
	query := `UPDATE public.order_list SET
		status = :status,
		updated_at = now()
	WHERE orderlist_uuid = :uuid`

	params := map[string]interface{}{
		"status":	data.Status,
		"uuid":     uuid,
	}

	_, err := r.NamedExec(query, params)
	if err != nil {
		return "", err
	}
	return "Data updated", nil
}