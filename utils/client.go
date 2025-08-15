package utils

import (
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/open-runtimes/types-for-go/v4/openruntimes"
	"os"
)

func CreateClientWithApiKey(ctx openruntimes.Context) *client.Client {
	adminClient := appwrite.NewClient(
		appwrite.WithProject(os.Getenv("APPWRITE_FUNCTION_PROJECT_ID")),
		appwrite.WithKey(ctx.Req.Headers["x-appwrite-key"]),
	)
	return &adminClient
}
