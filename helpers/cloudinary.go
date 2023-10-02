package helpers

import (
	"Hannon-app/app/config"
	"bytes"
	"context"
	"io/ioutil"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin/search"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

func UploadImage(c echo.Context) (string, string, error) {
	cfg := config.InitConfig()

	cld, errParam := cloudinary.NewFromParams(cfg.CLOUD_NAME, cfg.KEY_API, cfg.KEY_API_SECRET)
	if errParam != nil {
		return "", "", FailedRequest(c, "failed config "+errParam.Error(), nil)
	}
	// Mengambil file dari permintaan
	file, err := c.FormFile("profile_photo")
	if err != nil {
		// Menangani error ketika gagal mendapatkan file
		return "", "", FailedNotFound(c, "Failed to get profile photo file from request: "+err.Error(), nil)
	}

	// Mengambil file ktp_photo dari permintaan
	image, err := c.FormFile("image")
	if err != nil {
		// Menangani error ketika gagal mendapatkan file
		return "", "", FailedNotFound(c, "Failed to get ktp photo file from request: "+err.Error(), nil)
	}

	// Membuka file profile_photo
	fileSrc, err := file.Open()
	if err != nil {
		// Menangani error ketika gagal membuka file
		return "", "", FailedRequest(c, "Failed to open profile photo file", nil)
	}
	defer fileSrc.Close()

	// Membuka file ktp_photo
	ktpFileSrc, err := image.Open()
	if err != nil {
		// Menangani error ketika gagal membuka file
		return "", "", FailedRequest(c, "Failed to open ktp photo file", nil)
	}
	defer ktpFileSrc.Close()

	// Membaca isi file profile_photo sebagai byte array
	fileBytes, err := ioutil.ReadAll(fileSrc)
	if err != nil {
		// Menangani error ketika gagal membaca file
		return "", "", FailedRequest(c, "Failed to read profile photo file", nil)
	}

	// Membaca isi file ktp_photo sebagai byte array
	ktpFileBytes, err := ioutil.ReadAll(ktpFileSrc)
	if err != nil {
		// Menangani error ketika gagal membaca file
		return "", "", FailedRequest(c, "Failed to read ktp photo file", nil)
	}

	// Membuat buffer dari byte array
	fileBuffer := bytes.NewBuffer(fileBytes)

	// Membuat buffer dari byte array
	ktpFileBuffer := bytes.NewBuffer(ktpFileBytes)

	// Melakukan upload file profile_photo ke Cloudinary
	ctx := context.Background()
	profilePhotoUploadResult, err := cld.Upload.Upload(
		ctx,
		fileBuffer,
		uploader.UploadParams{PublicID: file.Filename},
	)
	if err != nil {
		// Menangani error ketika gagal melakukan upload file
		return "", "", FailedRequest(c, "Failed to upload profile photo file", nil)
	}
	log.Println(profilePhotoUploadResult.SecureURL)

	// Melakukan upload file ktp_photo ke Cloudinary
	ktpPhotoUploadResult, err := cld.Upload.Upload(
		ctx,
		ktpFileBuffer,
		uploader.UploadParams{PublicID: image.Filename},
	)
	if err != nil {
		// Menangani error ketika gagal melakukan upload file
		return "", "", FailedRequest(c, "Failed to upload ktp photo file", nil)
	}
	log.Println(ktpPhotoUploadResult.SecureURL)

	// log.Printf("Public ID: %v, URL: %v\n", asset.PublicID, asset.SecureURL)
	searchQuery := search.Query{
		Expression: "resource_type:image AND uploaded_at>1d AND bytes<1m",
		SortBy:     []search.SortByField{{"created_at": search.Descending}},
		MaxResults: 30,
	}

	_, errSearch := cld.Admin.Search(ctx, searchQuery)

	if errSearch != nil {
		return "", "", InternalError(c, "failed data result "+errSearch.Error(), nil)
	}

	return profilePhotoUploadResult.SecureURL, ktpPhotoUploadResult.SecureURL, nil
}
