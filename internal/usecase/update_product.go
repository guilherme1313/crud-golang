package usecase

import (
	"github.com/guilherme1313/crud-golang/internal/entity"
	"github.com/guilherme1313/crud-golang/internal/repository"
)

type UpdateProductInputDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductOutputDto struct {
	RowsAffected int64
}

type UpdateProductUseCase struct {
	ProductRepository repository.ProductRepository
}

func NewUpdateProductUseCase(productRepository repository.ProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{ProductRepository: productRepository}
}

func (u *UpdateProductUseCase) Execute(input UpdateProductInputDto, id string) (*UpdateProductOutputDto, error) {
	var product *entity.Product = &entity.Product{
		ID:    id,
		Name:  input.Name,
		Price: input.Price,
	}
	result, err := u.ProductRepository.Update(product, id)

	if err != nil {
		return nil, err
	}

	response := &UpdateProductOutputDto{RowsAffected: *result}

	return response, nil
}
