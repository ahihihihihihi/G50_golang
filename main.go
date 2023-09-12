package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

//go get gorm.io/gorm@v1.20.11
//go get gorm.io/driver/mysql@v1.0.3

func main() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3308)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//log.Println(db,err)

	newRestaurant := Restaurant{
		Name: "Tani",
		Addr: "9 Pham Van Hai",
	}

	if err := db.Create(&newRestaurant).Error; err != nil {
		log.Println(err)
	}

	log.Println("New ID Created: ", newRestaurant.Id)

	var myRestaurant Restaurant

	if err := db.Where("id = ?", 2).First(&myRestaurant).Error; err != nil {
		log.Println(err)
	}

	log.Println(myRestaurant)

	//myRestaurant.Name = "500lab"

	//if err := db.Where("id = ?", 2).Updates(&myRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(myRestaurant)

	update := "500lab"
	updateData := RestaurantUpdate{Name: &update}
	if err := db.Where("id = ?", 2).Updates(&updateData).Error; err != nil {
		log.Println(err)
	}

	log.Println(*updateData.Name)

	if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 5).Delete(nil).Error; err != nil {
		log.Println(err)
	}

}

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
