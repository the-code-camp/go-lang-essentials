package repository

import (
	"03-lab/model"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]model.Customer, error) {
	// var rows *sql.Rows
	var err error

	customers := make([]model.Customer, 0)
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		log.Println("Error while querying customers table " + err.Error())
		return nil, errors.New("unexpected database error")
	}

	// for rows.Next() {
	// 	var c model.Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	// 	if err != nil {
	// 		log.Println("Error while scanning customers " + err.Error())
	// 		return nil, errors.New("unexpected database error")
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil
}

func NewCustomerRepositoryDB(client *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client}
}
