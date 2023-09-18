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
		if f.UserId > 0 {
			db = db.Where("user_id = ?",f.UserId)
		}

		if len(f.Status) > 0 {
			db = db.Where("status in (?)",f.Status)
		}
	}

	if err := db.Count(&paging.Total).Error ; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor ; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?",uid.GetLocalID())
	} else {
		offset := (paging.Page - 1 ) * paging.Limit

		db = db.Offset(offset)
	}
	
	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		last := result[len(result) - 1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}
	
	return result, nil
}
