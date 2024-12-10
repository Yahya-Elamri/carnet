package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/labstack/echo/v4"
)

func UserAdd(c echo.Context) error {
	var user module.UserData
	if err := c.Bind(&user); err != nil {
		return err
	}

	EmailValidation := utils.IsFieldUnique("email", user.Email, module.Db, "user_data")
	NameValidation := utils.IsFieldUnique("username", user.Username, module.Db, "user_data")

	if !EmailValidation.Valid || !NameValidation.Valid {
		return c.HTML(200, fmt.Sprintf(`<form hx-post="/add-user" hx-target="this" hx-swap="outerHTML" class="w-full flex flex-col fade-in gap-5 justify-between items-start" method="POST">
            <div class="flex w-full gap-3 items-center flex-col md:flex-row">
                <div class="w-full flex flex-col gap-2">
                    <label class="font-semibold font-sans" for="fname">First Name</label>
                    <input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="FirstName" id="fname" required/>
                </div>
                <div class="w-full flex flex-col gap-2">
                    <label class="font-semibold font-sans" for="lname">Last Name</label>
                    <input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="LastName" id="lname" required/>
                </div>
            </div>
            <div class="w-full flex flex-col gap-2">
                <label class="font-semibold font-sans" for="username">Username</label>
                <input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Username" id="username" required/>
				<p class="%s">%s</p>
            </div>
            <div class="w-full flex flex-col gap-2">
                <label class="font-semibold font-sans" for="email">Email</label>
                <input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Email" id="email" required/>
				<p class="%s">%s</p>
            </div>
            <div class="w-full flex flex-col gap-2">
                <label class="font-semibold font-sans" for="password">Password</label>
                <input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="password" name="PasswordHash" id="password" required/>
            </div>
			<div class="w-full flex flex-col gap-2">
				<label class="poppins-regular" for="Phone">Numero de Telephone</label>
				<input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Phone" id="Phone" required/>
			</div>
			<div class="w-full flex flex-col gap-2">
				<label class="poppins-regular" for="Address">Address</label>
				<input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Address" id="Address" required/>
			</div>
            <div class="w-full flex flex-col items-center gap-2">  
                <button class="poppins-regular bg-green-500 px-4 py-[12px] w-full rounded-full text-white hover:bg-green-700" type="submit">Sign up</button>
                <h1 class="text-[#B4A7AF]">By creating an account you agree with our <a href="/signup" class="underline">Terms of Service</a> , <a href="/signup" class="underline">Privacy Policy</a></h1>
                <h1 class="text-[#B4A7AF]">Already have an account? <a href="/signup" class="underline text-black"> Sign In</a></h1>
            </div>
        </form>
	`, user.FirstName, user.LastName, user.Username, NameValidation.Class, NameValidation.Message, user.Email, EmailValidation.Class, EmailValidation.Message, user.PasswordHash, user.Phone, user.Address))
	}

	user.PasswordHash, _ = utils.Hashing(user.PasswordHash)
	user.CreatedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))
	user.UpdatedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))
	if err := module.Db.Create(&user).Error; err != nil {
		c.Response().Header().Set("HX-Redirect", "/404")
		return c.NoContent(404)
	}

	token, err := utils.GenerateJWT(user.UserId)

	if err != nil {
		c.Response().Header().Set("HX-Redirect", "/404")
		return c.NoContent(200)
	}

	expiration := time.Now().Add(24 * time.Hour)
	cookie := new(http.Cookie)
	cookie.Name = "Token"
	cookie.Value = token
	cookie.Expires = expiration
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	c.Response().Header().Set("HX-Redirect", "/")
	return c.NoContent(200)
}

func UpdateUserPasswd(c echo.Context) error {
	// Get user ID from JWT or session
	userId, _ := utils.Authorize(c) // Implement this function to get user ID
	Exestingpasswd := c.FormValue("user_password")
	// Fetch the existing user profile
	var user module.UserData
	if err := module.Db.First(&user, userId).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	// Bind the input to the struct
	updatedUser := new(module.UserData)
	if err := c.Bind(updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if Exestingpasswd != "" {
		if err := utils.CompairHash(Exestingpasswd, user.PasswordHash); err != nil {
			return c.HTML(200, `<form hx-post="/update-passwd" hx-target="this" hx-swap="outerHTML" enctype="multipart/form-data" class="w-full mt-3 flex flex-col fade-in gap-5 		justify-between items-start" method="POST">
				<div class="w-full flex flex-col items-start gap-2 bg-transparent text-black dark:text-white">
					<label class="poppins-regular" for="password">Ancien mot de passe</label>
					<input class="border w-full border-[#d5e0d5] dark:border-[#3F3F3F] bg-transparent text-black dark:text-white focus:outline-none hover:shadow-pink-300 hover:shadow-sm focus:shadow-pink-300 focus:shadow-sm px-[20px] py-[18px] rounded-xl" type="password" name="user_password" id="password" required/>
				</div>
				<p>Mot de passe incorrecte</p>
				<div class="w-full flex flex-col items-start gap-2">  
					<button class="poppins-regular bg-green-500 px-4 py-[12px] w-1/2 rounded-xl text-white hover:bg-green-700" type="submit">Changer mon mot de passe</button>
				</div>
			</form>`)
		}
	}
	return c.HTML(200, `<form hx-post="/update-profile" hx-target="this" hx-swap="outerHTML" enctype="multipart/form-data" class="w-full mt-3 flex flex-col fade-in gap-5 justify-between items-start" method="POST">
		<div class="w-full flex flex-col items-start gap-2 bg-transparent text-black dark:text-white">
			<label class="poppins-regular" for="password">Nouveau mot de passe</label>
			<input class="border w-full border-[#d5e0d5] dark:border-[#3F3F3F] bg-transparent text-black dark:text-white focus:outline-none hover:shadow-pink-300 hover:shadow-sm focus:shadow-pink-300 focus:shadow-sm px-[20px] py-[18px] rounded-xl" type="password" name="PasswordHash" id="password" required/>
		</div>
		<div class="w-full flex flex-col items-start gap-2 bg-transparent text-black dark:text-white">
			<label class="poppins-regular" for="password">Confirmez le mot de passe</label>
			<input class="border w-full border-[#d5e0d5] dark:border-[#3F3F3F] bg-transparent text-black dark:text-white focus:outline-none hover:shadow-pink-300 hover:shadow-sm focus:shadow-pink-300 focus:shadow-sm px-[20px] py-[18px] rounded-xl" type="password" id="password" required/>
		</div>
		<div class="w-full flex flex-col items-start gap-2">  
			<button class="poppins-regular bg-green-500 px-4 py-[12px] w-1/2 rounded-xl text-white hover:bg-green-700" type="submit">Changer mon mot de passe</button>
		</div>
	</form>
	`)
}

func UpdateUserProfile(c echo.Context) error {
	// Get user ID from JWT or session
	userId, _ := utils.Authorize(c) // Implement this function to get user ID

	// Fetch the existing user profile
	var user module.UserData
	if err := module.Db.First(&user, userId).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	// Bind the input to the struct
	updatedUser := new(module.UserData)
	if err := c.Bind(updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	// Prepare the base query
	query := module.Db.Model(&user)
	// Apply updates if fields are provided
	if updatedUser.Username != "" {
		query = query.Update("username", updatedUser.Username)
	}
	if updatedUser.Email != "" {
		EmailValidation := utils.IsFieldUnique("email", updatedUser.Email, module.Db, "user_data")
		if !EmailValidation.Valid {
			return c.HTML(200, fmt.Sprintf(`<form hx-post="/update-profile" hx-target="this" hx-swap="outerHTML" enctype="multipart/form-data" class="w-full mt-3 flex flex-col fade-in gap-5 justify-between items-start" method="POST">
				<div class="w-full flex flex-col items-start gap-2 bg-transparent text-black dark:text-white">
					<label class="poppins-regular" for="Email">Nouvel email</label>
					<input class="border w-full border-[#d5e0d5] dark:border-[#3F3F3F] bg-transparent text-black dark:text-white focus:outline-none hover:shadow-pink-300 hover:shadow-sm focus:shadow-pink-300 focus:shadow-sm px-[20px] py-[18px] rounded-xl" type="text" name="Email" id="Email" required/>
				</div>
				<p class="%s">%s</p>
				<div class="w-full flex flex-col items-start gap-2">  
					<button class="poppins-regular bg-green-500 px-4 py-[12px] w-1/2 rounded-xl text-white hover:bg-green-700" type="submit">mettre à jour</button>
				</div>
			</form>`, EmailValidation.Class, EmailValidation.Message))
		}
		query = query.Update("email", updatedUser.Email)
	}
	if updatedUser.Phone != "" {
		query = query.Update("phone", updatedUser.Phone)
	}
	if updatedUser.Address != "" {
		query = query.Update("address", updatedUser.Address)
	}
	if updatedUser.FirstName != "" {
		query = query.Update("first_name", updatedUser.FirstName)
	}
	if updatedUser.LastName != "" {
		query = query.Update("last_name", updatedUser.LastName)
	}
	form, err := c.MultipartForm()
	if err == nil {
		files := form.File["ProfileUrl"]
		if len(files) != 0 {
			imageUrls, err := utils.SaveOneUploadedFile(files[0])
			if imageUrls != "" && err == nil {
				updatedUser.ProfileUrl = imageUrls
				query = query.Update("profile_url", updatedUser.ProfileUrl)
			}
		}
	}
	if updatedUser.Description != "" {
		query = query.Update("description", updatedUser.Description)
	}
	if updatedUser.SocialLinks != "" {
		formValues, err := c.FormParams()
		if err != nil {
			return c.String(http.StatusBadRequest, "Failed to parse form data")
		}
		socialLinks := formValues["social_links[]"]
		for i := range socialLinks {
			socialLinks[i] = fmt.Sprintf("\"%s\"", socialLinks[i])
		}
		socialLinksString := fmt.Sprintf("[%s]", strings.Join(socialLinks, ","))
		query = query.Update("social_links", socialLinksString)
	}
	if updatedUser.Status != "" {
		query = query.Update("status", updatedUser.Status)
	}
	if updatedUser.PasswordHash != "" {
		hashed, err := utils.Hashing(updatedUser.PasswordHash)
		if err != nil {
			return err
		}
		query = query.Update("password_hash", hashed)
	}
	// Execute the query
	if err := query.Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update profile"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Profile updated successfully"})
}
