package controllers

import (
	"dts/learn_middleware/database"
	"dts/learn_middleware/helpers"
	"dts/learn_middleware/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()

	var product models.Product

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	if contentType != appJson {
		c.AbortWithError(400, errors.New("invalid content type"))
		return
	}

	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	userId := uint(userData["id"].(float64))
	product.UserId = userId
	err = db.Debug().Create(&product).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, product)

}
