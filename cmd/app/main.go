package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/guilherme1313/crud-golang/internal/infra/repository"
	"github.com/guilherme1313/crud-golang/internal/infra/web"
	"github.com/guilherme1313/crud-golang/internal/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3309)/db_go?charset=utf8")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listProductsUseCase := usecase.NewListProductsUseCase(repository)
	findOneProductUseCase := usecase.NewFindOneProductUseCase(repository)
	deleteProductUseCase := usecase.NewDeleteProductUseCase(repository)
	updateProductUseCase := usecase.NewUpdateProductUseCase(repository)

	productsHadlers := web.NewProductHandlers(createProductUseCase, listProductsUseCase, findOneProductUseCase, deleteProductUseCase, updateProductUseCase)

	r := chi.NewRouter()
	r.Post("/products", productsHadlers.CreateProductHandlers)
	r.Get("/products", productsHadlers.ListProductHandlers)
	r.Get("/products/{id}", productsHadlers.FindOneProductHandlers)
	r.Delete("/products/{id}", productsHadlers.DeleteProductHandlers)
	r.Put("/products/{id}", productsHadlers.UpdateProductHandlers)

	http.ListenAndServe(":3000", r)
}
