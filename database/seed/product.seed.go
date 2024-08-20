package seed

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedProducts(db *sqlx.DB) error {
	products := []models.Product{
		{
		Product_name	: "Cold Brew",
		Price			: 30000,
		Category		: "Coffee",
		Description		: "Cold brewing is a method of brewing that combines ground coffee and cool water and uses time instead of heat to extract the flavor.",
		Stock			: 200,
		},
		{
		Product_name	: "Veggie Tomato Mix",
		Price			: 34000,
		Category		: "Food",
		Description		: "Veggie tomato mix is a food and always be food.",
		Stock			: 200,
		},
		{
		Product_name	: "Hazelnut Latte",
		Price			: 25000,
		Category		: "Coffee",
		Description		: "Hazelnut Latte is a hazelnut and always be hazelnut.",
		Stock			: 200,
		},
		{
		Product_name	: "Summer Fried Rice",
		Price			: 32000,
		Category		: "Food",
		Description		: "Summer fried rice is a food and always be food.",
		Stock			: 200,
		},
		{
		Product_name	: "Creamy Ice Latte",
		Price			: 27000,
		Category		: "Coffee",
		Description		: "Creamy Ice Latte is a creamy and always be creamy.",
		Stock			: 200,
		},
		{
		Product_name	: "Drum Sticks",
		Price			: 30000,
		Category		: "Food",
		Description		: "Drum sticks is a food and always be food.",
		Stock			: 200,
		},
		{
		Product_name	: "Salty Rice",
		Price			: 20000,
		Category		: "Food",
		Description		: "Salty Rice is a food and always be food.",
		Stock			: 200,
		},
		{
		Product_name	: "Green Tea Latte",
		Price			: 35000,
		Category		: "Non coffee",
		Description		: "Smooth and creamy matcha is lightly sweetened and served with steamed milk.",
		Stock			: 200,
		},
		{
		Product_name	: "Black Tea Latte",
		Price			: 30000,
		Category		: "Non coffee",
		Description		: "A select blend of rich, full leaf black teas are lightly sweetened with liquid cane sugar and topped with steamed milk and a velvety foam.",
		Stock			: 200,
		},
		{
		Product_name	: "Chocolate",
		Price			: 45000,
		Category		: "Non coffee",
		Description		: "Chocolate is a sweet drink and always be chocolate.",
		Stock			: 200,
		},
	}

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

	for _, product := range products {
		_, err := db.NamedExec(query, product)
		if err != nil {
			return err
		}
	}

	log.Println("Seeding products completed successfully.")
	return nil
}