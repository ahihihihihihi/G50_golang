package restaurantbiz

import (
	"G05-food-delivery/common"
	restaurantmodel "G05-food-delivery/module/restaurant/model"
	"context"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}

func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	) ([]restaurantmodel.Restaurant, error) {

	result,err := biz.repo.ListRestaurant(context,filter,paging)

	if err != nil {
		return nil,common.ErrCannotListEntity(restaurantmodel.EntityName,err)
	}

	// list restaurants only have liked_count > 10

	return result,nil
}