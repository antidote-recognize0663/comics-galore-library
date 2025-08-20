package model

import (
	"fmt"
	"github.com/appwrite/sdk-for-go/models"
)

type Collection string

const (
	CollectionRevenue Collection = "revenue"
	CollectionNewUser Collection = "new-user"
)

type ChartData struct {
	Value      int64      `json:"value"`
	Label      string     `json:"label"`
	Collection Collection `json:"collection"`
}

// Validate validates that the Collection is within the allowed values
func (c Collection) Validate() error {
	switch c {
	case CollectionRevenue, CollectionNewUser:
		return nil
	default:
		return fmt.Errorf("invalid collection: %s", c)
	}
}

func NewChartData(value int64, label, collection string) (*ChartData, error) {
	data := ChartData{
		Value:      value,
		Label:      label,
		Collection: Collection(collection),
	}
	if err := data.Collection.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}
	return &data, nil
}

type Chart struct {
	*models.Document
	*ChartData
}

type ChartList struct {
	*models.DocumentList
	Charts []Chart `json:"documents"`
}
