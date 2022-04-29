package web

import (
	"github.com/gin-gonic/gin"
	Models "go-micro-gin-gateway/models"
	"go-micro-gin-gateway/web/handlers"
	"go-micro-gin-gateway/web/middlewares"
	"net/http"
)

func NewRouter(userListService Models.UserCommonService) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middlewares.UserMiddleware(userListService))
	ginRouter.Handle("GET", "/users/:size", handlers.GetUsersHandler)
	ginRouter.Handle(http.MethodPost, "/UserCommonService/GetUserDetail", handlers.GetUserDetailHandler)
	return ginRouter
}
