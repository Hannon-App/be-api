package helpers

import (
	"Hannon-app/app/config"
	"bytes"
	"context"
	"io/ioutil"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/admin/search"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

func UploadImage(c echo.Context) (string, error) {
	cfg := config.InitConfig()

	cld, errParam := cloudinary.NewFromParams(cfg.CLOUD_NAME, cfg.KEY_API, cfg.KEY_API_SECRET)
	if errParam != nil {
		return "", FailedRequest(c, "failed config "+errParam.Error(), nil)
	}
	// Mengambil file dari permintaan
	file, err := c.FormFile("image")
	if err != nil {
		// Menangani error ketika gagal mendapatkan file
		return "", FailedNotFound(c, "Failed to get image file from request"+err.Error(), nil)
	}

	// Membuka file
	fileSrc, err := file.Open()
	if err != nil {
		// Menangani error ketika gagal membuka file
		return "", FailedRequest(c, "Failed to open file", nil)
	}
	defer fileSrc.Close()

	// Membaca isi file sebagai byte array
	fileBytes, err := ioutil.ReadAll(fileSrc)
	if err != nil {
		// Menangani error ketika gagal membaca file
		return "", FailedRequest(c, "Failed to read file", nil)
	}

	// Membuat buffer dari byte array
	buffer := bytes.NewBuffer(fileBytes)

	// Melakukan upload file ke Cloudinary
	ctx := context.Background()
	uploadResult, err := cld.Upload.Upload(
		ctx,
		buffer,
		uploader.UploadParams{PublicID: file.Filename},
	)
	if err != nil {
		// Menangani error ketika gagal melakukan upload file
		return "", FailedRequest(c, "Failed to upload file", nil)
	}
	log.Println(uploadResult.SecureURL)

	asset, errAsset := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: file.Filename})
	if errAsset != nil {
		return "", FailedNotFound(c, "failed asset "+errAsset.Error(), nil)
	}

	// log.Printf("Public ID: %v, URL: %v\n", asset.PublicID, asset.SecureURL)
	searchQuery := search.Query{
		Expression: "resource_type:image AND uploaded_at>1d AND bytes<1m",
		SortBy:     []search.SortByField{{"created_at": search.Descending}},
		MaxResults: 30,
	}

	searchResult, errSearch := cld.Admin.Search(ctx, searchQuery)

	if errSearch != nil {
		return "", InternalError(c, "failed data result "+errSearch.Error(), nil)
	}

	log.Printf("Assets found: %v\n", searchResult.TotalCount)

	for _, asset := range searchResult.Assets {
		log.Printf("Public ID: %v, URL: %v\n", asset.PublicID, asset.SecureURL)
	}

	return asset.SecureURL, nil

}
