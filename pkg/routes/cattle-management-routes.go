package routes

import (
	"github.com/foekall/cattle-management/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterCattleManagementRoutes = func() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	v1.POST("/cattle", controllers.CreateCattle)
	v1.GET("/cattle", controllers.GetAllCattles)
	v1.GET("/cattle/:id", controllers.GetCattleById)
	v1.PUT("/cattle/:id", controllers.UpdateCattle)
	v1.DELETE("/cattle/:id", controllers.DeleteCattle)

	//routes for user relate
	v1.POST("/user", controllers.CreateUser)
	v1.GET("/user/:page/:size", controllers.GetAllUser)

	//routes for authentication
	v1.POST("/auth", controllers.Auth)

	r.Run(":8080")
}
