package usecase

import (
	"github.com/guilherme1313/crud-golang/internal/repository"
)

type DeleteProductOutputDto struct {
	Delete bool `json:"delete"`
}

type DeleteProductUseCase struct {
	ProductRepository repository.ProductRepository
}

func NewDeleteProductUseCase(productRepository repository.ProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{ProductRepository: productRepository}
}

func (u *DeleteProductUseCase) Execute(id string) (*DeleteProductOutputDto, error) {
	err := u.ProductRepository.Delete(id)

	if err != nil {
		return nil, err
	}

	response := DeleteProductOutputDto{Delete: true}

	return &response, nil
}
