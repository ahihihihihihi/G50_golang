package restaurantstorage

import (
	restaurantmodel "G05-food-delivery/module/restaurant/model"
	"context"
)

func (s *sqlStore) FindDataWithCondition(context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant

	if err := s.db.Where(condition).First(&data).Error; err != nil {

		return nil, err
	}
	return &data, nil
}
