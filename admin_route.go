package main

import (
	"G05-food-delivery/component/appctx"
	"G05-food-delivery/middleware"
	"G05-food-delivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
)

func setupAdminRoute(appContext appctx.AppContext, v1 *gin.RouterGroup)  {
	admin := v1.Group("/admin",
		middleware.RequireAuth(appContext),
		middleware.RoleRequired(appContext, "admin", ",mod"),
	)

	{
		admin.GET("/profile", ginuser.Profile(appContext))
	}
}
