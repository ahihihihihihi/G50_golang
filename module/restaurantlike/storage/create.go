package restaurantlikestorage

import (
	"G05-food-delivery/common"
	restaurantlikemodel "G05-food-delivery/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error ; err!= nil {
		return common.ErrDB(err)
	}

	return nil
}
