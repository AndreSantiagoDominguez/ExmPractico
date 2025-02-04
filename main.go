package main

import (
	"exmpractico/src/products/infraestructure"
	"exmpractico/src/products/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	infraestructure.Init()
	r := gin.Default()
	routes.ProductRouter(r)

	r.Run()
}