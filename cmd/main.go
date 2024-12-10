package main

import (
	"log"
	"net/http"
	"os"
	"time"

	middleware "github.com/Yahya-Elamri/signeitfaster/Middleware"
	"github.com/Yahya-Elamri/signeitfaster/handlers"
	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/Yahya-Elamri/signeitfaster/views/auth"
	"github.com/Yahya-Elamri/signeitfaster/views/components"
	"github.com/Yahya-Elamri/signeitfaster/views/profile"
	"github.com/Yahya-Elamri/signeitfaster/views/register"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	module.JwtSecret = []byte(os.Getenv("JWT_TOCKEN_KEY"))
	module.Db = module.ConnectDb(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	e.GET("/", handlers.Home)

	r1 := e.Group("")
	r1.Use(middleware.NotAuthMiddleware)
	r1.GET("/login", func(c echo.Context) error {
		return components.Page(auth.Login()).Render(c.Request().Context(), c.Response())
	})
	r1.GET("/signup", func(c echo.Context) error {
		return components.Page(auth.Signup()).Render(c.Request().Context(), c.Response())
	})
	r1.POST("/add-user", handlers.UserAdd)
	r1.POST("/auth-user", handlers.UserAuth)

	r2 := e.Group("")
	r2.Use(middleware.AuthMiddleware)
	r2.GET("/marketplace", handlers.MarketplaceHome)
	r2.GET("/marketplace/voiturelocation", handlers.VoitureLocationView)
	r2.GET("/marketplace/voiturelocation/:id", handlers.VoitureLocationOneView)
	r2.GET("/marketplace/voituresales", handlers.VoitureSalesView)
	r2.GET("/marketplace/voituresales/:id", handlers.VoitureSalesOneView)
	r2.GET("/marketplace/voitureparts", handlers.VoiturePartsView)
	r2.GET("/marketplace/voitureparts/:id", handlers.VoiturePartsOneView)
	r2.GET("/communautes", handlers.CommunautesView)
	r2.POST("/join-community", handlers.CommunautesJoin)
	r2.GET("/poster/post", func(c echo.Context) error {
		str, _ := utils.Authorize(c)
		var joinedCommunities []module.Communities
		var User struct {
			UserId     uint
			Username   string
			ProfileUrl string
		}

		module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
		err := module.Db.Joins("JOIN user_communities ON user_communities.community_id = communities.community_id").
			Where("user_communities.user_id = ?", str).
			Find(&joinedCommunities).Error
		if err != nil {
			log.Println("Error retrieving communities:", err)
		}
		return components.Page(register.Post(User, joinedCommunities)).Render(c.Request().Context(), c.Response())
	})
	r2.GET("/poster/thread", func(c echo.Context) error {
		str, _ := utils.Authorize(c)
		var joinedCommunities []module.Communities
		var User struct {
			UserId     uint
			Username   string
			ProfileUrl string
		}
		err := module.Db.Joins("JOIN user_communities ON user_communities.community_id = communities.community_id").
			Where("user_communities.user_id = ?", str).
			Find(&joinedCommunities).Error
		if err != nil {
			log.Println("Error retrieving communities:", err)
		}
		module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
		return components.Page(register.Thread(User, joinedCommunities)).Render(c.Request().Context(), c.Response())
	})
	r2.POST("/add-post", handlers.PostsAdd)
	r2.POST("/add-thread", handlers.ThreadAdd)
	r2.POST("/post/:post_id/vote", handlers.VoteOnPostHandler)
	r2.POST("/thread/:thread_id/vote", handlers.VoteOnThreadHandler)
	r2.POST("/comment/:comment_id/vote", handlers.VoteOnCommentHandler)
	r2.GET("/joined-communautes", handlers.JoinedCommunautes)
	r2.GET("/communautes/:name", handlers.CommunautesOneView)
	r2.GET("/communautes/creer", func(c echo.Context) error {
		str, _ := utils.Authorize(c)
		var User struct {
			UserId     uint
			Username   string
			ProfileUrl string
		}
		module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
		return components.Page(register.AddComunityView(User)).Render(c.Request().Context(), c.Response())
	})
	r2.GET("/thread", handlers.Thread)
	r2.GET("/video", handlers.Video)
	r2.GET("/tendance", handlers.PostTendenceHandler)
	r2.POST("/update-profile", handlers.UpdateUserProfile)
	r2.POST("/update-passwd", handlers.UpdateUserPasswd)
	r2.GET("/parametre", func(c echo.Context) error {
		str, _ := utils.Authorize(c)
		var User module.UserData
		module.Db.Model(&module.UserData{}).Where("user_id = ?", str).Find(&User)
		return components.Page(profile.Settings(User)).Render(c.Request().Context(), c.Response())
	})
	r2.GET("/profile", func(c echo.Context) error {
		str, _ := utils.Authorize(c)
		var posts []module.PostWithCommunity
		var comments []module.Comments
		var User module.UserData
		module.Db.Model(&module.UserData{}).Where("user_id = ?", str).Find(&User)
		postQuery := `
			SELECT p.*, c.name as community_name, c.communities_media as community_media,
				COALESCE(v.vote, 0) as user_vote
			FROM posts p
			LEFT JOIN communities c ON p.community_id = c.community_id
			LEFT JOIN post_vote v ON p.post_id = v.post_id AND v.user_id = ?
			where p.user_id = ?
		`

		if err := module.Db.Raw(postQuery, str, str).Scan(&posts).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to retrieve posts",
			})
		}

		overview := echo.Map{
			"posts":    posts,
			"comments": comments,
		}
		return components.Page(profile.Overview(User, overview)).Render(c.Request().Context(), c.Response())
	})
	r2.GET("/profile/:section", handlers.GetUserProfile)
	r2.GET("/parametre/compte", func(c echo.Context) error {
		str, _ := utils.Authorize(c)
		var User module.UserData
		module.Db.Model(&module.UserData{}).Where("user_id = ?", str).Find(&User)
		return components.Page(profile.SettingsAccount(User)).Render(c.Request().Context(), c.Response())
	})
	r2.GET("/register", func(c echo.Context) error {
		str, _ := utils.Authorize(c)
		var User struct {
			UserId     uint
			Username   string
			ProfileUrl string
		}
		module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
		return components.Page(register.AddNewPost(User)).Render(c.Request().Context(), c.Response())
	})
	r2.GET("/register/agence", func(c echo.Context) error {
		str, _ := utils.Authorize(c)
		var User struct {
			UserId     uint
			Username   string
			ProfileUrl string
		}
		module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
		return components.Page(register.AgenceForm(User)).Render(c.Request().Context(), c.Response())
	})
	r2.POST("/add-agency", handlers.AgencyAdd)
	r2.GET("/register/voiturelocation", func(c echo.Context) error {
		str, _ := utils.Authorize(c)
		var User struct {
			UserId     uint
			Username   string
			ProfileUrl string
		}
		module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
		var result []module.AgenciesData
		module.Db.Table("user_agencies").Select("agencies_data.*").
			Joins("left join agencies_data on agencies_data.id = user_agencies.agency_id").
			Where("user_agencies.user_id = ?", User.UserId).
			Scan(&result)
		if len(result) == 0 {
			return components.Page(register.VoitureRedirect(User)).Render(c.Request().Context(), c.Response())
		}
		return components.Page(register.VoitureLocation(User, result)).Render(c.Request().Context(), c.Response())
	})
	r2.GET("/register/voiturevendre", func(c echo.Context) error {
		str, _ := utils.Authorize(c)
		var User struct {
			UserId     uint
			Username   string
			ProfileUrl string
		}
		module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
		return components.Page(register.VoitureVendre(User)).Render(c.Request().Context(), c.Response())
	})
	r2.GET("/register/voitureparts", func(c echo.Context) error {
		str, _ := utils.Authorize(c)
		var User struct {
			UserId     uint
			Username   string
			ProfileUrl string
		}
		module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)
		return components.Page(register.VoitureParts(User)).Render(c.Request().Context(), c.Response())
	})
	r2.GET("/:type/:username/comments/:id", handlers.CommentsView)
	r2.GET("/:type/:username/comments/:id/replies/:comment", handlers.CommentsReplies)
	r2.POST("/add-voiture", handlers.VoitureLocation)
	r2.POST("/add-comment", handlers.CreateComment)
	r2.POST("/add-comunity", handlers.ComunityAdd)
	r2.POST("/add-car-parts", handlers.VoitureParts)
	r2.POST("/add-car-sales", handlers.VoitureSales)
	r2.GET("/disconnect", func(c echo.Context) error {
		cookie, err := c.Cookie("Token")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "No JWT cookie found"})
		}

		expiredCookie := &http.Cookie{
			Name:     cookie.Name,
			Value:    "",
			Expires:  time.Now().Add(-1 * time.Hour),
			Path:     "/",
			HttpOnly: true,
		}
		c.SetCookie(expiredCookie)

		return c.JSON(200, map[string]string{"message": "Successfully disconnected"})
	})
	e.Static("/assets", "assets")
	e.Logger.Fatal(e.Start(":3000"))
}
