package database

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/query"
	"github.com/appwrite/sdk-for-go/users"
	"github.com/open-runtimes/types-for-go/v4/openruntimes"
	"os"
)

func GetPayments(Context openruntimes.Context, client client.Client, limit int, offset int) (*model.PaymentList, error) {
	Context.Log("Retrieving payments ...")

	database := databases.New(client)
	documentList, err := database.ListDocuments(
		os.Getenv("APPWRITE_DATABASE_ID"),
		os.Getenv("APPWRITE_COLLECTION_ID_PAYMENTS"),
		database.WithListDocumentsQueries([]string{
			query.LessThan("payment_status", "paid"),
		}),
		database.WithListDocumentsQueries(
			[]string{
				query.Limit(limit),
				query.Offset(offset),
			}))
	if err != nil {
		Context.Error(fmt.Sprintf("Error retrieving payments: %v", err))
		return nil, fmt.Errorf("GetPayments error: %v", err)
	}
	if len(documentList.Documents) == 0 {
		Context.Log("GetPayments - No documents found.")
	}

	var paymentList model.PaymentList
	if err := documentList.Decode(&paymentList); err != nil {
		Context.Error(fmt.Sprintf("Error decoding payments: %v", err))
		return nil, fmt.Errorf("GetPayments decode error: %v", err)
	}

	return &paymentList, nil
}

// UpdatePayment updates an existing payment record.
func UpdatePayment(client client.Client, documentId string, paymentStatus *model.PaymentStatus) (*model.Payment, error) {
	if documentId == "" {
		return nil, fmt.Errorf("documentId is required to update payment")
	}

	updateData := map[string]interface{}{
		"payment_status": paymentStatus.PaymentStatus,
	}

	paymentDB := databases.New(client)
	document, err := paymentDB.UpdateDocument(
		os.Getenv("APPWRITE_DATABASE_ID"),
		os.Getenv("APPWRITE_COLLECTION_ID_PAYMENTS"), documentId,
		paymentDB.WithUpdateDocumentData(updateData))

	if err != nil {
		return nil, fmt.Errorf("UpdatePayment error for documentId '%s': %v", documentId, err)
	}

	var payment model.Payment
	if err := document.Decode(&payment); err != nil {
		return nil, fmt.Errorf("UpdatePayment decode error for documentId '%s': %v", documentId, err)
	}

	return &payment, nil
}

func UpdateLabels(client client.Client, userId string) (*models.User, error) {
	userDB := users.New(client)
	fetchedUser, err := userDB.Get(userId)
	if err != nil {
		return nil, fmt.Errorf("UpdateLabels error for userId '%s': %v", userId, err)
	}
	// Check if "subscriber" exists in the slice
	containsSubscriber := false
	for _, label := range fetchedUser.Labels {
		if label == "subscriber" {
			containsSubscriber = true
			break
		}
	}
	// Append "subscriber" if not already present
	if !containsSubscriber {
		user, err := userDB.UpdateLabels(userId, append(fetchedUser.Labels, "subscriber"))

		if err != nil {
			return nil, fmt.Errorf("UpdateLabels error for userId '%s': %v", userId, err)
		}
		return user, nil
	}
	return fetchedUser, nil
}
