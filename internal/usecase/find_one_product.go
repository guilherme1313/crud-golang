package usecase

import (
	"github.com/guilherme1313/crud-golang/internal/repository"
)

type FindOneProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type FindOneProductUseCase struct {
	ProductRepository repository.ProductRepository
}

func NewFindOneProductUseCase(productRepository repository.ProductRepository) *FindOneProductUseCase {
	return &FindOneProductUseCase{ProductRepository: productRepository}
}

func (u *FindOneProductUseCase) Execute(id string) (*FindOneProductOutputDto, error) {
	product, err := u.ProductRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return &FindOneProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
