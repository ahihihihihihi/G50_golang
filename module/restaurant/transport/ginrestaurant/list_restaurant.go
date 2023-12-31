package ginrestaurant

import (
	"G05-food-delivery/common"
	"G05-food-delivery/component/appctx"
	restaurantbiz "G05-food-delivery/module/restaurant/biz"
	restaurantmodel "G05-food-delivery/module/restaurant/model"
	restaurantrepo "G05-food-delivery/module/restaurant/repository"
	restaurantstorage "G05-food-delivery/module/restaurant/storage"
	restaurantlikestorage "G05-food-delivery/module/restaurantlike/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.Status = []int{1}

		store := restaurantstorage.NewSQLStore(db)
		likeStore := restaurantlikestorage.NewSQLStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store,likeStore)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		fmt.Println(result)

		for i := range result{
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
