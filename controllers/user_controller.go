package controllers

import (
	"dts/learn_middleware/database"
	"dts/learn_middleware/helpers"
	"dts/learn_middleware/models"
	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func UserRegister(c *gin.Context) {
	var user models.User

	db := database.GetDB()

	contentType := helpers.GetContentType(c)

	if contentType != appJson {
		return
	}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = db.Debug().Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func UserLogin(c *gin.Context) {
	var user models.User

	db := database.GetDB()

	contentType := helpers.GetContentType(c)

	if contentType != appJson {
		return
	}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	pwd := user.Password

	err = db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if !helpers.ComparePassword([]byte(user.Password), []byte(pwd)) {
		c.JSON(401, map[string]interface{}{
			"message": "invalid password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	c.JSON(200, map[string]interface{}{
		"token": token,
	})

}
