package ginrestaurant

import (
	"G05-food-delivery/common"
	"G05-food-delivery/component/appctx"
	restaurantbiz "G05-food-delivery/module/restaurant/biz"
	restaurantstorage "G05-food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appctx.Appcontext) gin.HandlerFunc  {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})

				return
			}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(),id) ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
