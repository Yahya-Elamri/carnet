package handlers

import (
	"fmt"
	"time"

	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/labstack/echo/v4"
)

func AgencyAdd(c echo.Context) error {
	var agency module.AgenciesData
	if err := c.Bind(&agency); err != nil {
		return err
	}

	EmailValidation := utils.IsFieldUnique("email", agency.Email, module.Db, "agencies_data")
	NameValidation := utils.IsFieldUnique("username", agency.Username, module.Db, "agencies_data")

	if !NameValidation.Valid {
		return c.HTML(200, fmt.Sprintf(`
			<form hx-post="/add-agency" hx-target="this" hx-swap="outerHTML" class="w-full flex flex-col fade-in gap-5 justify-between items-start" method="POST">
				<div class="w-full flex flex-col gap-2">
					<label class="font-semibold	 font-sans" for="Usename">Nom D'agence</label>
					<input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Usename" id="Usename" required/>
					<p class="%s">%s</p>
				</div>
				<div class="w-full flex flex-col gap-2">
					<label class="font-semibold	 font-sans" for="email">Email</label>
					<input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Email" id="email" required/>
					<p class="%s">%s</p>
				</div>
				<div class="w-full flex flex-col gap-2">
					<label class="font-semibold	 font-sans" for="Phone">Numero de Telephone</label>
					<input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Phone" id="Phone" required/>
				</div>
				<div class="w-full flex flex-col gap-2">
					<label class="font-semibold	 font-sans" for="Address">Address</label>
					<input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Address" id="Address" required/>
				</div>
				<div class="w-full flex flex-col items-center gap-2">  
					<button class="poppins-regular bg-green-500 px-4 py-[12px] w-full rounded-full text-white hover:bg-green-700" type="submit">ajouter agence</button>
				</div>
			</form>
		`, agency.Username, NameValidation.Class, NameValidation.Message, agency.Email, EmailValidation.Class, EmailValidation.Message, agency.Phone, agency.Address))
	}

	if !EmailValidation.Valid {
		return c.HTML(200, fmt.Sprintf(`
			<form hx-post="/add-agency" hx-target="this" hx-swap="outerHTML" class="w-full flex flex-col fade-in gap-5 justify-between items-start" method="POST">
                <div class="w-full flex flex-col gap-2">
                    <label class="font-semibold	 font-sans" for="Name">Nom D'agence</label>
                    <input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Name" id="Name" required/>
					<p class="%s">%s</p>
                </div>
                <div class="w-full flex flex-col gap-2">
                    <label class="font-semibold	 font-sans" for="email">Email</label>
                    <input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Email" id="email" required/>
					<p class="%s">%s</p>
                </div>
                <div class="w-full flex flex-col gap-2">
                    <label class="font-semibold	 font-sans" for="Phone">Numero de Telephone</label>
                    <input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Phone" id="Phone" required/>
                </div>
                <div class="w-full flex flex-col gap-2">
                    <label class="font-semibold	 font-sans" for="Address">Address</label>
                    <input value="%s" class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Address" id="Address" required/>
                </div>
                <div class="w-full flex flex-col items-center gap-2">  
                    <button class="poppins-regular bg-green-500 px-4 py-[12px] w-full rounded-full text-white hover:bg-green-700" type="submit">Sign up</button>
                </div>
            </form>
		`, agency.Username, NameValidation.Class, NameValidation.Message, agency.Email, EmailValidation.Class, EmailValidation.Message, agency.Phone, agency.Address))
	}

	agency.CreatedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))

	if err := module.Db.Create(&agency).Error; err != nil {
		c.Response().Header().Set("HX-Redirect", "/404")
		return c.NoContent(404)
	}

	id, _ := utils.Authorize(c)

	userAgency := module.UserAgencies{
		UserId:   id,
		AgencyId: agency.Id,
	}

	if err := module.Db.Create(&userAgency).Error; err != nil {
		c.Response().Header().Set("HX-Redirect", "/404")
		return c.NoContent(404)
	}

	c.Response().Header().Set("HX-Redirect", "/")
	return c.NoContent(200)
}
