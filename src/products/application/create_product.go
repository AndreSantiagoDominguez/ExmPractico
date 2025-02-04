package application

import (
	"exmpractico/src/products/domain/entities"
	"exmpractico/src/products/domain/repositories"
)

type CreateProduct struct{
	db repositories.IProduct
}

func NewCreateProduct(db repositories.IProduct) *CreateProduct {
	return &CreateProduct{db: db}
}

func (cp *CreateProduct) Run(product *entities.Product) (int64,error){
	return cp.db.CreateProduct(product)
}