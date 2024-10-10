package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/phuong/go-ecomerce/internal/controllers"
	"github.com/phuong/go-ecomerce/internal/middleware"
)

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before  --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before  --> CC")
	c.Next()
	fmt.Println("After --> CC")
}

func NewRouter() *gin.Engine {
	r := gin.Default()

	//use the middleware
	r.Use(middleware.AuthenMiddleware(), BB(), CC)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", c.NewUserController().GetUserById)
	}

	return r
}
