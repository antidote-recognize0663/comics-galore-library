package config

import "fmt"

type AppwriteConfig struct {
	ApiKey                     string
	Endpoint                   string
	ProjectID                  string
	DatabaseID                 string
	BucketIDImages             string
	BucketIDAvatars            string
	BucketIDBranding           string
	BucketIDArchives           string
	CollectionIDUsers          string
	RecoveryURL                string
	NowPaymentsNotificationURL string
	EmailVerificationURL       string
	CollectionIDCharts         string
	CollectionIDUploads        string
	CollectionIDArchives       string
	CollectionIDPayments       string
	CollectionIDBlogposts      string
	CollectionIDHeartbeats     string
	CollectionIDStatistics     string
	CollectionIDPostMetrics    string
	CounterDocumentID          string
}

func NewAppwriteConfig() *AppwriteConfig {
	url := GetEnv("APPWRITE_ENDPOINT", "appwrite.comics-galore.co/v1")
	return &AppwriteConfig{
		Endpoint:                   fmt.Sprintf("https://%s", url),
		RecoveryURL:                fmt.Sprintf("https://%s/auth/recovery", url),
		EmailVerificationURL:       fmt.Sprintf("https://%s/auth/verification", url),
		NowPaymentsNotificationURL: fmt.Sprintf("https://functions.%s", url),
		ApiKey:                     GetEnv("APPWRITE_API_KEY", ""),
		ProjectID:                  GetEnv("APPWRITE_PROJECT_ID", "6510a59f633f9d57fba2"),
		DatabaseID:                 GetEnv("APPWRITE_DATABASE_ID", "6510add9771bcf260b40"),
		BucketIDImages:             GetEnv("APPWRITE_BUCKET_ID_IMAGES", "68574e890011f6c911c3"),
		BucketIDAvatars:            GetEnv("APPWRITE_BUCKET_ID_AVATARS", "651b3476e4b9da11935f"),
		BucketIDBranding:           GetEnv("APPWRITE_BUCKET_ID_BRANDING", "653f0892519d5278d6e6"),
		BucketIDArchives:           GetEnv("APPWRITE_BUCKET_ID_ARCHIVES", "651b34b8d02e995f0cda"),
		CounterDocumentID:          GetEnv("APPWRITE_COUNTER_DOCUMENT_ID", "689e4a4a0015fd649ac1"),
		CollectionIDCharts:         GetEnv("APPWRITE_COLLECTION_ID_CHARTS", "689d17bb000013a8cf61"),
		CollectionIDUploads:        GetEnv("APPWRITE_COLLECTION_ID_UPLOADS", "6862e70800281295369c"),
		CollectionIDArchives:       GetEnv("APPWRITE_COLLECTION_ID_ARCHIVES", "657793fbd86713ea94ca"),
		CollectionIDPayments:       GetEnv("APPWRITE_COLLECTION_ID_PAYMENTS", "67806dd1003557f3794e"),
		CollectionIDBlogposts:      GetEnv("APPWRITE_COLLECTION_ID_BLOGPOSTS", "6510ae2ee8b7da6d715d"),
		CollectionIDHeartbeats:     GetEnv("APPWRITE_COLLECTION_ID_HEARTBEATS", "6625546a002bd9eb7ffe"),
		CollectionIDStatistics:     GetEnv("APPWRITE_COLLECTION_ID_STATISTICS", "689d116400217e4cd917"),
		CollectionIDPostMetrics:    GetEnv("APPWRITE_COLLECTION_ID_POST_METRICS", "67e928fc0018acb88f5b"),
	}
}
