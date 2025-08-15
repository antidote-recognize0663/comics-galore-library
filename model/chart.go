package model

import "github.com/appwrite/sdk-for-go/models"

type ChartData struct {
	Value int64
	Label string
}

type Chart struct {
	*models.Document
	*ChartData
}
