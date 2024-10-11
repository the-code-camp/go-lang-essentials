package main

import (
	"03-lab/model"
	"03-lab/repository"
	"03-lab/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	GetAllProducts() []model.Product
	GetProductById(id int) *model.Product
	AddProduct(name string, category model.Category, price float32) *model.Product
}

type ProductRouter struct {
	service ProductService
}

type CustomerRouter struct {
	// dependency of repository
	repo repository.CustomerRepositoryDb
}

func NewProductRouter(svc ProductService) ProductRouter {
	return ProductRouter{service: svc}
}

func NewCustomerRouter(repo repository.CustomerRepositoryDb) CustomerRouter {
	return CustomerRouter{repo}
}

func test_sql_enum_scan() {
	/**

		CREATE TABLE `products` (
		  `id` int(11) NOT NULL AUTO_INCREMENT,
		  `name` varchar(20) NOT NULL,
		  `category` varchar(20) NOT NULL,
		  PRIMARY KEY (`id`)
		) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

		insert into products (name, category) values
		("prd1", "BOOKS"),
	  ("prd2", "FOOD"),
	  ("prd3", "TOYS")

	**/

	dbClient := NewDbClient()

	getProducts := func() {
		var err error
		products := make([]model.Product, 0)
		findAllSql := "select id, name, category from products"
		err = dbClient.Select(&products, findAllSql)
		if err != nil {
			log.Fatal("Some error: ", err.Error())
		}
		for _, p := range products {
			fmt.Println(p)
		}
	}
	getProducts()
}

func main() {

	var repo service.ProductRepository // what should be the type ?

	// creating dependencies
	if os.Getenv("ENV") == "prod" {
		repo = repository.NewDbProductRepo()
	} else if os.Getenv("ENV") == "dev" {
		repo = repository.NewInMemoryProductRepo()
	} else {
		log.Fatal("environment variable not defined. sanity check failed....")
	}

	// wiring
	ps := service.NewProductService(repo)
	pr := NewProductRouter(&ps)

	dbClient := NewDbClient()
	customerRepo := repository.NewCustomerRepositoryDB(dbClient)
	cr := NewCustomerRouter(customerRepo)

	// router/multiplexer
	r := gin.Default()

	// registering routes
	r.GET("/ping", pingHandler)
	r.GET("/products", pr.productHandler)
	r.GET("/products/:id", pr.singleProductHandler)
	r.POST("/products", pr.newProductHandler)

	r.GET("/customers", cr.customersHandler) // implement cr.customersHandler

	r.Run()
}

func (pr *ProductRouter) newProductHandler(c *gin.Context) {
	var p model.Product
	if c.ShouldBindJSON(&p) == nil {
		newProduct := pr.service.AddProduct(p.Name, p.Category, p.Price)
		c.JSON(http.StatusCreated, newProduct)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error_msg": "bad request"})
	}
}

func (pr *ProductRouter) singleProductHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	productId, _ := strconv.ParseInt(id, 10, 16)

	product := pr.service.GetProductById(int(productId))
	if product != nil {
		c.JSON(http.StatusOK, product)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error_msg": "Product Id does not exist"})
	}
}

func (pr *CustomerRouter) customersHandler(c *gin.Context) {
	customers, err := pr.repo.FindAll("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error_msg": "something unexpected..."})
	} else {
		c.JSON(http.StatusOK, customers)
	}
}

func (pr *ProductRouter) productHandler(c *gin.Context) {
	products := pr.service.GetAllProducts()
	c.JSON(http.StatusOK, products)
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
