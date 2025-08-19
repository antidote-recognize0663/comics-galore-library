package model

import "github.com/appwrite/sdk-for-go/models"

type MetricsData struct {
	Rating        float32 `json:"rating"`
	LikeCount     int64   `json:"likes"`
	PageCount     int64   `json:"pages"`
	FileSize      int64   `json:"file_size"`
	DislikeCount  int64   `json:"dislikes"`
	CommentCount  int64   `json:"comments"`
	DownloadCount int64   `json:"downloads"`
	AuthViewCount int64   `json:"auth_views"`
	AnonViewCount int64   `json:"anon_views"`
	Post          Post    `json:"post"`
}

type Metrics struct {
	*models.Document
	*MetricsData
}
