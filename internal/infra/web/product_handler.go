package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/guilherme1313/crud-golang/internal/usecase"
)

type ProductHandlers struct {
	CreateProductUseCase  *usecase.CreateProductUseCase
	ListProductsUseCase   *usecase.ListProductsUseCase
	FindOneProductUseCase *usecase.FindOneProductUseCase
	DeleteProductUseCase  *usecase.DeleteProductUseCase
	UpdateProductUseCase  *usecase.UpdateProductUseCase
}

func NewProductHandlers(createproductUseCase *usecase.CreateProductUseCase, listProductsUseCase *usecase.ListProductsUseCase, findOneProductUseCase *usecase.FindOneProductUseCase, deleteProductUseCase *usecase.DeleteProductUseCase, updateProductUseCase *usecase.UpdateProductUseCase) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUseCase:  createproductUseCase,
		ListProductsUseCase:   listProductsUseCase,
		FindOneProductUseCase: findOneProductUseCase,
		DeleteProductUseCase:  deleteProductUseCase,
		UpdateProductUseCase:  updateProductUseCase,
	}
}

func (p *ProductHandlers) CreateProductHandlers(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := p.CreateProductUseCase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandlers) ListProductHandlers(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListProductsUseCase.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandlers) FindOneProductHandlers(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	output, err := p.FindOneProductUseCase.Execute(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandlers) DeleteProductHandlers(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	response, err := p.DeleteProductUseCase.Execute(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (p *ProductHandlers) UpdateProductHandlers(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var input usecase.UpdateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := p.UpdateProductUseCase.Execute(input, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
