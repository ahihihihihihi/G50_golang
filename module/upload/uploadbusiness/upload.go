package uploadbusiness

import (
	"G05-food-delivery/common"
	"G05-food-delivery/component/uploadprovider"
	"G05-food-delivery/module/upload/uploadmodel"
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type uploadBiz struct {
	provider uploadprovider.UploadProvider
}

func NewUploadBiz(provider uploadprovider.UploadProvider) *uploadBiz {
	return &uploadBiz{provider: provider}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, filename string) (*common.Image, error)  {
	filebytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(filebytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(filename) // img.jpg => ".jpg"
	filename = fmt.Sprintf("%d%s",time.Now().Nanosecond(),fileExt) // 986367348.jpg

	img, err := biz.provider.SaveFileUploaded(ctx,data,fmt.Sprintf("%s/%s",folder,filename))

	if err != nil {
		return nil, uploadmodel.ErrCanNotSaveFile(err)
	}

	img.Width = w
	img.Height = h

	img.Extension = fileExt

	return img, nil

}

func getImageDimension(reader io.Reader) (int,int,error)  {
	img, _ , err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ",err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
