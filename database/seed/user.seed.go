package seed

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/pkg"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedUsers(db *sqlx.DB) error {
	users := []models.User{
		{
		First_name	: "Test",
		Last_name	: "Example1",
		Phone		: "123456789",
		Address		: "123 Main St",
		Birth_date	: "2000-01-01",
		Email		: "test@example1.com",
		Password	: "password1",
		Role		: "admin",
		},
		{
		First_name	: "Test",
		Last_name	: "Example2",
		Phone		: "123456788",
		Address		: "1234 Main St",
		Birth_date	: "2000-01-02",
		Email		: "test@example2.com",
		Password	: "password2",
		Role		: "admin",
		},
		{
		First_name	: "Test",
		Last_name	: "Example3",
		Phone		: "123456787",
		Address		: "12345 Main St",
		Birth_date	: "2000-01-03",
		Email		: "test@example3.com",
		Password	: "password3",
		Role		: "admin",
		},
	}

	query := `INSERT INTO public.user(
		first_name,
		last_name,
		phone,
		address,
		birth_date,
		email,
		password,
		role
	) VALUES(
	 	:first_name,
		:last_name,
		:phone,
		:address,
		:birth_date,
		:email,
		:password,
		:role
	)`

	for i, user := range users {
        hashedPassword, err := pkg.HashPassword(user.Password)
        if err != nil {
            return err
        }
        users[i].Password = hashedPassword

        _, err = db.NamedExec(query, users[i])
        if err != nil {
            return err
        }
    }

	log.Println("Seeding users completed successfully.")
	return nil
}