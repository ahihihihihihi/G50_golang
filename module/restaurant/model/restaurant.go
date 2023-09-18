package restaurantmodel

import (
	"G05-food-delivery/common"
	"errors"
	"strings"
)

type RestaurantType string

const (
	TypeNormal  RestaurantType = "normal"
	TypePremium RestaurantType = "premium"

	EntityName = "Restaurant"
)

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image      `json:"logo" gorm:"logo;"`
	Cover           *common.Images     `json:"cover" gorm:"cover;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false"`
	UserId          int                `json:"-" gorm:"column:user_id;"`
	LikedCount      int                `json:"liked_count" gorm:"-"`
	//Type            RestaurantType `json:"type" gorm:"type"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int            `json:"-" gorm:"column:user_id;"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"logo;"`
	Cover           *common.Images `json:"cover" gorm:"cover;"`
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"addr" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"logo;"`
	Cover *common.Images `json:"cover" gorm:"cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

var (
	ErrNameIsEmpty = errors.New("Name cannot be empty")
)
