package repository

import "github.com/guilherme1313/crud-golang/internal/entity"

type ProductRepository interface {
	Create(product *entity.Product) error
	FindAll() ([]*entity.Product, error)
}
