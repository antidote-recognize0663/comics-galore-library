package payment

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/service/user"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/query"
	"log"
	"time"
)

type Payment interface {
	Delete(documentId string) error
	GetById(documentId string) (*model.Payment, error)
	Create(data *model.PaymentData) (*model.Payment, error)
	Update(data *model.NowPaymentsIPN) (*model.Payment, error)
	SaveOrUpdate(data *model.NowPaymentsIPN) (*model.Payment, error)
	WithQueryStatusNotEqual(status string) func([]string) []string
	WithQueryOrderBy(field string, ascending bool) func([]string) []string
	FetchList(secret, userID string, limit int, offset int, opts ...func([]string) []string) (*model.PaymentList, error)
	ManageSubscribers(limit int, label ...string) (int64, error)
}

type payment struct {
	database     *databases.Databases
	userService  user.User
	endpoint     string
	projectID    string
	databaseID   string
	collectionID string
}

func (p *payment) ManageSubscribers(limit int, label ...string) (int64, error) {
	if limit == 0 || limit < 0 {
		limit = 1000
	}
	if len(label) == 0 {
		label = append(label, "subscriber")
	} else if len(label) > 1 {
		return 0, fmt.Errorf("only one label is allowed")
	}
	var cursor string
	var totalProcessed int64 = 0
	for {
		queries := []string{
			query.Equal("expired", false),
			query.LessThan("expires_at", time.Now().UTC().Format(time.RFC3339)),
			query.Limit(limit),
			query.OrderDesc("$createdAt"),
			query.Select([]string{"$id", "user_id"}),
		}
		if cursor != "" {
			queries = append(queries, query.CursorAfter(cursor))
		}
		response, err := p.database.ListDocuments(p.databaseID, p.collectionID, p.database.WithListDocumentsQueries(queries))
		if err != nil {
			return totalProcessed, fmt.Errorf("could not list expired payments: %w", err)
		}
		if response.Total == 0 {
			break
		}
		var idsToUpdate []models.Document
		for _, doc := range response.Documents {
			expiredPayment, err := model.NewPayment(&doc)
			if err != nil {
				log.Printf("ManageSubscribers: Could not create payment struct: %w", err)
			}
			if _, err := p.userService.RemoveLabel(expiredPayment.UserID, label[0]); err != nil {
				log.Printf("ManageSubscribers: Could not remove label from user %s, skipping: %v", expiredPayment.UserID, err)
				continue
			}
			idsToUpdate = append(idsToUpdate, doc)
		}
		// Update the batch of payments to mark them as expired
		if len(idsToUpdate) > 0 {
			log.Printf("ManageSubscribers: Marking %d payments as expired.", len(idsToUpdate))
			_, err := p.database.UpdateDocuments(p.databaseID, p.collectionID, p.database.WithUpdateDocumentsData(
				Map(idsToUpdate, func(document models.Document) BulkUpdate {
					return BulkUpdate{Id: document.Id, Expired: true}
				})))
			if err != nil {
				log.Printf("Could not update users: " + err.Error())
			}
		}
		totalProcessed += int64(len(response.Documents))
		cursor = response.Documents[len(response.Documents)-1].Id
	}
	return totalProcessed, nil
}

func (p *payment) WithQueryStatusNotEqual(status string) func([]string) []string {
	return func(queries []string) []string {
		if status != "" {
			queries = append(queries, query.NotEqual("payment_status", status))
		}
		return queries
	}
}

func (p *payment) FetchList(secret, userID string, limit int, offset int, opts ...func([]string) []string) (*model.PaymentList, error) {
	database := p.getDatabases(secret)
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
		return nil, fmt.Errorf("FetchList error: %v", err)
	}
	var paymentList model.PaymentList
	if err := documentList.Decode(&paymentList); err != nil {
		return nil, fmt.Errorf("FetchList decode error: %v", err)
	}
	return &paymentList, nil
}

func (p *payment) SaveOrUpdate(data *model.NowPaymentsIPN) (*model.Payment, error) {
	if data == nil {
		return nil, fmt.Errorf("data is required to update payment")
	}
	//var saveOrUpdateData := &model.PaymentData{}
	//TODO: order_id must be unique in appwrite also (index)
	if data.OrderID == "" {
		data.OrderID = id.Unique()
	}
	document, err := p.database.UpsertDocument(
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
	document, err := p.database.UpdateDocument(
		p.databaseID,
		p.collectionID,
		data.OrderID, p.database.WithUpdateDocumentData(data))
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
		return nil, fmt.Errorf("data is required to create a new payment entry")
	}
	data.OrderID = id.Unique()
	document, err := p.database.CreateDocument(
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
	documents, err := p.database.GetDocument(p.databaseID, p.collectionID, documentId)
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
	_, err := p.database.DeleteDocument(p.databaseID, p.collectionID, documentId)
	if err != nil {
		return fmt.Errorf("failed to delete document with fileId '%s': %v", documentId, err)
	}
	return nil
}

func NewPayment(client *client.Client, userService user.User, opts ...Option) Payment {
	if client == nil {
		panic("appwrite client is required")
	}
	cfg := &Config{
		database:     appwrite.NewDatabases(*client),
		userService:  userService,
		endpoint:     "https://comics-galore.co/v1",
		projectID:    "6510a59f633f9d57fba2",
		databaseID:   "6510add9771bcf260b40",
		collectionID: "67806dd1003557f3794e",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return &payment{
		database:     cfg.database,
		userService:  cfg.userService,
		endpoint:     cfg.endpoint,
		projectID:    cfg.projectID,
		databaseID:   cfg.databaseID,
		collectionID: cfg.collectionID,
	}
}

func NewPaymentWithConfig(cfg *config.Config) Payment {
	adminClient := utils.NewAdminClient(cfg.Appwrite.ApiKey, utils.WithProject(cfg.Appwrite.ProjectID), utils.WithEndpoint(cfg.Appwrite.Endpoint))
	return &payment{
		database:     appwrite.NewDatabases(*adminClient),
		userService:  user.NewUser(adminClient),
		endpoint:     cfg.Appwrite.Endpoint,
		projectID:    cfg.Appwrite.ProjectID,
		databaseID:   cfg.Appwrite.DatabaseID,
		collectionID: cfg.Appwrite.CollectionIDPayments,
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
	database     *databases.Databases
	userService  user.User
	endpoint     string
	projectID    string
	databaseID   string
	collectionID string
}

type Option func(*Config)

func (p *payment) getDatabases(secret string) *databases.Databases {
	sessionClient := utils.NewSessionClient(secret, utils.WithProject(p.projectID), utils.WithEndpoint(p.endpoint))
	return appwrite.NewDatabases(*sessionClient)
}

type BulkUpdate struct {
	Id      string `json:"$id"`
	Expired bool   `json:"expired"`
}

func Map(input []models.Document, transform func(document models.Document) BulkUpdate) []BulkUpdate {
	result := make([]BulkUpdate, len(input))
	for i, v := range input {
		result[i] = transform(v)
	}
	return result
}
