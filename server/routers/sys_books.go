package routers

import (

	"github.com/gin-gonic/gin"

	"github.com/nj-jay/httpServer/controllers"

	"github.com/nj-jay/httpServer/middlewares"

)

func LoadBooks(e *gin.Engine) {

	e.GET("/api/v2/books", middlewares.JWTAuthMiddleware(), controllers.QueryAllData)

	e.POST("/api/v2/book/search", middlewares.JWTAuthMiddleware(), controllers.QueryDataByName)

	e.POST("/api/v2/book/delete", middlewares.JWTAuthMiddleware(), controllers.DeleteSingleDataById)

	e.POST("/api/v2/book/update", middlewares.JWTAuthMiddleware(), controllers.UpdateSingleDataById)

	e.POST("/api/v2/book/add", middlewares.JWTAuthMiddleware(), controllers.PostSingleData)

}