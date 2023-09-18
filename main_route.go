package main

import (
	"G05-food-delivery/component/appctx"
	"G05-food-delivery/middleware"
	"G05-food-delivery/module/restaurant/transport/ginrestaurant"
	ginupload2 "G05-food-delivery/module/upload/transport/ginupload"
	"G05-food-delivery/module/upload/uploadtransport/ginupload"
	"G05-food-delivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
)

func setupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup)  {
	v1.POST("/upload", ginupload.Upload(appContext))

	v1.POST("/uploadlocal", ginupload2.UploadImage(appContext))

	v1.POST("/register", ginuser.Register(appContext))

	v1.POST("/authenticate", ginuser.Login(appContext))

	v1.GET("/profile", middleware.RequireAuth(appContext), ginuser.Profile(appContext))

	restaurant := v1.Group("/restaurants", middleware.RequireAuth(appContext))

	restaurant.POST("", ginrestaurant.CreateRestaurant(appContext))

	restaurant.GET("", ginrestaurant.ListRestaurant(appContext))

	restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
}
