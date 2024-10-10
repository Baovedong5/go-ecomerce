package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/phuong/go-ecomerce/internal/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", c.NewUserController().GetUserById)
	}

	return r
}
