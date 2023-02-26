package routes

import (
	controller_blog "gdsc/controllers"

	"github.com/gin-gonic/gin"
)

func BlogRoute(router *gin.RouterGroup) {
router.POST("/blog", controller_blog.BlogPost())
}