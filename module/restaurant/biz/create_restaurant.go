package restaurantbiz

import (
	"G05-food-delivery/common"
	restaurantmodel "G05-food-delivery/module/restaurant/model"
	"context"
)

type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantmodel.RestaurantCreate)  error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz  {
	return &createRestaurantBiz{store: store}
}

func (biz createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error  {
	//if data.Name == "" {
	//	return errors.New("Name cannot be empty")
	//}

	if err := data.Validate() ; err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := biz.store.Create(context,data) ; err != nil {
		return common.ErrCanNotCreateEntity(restaurantmodel.EntityName,err)
	}

	return nil
}