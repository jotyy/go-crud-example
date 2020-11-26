package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"github.com/gin-gonic/gin"
	"github.com/jotyy/go-crud-example/api"
	"github.com/jotyy/go-crud-example/middleware"
	"github.com/jotyy/go-crud-example/model"
)

func main() {
	model.InitDB()

	gin.ForceConsoleColor()

	router := gin.Default()

	router.POST("/register", api.Register)
	router.POST("/login", api.Login)
	router.DELETE("/unregister", api.UnRegister)

	v1 := router.Group("/api/v1/")
	v1.Use(middleware.SetMiddlewareAuthentication())
	{
		v1.GET("articles", api.GetArticles)
		v1.POST("articles", api.AddArticles)
	}

	router.Run(":8999")
}
