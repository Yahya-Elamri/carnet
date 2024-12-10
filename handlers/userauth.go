package handlers

import (
	"net/http"
	"time"

	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/labstack/echo/v4"
)

func UserAuth(c echo.Context) error {
	u := new(module.UserData)
	if err := c.Bind(u); err != nil {
		c.Response().Header().Set("HX-Redirect", "/404")
		return c.NoContent(404)
	}

	var user module.UserData
	if err := module.Db.Where("username = ? OR email = ?", u.Username, u.Username).First(&user).Error; err != nil {
		return c.HTML(200, `<form class="w-full flex flex-col gap-5 justify-between items-start fade-in" hx-post="/auth-user" hx-target="this" hx-swap="outerHTML" method="POST">
                <div class="w-full flex flex-col gap-2">
                    <label class="font-semibold	 font-sans" for="username">Username or Email</label>
                    <input class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Username" id="username" required/>
                </div>
                <div class="w-full flex flex-col gap-2">
                    <label class="font-semibold	 font-sans" for="password">Password</label>
                    <input class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="password" name="PasswordHash" id="password" required/>
                </div>
				<p class="text-red-600 capitalize">Invalid username/email or password</p>
                <div class="w-full flex flex-col items-center gap-2">  
                    <button class="poppins-regular bg-green-500 px-4 py-[12px] w-full rounded-full text-white hover:bg-green-700" type="submit">Sign in</button>
                    <h1 class="text-[#B4A7AF]">Don't have an account? <a href="/signup" class="underline text-black"> Sign up</a></h1>
                </div>
            </form>`)
	}

	if err := utils.CompairHash(u.PasswordHash, user.PasswordHash); err != nil {
		return c.HTML(200, `<form class="w-full flex flex-col gap-5 justify-between items-start fade-in" hx-post="/auth-user" hx-target="this" hx-swap="outerHTML" method="POST">
                <div class="w-full flex flex-col gap-2">
                    <label class="font-semibold	 font-sans" for="username">Username or Email</label>
                    <input class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="text" name="Username" id="username" required/>
                </div>
                <div class="w-full flex flex-col gap-2">
                    <label class="font-semibold	 font-sans" for="password">Password</label>
                    <input class="border border-[#d5e0d5] focus:outline-none hover:shadow-pink-200 hover:shadow-md focus:shadow-pink-200 focus:shadow-md px-[20px] py-[18px] rounded-xl" type="password" name="PasswordHash" id="password" required/>
                </div>
				<p class="text-red-600 capitalize">Invalid username/email or password</p>
                <div class="w-full flex flex-col items-center gap-2">  
                    <button class="poppins-regular bg-green-500 px-4 py-[12px] w-full rounded-full text-white hover:bg-green-700" type="submit">Sign in</button>
                    <h1 class="text-[#B4A7AF]">Don't have an account? <a href="/signup" class="underline text-black"> Sign up</a></h1>
                </div>
            </form>`)
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
