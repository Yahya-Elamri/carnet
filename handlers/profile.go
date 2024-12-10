package handlers

import (
	"net/http"

	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/Yahya-Elamri/signeitfaster/views/components"
	"github.com/Yahya-Elamri/signeitfaster/views/profile"
	"github.com/labstack/echo/v4"
)

func GetUserProfile(c echo.Context) error {
	userId, _ := utils.Authorize(c)
	section := c.Param("section")

	var user module.UserData
	var communities []module.Communities
	var posts []module.PostWithCommunity
	var savedPosts []module.PostWithCommunity
	var comments []module.Comments
	var upvotes []module.PostWithCommunity
	var downvotes []module.PostWithCommunity

	if err := module.Db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}

	switch section {
	case "overview":
		postQuery := `
			SELECT p.*, c.name as community_name, c.communities_media as community_media,
				COALESCE(v.vote, 0) as user_vote
			FROM posts p
			LEFT JOIN communities c ON p.community_id = c.community_id
			LEFT JOIN post_vote v ON p.post_id = v.post_id AND v.user_id = ?
			where p.user_id = ?
		`

		if err := module.Db.Raw(postQuery, userId, userId).Scan(&posts).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to retrieve posts",
			})
		}

		overview := echo.Map{
			"posts":    posts,
			"comments": comments,
		}
		return components.Page(profile.Overview(user, overview)).Render(c.Request().Context(), c.Response())

	case "communities":
		if err := module.Db.Table("communities").
			Select("communities.*").
			Joins("JOIN user_communities ON user_communities.community_id = communities.community_id").
			Where("user_communities.user_id = ?", userId).
			Find(&communities).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error fetching communities"})
		}
		return components.Page(profile.Communities(user, communities)).Render(c.Request().Context(), c.Response())

	case "posts":
		if err := module.Db.Where("user_id = ?", userId).Find(&posts).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error fetching posts"})
		}
		return c.JSON(http.StatusOK, posts)

	case "saved":
		if err := module.Db.Table("posts").
			Select("posts.*").
			Joins("JOIN saved_posts ON saved_posts.post_id = posts.post_id").
			Where("saved_posts.user_id = ?", userId).
			Find(&savedPosts).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error fetching saved posts"})
		}
		return c.JSON(http.StatusOK, savedPosts)

	case "comments":
		if err := module.Db.Where("user_id = ?", userId).Find(&comments).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error fetching comments"})
		}
		return c.JSON(http.StatusOK, comments)

	case "upvotes":
		postQuery := `
			SELECT p.*, c.name as community_name, c.communities_media as community_media,
				COALESCE(v.vote, 0) as user_vote
			FROM posts p
			LEFT JOIN communities c ON p.community_id = c.community_id
			LEFT JOIN post_vote v ON p.post_id = v.post_id AND v.user_id = ?
			where v.vote = 1
		`

		if err := module.Db.Raw(postQuery, userId).Scan(&upvotes).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to retrieve posts",
			})
		}
		return components.Page(profile.Upvote(user, upvotes)).Render(c.Request().Context(), c.Response())
	case "downvotes":
		postQuery := `
			SELECT p.*, c.name as community_name, c.communities_media as community_media,
				COALESCE(v.vote, 0) as user_vote
			FROM posts p
			LEFT JOIN communities c ON p.community_id = c.community_id
			LEFT JOIN post_vote v ON p.post_id = v.post_id AND v.user_id = ?
			where v.vote = -1
		`
		if err := module.Db.Raw(postQuery, userId).Scan(&downvotes).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to retrieve posts",
			})
		}
		return components.Page(profile.Upvote(user, downvotes)).Render(c.Request().Context(), c.Response())

	default:
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid section"})
	}
}
