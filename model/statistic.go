package model

import "github.com/appwrite/sdk-for-go/models"

type StatisticData struct {
	TotalPosts           int64 `json:"total_posts"`
	TotalComments        int64 `json:"total_comments"`
	TotalArchives        int64 `json:"total_archives"`
	TotalPayments        int64 `json:"total_payments"`
	TotalActiveUsers     int64 `json:"total_active_users"`
	TotalSubscribedUsers int64 `json:"total_subscribed_users"`
	TotalRegisteredUsers int64 `json:"total_registered_users"`
}

type Statistic struct {
	*models.Document
	*StatisticData
}
