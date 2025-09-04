package statistic

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
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
	database     *databases.Databases
}

type Config struct {
	documentID   string
	databaseID   string
	collectionID string
	database     *databases.Databases
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

func NewCounter(client *client.Client, opts ...Option) Counter {
	cfg := &Config{
		database:     appwrite.NewDatabases(*client),
		databaseID:   "651213bf7705981232aa",
		documentID:   "689e4a4a0015fd649ac1",
		collectionID: "689d116400217e4cd917",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return &counter{
		database:     cfg.database,
		databaseID:   cfg.databaseID,
		documentID:   cfg.documentID,
		collectionID: cfg.collectionID,
	}
}

func NewCounterWithConfig(cfg *config.Config) Counter {
	adminClient := utils.NewAdminClient(
		cfg.Appwrite.ApiKey,
		utils.WithEndpoint(cfg.Appwrite.Endpoint),
		utils.WithProject(cfg.Appwrite.ProjectID))
	return &counter{
		databaseID:   cfg.Appwrite.DatabaseID,
		documentID:   cfg.Appwrite.CounterDocumentID,
		collectionID: cfg.Appwrite.CollectionIDStatistics,
		database:     appwrite.NewDatabases(*adminClient),
	}
}

func (s *counter) Increment(attributeName string, value ...int64) (int64, error) {
	var option databases.IncrementDocumentAttributeOption
	if len(value) == 1 {
		option = s.database.WithIncrementDocumentAttributeValue(float64(value[0]))
	} else {
		return 0, fmt.Errorf("invalid number of arguments only one optional argument is allowed")
	}
	document, err := s.database.IncrementDocumentAttribute(
		s.databaseID,
		s.collectionID,
		s.documentID,
		attributeName,
		option,
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
	var option databases.DecrementDocumentAttributeOption
	if len(value) == 1 {
		option = s.database.WithDecrementDocumentAttributeValue(float64(value[0]))
	} else {
		return 0, fmt.Errorf("invalid number of arguments only one optional argument is allowed")
	}
	document, err := s.database.DecrementDocumentAttribute(
		s.databaseID,
		s.collectionID,
		s.documentID,
		attributeName,
		option,
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
	document, err := s.database.GetDocument(
		s.databaseID,
		s.collectionID,
		s.documentID,
		s.database.WithGetDocumentQueries([]string{query.Select([]string{attributeName})}),
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
