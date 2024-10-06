package service

import (
	"03-lab/model"
	"03-lab/repository"
)

// Abstraction using ideomatic Go approach
// abstraction using oop approach

var repo repository.InMemoryProductRepo = repository.NewInMemoryProductRepo()

// var repo repository.DbProductRepo = repository.NewDbProductRepo()

type Finder interface {
	FindAll() []model.Product
}

func GetAllProducts(repo Finder) []model.Product {
	return repo.FindAll()
	// return repo.GetAllProducts()
}

func GetProductById(id int) *model.Product {
	// repo := repository.NewInMemoryProductRepo()
	return repo.FindBy(id)
	// return repo.GetProductById(id)
}

func AddProduct(name string, category model.Category, price float32) *model.Product {
	// repo := repository.NewInMemoryProductRepo()
	newProduct := model.Product{Name: name, Category: category, Price: price}

	np := repo.Save(newProduct)
	// np := repo.AddProduct(newProduct)
	return np
}
