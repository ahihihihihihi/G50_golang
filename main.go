package main

import (
	"G05-food-delivery/component/appctx"
	"G05-food-delivery/component/uploadprovider"
	"G05-food-delivery/middleware"
	//restaurantmodel "G05-food-delivery/module/restaurant/model"
	//"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	//"strconv"
)

func main() {

	dsn := os.Getenv("MYSQL_CONN_STRING")

	s3BucketName := os.Getenv("s3BucketName")
	s3APIKey := os.Getenv("s3APIKey")
	s3Region := os.Getenv("s3Region")
	s3SecretKey := os.Getenv("s3SecretKey")
	s3Domain := os.Getenv("s3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET") // ahihi

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// debug database mode
	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appContext := appctx.NewAppContext(db, s3Provider, secretKey)

	r := gin.Default()

	r.Use(middleware.Recover(appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Static("/ahihi", "./static")

	v1 := r.Group("/v1")

	setupRoute(appContext, v1)

	setupAdminRoute(appContext, v1)

	r.Run()

}

