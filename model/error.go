package model

// ErrorPageData holds the information to display on the error page
type ErrorPageData struct {
	ErrorCode       int    // e.g., 404, 500
	ErrorTitle      string // e.g., "Page Not Found", "Internal Server Error"
	ErrorMessage    string // User-friendly message
	IsUserError     bool   // True for 4xx errors, false for 5xx to slightly change tone/actions
	Timestamp       string // Optional: For logging/reporting
	RequestID       string // Optional: For tracking/reporting internal errors
	ShowSupportInfo bool   // Whether to show contact support or detailed request info
}
