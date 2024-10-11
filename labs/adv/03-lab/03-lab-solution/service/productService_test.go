package service

import (
	mocks "03-lab/mocks/service"
	"03-lab/model"
	"reflect"
	"testing"
)

func Test_should_ensure_the_price_of_product_is_increased_by_10_percent(t *testing.T) {
	// Arrange

	// mockRepo := ? // this comes from mocking framework
	mockRepo := mocks.NewProductRepository(t)
	mockRepo.On("FindAll").Return([]model.Product{{1, "prd1", model.BOOKS, 10.0}})
	service := NewProductService(mockRepo)
	expected := []model.Product{{1, "prd1", model.BOOKS, 11.0}}

	// Act & assert all together because this is an interaction test
	actual := service.GetAllProducts()

	// extra assert
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Price not increased by 10%")
	}

}
