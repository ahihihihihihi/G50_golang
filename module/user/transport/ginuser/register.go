package ginuser

import (
	"G05-food-delivery/common"
	"G05-food-delivery/component/appctx"
	"G05-food-delivery/component/hasher"
	userbiz "G05-food-delivery/module/user/biz"
	usermodel "G05-food-delivery/module/user/model"
	userstore "G05-food-delivery/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}
		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
