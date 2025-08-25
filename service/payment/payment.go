package payment

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/query"
)

type Payment interface {
	Delete(documentId string) error
	GetById(documentId string) (*model.Payment, error)
	Create(data *model.PaymentData) (*model.Payment, error)
	Update(data *model.NowPaymentsIPN) (*model.Payment, error)
	SaveOrUpdate(data *model.NowPaymentsIPN) (*model.Payment, error)
	WithQueryStatusNotEqual(status string) func([]string) []string
	WithQueryOrderBy(field string, ascending bool) func([]string) []string
	GetList(secret, userID string, limit int, offset int, opts ...func([]string) []string) (*model.PaymentList, error)
}

type payment struct {
	apiKey       string
	endpoint     string
	projectID    string
	databaseID   string
	collectionID string
}

func (p *payment) WithQueryStatusNotEqual(status string) func([]string) []string {
	return func(queries []string) []string {
		if status != "" {
			queries = append(queries, query.NotEqual("payment_status", status))
		}
		return queries
	}
}

func (p *payment) GetList(secret, userID string, limit int, offset int, opts ...func([]string) []string) (*model.PaymentList, error) {
	sessionClient := utils.NewSessionClient(secret, utils.WithProject(p.projectID), utils.WithEndpoint(p.endpoint))
	database := databases.New(*sessionClient)
	queries := []string{
		query.Limit(limit),
		query.Offset(offset),
		query.Equal("user_id", userID),
		query.OrderDesc("$updatedAt"),
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

func (p *payment) SaveOrUpdate(data *model.NowPaymentsIPN) (*model.Payment, error) {
	if data == nil {
		return nil, fmt.Errorf("data is required to update payment")
	}
	//var saveOrUpdateData := &model.PaymentData{}
	if data.OrderID == "" {
		data.OrderID = id.Unique()
	}
	sessionClient := utils.NewAdminClient(p.apiKey, utils.WithProject(p.projectID), utils.WithEndpoint(p.endpoint))
	database := databases.New(*sessionClient)
	document, err := database.UpsertDocument(
		p.databaseID,
		p.collectionID,
		data.OrderID,
		data)
	if err != nil {
		return nil, fmt.Errorf("saveOrUpdate error : %v", err)
	}
	var saveOrUpdatePayment model.Payment
	if err := document.Decode(&saveOrUpdatePayment); err != nil {
		return nil, fmt.Errorf("saveOrUpdate decode error : %v", err)
	}
	return &saveOrUpdatePayment, nil
}

func (p *payment) Update(data *model.NowPaymentsIPN) (*model.Payment, error) {
	if data == nil {
		return nil, fmt.Errorf("data is required to update payment")
	}
	sessionClient := utils.NewAdminClient(p.apiKey, utils.WithProject(p.projectID), utils.WithEndpoint(p.endpoint))
	database := databases.New(*sessionClient)
	document, err := database.UpdateDocument(
		p.databaseID,
		p.collectionID,
		data.OrderID, database.WithUpdateDocumentData(data))
	if err != nil {
		return nil, fmt.Errorf("saveOrUpdate error : %v", err)
	}
	var saveOrUpdatePayment model.Payment
	if err := document.Decode(&saveOrUpdatePayment); err != nil {
		return nil, fmt.Errorf("saveOrUpdate decode error : %v", err)
	}
	return &saveOrUpdatePayment, nil
}

func (p *payment) Create(data *model.PaymentData) (*model.Payment, error) {
	if data == nil {
		return nil, fmt.Errorf("data is required to create payment")
	}
	data.OrderID = id.Unique()
	adminClient := utils.NewAdminClient(p.apiKey, utils.WithProject(p.projectID), utils.WithEndpoint(p.endpoint))
	database := databases.New(*adminClient)
	document, err := database.CreateDocument(
		p.databaseID,
		p.collectionID,
		data.OrderID,
		data,
	)
	if err != nil {
		return nil, err
	}
	var createPayment model.Payment
	if err := document.Decode(&createPayment); err != nil {
		return nil, fmt.Errorf("create decode error for documentID '%s': %v", data.OrderID, err)
	}
	return &createPayment, nil
}

func (p *payment) GetById(documentId string) (*model.Payment, error) {
	adminClient := utils.NewAdminClient(p.apiKey, utils.WithProject(p.projectID), utils.WithEndpoint(p.endpoint))
	database := databases.New(*adminClient)
	documents, err := database.GetDocument(p.databaseID, p.collectionID, documentId)
	if err != nil {
		return nil, fmt.Errorf("failed to get document with fileId '%s': %v", documentId, err)
	}
	var getPayment model.Payment
	if err := documents.Decode(&getPayment); err != nil {
		return nil, fmt.Errorf("failed to decode document with fileId '%s': %v", documentId, err)
	}
	return &getPayment, nil
}

func (p *payment) Delete(documentId string) error {
	adminClient := utils.NewAdminClient(p.apiKey, utils.WithProject(p.projectID), utils.WithEndpoint(p.endpoint))
	database := databases.New(*adminClient)
	_, err := database.DeleteDocument(p.databaseID, p.collectionID, documentId)
	if err != nil {
		return fmt.Errorf("failed to delete document with fileId '%s': %v", documentId, err)
	}
	return nil
}

func NewPayment(opts ...Option) Payment {
	cfg := &Config{
		endpoint:     "https://comics-galore.co/v1",
		projectID:    "6510a59f633f9d57fba2",
		databaseID:   "6510add9771bcf260b40",
		collectionID: "67806dd1003557f3794e",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return &payment{
		apiKey:       cfg.apiKey,
		endpoint:     cfg.endpoint,
		projectID:    cfg.projectID,
		databaseID:   cfg.databaseID,
		collectionID: cfg.collectionID,
	}
}

func NewPaymentWithConfig(cfg *config.Config) Payment {
	return &payment{
		apiKey:    cfg.Appwrite.ApiKey,
		endpoint:  cfg.Appwrite.Endpoint,
		projectID: cfg.Appwrite.ProjectID,
	}
}

func WithApiKey(apiKey string) Option {
	return func(config *Config) {
		config.apiKey = apiKey
	}
}

func WithEndpoint(endpoint string) Option {
	return func(config *Config) {
		config.endpoint = endpoint
	}
}

func WithProjectID(projectID string) Option {
	return func(config *Config) {
		config.projectID = projectID
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

func (p *payment) WithQueryOrderBy(field string, ascending bool) func([]string) []string {
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
	apiKey       string
	endpoint     string
	projectID    string
	databaseID   string
	collectionID string
}

type Option func(*Config)
