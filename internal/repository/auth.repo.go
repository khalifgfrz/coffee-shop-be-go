package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type AuthRepositoryInterface interface {
	RegisterUser(body *models.Auth) (string, error)
	LoginUser(email string) (*models.Auth, error)
}

type RepoAuth struct {
	*sqlx.DB
}

func NewAuth(db *sqlx.DB) *RepoAuth {
	return &RepoAuth{db}
}

func (r *RepoAuth) RegisterUser(body *models.Auth) (string, error) {
	query := `INSERT INTO public.user(email, password, phone, role) VALUES (:email, :password, :phone, 'user')`
	_, err := r.NamedExec(query, body)
	if err != nil {
		return "", err
	}
	return "data created", nil
}

func (r *RepoAuth) LoginUser(email string) (*models.Auth, error) {
	result := models.Auth{}
	query := `SELECT user_id, email, password, role FROM public.user WHERE email=$1`
	err := r.Get(&result, query, email)
	if err != nil {
		return nil, err
	}
	return &result, nil
}