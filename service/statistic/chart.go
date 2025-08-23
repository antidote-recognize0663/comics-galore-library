package statistic

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/query"
)

type Chart interface {
	GetList(limit int, offset int, opts ...func([]string) []string) (*model.ChartList, error)
	AddData(data model.ChartData, opts ...DocumentOption) (*model.Chart, error)
}

type chart struct {
	endpoint     string
	projectID    string
	databaseID   string
	collectionID string
	client       *client.Client
}

func NewChart(opts ...Option) Chart {
	cfg := &Config{
		endpoint:     "https://fra.cloud.appwrite.io/v1",
		projectID:    "6510a59f633f9d57fba2",
		databaseID:   "6510add9771bcf260b40",
		collectionID: "",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return &chart{
		endpoint:     cfg.endpoint,
		projectID:    cfg.projectID,
		databaseID:   cfg.databaseID,
		collectionID: cfg.collectionID,
		client:       utils.NewAdminClient(cfg.apiKey, utils.WithEndpoint(cfg.endpoint), utils.WithProject(cfg.projectID)),
	}
}

func NewChartWithConfig(cfg *config.Config) Chart {
	return &chart{
		endpoint:   cfg.Appwrite.Endpoint,
		projectID:  cfg.Appwrite.ProjectID,
		databaseID: cfg.Appwrite.DatabaseID,
	}
}

func (p *chart) GetList(limit int, offset int, opts ...func([]string) []string) (*model.ChartList, error) {
	database := appwrite.NewDatabases(*p.client)
	queries := []string{
		query.Limit(limit),
		query.Offset(offset),
	}
	for _, opt := range opts {
		queries = opt(queries)
	}
	documentList, err := database.ListDocuments(
		p.databaseID,
		p.collectionID,
		database.WithListDocumentsQueries(queries),
	)
	if err != nil {
		return nil, fmt.Errorf("GetList error: %v", err)
	}
	var chartList model.ChartList
	if err := documentList.Decode(&chartList); err != nil {
		return nil, fmt.Errorf("GetList decode error: %v", err)
	}
	return &chartList, nil
}

type DocumentOption func(*documentOptions)

type documentOptions struct {
	documentID string
}

func (p *chart) AddData(data model.ChartData, opts ...DocumentOption) (*model.Chart, error) {
	options := &documentOptions{
		documentID: id.Unique(),
	}

	for _, opt := range opts {
		opt(options)
	}

	document, err := appwrite.NewDatabases(*p.client).CreateDocument(
		p.databaseID,
		p.collectionID,
		options.documentID,
		data,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create document: %v", err)
	}

	var chartData model.ChartData
	if err := document.Decode(&chartData); err != nil {
		return nil, fmt.Errorf("failed to decode created document: %v", err)
	}

	return &model.Chart{
		Document:  document,
		ChartData: &chartData,
	}, nil
}
