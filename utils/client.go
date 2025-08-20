package utils

import (
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
)

type endpointParams struct {
	Endpoint string
	Project  string
}

type ClientOption func(*endpointParams)

func WithEndpoint(endpoint string) ClientOption {
	return func(params *endpointParams) {
		params.Endpoint = endpoint
	}
}

func WithProject(project string) ClientOption {
	return func(params *endpointParams) {
		params.Project = project
	}
}

func NewSessionClient(secret string, opts ...ClientOption) *client.Client {
	params := endpointParams{
		Endpoint: "https://fra.cloud.appwrite.io/v1",
		Project:  "6510a59f633f9d57fba2",
	}
	for _, opt := range opts {
		opt(&params)
	}
	sessionClient := appwrite.NewClient(
		appwrite.WithEndpoint(params.Endpoint),
		appwrite.WithProject(params.Project),
		appwrite.WithSession(secret),
	)
	return &sessionClient
}

func NewAdminClient(apiKey string, opts ...ClientOption) *client.Client {
	params := endpointParams{
		Endpoint: "https://fra.cloud.appwrite.io/v1",
		Project:  "6510a59f633f9d57fba2",
	}
	for _, opt := range opts {
		opt(&params)
	}
	adminClient := appwrite.NewClient(
		appwrite.WithEndpoint(params.Endpoint),
		appwrite.WithProject(params.Project),
		appwrite.WithKey(apiKey),
	)
	return &adminClient
}
