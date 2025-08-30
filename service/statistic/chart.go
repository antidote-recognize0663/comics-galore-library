package statistic

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
)

type Chart interface {
	GetList(limit int, offset int, opts ...func([]string) []string) (*model.ChartList, error)
	AddData(data *model.ChartData, opts ...DocumentOption) (*model.Chart, error)
}

type chart struct {
	endpoint     string
	projectID    string
	databaseID   string
	collectionID string
	database     *databases.Databases
}

func NewChart(client *client.Client, opts ...Option) Chart {
	cfg := &Config{
		database:     appwrite.NewDatabases(*client),
		databaseID:   "6510add9771bcf260b40",
		collectionID: "689d17bb000013a8cf61",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return &chart{
		database:     cfg.database,
		databaseID:   cfg.databaseID,
		collectionID: cfg.collectionID,
	}
}

func NewChartWithConfig(cfg *config.Config) Chart {
	adminClient := utils.NewAdminClient(
		cfg.Appwrite.CollectionIDCharts,
		utils.WithProject(cfg.Appwrite.ProjectID),
		utils.WithEndpoint(cfg.Appwrite.Endpoint))
	return &chart{
		endpoint:     cfg.Appwrite.Endpoint,
		projectID:    cfg.Appwrite.ProjectID,
		databaseID:   cfg.Appwrite.DatabaseID,
		collectionID: cfg.Appwrite.CollectionIDCharts,
		database:     appwrite.NewDatabases(*adminClient),
	}
}

func (p *chart) GetList(limit int, offset int, opts ...func([]string) []string) (*model.ChartList, error) {
	queries := []string{
		query.Limit(limit),
		query.Offset(offset),
	}
	for _, opt := range opts {
		queries = opt(queries)
	}
	documentList, err := p.database.ListDocuments(
		p.databaseID,
		p.collectionID,
		p.database.WithListDocumentsQueries(queries),
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

func (p *chart) AddData(data *model.ChartData, opts ...DocumentOption) (*model.Chart, error) {
	options := &documentOptions{
		documentID: id.Unique(),
	}

	for _, opt := range opts {
		opt(options)
	}

	document, err := p.database.CreateDocument(
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
