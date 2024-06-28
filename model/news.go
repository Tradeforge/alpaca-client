package model

import (
	"time"
)

const (
	// PageLimit is the default limit of news items to be returned for given page.
	PageLimit = 50
)

type GetNewsParams struct {
	// List of symbols to retrieve news.
	Symbols string `query:"symbols,omitempty"`
	// (RFC 3339) Start date to start retrieving news articles.
	Since *time.Time `query:"start,omitempty"`
	// Sort articles by updated date. Options: DESC, ASC.
	Sort NewsSortParam `query:"sort,omitempty"`
	// (RFC 3339) End date of period to retrieve news articles.
	Until *time.Time `query:"end,omitempty"`
	// Limit of news items to be returned for given page.
	Limit *int `query:"limit,omitempty"`
	// Boolean indicator to include content for news aritcles (if available).
	IncludeContent *bool `query:"content,omitempty"`
	// Pagination token to continue on the next given page.
	PageToken *string `query:"page_token,omitempty"`
}

type NewsSortParam string

const (
	// NewsSortParamASC is the sort parameter for ascending order.
	NewsSortParamASC NewsSortParam = "ASC"
	// NewsSortParamDESC is the sort parameter for descending order.
	NewsSortParamDESC NewsSortParam = "DESC"
)

type GetNewsResponse struct {
	News          []News  `json:"news"`
	NextPageToken *string `json:"next_page_token,omitempty"`
}

type News struct {
	ID        int64       `json:"id"`
	Author    string      `json:"author"`
	Source    string      `json:"source"`
	Content   *string     `json:"content,omitempty"`
	Symbols   []string    `json:"symbols"`
	Images    []NewsImage `json:"images"`
	Headline  string      `json:"headline"`
	URL       *string     `json:"url,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type NewsImage struct {
	URL  string    `json:"url"`
	Size ImageSize `json:"size"`
}

type ImageSize string

const (
	ImageSizeThumbnail ImageSize = "thumbnail"
	ImageSizeSmall     ImageSize = "small"
	ImageSizeLarge     ImageSize = "large"
)
