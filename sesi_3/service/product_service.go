package service

import (
	"errors"
	"testify_test/entity"
	"testify_test/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (service ProductService) GetOneProductId(id string) (*entity.Product, error) {
	product := service.ProductRepository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (service ProductService) GetAllProduct() ([]*entity.Product, error) {
	product := service.ProductRepository.FindAll()
	if product == nil {
		return nil, errors.New("all product not found")
	}

	return product, nil
}
