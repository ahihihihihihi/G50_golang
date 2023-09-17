package ginuser

import (
	"G05-food-delivery/common"
	"G05-food-delivery/component/appctx"
	"G05-food-delivery/component/hasher"
	"G05-food-delivery/component/tokenprovider/jwt"
	userbiz "G05-food-delivery/module/user/biz"
	usermodel "G05-food-delivery/module/user/model"
	userstore "G05-food-delivery/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc  {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBindJSON(&loginUserData) ; err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store,tokenProvider,md5,60*60*24*30)
		account, err := business.Login(c.Request.Context(),&loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
