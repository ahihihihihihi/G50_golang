package restaurantstorage

import (
	"G05-food-delivery/common"
	restaurantmodel "G05-food-delivery/module/restaurant/model"
	"context"
)

func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate)  error {
	if err := s.db.Create(&data).Error; err != nil {

		return common.ErrDB(err)
	}
	return nil
}
