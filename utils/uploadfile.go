package utils

import (
	"context"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/masnann/plant_care/config"
	"time"
)

func ImageUploadHelper(input interface{}) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cloudName := config.InitConfig().CloudinaryCloudName
	apiKey := config.InitConfig().CloudinaryAPIKey
	apiSecret := config.InitConfig().CloudinaryAPISecret
	uploadFolder := config.InitConfig().Folder

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return "", err
	}

	uploadParams := uploader.UploadParams{
		Folder: uploadFolder,
	}

	uploadResult, err := cld.Upload.Upload(ctx, input, uploadParams)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
