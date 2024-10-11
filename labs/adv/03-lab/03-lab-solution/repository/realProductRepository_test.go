package repository

import (
	"03-lab/model"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("thecodecamp/customers-db", "latest", []string{"MYSQL_ROOT_PASSWORD=thecodecamp"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sqlx.Open("mysql", fmt.Sprintf("root:thecodecamp@(localhost:%s)/banking", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func Test_should_ensure_string_values_get_converted_to_enum(t *testing.T) {
	// Arrange
	create := `CREATE TABLE products (
		  id int(11) NOT NULL AUTO_INCREMENT,
		  name varchar(20) NOT NULL,
		  category varchar(20) NOT NULL,
		  PRIMARY KEY (id)
		) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1`
	db.Exec(create)
	insert := `insert into products (name, category) values
		("prd1", "BOOKS"),
	  ("prd2", "FOOD"),
	  ("prd3", "TOYS")`
	db.Exec(insert)

	prodRepo := NewProductRepository(db)
	expected := []model.Product{
		{1, "prd1", 1, 0.0},
		{2, "prd2", 0, 0.0},
		{3, "prd3", model.TOYS, 0.0},
	}
	// Act
	actual, err := prodRepo.FindAll()
	if err != nil {
		t.Error("Cannot fetch products: ", err.Error())
	}

	// Assert
	if !reflect.DeepEqual(expected, actual) {
		t.Error("expected product list is not matching actual")
	}
}
