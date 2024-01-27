package products

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	getProducts() ([]ProductResponse, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) getProducts() ([]ProductResponse, error) {
	fmt.Println(" Get Products function")
	var allproducts []ProductResponse
	err := r.db.Find(&allproducts).Error
	fmt.Println(allproducts, " get products from db")
	return allproducts, err
}
