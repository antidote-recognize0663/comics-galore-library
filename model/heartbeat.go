package model

import "github.com/appwrite/sdk-for-go/models"

type HeartbeatData struct {
	Label  string `json:"label"`
	UserID string `json:"user_id"`
}

type Heartbeat struct {
	*models.Document
	*HeartbeatData
}

type HeartbeatList struct {
	*models.DocumentList
	Heartbeats []Heartbeat `json:"documents"`
}
