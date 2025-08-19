package payment

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/query"
)

type Service interface {
	WithQueryOrderBy(field string, ascending bool) func([]string) []string
	WithQueryStatusNotEqual(status string) func([]string) []string
	GetList(limit int, offset int, opts ...func([]string) []string) (*model.PaymentList, error)
	Update(documentId string, notification map[string]interface{}) (*model.Payment, error)
}

type service struct {
	databaseID   string
	collectionID string
	client       *client.Client
}

func (p *service) WithQueryStatusNotEqual(status string) func([]string) []string {
	return func(queries []string) []string {
		if status != "" {
			queries = append(queries, query.NotEqual("payment_status", status))
		}
		return queries
	}
}

func (p *service) GetList(limit int, offset int, opts ...func([]string) []string) (*model.PaymentList, error) {
	database := databases.New(*p.client)
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
	var paymentList model.PaymentList
	if err := documentList.Decode(&paymentList); err != nil {
		return nil, fmt.Errorf("GetList decode error: %v", err)
	}
	return &paymentList, nil
}

func (p *service) Update(documentID string, notification map[string]interface{}) (*model.Payment, error) {
	if documentID == "" {
		return nil, fmt.Errorf("documentID is required to update payment")
	}
	paymentDB := databases.New(*p.client)
	document, err := paymentDB.UpdateDocument(
		p.databaseID,
		p.collectionID,
		documentID,
		paymentDB.WithUpdateDocumentData(notification))
	if err != nil {
		return nil, fmt.Errorf("Update error for documentID '%s': %v", documentID, err)
	}
	var payment model.Payment
	if err := document.Decode(&payment); err != nil {
		return nil, fmt.Errorf("Update decode error for documentID '%s': %v", documentID, err)
	}
	return &payment, nil
}

func NewService(client *client.Client, opts ...Option) Service {
	config := &Config{
		databaseID:   "6510add9771bcf260b40",
		collectionID: "67806dd1003557f3794e",
	}
	for _, opt := range opts {
		opt(config)
	}
	return &service{
		client:       client,
		databaseID:   config.databaseID,
		collectionID: config.collectionID,
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

func (p *service) WithQueryOrderBy(field string, ascending bool) func([]string) []string {
	return func(queries []string) []string {
		if field != "" {
			if ascending {
				queries = append(queries, query.OrderAsc(field))
			} else {
				queries = append(queries, query.OrderDesc(field))
			}
		}
		return queries
	}
}

type Config struct {
	databaseID   string
	collectionID string
}

type Option func(*Config)
