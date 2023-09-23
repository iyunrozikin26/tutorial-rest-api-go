package route

import (
	"github.com/gin-gonic/gin"
	user "github.com/iyunrozikin26/tutorial-rest-api-go.git/src/models/user"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func Api(router *gin.Engine, db *gorm.DB) {
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userController := user.NewUserController(userService, ctx)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", userController.Index)
		v1.GET("/users/:id", userController.GetById)
		v1.POST("/users", userController.Create)
		v1.PATCH("/users/:id", userController.Update)
		v1.DELETE("/users/:id", userController.Delete)
	}
}
