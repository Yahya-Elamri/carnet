package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/labstack/echo/v4"
)

func VoitureLocation(c echo.Context) error {
	carListing := new(module.CarListings)
	if err := c.Bind(carListing); err != nil {
		return err
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["images"]

	imageUrls, err := utils.SaveUploadedFiles(files)
	if err != nil {
		return err
	}

	imageUrlsJSON, err := json.Marshal(imageUrls)
	if err != nil {
		return err
	}
	carListing.ImageURL = string(imageUrlsJSON)
	carListing.CreatedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))
	carListing.Category = "voiturelocation"
	module.Db.Create(carListing)
	return c.JSON(http.StatusOK, carListing)
}

func VoitureParts(c echo.Context) error {
	CarPartsListing := new(module.CarPartsListings)
	if err := c.Bind(CarPartsListing); err != nil {
		return err
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["images"]

	imageUrls, err := utils.SaveUploadedFiles(files)
	if err != nil {
		return err
	}

	imageUrlsJSON, err := json.Marshal(imageUrls)
	if err != nil {
		return err
	}
	str, _ := utils.Authorize(c)
	var User struct {
		UserId   uint
		Username string
	}
	module.Db.Model(&module.UserData{}).Select("user_id,username").Where("user_id = ?", str).Find(&User)
	CarPartsListing.Username = User.Username
	CarPartsListing.ImageURL = string(imageUrlsJSON)
	CarPartsListing.CreatedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))
	CarPartsListing.UpdatedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))
	CarPartsListing.Category = "voitureparts"
	module.Db.Create(CarPartsListing)
	return c.JSON(http.StatusOK, CarPartsListing)
}

func VoitureSales(c echo.Context) error {
	CarSalesListing := new(module.CarSalesListings)
	if err := c.Bind(CarSalesListing); err != nil {
		return err
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["images"]

	imageUrls, err := utils.SaveUploadedFiles(files)
	if err != nil {
		return err
	}

	imageUrlsJSON, err := json.Marshal(imageUrls)
	if err != nil {
		return err
	}
	str, _ := utils.Authorize(c)
	var User struct {
		UserId   uint
		Username string
	}
	module.Db.Model(&module.UserData{}).Select("user_id,username").Where("user_id = ?", str).Find(&User)
	CarSalesListing.Username = User.Username
	CarSalesListing.ImageURL = string(imageUrlsJSON)
	CarSalesListing.CreatedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))
	CarSalesListing.Category = "voituresales"
	module.Db.Create(CarSalesListing)
	return c.JSON(http.StatusOK, CarSalesListing)
}
