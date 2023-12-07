package repository

import (
	"database/sql"

	"github.com/guilherme1313/crud-golang/internal/entity"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{DB: db}
}

func (r *ProductRepositoryMysql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("INSERT into products (id, name, price) values (?,?,?)", product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, price from products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product

		err = rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (r *ProductRepositoryMysql) FindOne(id string) (*entity.Product, error) {
	row, err := r.DB.Query("SELECT * from products where id=?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var product entity.Product

	for row.Next() {
		err = row.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
	}
	return &product, nil
}

func (r *ProductRepositoryMysql) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM products WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryMysql) Update(product *entity.Product, id string) (*int64, error) {
	result, err := r.DB.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, id)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &rowsAffected, nil
}
