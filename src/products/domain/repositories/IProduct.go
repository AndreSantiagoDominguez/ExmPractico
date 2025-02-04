package repositories

type IProduct interface {
	createProduct(product string)
	GetAllProduct() ([]IProduct, error)
}
