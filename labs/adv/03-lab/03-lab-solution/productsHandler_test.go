package main

import (
	"03-lab/mocks"
	"03-lab/model"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_should_test_get_all_products_return_200_OK(t *testing.T) {
	// arrange
	mockSvc := mocks.NewProductService(t)
	mockSvc.On("GetAllProducts").Return([]model.Product{{1, "prd1", model.BOOKS, 10.0}})
	pr := NewProductRouter(mockSvc)

	r := gin.Default()
	r.GET("/products", pr.productHandler)

	request, _ := http.NewRequest(http.MethodGet, "/products", nil)
	recorder := httptest.NewRecorder()

	// act
	r.ServeHTTP(recorder, request)

	//assert
	if recorder.Result().StatusCode != http.StatusOK {
		t.Error("failed, because status code is not 200")
	}
}

func Test_addProduct_serialize_incoming_request(t *testing.T) {
	// arrange
	mockSvc := mocks.NewProductService(t)

	pr := NewProductRouter(mockSvc)

	r := gin.Default()
	r.POST("/products", pr.newProductHandler)

	badRequest := bytes.NewReader([]byte("some junk"))
	request, _ := http.NewRequest(http.MethodPost, "/products", badRequest)
	recorder := httptest.NewRecorder()

	// act
	r.ServeHTTP(recorder, request)

	//assert
	if recorder.Result().StatusCode != http.StatusBadRequest {
		t.Error("failed, because status code is not 400")
	}
}
