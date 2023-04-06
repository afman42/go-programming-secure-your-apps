package repository

import "testify_test/entity"

//go:generate mockery --name entity.Product
type ProductRepository interface {
	FindById(id string) *entity.Product
	FindAll() []*entity.Product
}
