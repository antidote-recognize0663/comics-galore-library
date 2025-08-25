package config

import "github.com/antidote-recognize0663/comics-galore-library/config/utils"

type AppwriteConfig struct {
	ApiKey                  string
	Endpoint                string
	ProjectID               string
	DatabaseID              string
	BucketIDImages          string
	BucketIDAvatars         string
	BucketIDBranding        string
	BucketIDArchives        string
	CollectionIDUsers       string
	RecoveryURL             string
	NotificationURL         string
	EmailVerificationURL    string
	CollectionIDUploads     string
	CollectionIDArchives    string
	CollectionIDPayments    string
	CollectionIDBlogposts   string
	CollectionIDHeartbeats  string
	CollectionIDStatistics  string
	CollectionIDPostMetrics string
	CounterDocumentID       string
}

func NewAppwriteConfig() *AppwriteConfig {
	return &AppwriteConfig{
		NotificationURL:         utils.GetEnv("APPWRITE_NOTIFICATION_URL", ""),
		ApiKey:                  utils.GetEnv("APPWRITE_API_KEY", ""),
		Endpoint:                utils.GetEnv("APPWRITE_ENDPOINT", "https://fra.cloud.appwrite.io/v1"),
		ProjectID:               utils.GetEnv("APPWRITE_PROJECT_ID", "6510a59f633f9d57fba2"),
		DatabaseID:              utils.GetEnv("APPWRITE_DATABASE_ID", "6510add9771bcf260b40"),
		RecoveryURL:             utils.GetEnv("APPWRITE_RECOVERY_URL", "https://fra.cloud.appwrite.io/v1/auth/recovery"),
		BucketIDImages:          utils.GetEnv("APPWRITE_BUCKET_ID_IMAGES", "68574e890011f6c911c3"),
		BucketIDAvatars:         utils.GetEnv("APPWRITE_BUCKET_ID_AVATARS", "651b3476e4b9da11935f"),
		BucketIDBranding:        utils.GetEnv("APPWRITE_BUCKET_ID_BRANDING", "653f0892519d5278d6e6"),
		BucketIDArchives:        utils.GetEnv("APPWRITE_BUCKET_ID_ARCHIVES", "651b34b8d02e995f0cda"),
		CounterDocumentID:       utils.GetEnv("APPWRITE_COUNTER_DOCUMENT_ID", "689e4a4a0015fd649ac1"),
		EmailVerificationURL:    utils.GetEnv("APPWRITE_EMAIL_VERIFICATION_URL", "https://fra.cloud.appwrite.io/v1/auth/verification"),
		CollectionIDUploads:     utils.GetEnv("APPWRITE_COLLECTION_ID_UPLOADS", "6862e70800281295369c"),
		CollectionIDArchives:    utils.GetEnv("APPWRITE_COLLECTION_ID_ARCHIVES", "657793fbd86713ea94ca"),
		CollectionIDPayments:    utils.GetEnv("APPWRITE_COLLECTION_ID_PAYMENTS", "67806dd1003557f3794e"),
		CollectionIDBlogposts:   utils.GetEnv("APPWRITE_COLLECTION_ID_BLOGPOSTS", "6510ae2ee8b7da6d715d"),
		CollectionIDHeartbeats:  utils.GetEnv("APPWRITE_COLLECTION_ID_HEARTBEATS", "6625546a002bd9eb7ffe"),
		CollectionIDStatistics:  utils.GetEnv("APPWRITE_COLLECTION_ID_STATISTICS", "689d116400217e4cd917"),
		CollectionIDPostMetrics: utils.GetEnv("APPWRITE_COLLECTION_ID_POST_METRICS", "67e928fc0018acb88f5b"),
	}
}
