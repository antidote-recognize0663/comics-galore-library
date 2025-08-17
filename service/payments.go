package services

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/query"
)

type PaymentService interface {
	GetPayments(limit int, offset int, opts ...func([]string) []string) (*model.PaymentList, error)
	UpdatePayment(documentId string, notification map[string]interface{}) (*model.Payment, error)
}

type paymentService struct {
	databaseID   string
	collectionID string
	client       *client.Client
}

func WithQueryPaymentStatusNotEqual(status string) func([]string) []string {
	return func(queries []string) []string {
		if status != "" {
			queries = append(queries, query.NotEqual("payment_status", status))
		}
		return queries
	}
}

func WithQueryPaymentOrderBy(field string, ascending bool) func([]string) []string {
	return func(queries []string) []string {
		if field != "" {
			queries = append(queries, query.OrderDesc(field))
			if ascending {
				queries = append(queries, query.OrderAsc(field))
			}
		}
		return queries
	}
}

func (p paymentService) GetPayments(limit int, offset int, opts ...func([]string) []string) (*model.PaymentList, error) {
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
		return nil, fmt.Errorf("GetPayments error: %v", err)
	}
	var paymentList model.PaymentList
	if err := documentList.Decode(&paymentList); err != nil {
		return nil, fmt.Errorf("GetPayments decode error: %v", err)
	}
	return &paymentList, nil
}

func (p paymentService) UpdatePayment(documentID string, notification map[string]interface{}) (*model.Payment, error) {
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
		return nil, fmt.Errorf("UpdatePayment error for documentID '%s': %v", documentID, err)
	}
	var payment model.Payment
	if err := document.Decode(&payment); err != nil {
		return nil, fmt.Errorf("UpdatePayment decode error for documentID '%s': %v", documentID, err)
	}
	return &payment, nil
}

func NewPaymentService(client *client.Client, opts ...PaymentOption) PaymentService {
	config := &PaymentConfig{
		databaseID:   "6510add9771bcf260b40",
		collectionID: "67806dd1003557f3794e",
	}
	for _, opt := range opts {
		opt(config)
	}
	return &paymentService{
		client:       client,
		databaseID:   config.databaseID,
		collectionID: config.collectionID,
	}
}

type PaymentConfig struct {
	databaseID   string
	collectionID string
}

type PaymentOption func(*PaymentConfig)

func WithPaymentCollectionID(collectionID string) PaymentOption {
	return func(config *PaymentConfig) {
		config.collectionID = collectionID
	}
}

func WithPaymentDatabaseID(databaseID string) PaymentOption {
	return func(config *PaymentConfig) {
		config.databaseID = databaseID
	}
}
