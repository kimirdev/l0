package db

import (
	"l0/models"

	"github.com/jmoiron/sqlx"
)

type Order interface {
	GetAll() ([]models.OrderDTO, error)
	Create(data models.OrderDTO) error
	GetById(uid string) (models.OrderDTO, error)
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: NewOrderSql(db),
	}
}
