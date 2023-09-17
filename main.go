package main

import (
	"G05-food-delivery/component/appctx"
	"G05-food-delivery/component/uploadprovider"
	"G05-food-delivery/middleware"
	ginupload2 "G05-food-delivery/module/upload/transport/ginupload"
	"G05-food-delivery/module/upload/uploadtransport/ginupload"
	"G05-food-delivery/module/user/transport/ginuser"

	//restaurantmodel "G05-food-delivery/module/restaurant/model"
	"G05-food-delivery/module/restaurant/transport/ginrestaurant"
	//"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	//"strconv"
)



//go get gorm.io/gorm@v1.20.11
//go get gorm.io/driver/mysql@v1.0.3
//go get -u github.com/gin-gonic/gin@v1.7.1
//go get -u github.com/btcsuite/btcutil/base58

func main() {

	////json.Marshal & json.Unmarshal
	//
	//test := restaurantmodel.Restaurant{
	//	Name: "200lab",
	//	Addr: "somewhere",
	//}
	//
	//jsByte, err := json.Marshal(test) // {"id":1,"name":"200lab","addr":"somewhere"}
	//log.Println(string(jsByte), err)
	//json.Unmarshal([]byte("{\"name\":\"200lab Edu\",\"addr\":\"somewhere\"}"), &test)
	//log.Println(test)
	//log.Println(test.Name)
	//log.Println(test.Addr)




	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3308)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
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

	s3Provider := uploadprovider.NewS3Provider(s3BucketName,s3Region,s3APIKey,s3SecretKey,s3Domain)

	appContext := appctx.NewAppContext(db,s3Provider,secretKey)

	r := gin.Default()

	r.Use(middleware.Recover(appContext))


	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Static("/ahihi","./static")

	v1 := r.Group("/v1")

	v1.POST("/upload",ginupload.Upload(appContext))

	v1.POST("/uploadlocal",ginupload2.UploadImage(appContext))

	v1.POST("/register",ginuser.Register(appContext))

	v1.POST("/authenticate",ginuser.Login(appContext))

	restaurant := v1.Group("/restaurants")

	restaurant.POST("", ginrestaurant.CreateRestaurant(appContext))

	//restaurant.GET("/:id", func(c *gin.Context) {
	//	id, err := strconv.Atoi(c.Param("id"))
	//
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"err": err.Error(),
	//		})
	//
	//		return
	//	}
	//
	//	var data Restaurant
	//
	//	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"err": err.Error(),
	//		})
	//
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"data": data,
	//	})
	//})
	//
	restaurant.GET("", ginrestaurant.ListRestaurant(appContext))

	//restaurant.PATCH("/:id", func(c *gin.Context) {
	//	id, err := strconv.Atoi(c.Param("id"))
	//
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"err": err.Error(),
	//		})
	//
	//		return
	//	}
	//
	//	var data RestaurantUpdate
	//
	//	if err := c.ShouldBindJSON(&data); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"err": err.Error(),
	//		})
	//
	//		return
	//	}
	//
	//	if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"err": err.Error(),
	//		})
	//
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"data": true,
	//	})
	//})
	//
	restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	r.Run()

	//log.Println(db,err)

	//newRestaurant := Restaurant{
	//	Name: "Tani",
	//	Addr: "9 Pham Van Hai",
	//}
	//
	//if err := db.Create(&newRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println("New ID Created: ", newRestaurant.Id)
	//
	//var myRestaurant Restaurant
	//
	//if err := db.Where("id = ?", 2).First(&myRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(myRestaurant)

	//myRestaurant.Name = "500lab"

	//if err := db.Where("id = ?", 2).Updates(&myRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(myRestaurant)

	//update := "500lab"
	//updateData := RestaurantUpdate{Name: &update}
	//if err := db.Where("id = ?", 2).Updates(&updateData).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(*updateData.Name)
	//
	//if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 5).Delete(nil).Error; err != nil {
	//	log.Println(err)
	//}

}

//change CHARACTER & COLLATE table
//ALTER TABLE restaurants CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

//CREATE TABLE `restaurants` (
//`id` int(11) NOT NULL AUTO_INCREMENT,
//`owner_id` int(11) NULL,
//`name` varchar(50) NOT NULL,
//`addr` varchar(255) NOT NULL,
//`city_id` int(11) DEFAULT NULL,
//`lat` double DEFAULT NULL,
//`lng` double DEFAULT NULL,
//`cover` json NULL,
//`logo` json NULL,
//`shipping_fee_per_km` double DEFAULT '0',
//`status` int(11) NOT NULL DEFAULT '1',
//`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//PRIMARY KEY (`id`),
//KEY `owner_id` (`owner_id`) USING BTREE,
//KEY `city_id` (`city_id`) USING BTREE,
//KEY `status` (`status`) USING BTREE
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;
