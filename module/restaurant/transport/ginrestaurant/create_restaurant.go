package ginrestaurant

import (
	"G05-food-delivery/common"
	"G05-food-delivery/component/appctx"
	restaurantbiz "G05-food-delivery/module/restaurant/biz"
	restaurantmodel "G05-food-delivery/module/restaurant/model"
	restaurantstorage "G05-food-delivery/module/restaurant/storage"
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc  {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data restaurantmodel.RestaurantCreate

		// test middleware
		//a := []int{}
		//fmt.Println(a[0])

		// test AppRecover
		//go func() {
		//	defer common.AppRecover()
		//	a := []int{}
		//	fmt.Println(a[0])
		//}()


		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(),&data) ; err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
