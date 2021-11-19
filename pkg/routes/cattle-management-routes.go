package routes

import (
	"github.com/foekall/cattle-management/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterCattleManagementRoutes = func() {
	r := gin.Default()

	// r.HandleContext(c * gin.Context)
	//routes for user relate
	r.POST("/user", controllers.CreateUser)

	//routes for authentication
	r.POST("/auth", controllers.Auth)

	authRoutes := r.Group("api/v1")
	authRoutes.Use(controllers.TokenAuthMiddleware())
	authRoutes.GET("/user/:page/:size", controllers.GetAllUser)
	authRoutes.POST("/cattle", controllers.CreateCattle)
	authRoutes.GET("/cattle", controllers.GetAllCattles)
	authRoutes.GET("/cattle/:id", controllers.GetCattleById)
	authRoutes.PUT("/cattle/:id", controllers.UpdateCattle)
	authRoutes.DELETE("/cattle/:id", controllers.DeleteCattle)

	r.Run(":8080")
}
