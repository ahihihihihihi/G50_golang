package restaurantstorage

import (
	"G05-food-delivery/common"
	restaurantmodel "G05-food-delivery/module/restaurant/model"
	"context"
)

func (s *sqlStore) ListDataWithCondition(context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	//db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1)")

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	if f := filter ; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?",f.OwnerId)
		}

		if len(f.Status) > 0 {
			db = db.Where("status in (?)",f.Status)
		}
	}


	if err := db.Count(&paging.Total).Error ; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1 ) * paging.Limit

	if err := db.Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {

		return nil, common.ErrDB(err)
	}
	return result, nil
}
