package service

import (
	"product-app/domain"
	"product-app/persistence"
	"product-app/service/model"
)

type IProductService interface {
	Add(productCreate model.ProductCreate) error
	DeleteById(productId int64) error
	GetById(productId int64) (domain.Product, error)
	UpdatePrice(productId int64, newPrice float32) error
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func (p ProductService) Add(productCreate model.ProductCreate) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) DeleteById(productId int64) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) GetById(productId int64) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) UpdatePrice(productId int64, newPrice float32) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) GetAllProducts() []domain.Product {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) GetAllProductsByStore(storeName string) []domain.Product {
	//TODO implement me
	panic("implement me")
}
