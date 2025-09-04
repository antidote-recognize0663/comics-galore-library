package post

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/id"
)

type Post interface {
	GetByID(secret, documentId string) (*model.Post, error)
	FetchList(secret string, queries []string) (*model.PostList, error)
	Create(secret string, data *model.CreatePost, opts ...DocumentOption) (*model.Post, error)
	Update(secret, documentId string, data model.PostData) (*model.Post, error)
	Delete(secret, documentId string) error
}

type post struct {
	endpoint     string
	projectID    string
	databaseID   string
	collectionID string
}

func (p post) GetByID(secret, documentId string) (*model.Post, error) {
	client := utils.NewSessionClient(secret, utils.WithEndpoint(p.endpoint), utils.WithProject(p.projectID))
	document, err := appwrite.NewDatabases(*client).GetDocument(
		p.databaseID,
		p.collectionID,
		documentId,
	)
	if err != nil {
		return nil, fmt.Errorf("error fetching document documentId %s: %v", documentId, err)
	}
	var postData model.PostData
	if err := document.Decode(&postData); err != nil {
		return nil, fmt.Errorf("failed decoding document documentId %s: %v", documentId, err)
	}
	return &model.Post{
		Document: document,
		PostData: &postData,
	}, nil
}

func (p post) FetchList(secret string, queries []string) (*model.PostList, error) {
	client := utils.NewSessionClient(secret, utils.WithEndpoint(p.endpoint), utils.WithProject(p.projectID))
	databases := appwrite.NewDatabases(*client)
	documents, err := databases.ListDocuments(
		p.databaseID,
		p.collectionID,
		databases.WithListDocumentsQueries(queries))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch documents: %v", err)
	}
	if len(documents.Documents) == 0 {
		return &model.PostList{
			DocumentList: documents,
			Posts:        []model.Post{},
		}, nil
	}
	var postList model.PostList
	if err := documents.Decode(&postList); err != nil {
		return nil, fmt.Errorf("failed to decode documents: %v", err)
	}
	return &postList, nil
}

type DocumentOption func(*documentOptions)

type documentOptions struct {
	documentID string
}

func WithDocumentID(documentID string) DocumentOption {
	return func(o *documentOptions) {
		o.documentID = documentID
	}
}

func (p post) Create(secret string, data *model.CreatePost, opts ...DocumentOption) (*model.Post, error) {
	options := &documentOptions{
		documentID: id.Unique(),
	}

	for _, opt := range opts {
		opt(options)
	}

	client := utils.NewSessionClient(secret, utils.WithEndpoint(p.endpoint), utils.WithProject(p.projectID))

	document, err := appwrite.NewDatabases(*client).CreateDocument(
		p.databaseID,
		p.collectionID,
		options.documentID,
		data,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create document: %v", err)
	}

	var postData model.PostData
	if err := document.Decode(&postData); err != nil {
		return nil, fmt.Errorf("failed to decode created document: %v", err)
	}

	return &model.Post{
		Document: document,
		PostData: &postData,
	}, nil
}

func (p post) Update(secret, documentId string, data model.PostData) (*model.Post, error) {
	client := utils.NewSessionClient(secret, utils.WithEndpoint(p.endpoint), utils.WithProject(p.projectID))
	databases := appwrite.NewDatabases(*client)
	document, err := databases.UpdateDocument(
		p.databaseID,
		p.collectionID,
		documentId,
		databases.WithUpdateDocumentData(data),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update document documentId %s: %v", documentId, err)
	}

	var postData model.PostData
	if err := document.Decode(&postData); err != nil {
		return nil, fmt.Errorf("failed to decode updated document documentId %s: %v", documentId, err)
	}

	return &model.Post{
		Document: document,
		PostData: &postData,
	}, nil
}

func (p post) Delete(secret, documentId string) error {
	client := utils.NewSessionClient(secret, utils.WithEndpoint(p.endpoint), utils.WithProject(p.projectID))
	_, err := appwrite.NewDatabases(*client).DeleteDocument(
		p.databaseID,
		p.collectionID,
		documentId,
	)
	if err != nil {
		return fmt.Errorf("failed to delete document with documentId %s: %v", documentId, err)
	}
	return nil
}

func NewPostWithConfig(config *config.Config) Post {
	return &post{
		endpoint:     config.Appwrite.Endpoint,
		projectID:    config.Appwrite.ProjectID,
		databaseID:   config.Appwrite.DatabaseID,
		collectionID: config.Appwrite.CollectionIDBlogposts,
	}
}

func NewPost(opts ...Option) Post {
	cfg := &Config{
		endpoint:     "https://fra.cloud.appwrite.io/v1",
		projectID:    "6512130e80992b6c3e11",
		databaseID:   "651213bf7705981232aa",
		collectionID: "65121414e190acfc7abd",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return &post{
		endpoint:     cfg.endpoint,
		projectID:    cfg.projectID,
		databaseID:   cfg.databaseID,
		collectionID: cfg.collectionID,
	}
}

func WithProjectID(projectID string) Option {
	return func(c *Config) {
		c.projectID = projectID
	}
}

func WithEndpoint(endpoint string) Option {
	return func(c *Config) {
		c.endpoint = endpoint
	}
}

type Config struct {
	endpoint     string
	projectID    string
	databaseID   string
	collectionID string
}

type Option func(*Config)
