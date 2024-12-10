package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/Yahya-Elamri/signeitfaster/views/components"
	"github.com/Yahya-Elamri/signeitfaster/views/home"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func getComments(Type, id string, userID uint) ([]module.CommentWithUser, error) {
	var comments []module.CommentWithUser
	if Type == "p" {
		query := `
			SELECT
				comments.comment_id,
				comments.content,
				comments.created_at,
				COALESCE(vote_data.vote, 0) AS user_vote,
				user_data.username,
				user_data.user_id,
				user_data.profile_url,
				comments.votes,
				COALESCE(reply_data.reply_total, 0) AS reply_count
			FROM comments
			LEFT JOIN user_data ON comments.user_id = user_data.user_id
			LEFT JOIN comment_vote vote_data ON comments.comment_id = vote_data.comment_id AND vote_data.user_id = ?
			LEFT JOIN (
				SELECT parent_comment_id, COUNT(*) AS reply_total
				FROM comments
				WHERE parent_comment_id IS NOT NULL
				GROUP BY parent_comment_id
			) AS reply_data ON comments.comment_id = reply_data.parent_comment_id
			WHERE comments.post_id = ? AND comments.parent_comment_id IS NULL
			ORDER BY comments.created_at ASC;
		`
		if err := module.Db.Raw(query, userID, id).Scan(&comments).Error; err != nil {
			return nil, err
		}
	}
	if Type == "t" {
		query := `
			SELECT
				comments.comment_id,
				comments.content,
				comments.created_at,
				COALESCE(vote_data.vote, 0) AS user_vote,
				user_data.username,
				user_data.user_id,
				user_data.profile_url,
				comments.votes,
				COALESCE(reply_data.reply_total, 0) AS reply_count
			FROM comments
			LEFT JOIN user_data ON comments.user_id = user_data.user_id
			LEFT JOIN comment_vote vote_data ON comments.comment_id = vote_data.comment_id AND vote_data.user_id = ?
			LEFT JOIN (
				SELECT parent_comment_id, COUNT(*) AS reply_total
				FROM comments
				WHERE parent_comment_id IS NOT NULL
				GROUP BY parent_comment_id
			) AS reply_data ON comments.comment_id = reply_data.parent_comment_id
			WHERE comments.thread_id = ? AND comments.parent_comment_id IS NULL
			ORDER BY comments.created_at ASC;
		`
		if err := module.Db.Raw(query, userID, id).Scan(&comments).Error; err != nil {
			return nil, err
		}
	}
	return comments, nil
}

func CommentsReplies(c echo.Context) error {
	userID, _ := utils.Authorize(c)
	Type := c.Param("type")
	id := c.Param("id")
	CommentId := c.Param("comment")
	posts := 0
	threads := 0
	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	if err := module.Db.Model(&module.UserData{}).Select("user_id, username, profile_url").Where("user_id = ?", userID).Find(&User).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve user data",
		})
	}

	var comments []module.CommentWithUser
	if Type == "p" {
		posts, _ = strconv.Atoi(id)
		query := `
			SELECT
				comments.comment_id,
				comments.content,
				comments.created_at,
				COALESCE(vote_data.vote, 0) AS user_vote,
				user_data.username,
				user_data.user_id,
				user_data.profile_url,
				comments.votes,
				COALESCE(reply_data.reply_total, 0) AS reply_count
			FROM comments
			LEFT JOIN user_data ON comments.user_id = user_data.user_id
			LEFT JOIN comment_vote vote_data ON comments.comment_id = vote_data.comment_id AND vote_data.user_id = ?
			LEFT JOIN (
				SELECT parent_comment_id, COUNT(*) AS reply_total
				FROM comments
				WHERE parent_comment_id IS NOT NULL
				GROUP BY parent_comment_id
			) AS reply_data ON comments.comment_id = reply_data.parent_comment_id
			WHERE comments.post_id = ? AND comments.parent_comment_id = ?
			ORDER BY comments.created_at ASC;
		`
		if err := module.Db.Raw(query, userID, id, CommentId).Scan(&comments).Error; err != nil {
			return err
		}
	}
	if Type == "t" {
		threads, _ = strconv.Atoi(id)
		query := `
			SELECT
				comments.comment_id,
				comments.content,
				comments.created_at,
				COALESCE(vote_data.vote, 0) AS user_vote,
				user_data.username,
				user_data.user_id,
				user_data.profile_url,
				comments.votes,
				COALESCE(reply_data.reply_total, 0) AS reply_count
			FROM comments
			LEFT JOIN user_data ON comments.user_id = user_data.user_id
			LEFT JOIN comment_vote vote_data ON comments.comment_id = vote_data.comment_id AND vote_data.user_id = ?
			LEFT JOIN (
				SELECT parent_comment_id, COUNT(*) AS reply_total
				FROM comments
				WHERE parent_comment_id IS NOT NULL
				GROUP BY parent_comment_id
			) AS reply_data ON comments.comment_id = reply_data.parent_comment_id
			WHERE comments.thread_id = ? AND comments.parent_comment_id = ?
			ORDER BY comments.created_at ASC;
		`
		if err := module.Db.Raw(query, userID, id, CommentId).Scan(&comments).Error; err != nil {
			return err
		}
	}

	return components.Comment(User, comments, posts, threads).Render(c.Request().Context(), c.Response())
}

func CommentsView(c echo.Context) error {
	userID, _ := utils.Authorize(c)
	Type := c.Param("type")
	id := c.Param("id")
	// username := c.Param("username")

	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	if err := module.Db.Model(&module.UserData{}).Select("user_id, username, profile_url").Where("user_id = ?", userID).Find(&User).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve user data",
		})
	}

	var communitiesSugg []module.Communities
	query := `
		(
			SELECT * FROM communities
			ORDER BY total_joined DESC
			LIMIT 1
		)
		UNION
		(
			SELECT * FROM communities
			ORDER BY RAND()
			LIMIT 1
		)
		UNION
		(
			SELECT c.* FROM communities c
			JOIN user_communities uc ON uc.community_id = c.community_id
			WHERE uc.user_id = ?
			LIMIT 2
		)
	`

	if err := module.Db.Raw(query, User.UserId).Scan(&communitiesSugg).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve community suggestions",
		})
	}

	comments, _ := getComments(Type, id, userID)

	if Type == "p" {
		var posts *module.PostWithCommunity
		postQuery := `
			SELECT p.*, 
			       c.name AS community_name, 
			       c.communities_media AS community_media,
			       COALESCE(v.vote, 0) AS user_vote
			FROM posts p
			LEFT JOIN communities c ON p.community_id = c.community_id
			LEFT JOIN post_vote v ON p.post_id = v.post_id AND v.user_id = ?
			LEFT JOIN user_data u ON p.user_id = u.user_id 
			WHERE p.post_id = ?
		`
		if err := module.Db.Raw(postQuery, userID, id).Scan(&posts).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to retrieve posts",
			})
		}

		return components.Page(home.TheOneView(User, posts, nil, comments, communitiesSugg)).Render(c.Request().Context(), c.Response())
	}

	if Type == "t" {
		var threads *module.ThreadWithCommunity
		postQuery := `
			SELECT p.*, c.name as community_name, c.communities_media as community_media,
				COALESCE(v.vote, 0) as user_vote , u.username , u.profile_url
			FROM threads p
			LEFT JOIN communities c ON p.community_id = c.community_id
			LEFT JOIN thread_vote v ON p.thread_id = v.thread_id AND v.user_id = ?
			LEFT JOIN user_data u ON p.user_id = u.user_id
			WHERE p.thread_id = ?
		`

		if err := module.Db.Raw(postQuery, userID, id).Scan(&threads).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to retrieve threads",
			})
		}

		return components.Page(home.TheOneView(User, nil, threads, comments, communitiesSugg)).Render(c.Request().Context(), c.Response())
	}

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": "Failed",
	})
}

// CreateComment handles POST requests to create a new comment
func CreateComment(c echo.Context) error {
	// Bind the request body to the Comment struct
	newComment := new(module.Comments)
	if err := c.Bind(newComment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	newComment.CreatedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))
	newComment.UpdatedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))
	if err := module.Db.Create(&newComment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add user to community"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Comment created successfully",
	})
}

func VoteOnCommentHandler(c echo.Context) error {
	CommentID := c.Param("comment_id") // Retrieve post ID from the request parameters
	UserId, _ := utils.Authorize(c)
	voteType := c.FormValue("vote") // Get the vote type from the form

	var vote int
	if voteType == "upvote" {
		vote = 1
	} else if voteType == "downvote" {
		vote = -1
	} else {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid vote"})
	}

	// Check for existing vote
	var existingVote module.CommentVote
	CommentIDUint, _ := strconv.ParseUint(CommentID, 10, 32)
	result := module.Db.Table("comment_vote").Where("comment_id = ? AND user_id = ?", CommentIDUint, UserId).First(&existingVote)
	var Comment module.Comments
	if err := module.Db.First(&Comment, CommentIDUint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error fetching post"})
	}

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error checking vote"})
	}

	// Case 1: User has voted before
	if existingVote.CommentID != 0 && existingVote.UserID != 0 {
		// If user clicks the same vote (e.g., upvote after already upvoting), remove the vote
		if existingVote.Vote == vote {
			Comment.Votes -= vote // Revert the vote count by removing the previous vote
			if err := module.Db.Table("comment_vote").Delete(&existingVote).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error removing vote"})
			}
			// Set classes to neutral state after removing vote
			existingVote.Vote = 0
		} else {
			// Change vote (from upvote to downvote or vice versa)
			Comment.Votes += 2 * vote // Adjust the vote count (from -1 to +1 or +1 to -1)
			existingVote.Vote = vote
			existingVote.VotedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))
			if err := module.Db.Table("comment_vote").Save(&existingVote).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error updating vote"})
			}
		}
	} else {
		// Case 2: User is voting for the first time
		newVote := module.CommentVote{
			CommentID: uint(CommentIDUint),
			UserID:    UserId,
			Vote:      vote,
			VotedAt:   []uint8(time.Now().Format("2006-01-02 15:04:05")),
		}
		Comment.Votes += vote // Increment or decrement the vote count
		if err := module.Db.Table("comment_vote").Create(&newVote).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error casting vote"})
		}
		existingVote.Vote = vote // Set the vote for class determination
	}

	if err := module.Db.Save(&Comment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error updating post votes"})
	}

	// Determine the CSS classes based on the vote status
	var upvoteClass, downvoteClass string
	if existingVote.Vote == 1 {
		upvoteClass = "stroke-green-500"
		downvoteClass = "stroke-gray-800 dark:stroke-white  hover:stroke-red-500 transition-colors duration-200"
	} else if existingVote.Vote == -1 {
		downvoteClass = "stroke-red-500"
		upvoteClass = "stroke-gray-800 dark:stroke-white hover:stroke-green-500 transition-colors duration-300"
	} else {
		upvoteClass = "stroke-gray-800 dark:stroke-white hover:stroke-green-500 transition-colors duration-300"
		downvoteClass = "stroke-gray-800 dark:stroke-white hover:stroke-red-500 transition-colors duration-200"
	}

	// Generate updated HTML
	html := fmt.Sprintf(`
		<div hx-target="this" hx-swap="outerHTML" class="flex bg-[#eaedef] dark:bg-[#060809] items-center rounded-full">
			<form hx-post="/comment/%d/vote" hx-include="this" class="flex hover:bg-gray-200 dark:bg-[#060809] text-black dark:text-white items-center rounded-full p-2">
				<input type="hidden" name="vote" value="upvote">
				<button type="submit">
					<svg width="20px" class="%s cursor-pointer" viewBox="-0.5 0 25 25" fill="none" xmlns="http://www.w3.org/2000/svg">
						<path d="M8 13.8599L10.87 10.8C11.0125 10.6416 11.1868 10.5149 11.3815 10.4282C11.5761 10.3415 11.7869 10.2966 12 10.2966C12.2131 10.2966 12.4239 10.3415 12.6185 10.4282C12.8132 10.5149 12.9875 10.6416 13.13 10.8L16 13.8599" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
						<path d="M3 7.41992L3 17.4199C3 19.6291 4.79086 21.4199 7 21.4199H17C19.2091 21.4199 21 19.6291 21 17.4199V7.41992C21 5.21078 19.2091 3.41992 17 3.41992H7C4.79086 3.41992 3 5.21078 3 7.41992Z" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
					</svg>
				</button>
			</form>
			
			<p class="poppins-regular text-md">%d</p>
			
			<form hx-post="/comment/%d/vote" hx-include="this" class="hover:bg-gray-200 dark:bg-[#060809] text-black dark:text-white flex items-center rounded-full p-2">
				<input type="hidden" name="vote" value="downvote">
				<button type="submit">
					<svg width="20px" class="rotate-180 cursor-pointer %s" viewBox="-0.5 0 25 25" fill="none" xmlns="http://www.w3.org/2000/svg">
						<path d="M8 13.8599L10.87 10.8C11.0125 10.6416 11.1868 10.5149 11.3815 10.4282C11.5761 10.3415 11.7869 10.2966 12 10.2966C12.2131 10.2966 12.4239 10.3415 12.6185 10.4282C12.8132 10.5149 12.9875 10.6416 13.13 10.8L16 13.8599" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
						<path d="M3 7.41992L3 17.4199C3 19.6291 4.79086 21.4199 7 21.4199H17C19.2091 21.4199 21 19.6291 21 17.4199V7.41992C21 5.21078 19.2091 3.41992 17 3.41992H7C4.79086 3.41992 3 5.21078 3 7.41992Z" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
					</svg>
				</button>
			</form>
		</div>
	`, CommentIDUint, upvoteClass, Comment.Votes, CommentIDUint, downvoteClass)

	return c.HTML(http.StatusOK, html)
}
