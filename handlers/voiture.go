package handlers

import (
	"net/http"
	"strconv"

	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/Yahya-Elamri/signeitfaster/views/components"
	"github.com/Yahya-Elamri/signeitfaster/views/voiture"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func VoitureLocationView(c echo.Context) error {
	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	var cars []module.CarListings
	var totalCount int64

	str, _ := utils.Authorize(c)
	module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)

	baseQuery := module.Db.Model(&module.CarListings{}).Joins("INNER JOIN agencies_data ON agencies_data.id = car_listings.agency_id")

	if address := c.QueryParam("address"); address != "" {
		baseQuery = baseQuery.Where("agencies_data.address LIKE ?", "%"+address+"%")
	}
	if name := c.QueryParam("name"); name != "" {
		baseQuery = baseQuery.Where("agencies_data.name LIKE ?", "%"+name+"%")
	}
	if make := c.QueryParam("make"); make != "" {
		baseQuery = baseQuery.Where("car_listings.make LIKE ?", "%"+make+"%")
	}
	if minPrice := c.QueryParam("min_price"); minPrice != "" {
		if min, err := strconv.ParseFloat(minPrice, 64); err == nil {
			baseQuery = baseQuery.Where("car_listings.price >= ?", min)
		}
	}
	if maxPrice := c.QueryParam("max_price"); maxPrice != "" {
		if max, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			baseQuery = baseQuery.Where("car_listings.price <= ?", max)
		}
	}
	if transmissions := c.QueryParams()["transmission[]"]; len(transmissions) > 0 {
		baseQuery = baseQuery.Where("car_listings.transmission IN ?", transmissions)
	}
	if seatParams := c.QueryParams()["seats[]"]; len(seatParams) > 0 {
		var seats []int
		var sevenIncluded bool

		for _, seatParam := range seatParams {
			if seatCount, err := strconv.Atoi(seatParam); err == nil {
				if seatCount == 7 {
					sevenIncluded = true
				} else {
					seats = append(seats, seatCount)
				}
			}
		}

		if sevenIncluded {
			baseQuery = baseQuery.Where("car_listings.seats > ?", 7)
		}

		if len(seats) > 0 {
			baseQuery = baseQuery.Where("car_listings.seats IN ?", seats)
		}
	}
	if fuelTypes := c.QueryParams()["fuel_type[]"]; len(fuelTypes) > 0 {
		baseQuery = baseQuery.Where("car_listings.fuel_type IN ?", fuelTypes)
	}

	if err := baseQuery.Count(&totalCount).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to count car listings"})
	}

	page := 1
	limit := 3
	if p := c.QueryParam("page"); p != "" {
		if pNum, err := strconv.Atoi(p); err == nil {
			page = pNum
		}
	}
	offset := (page - 1) * limit
	paginatedQuery := baseQuery.Limit(limit).Offset(offset)

	if err := paginatedQuery.Find(&cars).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch car listings"})
	}

	return components.Page(voiture.VoitureLocation(User, cars, (totalCount+int64(limit)-1)/int64(limit))).Render(c.Request().Context(), c.Response())
}

func VoitureLocationOneView(c echo.Context) error {
	id := c.Param("id")
	str, _ := utils.Authorize(c)
	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
	var carListings module.MixedListing
	buildBaseQuery := func(table, joins string) *gorm.DB {
		return module.Db.Table(table).
			Joins(joins)
	}
	carListingsDataQuery := buildBaseQuery("car_listings",
		"left join agencies_data on car_listings.agency_id = agencies_data.id left join user_data on agencies_data.id = user_data.user_id").Where("car_listings.id=?", id)
	carListingsDataQuery.Select("car_listings.*, 'rent' as type, user_data.user_id, agencies_data.email, agencies_data.address, agencies_data.phone, user_data.first_name, user_data.last_name, agencies_data.username, user_data.profile_url").
		First(&carListings)

	return components.Page(voiture.VoitureOnly(User, carListings)).Render(c.Request().Context(), c.Response())
}

func VoitureSalesView(c echo.Context) error {
	var cars []module.CarSalesListings
	var totalCount int64
	str, _ := utils.Authorize(c)
	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
	baseQuery := module.Db.Model(&module.CarSalesListings{}).Joins("INNER JOIN user_data ON user_data.username = car_sales_listings.username")

	if address := c.QueryParam("address"); address != "" {
		baseQuery = baseQuery.Where("user_data.address LIKE ?", "%"+address+"%")
	}
	if make := c.QueryParam("make"); make != "" {
		baseQuery = baseQuery.Where("car_sales_listings.model LIKE ?", "%"+make+"%")
	}
	if minPrice := c.QueryParam("min_price"); minPrice != "" {
		if min, err := strconv.ParseFloat(minPrice, 64); err == nil {
			baseQuery = baseQuery.Where("car_sales_listings.price >= ?", min)
		}
	}
	if maxPrice := c.QueryParam("max_price"); maxPrice != "" {
		if max, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			baseQuery = baseQuery.Where("car_sales_listings.price <= ?", max)
		}
	}
	if transmissions := c.QueryParams()["transmission[]"]; len(transmissions) > 0 {
		baseQuery = baseQuery.Where("car_sales_listings.transmission IN ?", transmissions)
	}
	if seatParams := c.QueryParams()["seats[]"]; len(seatParams) > 0 {
		var seats []int
		var sevenIncluded bool

		for _, seatParam := range seatParams {
			if seatCount, err := strconv.Atoi(seatParam); err == nil {
				if seatCount == 7 {
					sevenIncluded = true
				} else {
					seats = append(seats, seatCount)
				}
			}
		}

		if sevenIncluded {
			baseQuery = baseQuery.Where("car_sales_listings.seats > ?", 7)
		}

		if len(seats) > 0 {
			baseQuery = baseQuery.Where("car_sales_listings.seats IN ?", seats)
		}
	}
	if fuelTypes := c.QueryParams()["fuel_type[]"]; len(fuelTypes) > 0 {
		baseQuery = baseQuery.Where("car_sales_listings.fuel_type IN ?", fuelTypes)
	}

	if err := baseQuery.Count(&totalCount).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to count car listings"})
	}

	page := 1
	limit := 3
	if p := c.QueryParam("page"); p != "" {
		if pNum, err := strconv.Atoi(p); err == nil {
			page = pNum
		}
	}
	offset := (page - 1) * limit
	paginatedQuery := baseQuery.Limit(limit).Offset(offset)

	if err := paginatedQuery.Find(&cars).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch car listings"})
	}

	return components.Page(voiture.VoitureSales(User, cars, (totalCount+int64(limit)-1)/int64(limit))).Render(c.Request().Context(), c.Response())
}

func VoitureSalesOneView(c echo.Context) error {
	id := c.Param("id")
	str, _ := utils.Authorize(c)
	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
	var carListings module.MixedListing
	buildBaseQuery := func(table, joins string) *gorm.DB {
		return module.Db.Table(table).
			Joins(joins)
	}
	carSalesDataQuery := buildBaseQuery("car_sales_listings", "left join user_data on car_sales_listings.username = user_data.username")
	carSalesDataQuery.Select("car_sales_listings.*, 'sale' as type, user_data.user_id, user_data.email, user_data.address, user_data.phone, user_data.first_name, user_data.last_name, user_data.profile_url").Where("car_sales_listings.id=?", id).
		First(&carListings)

	return components.Page(voiture.VoitureOnly(User, carListings)).Render(c.Request().Context(), c.Response())
}

func VoiturePartsView(c echo.Context) error {
	var cars []module.CarPartsListings
	var totalCount int64
	str, _ := utils.Authorize(c)
	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
	baseQuery := module.Db.Model(&module.CarPartsListings{}).Joins("INNER JOIN user_data ON user_data.username = car_parts_listings.username")

	if address := c.QueryParam("address"); address != "" {
		baseQuery = baseQuery.Where("user_data.address LIKE ?", "%"+address+"%")
	}
	if PartName := c.QueryParam("part_name"); PartName != "" {
		baseQuery = baseQuery.Where("car_parts_listings.part_name LIKE ?", "%"+PartName+"%")
	}
	if make := c.QueryParam("make"); make != "" {
		baseQuery = baseQuery.Where("car_parts_listings.make LIKE ?", "%"+make+"%")
	}
	if etat := c.QueryParam("etat"); etat != "" {
		baseQuery = baseQuery.Where("car_parts_listings.etat LIKE ?", "%"+etat+"%")
	}
	if minPrice := c.QueryParam("min_price"); minPrice != "" {
		if min, err := strconv.ParseFloat(minPrice, 64); err == nil {
			baseQuery = baseQuery.Where("car_parts_listings.price >= ?", min)
		}
	}
	if maxPrice := c.QueryParam("max_price"); maxPrice != "" {
		if max, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			baseQuery = baseQuery.Where("car_parts_listings.price <= ?", max)
		}
	}

	if err := baseQuery.Count(&totalCount).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to count car listings"})
	}

	page := 1
	limit := 3
	if p := c.QueryParam("page"); p != "" {
		if pNum, err := strconv.Atoi(p); err == nil {
			page = pNum
		}
	}
	offset := (page - 1) * limit
	paginatedQuery := baseQuery.Limit(limit).Offset(offset)

	if err := paginatedQuery.Find(&cars).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch car listings"})
	}

	return components.Page(voiture.VoitureParts(User, cars, (totalCount+int64(limit)-1)/int64(limit))).Render(c.Request().Context(), c.Response())
}

func VoiturePartsOneView(c echo.Context) error {
	id := c.Param("id")
	str, _ := utils.Authorize(c)
	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
	var carParts module.MixedListing
	buildBaseQuery := func(table, joins string) *gorm.DB {
		return module.Db.Table(table).
			Joins(joins)
	}
	carPartsDataQuery := buildBaseQuery("car_parts_listings", "left join user_data on car_parts_listings.username = user_data.username")
	carPartsDataQuery.Select("car_parts_listings.*, 'part' as type, user_data.user_id, user_data.email, user_data.address, user_data.phone, user_data.first_name, user_data.last_name,user_data.profile_url").Where("car_parts_listings.id=?", id).
		First(&carParts)
	return components.Page(voiture.VoitureOnly(User, carParts)).Render(c.Request().Context(), c.Response())
}
