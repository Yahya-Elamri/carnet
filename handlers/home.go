package handlers

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/Yahya-Elamri/signeitfaster/utils"
	"github.com/Yahya-Elamri/signeitfaster/views/components"
	"github.com/Yahya-Elamri/signeitfaster/views/home"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const dbTimeout = 5 * time.Second

func Home(c echo.Context) error {
	db := module.Db

	// Retrieve the user ID from the context or JWT
	userID, stm := utils.Authorize(c)

	// Context with timeout for database operations
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Channels to receive posts and threads
	postsCh := make(chan []module.PostWithCommunity)
	threadsCh := make(chan []module.ThreadWithCommunity)
	errCh := make(chan error, 2)

	// Fetch posts in a goroutine
	go func() {
		var posts []module.PostWithCommunity
		postQuery := `
			SELECT p.*, 
			       c.name AS community_name, 
			       c.communities_media AS community_media,
			       COALESCE(v.vote, 0) AS user_vote
			FROM posts p
			LEFT JOIN communities c ON p.community_id = c.community_id
			LEFT JOIN post_vote v ON p.post_id = v.post_id AND v.user_id = ?
		`
		if err := db.WithContext(ctx).Raw(postQuery, userID).Scan(&posts).Error; err != nil {
			errCh <- fmt.Errorf("error fetching posts: %v", err)
			return
		}
		postsCh <- posts
	}()

	// Fetch threads in a goroutine
	go func() {
		var threads []module.ThreadWithCommunity
		threadQuery := `
			SELECT t.*, 
			       c.name AS community_name, 
			       c.communities_media AS community_media,
			       COALESCE(v.vote, 0) AS user_vote
			FROM threads t
			LEFT JOIN communities c ON t.community_id = c.community_id
			LEFT JOIN thread_vote v ON t.thread_id = v.thread_id AND v.user_id = ?
		`
		if err := db.WithContext(ctx).Raw(threadQuery, userID).Scan(&threads).Error; err != nil {
			errCh <- fmt.Errorf("error fetching threads: %v", err)
			return
		}
		threadsCh <- threads
	}()

	// Wait for results or errors
	var posts []module.PostWithCommunity
	var threads []module.ThreadWithCommunity

	select {
	case err := <-errCh:
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	case posts = <-postsCh:
		// Received posts
	}

	select {
	case err := <-errCh:
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	case threads = <-threadsCh:
		// Received threads
	}

	mixedItems := make([]module.MixedItem, 0, len(posts)+len(threads))

	// Append posts and threads
	for _, post := range posts {
		mixedItems = append(mixedItems, module.MixedItem{Type: "post", Content: post})
	}
	for _, thread := range threads {
		mixedItems = append(mixedItems, module.MixedItem{Type: "thread", Content: thread})
	}

	// Shuffle the mixed items to randomize the order
	rand.Shuffle(len(mixedItems), func(i, j int) { mixedItems[i], mixedItems[j] = mixedItems[j], mixedItems[i] })

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

	if stm {
		return components.Page(home.TheRealHome(User, mixedItems, communitiesSugg)).Render(c.Request().Context(), c.Response())
	}

	return components.Page(home.Home()).Render(c.Request().Context(), c.Response())
}

func Thread(c echo.Context) error {
	var threads []module.ThreadWithCommunity

	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	str, stm := utils.Authorize(c)
	if err := module.Db.Model(&module.UserData{}).Select("user_id, username, profile_url").Where("user_id = ?", str).Find(&User).Error; err != nil {
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

	// Fetch posts along with community data and user vote
	postQuery := `
		SELECT p.*, c.name as community_name, c.communities_media as community_media,
		       COALESCE(v.vote, 0) as user_vote
		FROM threads p
		LEFT JOIN communities c ON p.community_id = c.community_id
		LEFT JOIN thread_vote v ON p.thread_id = v.thread_id AND v.user_id = ?
	`

	if err := module.Db.Raw(postQuery, User.UserId).Scan(&threads).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve posts",
		})
	}

	if stm {
		return components.Page(home.HomeLogged(User, nil, threads, communitiesSugg)).Render(c.Request().Context(), c.Response())
	}

	return components.Page(home.Home()).Render(c.Request().Context(), c.Response())
}

func Video(c echo.Context) error {
	var posts []module.PostWithCommunity

	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	str, stm := utils.Authorize(c)
	if err := module.Db.Model(&module.UserData{}).Select("user_id, username, profile_url").Where("user_id = ?", str).Find(&User).Error; err != nil {
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

	// Fetch posts along with community data and user vote
	postQuery := `
		SELECT p.*, c.name as community_name, c.communities_media as community_media,
		       COALESCE(v.vote, 0) as user_vote
		FROM posts p
		LEFT JOIN communities c ON p.community_id = c.community_id
		LEFT JOIN post_vote v ON p.post_id = v.post_id AND v.user_id = ?
		WHERE posts_media_extention = ?
	`

	if err := module.Db.Raw(postQuery, User.UserId, "[\".mp4\"]").Scan(&posts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve posts",
		})
	}

	if stm {
		return components.Page(home.HomeLogged(User, posts, nil, communitiesSugg)).Render(c.Request().Context(), c.Response())
	}

	return components.Page(home.Home()).Render(c.Request().Context(), c.Response())
}

func getListings(limit, offset int) ([]module.MixedListing, int64, error) {
	var mixedListings []module.MixedListing
	var totalCount1, totalCount2, totalCount3 int64

	// Helper function to build the base query
	buildBaseQuery := func(table, joins string) *gorm.DB {
		return module.Db.Table(table).
			Joins(joins)
	}

	// Car Parts Listings
	var carParts []module.MixedListing
	carPartsCountQuery := buildBaseQuery("car_parts_listings", "left join user_data on car_parts_listings.username = user_data.username")
	carPartsCountQuery.Count(&totalCount1)
	carPartsDataQuery := buildBaseQuery("car_parts_listings", "left join user_data on car_parts_listings.username = user_data.username")
	carPartsDataQuery.Select("car_parts_listings.*, 'part' as type, user_data.user_id, user_data.email, user_data.address, user_data.phone, user_data.first_name, user_data.last_name, user_data.profile_url").
		Limit(limit).Offset(offset).
		Scan(&carParts)
	mixedListings = append(mixedListings, carParts...)

	// Car Sales Listings
	var carSales []module.MixedListing
	carSalesCountQuery := buildBaseQuery("car_sales_listings", "left join user_data on car_sales_listings.username = user_data.username")
	carSalesCountQuery.Count(&totalCount2)
	carSalesDataQuery := buildBaseQuery("car_sales_listings", "left join user_data on car_sales_listings.username = user_data.username")
	carSalesDataQuery.Select("car_sales_listings.*, 'sale' as type, user_data.user_id, user_data.email, user_data.address, user_data.phone, user_data.first_name, user_data.last_name, user_data.profile_url").
		Limit(limit).Offset(offset).
		Scan(&carSales)
	mixedListings = append(mixedListings, carSales...)

	// Car Rentals Listings
	var carListings []module.MixedListing
	carListingsCountQuery := buildBaseQuery("car_listings",
		"left join agencies_data on car_listings.agency_id = agencies_data.id left join user_data on agencies_data.id = user_data.user_id")
	carListingsCountQuery.Count(&totalCount3)
	carListingsDataQuery := buildBaseQuery("car_listings",
		"left join agencies_data on car_listings.agency_id = agencies_data.id left join user_data on agencies_data.id = user_data.user_id")
	carListingsDataQuery.Select("car_listings.*, 'rent' as type, user_data.user_id, agencies_data.email, agencies_data.address, agencies_data.phone, user_data.first_name, user_data.last_name, agencies_data.username, user_data.profile_url").
		Limit(limit).Offset(offset).
		Scan(&carListings)
	mixedListings = append(mixedListings, carListings...)

	// Calculate total count
	totalCount := totalCount1 + totalCount2 + totalCount3

	return mixedListings, totalCount, nil
}

func MarketplaceHome(c echo.Context) error {
	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	// Default page and limit
	page := 1
	limit := 3

	// Get page parameter from query
	if p := c.QueryParam("page"); p != "" {
		if pNum, err := strconv.Atoi(p); err == nil && pNum > 0 {
			page = pNum
		}
	}

	// Calculate offset based on page and limit
	offset := (page - 1) * limit

	// Get mixed listings and total count
	mixedListings, totalCount, _ := getListings(limit, offset)

	// Sort by timestamp
	mixedListings, err := utils.SortByTimestamp(mixedListings)
	if err != nil {
		return err
	}

	// Calculate total pages for pagination
	totalPages := (totalCount + int64(limit) - 1) / int64(limit)

	// Authorize user
	str, stm := utils.Authorize(c)
	module.Db.Model(&module.UserData{}).Select("username, profile_url").Where("user_id = ?", str).Find(&User)
	// Render the page with pagination
	if stm {
		return components.Page(home.HomeLoggedMarketplace(User, mixedListings, totalPages)).Render(c.Request().Context(), c.Response())
	}

	return components.Page(home.Home()).Render(c.Request().Context(), c.Response())
}

func PostTendenceHandler(c echo.Context) error {
	var posts []module.PostWithCommunity
	var User struct {
		UserId     uint
		Username   string
		ProfileUrl string
	}
	userId, _ := utils.Authorize(c)
	if err := module.Db.Model(&module.UserData{}).Select("user_id, username, profile_url").Where("user_id = ?", userId).Find(&User).Error; err != nil {
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
	// Fetch posts with votes, community data, and created_at from the database
	err := module.Db.Table("posts").
		Select(`
            posts.post_id, posts.community_id, posts.user_id, posts.content, posts.posts_media,
            posts.posts_media_extention, posts.votes, posts.created_at, posts.edited_at,
            communities.name AS community_name, communities.communities_media AS community_media,
            COALESCE(post_vote.vote, 0) AS user_vote
        `).
		Joins("LEFT JOIN communities ON posts.community_id = communities.community_id").
		Joins("LEFT JOIN post_vote ON posts.post_id = post_vote.post_id AND post_vote.user_id = ?", userId).
		Scan(&posts).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve posts"})
	}

	// Calculate trend score for each post
	for i, post := range posts {
		// Convert the created_at timestamp (assumed []uint8) to time.Time
		postCreationTime, err := time.Parse("2006-01-02 15:04:05", string(post.CreatedAt))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid timestamp format"})
		}

		// Calculate time difference in hours
		hoursSinceCreation := time.Since(postCreationTime).Hours()

		// Apply the trend score formula
		trendScore := float64(post.Votes) / math.Pow(hoursSinceCreation+2, 1.5)

		// Add trend score to each post (you can add a field in PostWithCommunity to store this)
		posts[i].TrendScore = trendScore
	}

	// Sort posts by trend score in descending order
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].TrendScore > posts[j].TrendScore
	})

	return components.Page(home.TendanceView(User, posts, communitiesSugg)).Render(c.Request().Context(), c.Response())
}
