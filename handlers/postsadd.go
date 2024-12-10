package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PostsAdd(c echo.Context) error {
	post := new(module.Posts)
	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to bind community data",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	//handling profile imgs
	files := form.File["posts_media"]
	imageUrls, imageExt, err := utils.SaveUploadedFilesExtentions(files)
	if err != nil {
		return err
	}
	imageUrlsJSON, err := json.Marshal(imageUrls)
	if err != nil {
		return err
	}
	imageExtJSON, err := json.Marshal(imageExt)
	if err != nil {
		return err
	}
	post.PostsMedia = string(imageUrlsJSON)
	post.PostsMediaExtention = string(imageExtJSON)
	post.CreatedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))
	post.EditedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))

	if err := module.Db.Create(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create community in the database",
		})
	}
	c.Response().Header().Set("HX-Redirect", "")
	return c.NoContent(200)
}

func VoteOnPostHandler(c echo.Context) error {
	postID := c.Param("post_id") // Retrieve post ID from the request parameters
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
	var existingVote module.PostVote
	postIDUint, _ := strconv.ParseUint(postID, 10, 32)
	result := module.Db.Table("post_vote").Where("post_id = ? AND user_id = ?", postIDUint, UserId).First(&existingVote)
	var post module.Posts
	if err := module.Db.First(&post, postIDUint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error fetching post"})
	}

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error checking vote"})
	}

	// Case 1: User has voted before
	if existingVote.PostID != 0 && existingVote.UserID != 0 {
		// If user clicks the same vote (e.g., upvote after already upvoting), remove the vote
		if existingVote.Vote == vote {
			post.Votes -= vote // Revert the vote count by removing the previous vote
			if err := module.Db.Table("post_vote").Delete(&existingVote).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error removing vote"})
			}
			// Set classes to neutral state after removing vote
			existingVote.Vote = 0
		} else {
			// Change vote (from upvote to downvote or vice versa)
			post.Votes += 2 * vote // Adjust the vote count (from -1 to +1 or +1 to -1)
			existingVote.Vote = vote
			existingVote.VotedAt = []uint8(time.Now().Format("2006-01-02 15:04:05"))
			if err := module.Db.Table("post_vote").Save(&existingVote).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error updating vote"})
			}
		}
	} else {
		// Case 2: User is voting for the first time
		newVote := module.PostVote{
			PostID:  uint(postIDUint),
			UserID:  UserId,
			Vote:    vote,
			VotedAt: []uint8(time.Now().Format("2006-01-02 15:04:05")),
		}
		post.Votes += vote // Increment or decrement the vote count
		if err := module.Db.Table("post_vote").Create(&newVote).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error casting vote"})
		}
		existingVote.Vote = vote // Set the vote for class determination
	}

	// Save the updated vote count to the post
	if err := module.Db.Save(&post).Error; err != nil {
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
			<form hx-post="/post/%d/vote" hx-include="this" class="flex hover:bg-gray-200 dark:bg-[#060809] text-black dark:text-white items-center rounded-full p-2">
				<input type="hidden" name="vote" value="upvote">
				<button type="submit">
					<svg width="28px" class="%s cursor-pointer" viewBox="-0.5 0 25 25" fill="none" xmlns="http://www.w3.org/2000/svg">
						<path d="M8 13.8599L10.87 10.8C11.0125 10.6416 11.1868 10.5149 11.3815 10.4282C11.5761 10.3415 11.7869 10.2966 12 10.2966C12.2131 10.2966 12.4239 10.3415 12.6185 10.4282C12.8132 10.5149 12.9875 10.6416 13.13 10.8L16 13.8599" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
						<path d="M3 7.41992L3 17.4199C3 19.6291 4.79086 21.4199 7 21.4199H17C19.2091 21.4199 21 19.6291 21 17.4199V7.41992C21 5.21078 19.2091 3.41992 17 3.41992H7C4.79086 3.41992 3 5.21078 3 7.41992Z" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
					</svg>
				</button>
			</form>
			
			<p class="poppins-regular text-md">%d</p>
			
			<form hx-post="/post/%d/vote" hx-include="this" class="hover:bg-gray-200 dark:bg-[#060809] text-black dark:text-white flex items-center rounded-full p-2">
				<input type="hidden" name="vote" value="downvote">
				<button type="submit">
					<svg width="28px" class="rotate-180 cursor-pointer %s" viewBox="-0.5 0 25 25" fill="none" xmlns="http://www.w3.org/2000/svg">
						<path d="M8 13.8599L10.87 10.8C11.0125 10.6416 11.1868 10.5149 11.3815 10.4282C11.5761 10.3415 11.7869 10.2966 12 10.2966C12.2131 10.2966 12.4239 10.3415 12.6185 10.4282C12.8132 10.5149 12.9875 10.6416 13.13 10.8L16 13.8599" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
						<path d="M3 7.41992L3 17.4199C3 19.6291 4.79086 21.4199 7 21.4199H17C19.2091 21.4199 21 19.6291 21 17.4199V7.41992C21 5.21078 19.2091 3.41992 17 3.41992H7C4.79086 3.41992 3 5.21078 3 7.41992Z" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
					</svg>
				</button>
			</form>
		</div>
	`, postIDUint, upvoteClass, post.Votes, postIDUint, downvoteClass)

	return c.HTML(http.StatusOK, html)
}
