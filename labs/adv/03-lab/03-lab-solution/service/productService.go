package service

import (
	"03-lab/model"
)

// abstractions using ideomatic go
// abstraction using oop way

// goal is to get rid of public repo
// var repo repository.InMemoryProductRepo = repository.NewInMemoryProductRepo()

type Finder interface {
	FindAll() []model.Product
}

type SingleFinder interface {
	FindBy(id int) *model.Product
}

type ProductSaver interface {
	Save(newProduct model.Product) *model.Product
}

type ProductRepository interface {
	Finder
	SingleFinder
	ProductSaver
}

type DefaultProductService struct {
	repo ProductRepository
}

func (ps *DefaultProductService) GetAllProducts() []model.Product {
	products := ps.repo.FindAll()
	for idx, p := range products {
		products[idx].Price = p.Price * 1.1
	}
	return products
	// return ps.repo.FindAll()
}

func (ps *DefaultProductService) GetProductById(id int) *model.Product {
	return ps.repo.FindBy(id)
}

func (ps *DefaultProductService) AddProduct(name string, category model.Category, price float32) *model.Product {
	newProduct := model.Product{Name: name, Category: category, Price: price}

	return ps.repo.Save(newProduct)
}

func NewProductService(repo ProductRepository) DefaultProductService {
	return DefaultProductService{repo: repo}
}

// what is the dependency for this?
// func GetAllProducts(repo Finder) []model.Product {
// 	return repo.FindAll()
// }

// func GetProductById(repo SingleFinder, id int) *model.Product {

// 	return repo.FindBy(id)
// }

// func AddProduct(repo ProductSaver, name string, category model.Category, price float32) *model.Product {
// 	newProduct := model.Product{Name: name, Category: category, Price: price}

// 	return repo.Save(newProduct)
// }
