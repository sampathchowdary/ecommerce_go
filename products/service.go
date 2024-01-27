package products

import (
	"fmt"
)

type Service interface {
	getProducts() ([]ProductResponse, error)
}

type service struct {
	repository Repository
}

func (s *service) getProducts() ([]ProductResponse, error) {
	fmt.Println(" Get Products Service function")
	products, err := s.repository.getProducts()
	return products, err

}
