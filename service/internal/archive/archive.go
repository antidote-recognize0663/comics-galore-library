package archive

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/file"
)

type Archive interface {
	GeFileDownload(secret string, fileId string) (*[]byte, error)
	GetFile(secret string, fileId string) (*model.File, error)
	DeleteFile(secret string, fileId string) error
	CreateFile(secret string, fileId string, file file.InputFile) (*model.File, error)
}

type archive struct {
	endpoint  string
	bucketID  string
	projectID string
}

func (a *archive) GetFile(secret string, fileId string) (*model.File, error) {
	client := utils.NewSessionClient(secret, utils.WithProject(a.projectID), utils.WithEndpoint(a.endpoint))
	storage := appwrite.NewStorage(*client)
	getFile, err := storage.GetFile(a.bucketID, fileId)
	if err != nil {
		return nil, fmt.Errorf("failed to get file: %w", err)
	}
	var fileData model.FileData
	if err := getFile.Decode(&fileData); err != nil {
		return nil, fmt.Errorf("failed to decode file: %w", err)
	}
	return &model.File{
		File:     getFile,
		FileData: &fileData,
	}, nil
}

func (a *archive) GeFileDownload(secret string, fileId string) (*[]byte, error) {
	client := utils.NewSessionClient(secret, utils.WithProject(a.projectID), utils.WithEndpoint(a.endpoint))
	storage := appwrite.NewStorage(*client)
	fileDownload, err := storage.GetFileDownload(a.bucketID, fileId)
	if err != nil {
		return nil, fmt.Errorf("failed to get file data: %w", err)
	}
	return fileDownload, nil
}

func (a *archive) DeleteFile(secret string, fileId string) error {
	client := utils.NewSessionClient(secret, utils.WithProject(a.projectID), utils.WithEndpoint(a.endpoint))
	storage := appwrite.NewStorage(*client)
	_, err := storage.DeleteFile(a.bucketID, fileId)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}

func (a *archive) CreateFile(secret string, fileId string, file file.InputFile) (*model.File, error) {
	client := utils.NewSessionClient(secret, utils.WithProject(a.projectID), utils.WithEndpoint(a.endpoint))
	storage := appwrite.NewStorage(*client)
	createFile, err := storage.CreateFile(a.bucketID, fileId, file)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	var fileData model.FileData
	if err := createFile.Decode(&fileData); err != nil {
		return nil, fmt.Errorf("failed to decode file: %w", err)
	}
	return &model.File{
		File:     createFile,
		FileData: &fileData,
	}, nil
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

func WithBucketID(bucketID string) Option {
	return func(config *Config) {
		config.bucketID = bucketID
	}
}

func NewArchive(opts ...Option) Archive {
	_config := &Config{
		endpoint:  "https://fra.cloud.appwrite.io/v1",
		projectID: "6510a59f633f9d57fba2",
		bucketID:  "651b34b8d02e995f0cda",
	}
	for _, opt := range opts {
		opt(_config)
	}

	return &archive{
		endpoint:  _config.endpoint,
		bucketID:  _config.bucketID,
		projectID: _config.projectID,
	}
}

func NewArchiveWithConfig(config *config.Config) Archive {
	return &archive{
		endpoint:  config.Appwrite.Endpoint,
		projectID: config.Appwrite.ProjectID,
		bucketID:  config.Appwrite.BucketIDArchives,
	}
}

type Config struct {
	endpoint  string
	projectID string
	bucketID  string
}

type Option func(*Config)
