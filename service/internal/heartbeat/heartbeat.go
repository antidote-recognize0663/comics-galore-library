package heartbeat

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/id"
)

type Heartbeat interface {
	Upsert(userId, label string) (*model.Heartbeat, error)
}

type heartbeat struct {
	databaseID   string
	collectionID string
	client       *client.Client
}

func (h *heartbeat) Upsert(userId, label string) (*model.Heartbeat, error) {
	database := appwrite.NewDatabases(*h.client)
	upsertDocument, err := database.UpsertDocument(h.databaseID, h.collectionID, id.Unique(), map[string]interface{}{
		"userId": userId,
		"label":  label,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upsert document: %w", err)
	}
	var heartbeatData model.HeartbeatData
	if err := upsertDocument.Decode(&heartbeatData); err != nil {
		return nil, fmt.Errorf("failed to decode updated document documentId %s: %v", upsertDocument.Id, err)
	}
	return &model.Heartbeat{
		Document:      upsertDocument,
		HeartbeatData: &heartbeatData,
	}, nil
}

type Config struct {
	apiKey       string
	endpoint     string
	projectID    string
	databaseID   string
	collectionID string
}

type Option func(config *Config)

func WithApiKey(apiKey string) Option {
	return func(c *Config) {
		c.apiKey = apiKey
	}
}

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

func WithEndpoint(endpoint string) Option {
	return func(c *Config) {
		c.endpoint = endpoint
	}
}

func WithProject(projectID string) Option {
	return func(c *Config) {
		c.projectID = projectID
	}
}

func NewHeartbeatWithConfig(config *config.Config) Heartbeat {
	return &heartbeat{
		databaseID:   config.Appwrite.DatabaseID,
		collectionID: config.Appwrite.CollectionIDHeartbeats,
		client:       utils.NewAdminClient(config.Appwrite.ApiKey, utils.WithEndpoint(config.Appwrite.Endpoint), utils.WithProject(config.Appwrite.ProjectID)),
	}
}

func NewHeartbeat(options ...Option) Heartbeat {
	_config := &Config{
		endpoint:     "https://fra.cloud.appwrite.io/v1",
		projectID:    "6510a59f633f9d57fba2",
		databaseID:   "6510add9771bcf260b40",
		collectionID: "6625546a002bd9eb7ffe",
	}
	for _, option := range options {
		option(_config)
	}
	return &heartbeat{
		databaseID:   _config.databaseID,
		collectionID: _config.collectionID,
		client:       utils.NewAdminClient(_config.apiKey, utils.WithEndpoint(_config.endpoint), utils.WithProject(_config.projectID)),
	}
}
