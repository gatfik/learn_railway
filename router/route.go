package router

import (
	"dts/learn_middleware/controllers"
	"dts/learn_middleware/middleware"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	{
		userRouter.POST("/", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/product")
	{
		productRouter.Use(middleware.AuthVerify())
		productRouter.POST("/", controllers.CreateProduct)
	}
	return r
}
