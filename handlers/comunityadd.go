package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/labstack/echo/v4"
)

func ComunityAdd(c echo.Context) error {
	community := new(module.Communities)
	if err := c.Bind(community); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to bind community data",
		})
	}

	NameValidation := utils.IsFieldUnique("name", community.Name, module.Db, "communities")

	if !NameValidation.Valid {
		return c.HTML(200, fmt.Sprintf(`<form hx-post="/add-comunity" hx-target="this" hx-swap="outerHTML" class="w-full flex flex-col fade-in gap-5 justify-between items-start" method="POST">
			<div class="w-full flex flex-col gap-2">
				<label class="poppins-regular" for="name">Nom du communautés</label>
				<input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="name" id="name" required/>
				<p class="%s">%s</p>
			</div>
			<div class="w-full flex flex-col gap-2">
				<label class="poppins-regular" for="description">Description</label>
				<input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="description" id="description" required/>
			</div>
			<div class="w-full flex flex-col gap-2">
				<label for="file-upload1" class="poppins-regular">Photo du communautes</label>
				<div id="dropzone1" class="dropzone mt-2 p-6 bg-white flex items-center justify-center">
					<p class="poppins-regular text-gray-500">Drag and drop images here, or click to select files</p>
					<input value="%s" id="file-upload1" name="communities_media" type="file" class="hidden" accept="image/*" multiple>
				</div>
			</div>
			<div class="w-full flex flex-col gap-2">
				<label for="file-upload2" class="poppins-regular">Banner du communautes</label>
				<div id="dropzone2" class="dropzone mt-2 p-6 bg-white flex items-center justify-center">
					<p class="poppins-regular text-gray-500">Drag and drop images here, or click to select files</p>
					<input value="%s" id="file-upload2" name="communities_banner" type="file" class="hidden" accept="image/*" multiple>
				</div>
			</div>
			<div class="w-full flex flex-col items-start gap-2">  
				<button class="poppins-regular bg-green-500 px-4 py-[12px] w-1/2 rounded-xl text-white hover:bg-green-700" type="submit">Soumettre</button>
			</div>
		</form>`, community.Name, NameValidation.Class, NameValidation.Message, community.Description, community.CommunitiesMedia, community.CommunitiesBanner))
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	//handling profile imgs
	files := form.File["communities_media"]
	imageUrls, err := utils.SaveUploadedFiles(files)
	if err != nil {
		return err
	}
	imageUrlsJSON, err := json.Marshal(imageUrls)
	if err != nil {
		return err
	}
	community.CommunitiesMedia = string(imageUrlsJSON)

	//handling banner imgs
	files = form.File["communities_banner"]
	imageUrls, err = utils.SaveUploadedFiles(files)
	if err != nil {
		return err
	}
	imageUrlsJSON, err = json.Marshal(imageUrls)
	if err != nil {
		return err
	}
	community.CommunitiesBanner = string(imageUrlsJSON)

	UserId, _ := utils.Authorize(c)

	community.UserId = UserId
	community.CreatedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))
	// Insert the community into the database
	if err := module.Db.Create(&community).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create community in the database",
		})
	}

	userCommunity := module.UserCommunities{
		UserId:      UserId,
		CommunityID: community.CommunityID,
		Role:        "owner", // or another role as needed
		JoinedAt:    []uint8(time.Now().Format("2006-01-02 15:04:05")),
	}

	if err := module.Db.Create(&userCommunity).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add user to community"})
	}
	// Return a success message or redirect
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Community successfully created",
	})
}
