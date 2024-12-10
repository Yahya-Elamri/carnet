package module

import (
	"gorm.io/gorm"
)

type (
	UserData struct {
		UserId       uint    `gorm:"primaryKey"`
		Username     string  `gorm:"unique;not null" form:"Username" json:"username"`
		PasswordHash string  `gorm:"not null" form:"PasswordHash"`
		Email        string  `gorm:"unique;not null" form:"Email"`
		Phone        string  `form:"Phone"`
		Address      string  `gorm:"type:text" form:"address"`
		FirstName    string  `gorm:"not null" form:"FirstName"`
		LastName     string  `gorm:"not null" form:"LastName"`
		Description  string  `gorm:"type:text" form:"Description"`
		SocialLinks  string  `gorm:"type:text" form:"social_links[]"`
		ProfileUrl   string  `form:"ProfileUrl" gorm:"default:'/profile/Default-profile-pic.jpg'"`
		CreatedAt    []uint8 `json:"created_at"`
		UpdatedAt    []uint8 `json:"updated_at"`
		Status       string  `gorm:"type:enum('active', 'inactive', 'banned');default:'active'" form:"Status"`
	}
	AgenciesData struct {
		Id        uint    `gorm:"primaryKey"`
		Username  string  `gorm:"unique;not null" form:"Username" json:"username"`
		Email     string  `gorm:"unique;not null" form:"Email"`
		Address   string  `form:"Address"`
		Phone     string  `form:"Phone"`
		CreatedAt []uint8 `json:"created_at"`
	}

	CarListings struct {
		Id           int     `gorm:"primaryKey" form:"id"`
		AgencyID     int     `gorm:"not null" form:"agency_id"`
		Make         string  `gorm:"size:50;not null" form:"make"`
		Model        string  `gorm:"size:50;not null" form:"model"`
		Year         int     `gorm:"not null" form:"year"`
		Price        float64 `gorm:"type:decimal(10,2);not null" form:"daily_rate"`
		ImageURL     string  `gorm:"size:255" form:"images"`
		Description  string  `gorm:"type:text" form:"description"`
		FuelType     string  `gorm:"size:20" form:"fuel_type"`
		Transmission string  `gorm:"size:20" form:"transmission"`
		Mileage      int     `form:"mileage"`
		Seats        int     `form:"seats"`
		Color        string  `gorm:"size:30" form:"color"`
		Status       string  `gorm:"size:20;not null;default:'available'"`
		Category     string
		CreatedAt    []uint8 `json:"created_at"`
	}

	CarSalesListings struct {
		Id           int     `gorm:"primaryKey" form:"id"`
		Username     string  `gorm:"not null"`
		Make         string  `gorm:"size:50;not null" form:"make"`
		Model        string  `gorm:"size:50;not null" form:"model"`
		Year         int     `gorm:"not null" form:"year"`
		Price        float64 `gorm:"type:decimal(10,2);not null" form:"Price"`
		ImageURL     string  `gorm:"size:255" form:"images"`
		Description  string  `gorm:"type:text" form:"description"`
		FuelType     string  `gorm:"size:20" form:"fuel_type"`
		Transmission string  `gorm:"size:20" form:"transmission"`
		Mileage      int     `form:"mileage"`
		Seats        int     `form:"seats"`
		Color        string  `gorm:"size:30" form:"color"`
		Status       string  `gorm:"size:20;not null;default:'available'"`
		Category     string
		CreatedAt    []uint8 `json:"created_at"`
	}

	CarPartsListings struct {
		Id          uint    `gorm:"primaryKey" `
		Username    string  `gorm:"not null"`
		PartName    string  `gorm:"size:100;not null" form:"Parts"`
		Make        string  `gorm:"size:255;not null" form:"make"`
		Etat        string  `gorm:"size:255;not null" form:"etat"`
		Description string  `gorm:"type:text" form:"Description"`
		Price       float64 `gorm:"type:decimal(10,2);not null" form:"Price"`
		Stock       int     `gorm:"not null" form:"Stock"`
		ImageURL    string  `gorm:"size:255" form:"images"`
		Status      string  `gorm:"size:20;not null;default:'available'"`
		Category    string
		CreatedAt   []uint8 `json:"created_at"`
		UpdatedAt   []uint8 `json:"updated_at"`
	}

	UserAgencies struct {
		UserId   uint
		AgencyId uint
	}

	Communities struct {
		CommunityID       uint    `gorm:"primaryKey" json:"community_id" form:"community_id"`
		UserId            uint    `gorm:"not null" json:"user_id" form:"user_id"`
		Name              string  `gorm:"unique;not null" json:"name" form:"name"`
		Description       string  `gorm:"type:text" json:"description" form:"description"`
		CommunitiesMedia  string  `gorm:"default:'/communities/Default-communities-pic.jpg'" json:"communities_media" form:"communities_media"`
		CommunitiesBanner string  `gorm:"default:'/communities/Default-communities-banner.jpg'" json:"communities_banner" form:"communities_banner"`
		TotalJoined       int     `gorm:"default:1" json:"total_joined" form:"total_joined"`
		Category          string  `gorm:"-"`
		CreatedAt         []uint8 `gorm:"not null" json:"created_at" form:"created_at"`
	}

	UserCommunities struct {
		UserId      uint    `gorm:"primaryKey;not null" json:"user_id"`
		CommunityID uint    `gorm:"primaryKey;not null" json:"community_id"`
		Role        string  `gorm:"type:enum('user','admin','owner','banned');default:'user'" json:"role"`
		JoinedAt    []uint8 `gorm:"not null;default:CURRENT_TIMESTAMP" json:"joined_at"`
	}

	Threads struct {
		ThreadID    uint    `gorm:"primaryKey" json:"thread_id" form:"thread_id"`
		UserID      uint    `gorm:"not null" json:"user_id" form:"user_id"`
		CommunityID uint    `gorm:"not null" json:"community_id" form:"community_id"`
		Content     string  `gorm:"type:text" json:"content" form:"content"`
		Votes       int     `gorm:"not null;default:0" json:"votes" form:"votes"`
		Tags        string  `gorm:"not null" json:"tags" form:"tags"`
		CreatedAt   []uint8 `json:"created_at" form:"created_at"`
		Status      string  `gorm:"type:enum('active', 'inactive', 'banned');default:'active'" form:"Status"`
	}

	Posts struct {
		PostID              uint     `gorm:"primaryKey" json:"post_id" form:"post_id"`
		CommunityID         uint     `gorm:"not null" json:"community_id" form:"community_id"`
		UserID              uint     `gorm:"not null" json:"user_id" form:"user_id"`
		Content             string   `gorm:"type:text;not null" json:"content" form:"content"`
		PostsMedia          string   `json:"posts_media" form:"posts_media"`
		PostsMediaExtention string   `json:"posts_media_extention" form:"posts_media_extention"`
		Votes               int      `gorm:"not null;default:0" json:"votes" form:"votes"`
		User                UserData `json:"user" gorm:"foreignKey:UserID"`
		CreatedAt           []uint8  `json:"created_at" form:"created_at"`
		EditedAt            []uint8  `json:"edited_at" form:"edited_at"`
	}

	PostVote struct {
		UserID  uint     `gorm:"primaryKey;not null" json:"user_id"`
		PostID  uint     `gorm:"primaryKey;not null" json:"post_id"`
		Vote    int      `json:"vote"`
		VotedAt []uint8  `gorm:"autoCreateTime" json:"voted_at"`
		User    UserData `gorm:"foreignKey:UserID" json:"user"`
		Post    Posts    `gorm:"foreignKey:PostID" json:"post"`
	}

	ThreadVote struct {
		UserID   uint     `gorm:"primaryKey;not null" json:"user_id"`
		ThreadID uint     `gorm:"primaryKey;not null" json:"thread_id"`
		Vote     int      `json:"vote"`
		VotedAt  []uint8  `gorm:"autoCreateTime" json:"voted_at"`
		User     UserData `gorm:"foreignKey:UserID" json:"user"`
		Thread   Threads  `gorm:"foreignKey:ThreadID" json:"thread"`
	}

	PostWithCommunity struct {
		PostID              uint    `json:"post_id"`
		CommunityID         uint    `json:"community_id"`
		UserID              uint    `json:"user_id"`
		Content             string  `json:"content"`
		PostsMedia          string  `json:"posts_media"`
		PostsMediaExtention string  `json:"posts_media_extention"`
		Votes               int     `gorm:"not null;default:0" json:"votes"`
		CreatedAt           []uint8 `json:"created_at"`
		EditedAt            []uint8 `json:"edited_at"`
		CommunityName       string  `json:"community_name"`
		CommunityMedia      string  `json:"community_media"`
		UserVote            int     `json:"user_vote"`
		TrendScore          float64 `json:"trend_score"`
		Username            string  `json:"username"`
		ProfileUrl          string  `json:"profile_url"`
	}

	ThreadWithCommunity struct {
		ThreadID       uint    `json:"thread_id"`
		UserID         uint    `json:"user_id"`
		CommunityID    uint    `json:"community_id"`
		Content        string  `json:"content"`
		Votes          int     `gorm:"not null;default:0" json:"votes"`
		CreatedAt      []uint8 `json:"created_at"`
		CommunityName  string  `json:"community_name"`
		CommunityMedia string  `json:"community_media"`
		UserVote       int     `json:"user_vote"`
		TrendScore     float64 `json:"trend_score"`
		Username       string  `json:"username"`
		ProfileUrl     string  `json:"profile_url"`
	}

	MixedItem struct {
		Type    string      `json:"type"`
		Content interface{} `json:"content"`
	}

	Comments struct {
		CommentID       uint    `gorm:"primaryKey" json:"comment_id" db:"comment_id"`
		UserID          uint    `json:"user_id" db:"user_id" form:"user_id"`
		ThreadID        uint    `json:"thread_id,omitempty" db:"thread_id" form:"thread_id"`                         // Nullable if it's not a thread comment
		PostID          uint    `json:"post_id,omitempty" db:"post_id" form:"post_id"`                               // Nullable if it's not a post comment
		ParentCommentID uint    `json:"parent_comment_id,omitempty" db:"parent_comment_id" form:"parent_comment_id"` // Nullable for root comments
		Content         string  `json:"content" db:"content" form:"content"`
		Votes           int     `gorm:"not null;default:0" json:"votes"`
		CreatedAt       []uint8 `json:"created_at" db:"created_at"`
		UpdatedAt       []uint8 `json:"updated_at" db:"updated_at"`
	}

	CommentWithUser struct {
		CommentID  int64   `json:"comment_id"`
		Content    string  `json:"content"`
		CreatedAt  []uint8 `json:"created_at"`
		Username   string  `json:"username"`
		UserID     int64   `json:"user_id"`
		Votes      int     `gorm:"not null;default:0" json:"votes"`
		UserVote   int     `json:"user_vote"`
		ProfileUrl string  `json:"profile_url"`
		ReplyCount int     `json:"reply_count"`
	}

	CommentVote struct {
		UserID    uint     `gorm:"primaryKey;not null" json:"user_id"`
		CommentID uint     `gorm:"primaryKey;not null" json:"comment_id"`
		Vote      int      `json:"vote"`
		VotedAt   []uint8  `gorm:"autoCreateTime" json:"voted_at"`
		User      UserData `gorm:"foreignKey:UserID" json:"user"`
		Comment   Comments `gorm:"foreignKey:CommentID;references:CommentID" json:"comment"`
	}

	MixedListing struct {
		Type         string
		Id           int
		Username     string
		Make         string
		Model        string
		Year         int
		Price        float64
		ImageURL     string
		Description  string
		Etat         string
		Phone        string
		Address      string
		FuelType     string
		Transmission string
		Mileage      int
		Seats        int
		Color        string
		Stock        int
		PartName     string
		Status       string
		Category     string
		CreatedAt    []uint8
		UpdatedAt    []uint8
		UserID       uint
		Email        string
		FirstName    string
		LastName     string
		ProfileUrl   string
	}
)

var (
	Db        *gorm.DB
	JwtSecret []byte
)
