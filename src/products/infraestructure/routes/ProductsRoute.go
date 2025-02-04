package routes

import (
	"exmpractico/src/products/infraestructure"

	"github.com/gin-gonic/gin"
)

func ProductRouter(router *gin.Engine) {

	// Nombre del recurso
	routes := router.Group("/products")

	// Instancias de los controladores
	createProductController := infraestructure.CreateProductController().Run
	
	// Rutas
	routes.POST("", createProductController)

}