package routes

import (
	"crud_go/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/users", controllers.GetUsers)
	api.GET("/users/:id", controllers.GetUserDetails)
	api.POST("/users", controllers.AddUsers)
	api.PUT("/users/:id", controllers.EditUsers)
	api.DELETE("/users/:id", controllers.DeleteUsers)

	api.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Halo kamu berhasil lagi"})
	})
}
