package ginrestaurant

import (
	"G05-food-delivery/common"
	"G05-food-delivery/component/appctx"
	restaurantbiz "G05-food-delivery/module/restaurant/biz"
	restaurantstorage "G05-food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strconv"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc  {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		//id, err := strconv.Atoi(c.Param("id"))

		uid, err := common.FromBase58(c.Param("id"))

			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(),int(uid.GetLocalID())) ; err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
