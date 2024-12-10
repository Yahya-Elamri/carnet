package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/Yahya-Elamri/signeitfaster/views/communities"
	"github.com/Yahya-Elamri/signeitfaster/views/components"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func detectCategory(title, description string) string {
	title = strings.ToLower(title)
	description = strings.ToLower(description)

	// Car Repair
	repairKeywords := []string{"tsli7", "mekanik", "garage", "service", "motor", "tomobil", "wrsha", "msla7",
		"repair", "mechanic", "car", "engine", "workshop", "fix",
		"réparation", "mécanique", "voiture", "moteur", "atelier", "réparer"}
	for _, keyword := range repairKeywords {
		if strings.Contains(title, keyword) || strings.Contains(description, keyword) {
			return "Islah Tomobilat / Réparation Automobile"
		}
	}

	// Car Meetings
	meetingKeywords := []string{"tajma3", "ijtima3", "l9a", "group", "7afl", "7da", "khrja",
		"meeting", "gathering", "event", "group", "party", "outing",
		"réunion", "rassemblement", "événement", "groupe", "fête", "sortie"}
	for _, keyword := range meetingKeywords {
		if strings.Contains(title, keyword) || strings.Contains(description, keyword) {
			return "Tajma3at Tomobilat / Rencontres Automobiles"
		}
	}

	// Car Sales
	salesKeywords := []string{"biya3", "chra", "tomobil lbiya3", "tfahm", "3roud", "thmn", "jdida",
		"sale", "buy", "car for sale", "deal", "offers", "price", "new",
		"vente", "achat", "voiture à vendre", "accord", "offres", "prix", "nouvelle"}
	for _, keyword := range salesKeywords {
		if strings.Contains(title, keyword) || strings.Contains(description, keyword) {
			return "Biya3 Tomobilat / Vente de Voitures"
		}
	}

	// Car Parts
	partsKeywords := []string{"piyass", "9ta3", "accessoire", "m7rak", "khaddam", "parts", "9i3a",
		"parts", "accessories", "engine", "spare",
		"pièces", "accessoires", "moteur", "pièces détachées"}
	for _, keyword := range partsKeywords {
		if strings.Contains(title, keyword) || strings.Contains(description, keyword) {
			return "Piyass Tomobilat / Pièces Automobiles"
		}
	}

	// Car Rentals
	rentalKeywords := []string{"kra", "tkri", "tomobil lkra", "louer", "ijar",
		"rent", "rental", "car for rent", "hire", "leasing",
		"louer", "location", "voiture de location", "bail"}
	for _, keyword := range rentalKeywords {
		if strings.Contains(title, keyword) || strings.Contains(description, keyword) {
			return "Kra Tomobilat / Location de Voitures"
		}
	}

	// Car News
	newsKeywords := []string{"jdid", "akhbar", "mostajdat", "tomobil jdida", "akhbar sayarat",
		"news", "latest", "updates", "new car", "automotive news",
		"nouvelles", "dernières", "mises à jour", "nouvelle voiture", "actualités automobiles"}
	for _, keyword := range newsKeywords {
		if strings.Contains(title, keyword) || strings.Contains(description, keyword) {
			return "Akhbar Tomobilat / Actualités Automobiles"
		}
	}

	return "Autre" // Default category if no match is found
}

func CommunautesView(c echo.Context) error {
	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}

	// Fetch the authenticated user
	str, _ := utils.Authorize(c)
	if err := module.Db.Model(&module.UserData{}).
		Select("user_id, username, profile_url").
		Where("user_id = ?", str).
		Find(&User).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve user data",
		})
	}

	// Fetch all communities
	var communitie []module.Communities
	if err := module.Db.Find(&communitie).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve communities",
		})
	}

	// Map to store categorized communities
	categorizedCommunities := map[string][]module.Communities{}

	// Categorize communities in a single loop
	for _, community := range communitie {
		category := detectCategory(community.Name, community.Description)
		if _, exists := categorizedCommunities[category]; !exists {
			categorizedCommunities[category] = []module.Communities{}
		}
		categorizedCommunities[category] = append(categorizedCommunities[category], community)
	}

	// Filter out empty categories
	filteredCategories := make(map[string][]module.Communities)
	for category, coms := range categorizedCommunities {
		if len(coms) > 0 {
			filteredCategories[category] = coms
		}
	}

	// Render the view with filtered categories
	return components.Page(communities.CommunautesView(User, filteredCategories)).
		Render(c.Request().Context(), c.Response())
}

func CommunautesOneView(c echo.Context) error {
	name := c.Param("name")
	var posts []module.PostWithCommunity

	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	str, _ := utils.Authorize(c)
	module.Db.Model(&module.UserData{}).Select("user_id,username, profile_url").Where("user_id = ?", str).Find(&User)

	var communitie module.Communities
	var userCommunitie module.UserCommunities

	// Query to get all communities
	if err := module.Db.Where("name = ?", name).First(&communitie).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Community not found")
	}

	if err := module.Db.Where("user_id = ? AND community_id = ?", User.UserId, communitie.CommunityID).First(&userCommunitie).Error; err != nil {
		userCommunitie.CommunityID = 0
		userCommunitie.UserId = 0
	}

	postQuery := `
		SELECT p.*, c.name as community_name, c.communities_media as community_media,
		       COALESCE(v.vote, 0) as user_vote, u.username, u.profile_url
		FROM posts p
		LEFT JOIN communities c ON p.community_id = c.community_id
		LEFT JOIN post_vote v ON p.post_id = v.post_id AND v.user_id = ?
		LEFT JOIN user_data u ON p.user_id = u.user_id
		WHERE p.community_id = ?
	`

	if err := module.Db.Raw(postQuery, User.UserId, communitie.CommunityID).Scan(&posts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve posts",
		})
	}

	return components.Page(communities.CommunautesOneView(User, communitie, userCommunitie, posts)).Render(c.Request().Context(), c.Response())
}

func CommunautesJoin(c echo.Context) error {
	user_id := c.FormValue("user_id")
	community_id := c.FormValue("community_id")
	u1, _ := strconv.ParseUint(user_id, 10, 32)
	u2, _ := strconv.ParseUint(community_id, 10, 32)

	userCommunity := module.UserCommunities{
		UserId:      uint(u1),
		CommunityID: uint(u2),
		Role:        "user", // or another role as needed
		JoinedAt:    []uint8(time.Now().Format("2006-01-02 15:04:05")),
	}

	if err := module.Db.Create(&userCommunity).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add user to community"})
	}

	err := module.Db.Model(&module.Communities{}).Where("community_id = ?", community_id).UpdateColumn("total_joined", gorm.Expr("total_joined + ?", 1)).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to join"})
	}

	c.Response().Header().Set("HX-Redirect", "")
	return c.NoContent(200)
}

func JoinedCommunautes(c echo.Context) error {
	type CommunityInfo struct {
		CommunityID      uint   `json:"community_id"`
		Name             string `json:"name"`
		CommunitiesMedia string `json:"communities_media"`
	}
	userID, _ := utils.Authorize(c)
	var joinedCommunities []CommunityInfo
	if err := module.Db.Table("communities").
		Select("communities.community_id, communities.name, communities.communities_media").
		Joins("join user_communities on user_communities.community_id = communities.community_id").
		Where("user_communities.user_id = ?", userID).
		Scan(&joinedCommunities).Limit(4).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	htmlResponse := ""
	if len(joinedCommunities) > 0 {
		for _, community := range joinedCommunities {
			htmlResponse += "<li class='flex items-center gap-2 hover:bg-[#f1f1f1] dark:hover:bg-[#060809] rounded-md w-full px-2 py-2'>"
			htmlResponse += "<img src=\"/assets" + utils.Transform(community.CommunitiesMedia)[0] + "\" alt=\"" + community.Name + "\" class='rounded-full' style='width: 43px; height: 43px;'>"
			htmlResponse += "<h2 class='poppins-regular text-md'>" + community.Name + "</h2>"
			htmlResponse += "</li>"
		}
	} else {
		htmlResponse += "<li>No communities joined.</li>"
	}

	return c.HTML(200, htmlResponse)
}
