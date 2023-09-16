package ginupload

import (
	"G05-food-delivery/common"
	"G05-food-delivery/component/appctx"
	"G05-food-delivery/module/upload/uploadbusiness"
	"github.com/gin-gonic/gin"
)

func Upload(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder","img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // we can close here

		dataBytes := make([]byte,fileHeader.Size)

		if _, err := file.Read(dataBytes) ; err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := uploadbusiness.NewUploadBiz(appCtx.UploadProvider())
		img, err := biz.Upload(c.Request.Context(),dataBytes,folder,fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		c.JSON(200,common.SimpleSuccessResponse(img))

	}
}
