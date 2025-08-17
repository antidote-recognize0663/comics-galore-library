package services

import (
	"fmt"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/query"
)

type StatisticService interface {
	Increment(attributeName string) (int64, error)
	Decrement(attributeName string) (int64, error)
	GetValue(attributeName string) (int64, error)
}

type statisticService struct {
	databaseID   string
	documentID   string
	collectionID string
	client       *client.Client
}

type StatisticConfig struct {
	documentID   string
	databaseID   string
	collectionID string
}

type StatisticOption func(*StatisticConfig)

func WithDocumentID(documentID string) StatisticOption {
	return func(config *StatisticConfig) {
		config.documentID = documentID
	}
}

func WithCollectionID(collectionID string) StatisticOption {
	return func(config *StatisticConfig) {
		config.collectionID = collectionID
	}
}

func WithDatabaseID(databaseID string) StatisticOption {
	return func(config *StatisticConfig) {
		config.databaseID = databaseID
	}
}

func NewStatisticService(client *client.Client, opts ...StatisticOption) StatisticService {
	config := &StatisticConfig{
		databaseID:   "6510add9771bcf260b40",
		documentID:   "689e4a4a0015fd649ac1",
		collectionID: "689d116400217e4cd917",
	}
	for _, opt := range opts {
		opt(config)
	}
	return &statisticService{
		client:       client,
		databaseID:   config.databaseID,
		documentID:   config.documentID,
		collectionID: config.collectionID,
	}
}

func (s statisticService) Increment(attributeName string) (int64, error) {
	db := databases.New(*s.client)
	db.WithGetDocumentQueries([]string{query.Select([]string{attributeName})})
	document, err := db.IncrementDocumentAttribute(
		s.databaseID,
		s.collectionID,
		s.documentID,
		attributeName,
	)
	if err != nil {
		return 0, fmt.Errorf("appwrite API error while incrementing '%s': %v", attributeName, err)
	}
	var data map[string]interface{}
	if err := document.Decode(&data); err != nil {
		return 0, fmt.Errorf("failed to decode Appwrite response: %v", err)
	}
	newValue, ok := data[attributeName].(float64) // Appwrite returns numbers from JSON as float64
	if !ok {
		return 0, fmt.Errorf("attribute '%s' not found or not a number in response", attributeName)
	}
	return int64(newValue), nil
}

func (s statisticService) Decrement(attributeName string) (int64, error) {
	db := databases.New(*s.client)
	db.WithGetDocumentQueries([]string{query.Select([]string{attributeName})})
	document, err := db.DecrementDocumentAttribute(
		s.databaseID,
		s.collectionID,
		s.documentID,
		attributeName,
	)
	if err != nil {
		return 0, fmt.Errorf("appwrite API error while decrementing '%s': %v", attributeName, err)
	}
	var data map[string]interface{}
	if err := document.Decode(&data); err != nil {
		return 0, fmt.Errorf("failed to decode Appwrite response: %v", err)
	}
	newValue, ok := data[attributeName].(float64) // Appwrite returns numbers from JSON as float64
	if !ok {
		return 0, fmt.Errorf("attribute '%s' not found or not a number in response", attributeName)
	}
	return int64(newValue), nil
}

func (s statisticService) GetValue(attributeName string) (int64, error) {
	db := databases.New(*s.client)
	db.WithGetDocumentQueries([]string{query.Select([]string{attributeName})})
	document, err := db.GetDocument(
		s.databaseID,
		s.collectionID,
		s.documentID,
	)
	if err != nil {
		return 0, fmt.Errorf("appwrite API error while fetching : %v", err)
	}
	var data map[string]interface{}
	if err := document.Decode(&data); err != nil {
		return 0, fmt.Errorf("failed to decode Appwrite response: %v", err)
	}
	newValue, ok := data[attributeName].(float64)
	if !ok {
		return 0, fmt.Errorf("attribute '%s' not found or not a number in response", attributeName)

	}
	return int64(newValue), nil
}
