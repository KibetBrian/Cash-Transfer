package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/KibetBrian/fisa/controllers"
)

var AcccountRoutes = func(router *gin.Engine) {
	router.POST("/account/create", controllers.CreateAccount)
	router.GET("/account/delete", controllers.DeleteAccount)
}