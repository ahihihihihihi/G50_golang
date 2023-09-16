package uploadprovider

import (
	"G05-food-delivery/common"
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
