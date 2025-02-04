package repositories

import "exmpractico/src/products/domain/entities"

type IProduct interface {
	CreateProduct(product *entities.Product) (int64, error)
}
