package repository

import (
	"03-lab/model"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	client *sqlx.DB
}

func (p ProductRepository) FindAll() ([]model.Product, error) {
	var err error

	products := make([]model.Product, 0)
	findAllSql := "SELECT id, name, category FROM products"
	err = p.client.Select(&products, findAllSql)

	if err != nil {
		log.Println("Error while querying products table " + err.Error())
		return nil, errors.New("unexpected database error")
	}

	return products, nil
}

func NewProductRepository(client *sqlx.DB) ProductRepository {
	return ProductRepository{client}
}
