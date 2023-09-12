package ginrestaurant

import (
	"G05-food-delivery/common"
	"G05-food-delivery/component/appctx"
	restaurantbiz "G05-food-delivery/module/restaurant/biz"
	restaurantmodel "G05-food-delivery/module/restaurant/model"
	restaurantstorage "G05-food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.Appcontext) gin.HandlerFunc  {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(),&data) ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
