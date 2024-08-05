package repository

import (
	"database/sql"
	"fmt"
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
	return err
}

func (r *RepoProduct) GetAllProduct(que *models.ProductQuery) (*models.Products, error) {
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
	
	rows, err := r.DB.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data models.Products
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.Product_id, 
			&product.Product_uuid, 
			&product.Product_name, 
			&product.Price, 
			&product.Category, 
			&product.Description, 
			&product.Stock, 
			&product.Created_at, 
			&product.Updated_at,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoProduct) GetDetailProduct(id int) (*models.Product, error) {
	query := `SELECT * FROM public.product WHERE product_id = :product_id`
	data := models.Product{}

	rows, err := r.DB.NamedQuery(query, map[string]interface{}{
		"product_id": id,
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
func (r *RepoProduct) DeleteProduct(id int) error {
	query := `DELETE FROM public.product WHERE product_id = :product_id`

	_, err := r.DB.NamedExec(query, map[string]interface{}{
		"product_id": id,
	})
	return err
}

func (r *RepoProduct) UpdateProduct(data *models.Product, id int) (*models.Product, error) {
	query := `UPDATE public.product SET`
	var values []interface{}
	condition := false

	if data.Product_name != "" {
		query += fmt.Sprintf(` product_name = $%d`, len(values)+1)
		values = append(values, data.Product_name)
		condition = true
	}
	if data.Price != 0 {
		if condition {
			query += ","
		}
		query += fmt.Sprintf(` price = $%d`, len(values)+1)
		values = append(values, data.Price)
		condition = true
	}
	if data.Category != "" {
		if condition {
			query += ","
		}
		query += fmt.Sprintf(` category = $%d`, len(values)+1)
		values = append(values, data.Category)
		condition = true
	}
	if data.Description != "" {
		if condition {
			query += ","
		}
		query += fmt.Sprintf(` description = $%d`, len(values)+1)
		values = append(values, data.Description)
		condition = true
	}
	if data.Stock != 0 {
		if condition {
			query += ","
		}
		query += fmt.Sprintf(` stock = $%d`, len(values)+1)
		values = append(values, data.Stock)
		condition = true
	}

	if !condition {
		return nil, fmt.Errorf("no fields to update")
	}

	query += fmt.Sprintf(`, updated_at = now() WHERE product_id = $%d RETURNING *`, len(values)+1)
	values = append(values, id)

	row := r.DB.QueryRow(query, values...)
	var product models.Product
	err := row.Scan(
		&product.Product_id,
		&product.Product_uuid,
		&product.Product_name,
		&product.Price,
		&product.Category,
		&product.Description,
		&product.Stock,
		&product.Created_at,
		&product.Updated_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(`product with id %d not found`, id)
		}
		return nil, fmt.Errorf(`query execution error: %w`, err)
	}

	return &product, nil
}