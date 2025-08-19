package model

import "github.com/appwrite/sdk-for-go/models"

type HeartbeatData struct {
	UserID string   `json:"user_id"`
	Labels []string `json:"labels"`
}

type Heartbeat struct {
	*models.Document
	*HeartbeatData
}

type HeartbeatList struct {
	*models.DocumentList
	Heartbeats []Heartbeat `json:"documents"`
}
