package router

import (
	"jwt-assignment/controller"
	"jwt-assignment/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
	}
	ProductRouter := r.Group("/products")
	{
		ProductRouter.Use(middleware.Authentication())
		ProductRouter.POST("/", controller.CreateProduct)
		ProductRouter.POST("/:id", middleware.ProductAuthorization(), controller.UpdateProduct)
	}
	return r
}
