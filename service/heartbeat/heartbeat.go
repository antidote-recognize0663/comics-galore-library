package heartbeat

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/query"
	"time"
)

type Heartbeat interface {
	Upsert(userId, label string) (*model.Heartbeat, error)
	GetActiveUsers(duration ...time.Duration) (*model.HeartbeatList, error)
}

type heartbeat struct {
	databaseID   string
	collectionID string
	database     *databases.Databases
}

// GetActiveUsers retrieve the list of active users
// Example of usage:
//
//	Fetches users active in the last 15 minutes
//		fifteenMinutes := 15 * time.Minute
//		activeUsers, err := myHeartbeatService.GetActiveUsers(fifteenMinutes)
//	Fetches users active in the last 72 hours
//		seventyTwoHours := 72 * time.Hour
//		activeUsers, err = myHeartbeatService.GetActiveUsers(seventyTwoHours)
func (h *heartbeat) GetActiveUsers(duration ...time.Duration) (*model.HeartbeatList, error) {
	sinceDuration := 1 * time.Hour
	if len(duration) > 0 {
		sinceDuration = duration[0]
	}
	startTime := time.Now().UTC().Add(-sinceDuration).Format(time.RFC3339)
	queries := h.database.WithListDocumentsQueries(
		[]string{
			query.GreaterThanEqual("$updatedAt", startTime),
		},
	)
	documents, err := h.database.ListDocuments(h.databaseID, h.collectionID, queries)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch documents: %v", err)
	}
	if len(documents.Documents) == 0 {
		return &model.HeartbeatList{
			DocumentList: documents,
			Heartbeats:   []model.Heartbeat{},
		}, nil
	}
	var heartbeatList model.HeartbeatList
	if err := documents.Decode(&heartbeatList); err != nil {
		return nil, fmt.Errorf("failed to decode documents: %v", err)
	}
	return &heartbeatList, nil
}

func (h *heartbeat) Upsert(userId, label string) (*model.Heartbeat, error) {
	upsertData := []interface{}{
		map[string]interface{}{
			"$id":     id.Unique(),
			"label":   label,
			"user_id": userId,
		},
	}
	upsertDocumentList, err := h.database.UpsertDocuments(h.databaseID, h.collectionID, upsertData)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert document: %w", err)
	}
	if len(upsertDocumentList.Documents) == 0 {
		return nil, fmt.Errorf("upsert operation did not return any documents")
	}
	upsertDocument := upsertDocumentList.Documents[0]
	var heartbeatData model.HeartbeatData
	if err := upsertDocument.Decode(&heartbeatData); err != nil {
		return nil, fmt.Errorf("failed to decode upserted document with id %s: %v", upsertDocument.Id, err)
	}
	return &model.Heartbeat{
		Document:      &upsertDocument,
		HeartbeatData: &heartbeatData,
	}, nil
}

type Config struct {
	database     *databases.Databases
	databaseID   string
	collectionID string
}

type Option func(config *Config)

func WithDatabaseID(databaseID string) Option {
	return func(c *Config) {
		c.databaseID = databaseID
	}
}

func WithCollectionID(collectionID string) Option {
	return func(c *Config) {
		c.collectionID = collectionID
	}
}

func NewHeartbeatWithConfig(config *config.Config) Heartbeat {
	adminClient := utils.NewAdminClient(config.Appwrite.ApiKey, utils.WithEndpoint(config.Appwrite.Endpoint), utils.WithProject(config.Appwrite.ProjectID))
	return &heartbeat{
		database:     appwrite.NewDatabases(*adminClient),
		databaseID:   config.Appwrite.DatabaseID,
		collectionID: config.Appwrite.CollectionIDHeartbeats,
	}
}

func NewHeartbeat(client *client.Client, options ...Option) Heartbeat {
	cfg := &Config{
		database:     appwrite.NewDatabases(*client),
		databaseID:   "6510add9771bcf260b40",
		collectionID: "6625546a002bd9eb7ffe",
	}
	for _, option := range options {
		option(cfg)
	}
	return &heartbeat{
		database:     cfg.database,
		databaseID:   cfg.databaseID,
		collectionID: cfg.collectionID,
	}
}
