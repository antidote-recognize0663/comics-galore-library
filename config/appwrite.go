package config

import "github.com/antidote-recognize0663/comics-galore-library/config/utils"

type AppwriteConfig struct {
	ApiKey         string
	Endpoint       string
	ProjectID      string
	DatabaseID     string
	RecoveryURL    string
	IPNFunctionID  string
	BucketIDImages string
	//BucketIDCovers          string
	BucketIDAvatars  string
	BucketIDBranding string
	//BucketIDPreviews        string
	BucketIDArchives     string
	CollectionIDUsers    string
	EmailVerificationURL string
	CollectionIDUploads  string
	CollectionIDArchives string
	CollectionIDPayments string
	//CollectionIDFavorites   string
	CollectionIDBlogposts   string
	CollectionIDHeartbeats  string
	CollectionIDStatistics  string
	CollectionIDPostMetrics string
	//CollectionIDCounters    string
	CounterDocumentID string
}

func NewAppwriteConfig() *AppwriteConfig {
	return &AppwriteConfig{
		ApiKey:         utils.GetEnv("APPWRITE_API_KEY", ""),
		Endpoint:       utils.GetEnv("APPWRITE_ENDPOINT", "https://fra.cloud.appwrite.io/v1"),
		ProjectID:      utils.GetEnv("APPWRITE_PROJECT_ID", "6510a59f633f9d57fba2"),
		DatabaseID:     utils.GetEnv("APPWRITE_DATABASE_ID", "6510add9771bcf260b40"),
		RecoveryURL:    utils.GetEnv("APPWRITE_RECOVERY_URL", "https://fra.cloud.appwrite.io/v1/auth/recovery"),
		IPNFunctionID:  utils.GetEnv("APPWRITE_IPN_FUNCTION_ID", "6796205f0003681ff4c2"),
		BucketIDImages: utils.GetEnv("APPWRITE_BUCKET_ID_IMAGES", "68574e890011f6c911c3"),
		//BucketIDCovers:          utils.GetEnv("APPWRITE_BUCKET_ID_COVERS", "651b348b1f922cb08ea3"),
		BucketIDAvatars:  utils.GetEnv("APPWRITE_BUCKET_ID_AVATARS", "651b3476e4b9da11935f"),
		BucketIDBranding: utils.GetEnv("APPWRITE_BUCKET_ID_BRANDING", "653f0892519d5278d6e6"),
		//BucketIDPreviews:        utils.GetEnv("APPWRITE_BUCKET_ID_PREVIEWS", "651b34a9c7484267c929"),
		BucketIDArchives: utils.GetEnv("APPWRITE_BUCKET_ID_ARCHIVES", "651b34b8d02e995f0cda"),
		//CollectionIDUsers:       utils.GetEnv("APPWRITE_COLLECTION_ID_USERS", "65273588acc895a84389"),
		EmailVerificationURL: utils.GetEnv("APPWRITE_EMAIL_VERIFICATION_URL", "https://fra.cloud.appwrite.io/v1/auth/verification"),
		CollectionIDUploads:  utils.GetEnv("APPWRITE_COLLECTION_ID_UPLOADS", "6862e70800281295369c"),
		CollectionIDArchives: utils.GetEnv("APPWRITE_COLLECTION_ID_ARCHIVES", "657793fbd86713ea94ca"),
		CollectionIDPayments: utils.GetEnv("APPWRITE_COLLECTION_ID_PAYMENTS", "67806dd1003557f3794e"),
		//CollectionIDFavorites:   utils.GetEnv("APPWRITE_COLLECTION_ID_FAVORITES", "652738a445707a17efa3"),
		CollectionIDBlogposts:   utils.GetEnv("APPWRITE_COLLECTION_ID_BLOGPOSTS", "6510ae2ee8b7da6d715d"),
		CollectionIDHeartbeats:  utils.GetEnv("APPWRITE_COLLECTION_ID_HEARTBEATS", "6625546a002bd9eb7ffe"),
		CollectionIDStatistics:  utils.GetEnv("APPWRITE_COLLECTION_ID_STATISTICS", "689d116400217e4cd917"),
		CollectionIDPostMetrics: utils.GetEnv("APPWRITE_COLLECTION_ID_POST_METRICS", "67e928fc0018acb88f5b"),
		//CollectionIDCounters:    utils.GetEnv("APPWRITE_COLLECTION_ID_COUNTERS", "689d116400217e4cd917"),
		CounterDocumentID: utils.GetEnv("APPWRITE_COUNTER_DOCUMENT_ID", "689e4a4a0015fd649ac1"),
	}
}
