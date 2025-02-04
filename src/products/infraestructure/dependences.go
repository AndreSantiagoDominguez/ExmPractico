package infraestructure

import (
	"exmpractico/src/products/infraestructure/controllers"
	"exmpractico/src/products/application"
)

var (
	mysql *MySQL
)

func Init() {
	mysql = NewMySQL()
}

// Casos de uso
func CreateProductController() *controllers.CreateProductController {
	useCaseCreateProduct := application.NewCreateProduct(mysql)
	
	return controllers.NewCreateProductController(useCaseCreateProduct)
}