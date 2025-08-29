package statistic

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/query"
)

type Counter interface {
	Increment(attributeName string, value ...int64) (int64, error)
	Decrement(attributeName string, value ...int64) (int64, error)
	GetValue(attributeName string) (int64, error)
}

type counter struct {
	databaseID   string
	documentID   string
	collectionID string
	client       *client.Client
}

type Config struct {
	apiKey       string
	endpoint     string
	projectID    string
	documentID   string
	databaseID   string
	collectionID string
}

type Option func(*Config)

func WithDocumentID(documentID string) Option {
	return func(config *Config) {
		config.documentID = documentID
	}
}

func WithCollectionID(collectionID string) Option {
	return func(config *Config) {
		config.collectionID = collectionID
	}
}

func WithDatabaseID(databaseID string) Option {
	return func(config *Config) {
		config.databaseID = databaseID
	}
}

func WithApiKey(apiKey string) Option {
	return func(config *Config) {
		config.apiKey = apiKey
	}
}

func NewCounter(opts ...Option) Counter {
	cfg := &Config{
		endpoint:     "https://fra.cloud.appwrite.io/v1",
		projectID:    "6510a59f633f9d57fba2",
		databaseID:   "6510add9771bcf260b40",
		documentID:   "689e4a4a0015fd649ac1",
		collectionID: "689d116400217e4cd917",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return &counter{
		databaseID:   cfg.databaseID,
		documentID:   cfg.documentID,
		collectionID: cfg.collectionID,
		client:       utils.NewAdminClient(cfg.apiKey, utils.WithEndpoint(cfg.endpoint), utils.WithProject(cfg.projectID)),
	}
}

func NewCounterWithConfig(cfg *config.Config) Counter {
	return &counter{
		databaseID:   cfg.Appwrite.DatabaseID,
		documentID:   cfg.Appwrite.CounterDocumentID,
		collectionID: cfg.Appwrite.CollectionIDStatistics,
		client:       utils.NewAdminClient(cfg.Appwrite.ApiKey, utils.WithEndpoint(cfg.Appwrite.Endpoint), utils.WithProject(cfg.Appwrite.ProjectID)),
	}
}

func (s *counter) Increment(attributeName string, value ...int64) (int64, error) {
	db := appwrite.NewDatabases(*s.client)
	db.WithGetDocumentQueries([]string{query.Select([]string{attributeName})})
	if len(value) == 1 {
		db.WithIncrementDocumentAttributeValue(float64(value[0]))
	} else {
		return 0, fmt.Errorf("invalid number of arguments only one optional argument is allowed")
	}
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

func (s *counter) Decrement(attributeName string, value ...int64) (int64, error) {
	db := appwrite.NewDatabases(*s.client)
	db.WithGetDocumentQueries([]string{query.Select([]string{attributeName})})
	if len(value) == 1 {
		db.WithDecrementDocumentAttributeValue(float64(value[0]))
	} else {
		return 0, fmt.Errorf("invalid number of arguments only one optional argument is allowed")
	}
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

func (s *counter) GetValue(attributeName string) (int64, error) {
	db := appwrite.NewDatabases(*s.client)
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
