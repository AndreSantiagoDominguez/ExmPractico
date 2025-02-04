package controllers

import (
	"exmpractico/src/products/application"
	"exmpractico/src/products/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	useCase *application.CreateProduct
}

func NewCreateProductController(useCase *application.CreateProduct) *CreateProductController {
	return &CreateProductController{useCase: useCase}
}

func (cp_c *CreateProductController) Run(ctx *gin.Context) {
	var product entities.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if product.Name == "" && product.Quantity <= 0 && product.Bcode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Campos vacíos o no válidos"})
		return
	}

	_, err := cp_c.useCase.Run(&product)


	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Message": "Created Product"})


}