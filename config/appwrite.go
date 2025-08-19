package config

import "github.com/antidote-recognize0663/comics-galore-library/config/utils"

type AppwriteConfig struct {
	ApiKey                  string
	Endpoint                string
	ProjectID               string
	DatabaseID              string
	RecoveryURL             string
	IPNFunctionID           string
	BucketIDImages          string
	BucketIDCovers          string
	BucketIDAvatars         string
	BucketIDBranding        string
	BucketIDPreviews        string
	BucketIDArchives        string
	CollectionIDUsers       string
	EmailVerificationURL    string
	CollectionIDArchives    string
	CollectionIDPayments    string
	CollectionIDFavorites   string
	CollectionIDBlogposts   string
	CollectionIDHeartbeats  string
	CollectionIDStatistics  string
	CollectionIDPostMetrics string
}

func NewAppwriteConfig() *AppwriteConfig {
	return &AppwriteConfig{
		ApiKey:                  utils.GetEnv("APPWRITE_API_KEY", ""),
		Endpoint:                utils.GetEnv("APPWRITE_ENDPOINT", ""),
		ProjectID:               utils.GetEnv("APPWRITE_PROJECT_ID", ""),
		DatabaseID:              utils.GetEnv("APPWRITE_DATABASE_ID", ""),
		RecoveryURL:             utils.GetEnv("APPWRITE_RECOVERY_URL", ""),
		IPNFunctionID:           utils.GetEnv("APPWRITE_IPN_FUNCTION_ID", ""),
		BucketIDImages:          utils.GetEnv("APPWRITE_BUCKET_ID_IMAGES", ""),
		BucketIDCovers:          utils.GetEnv("APPWRITE_BUCKET_ID_COVERS", ""),
		BucketIDAvatars:         utils.GetEnv("APPWRITE_BUCKET_ID_AVATARS", ""),
		BucketIDBranding:        utils.GetEnv("APPWRITE_BUCKET_ID_BRANDING", ""),
		BucketIDPreviews:        utils.GetEnv("APPWRITE_BUCKET_ID_PREVIEWS", ""),
		BucketIDArchives:        utils.GetEnv("APPWRITE_BUCKET_ID_ARCHIVES", ""),
		CollectionIDUsers:       utils.GetEnv("APPWRITE_COLLECTION_ID_USERS", ""),
		EmailVerificationURL:    utils.GetEnv("APPWRITE_EMAIL_VERIFICATION_URL", ""),
		CollectionIDArchives:    utils.GetEnv("APPWRITE_COLLECTION_ID_ARCHIVES", ""),
		CollectionIDPayments:    utils.GetEnv("APPWRITE_COLLECTION_ID_PAYMENTS", ""),
		CollectionIDFavorites:   utils.GetEnv("APPWRITE_COLLECTION_ID_FAVORITES", ""),
		CollectionIDBlogposts:   utils.GetEnv("APPWRITE_COLLECTION_ID_BLOGPOSTS", ""),
		CollectionIDHeartbeats:  utils.GetEnv("APPWRITE_COLLECTION_ID_HEARTBEATS", ""),
		CollectionIDStatistics:  utils.GetEnv("APPWRITE_COLLECTION_ID_STATISTICS", ""),
		CollectionIDPostMetrics: utils.GetEnv("APPWRITE_COLLECTION_ID_POST_METRICS", ""),
	}
}
