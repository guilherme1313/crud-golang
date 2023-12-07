package repository

import (
	"github.com/guilherme1313/crud-golang/internal/entity"
)

type ProductRepository interface {
	Create(product *entity.Product) error
	FindAll() ([]*entity.Product, error)
	FindOne(id string) (*entity.Product, error)
	Delete(id string) error
	Update(product *entity.Product, id string) (*int64, error)
}
