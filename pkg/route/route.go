package route

import (

	controller "github.com/adefemi171/postgres-go/pkg/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", controller.Welcome)
	router.POST("/user", controller.CreateUser)
	router.GET("/users", controller.GetUser)
	router.GET("/user/:userId", controller.GetUserByID)
	// router.PUT("/user/:userId", controller.EditUser)
	router.DELETE("/user/:userId", controller.DeleteUser)
	router.NoRoute(controller.NotFound)
}
