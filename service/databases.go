package service

import (
	"github.com/appwrite/sdk-for-go/models"
)

type ListOptions struct {
	Queries        []string
	Search         string
	enabledSetters map[string]bool
}

type CreateOptions struct{ Enabled bool }
type UpdateOptions struct{ Enabled bool }
type ListCollectionsOptions struct {
	Queries []string
	Search  string
}
type CreateCollectionOptions struct {
	Permissions      []string
	DocumentSecurity bool
	Enabled          bool
}
type UpdateCollectionOptions struct {
	Permissions      []string
	DocumentSecurity bool
	Enabled          bool
}
type ListAttributesOptions struct{ Queries []string }
type CreateBooleanAttributeOptions struct {
	Default bool
	Array   bool
}
type UpdateBooleanAttributeOptions struct{ NewKey string }
type CreateDatetimeAttributeOptions struct {
	Default string
	Array   bool
}
type UpdateDatetimeAttributeOptions struct{ NewKey string }
type CreateEmailAttributeOptions struct {
	Default string
	Array   bool
}
type UpdateEmailAttributeOptions struct{ NewKey string }
type CreateEnumAttributeOptions struct {
	Default string
	Array   bool
}
type UpdateEnumAttributeOptions struct{ NewKey string }
type CreateFloatAttributeOptions struct {
	Min     float64
	Max     float64
	Default float64
	Array   bool
}
type UpdateFloatAttributeOptions struct {
	Min    float64
	Max    float64
	NewKey string
}
type CreateIntegerAttributeOptions struct {
	Min     int
	Max     int
	Default int
	Array   bool
}
type UpdateIntegerAttributeOptions struct {
	Min    int
	Max    int
	NewKey string
}
type CreateIpAttributeOptions struct {
	Default string
	Array   bool
}
type UpdateIpAttributeOptions struct{ NewKey string }
type CreateRelationshipAttributeOptions struct {
	TwoWay    bool
	Key       string
	TwoWayKey string
	OnDelete  string
}
type UpdateRelationshipAttributeOptions struct {
	OnDelete string
	NewKey   string
}
type CreateStringAttributeOptions struct {
	Default string
	Array   bool
	Encrypt bool
}
type UpdateStringAttributeOptions struct {
	Size   int
	NewKey string
}
type CreateUrlAttributeOptions struct {
	Default string
	Array   bool
}
type UpdateUrlAttributeOptions struct{ NewKey string }
type ListIndexesOptions struct{ Queries []string }
type CreateIndexOptions struct {
	Orders  []string
	Lengths []int
}
type ListDocumentsOptions struct{ Queries []string }
type CreateDocumentOptions struct{ Permissions []string }
type GetDocumentOptions struct{ Queries []string }
type UpsertDocumentOptions struct{ Permissions []string }
type UpdateDocumentOptions struct {
	Data        interface{}
	Permissions []string
}
type UpdateDocumentsOptions struct {
	Data           interface{}
	Queries        []string
	enabledSetters map[string]bool
}
type DeleteDocumentsOptions struct{ Queries []string }

type DecrementDocumentAttributeOptions struct {
	Value float64
	Min   float64
}
type IncrementDocumentAttributeOptions struct {
	Value float64
	Max   float64
}

type ListOption func(*ListOptions)
type CreateOption func(*CreateOptions)
type UpdateOption func(*UpdateOptions)
type ListCollectionsOption func(*ListCollectionsOptions)
type CreateCollectionOption func(*CreateCollectionOptions)
type UpdateCollectionOption func(*UpdateCollectionOptions)
type ListAttributesOption func(*ListAttributesOptions)
type CreateBooleanAttributeOption func(*CreateBooleanAttributeOptions)
type UpdateBooleanAttributeOption func(*UpdateBooleanAttributeOptions)
type CreateDatetimeAttributeOption func(*CreateDatetimeAttributeOptions)
type UpdateDatetimeAttributeOption func(*UpdateDatetimeAttributeOptions)
type CreateEmailAttributeOption func(*CreateEmailAttributeOptions)
type UpdateEmailAttributeOption func(*UpdateEmailAttributeOptions)
type CreateEnumAttributeOption func(*CreateEnumAttributeOptions)
type UpdateEnumAttributeOption func(*UpdateEnumAttributeOptions)
type CreateFloatAttributeOption func(*CreateFloatAttributeOptions)
type UpdateFloatAttributeOption func(*UpdateFloatAttributeOptions)
type CreateIntegerAttributeOption func(*CreateIntegerAttributeOptions)
type UpdateIntegerAttributeOption func(*UpdateIntegerAttributeOptions)
type CreateIpAttributeOption func(*CreateIpAttributeOptions)
type UpdateIpAttributeOption func(*UpdateIpAttributeOptions)
type CreateRelationshipAttributeOption func(*CreateRelationshipAttributeOptions)
type UpdateRelationshipAttributeOption func(*UpdateRelationshipAttributeOptions)
type CreateStringAttributeOption func(*CreateStringAttributeOptions)
type UpdateStringAttributeOption func(*UpdateStringAttributeOptions)
type CreateUrlAttributeOption func(*CreateUrlAttributeOptions)
type UpdateUrlAttributeOption func(*UpdateUrlAttributeOptions)
type ListIndexesOption func(*ListIndexesOptions)
type CreateIndexOption func(*CreateIndexOptions)
type ListDocumentsOption func(*ListDocumentsOptions)
type CreateDocumentOption func(*CreateDocumentOptions)
type GetDocumentOption func(*GetDocumentOptions)
type UpsertDocumentOption func(*UpsertDocumentOptions)
type UpdateDocumentOption func(*UpdateDocumentOptions)
type UpdateDocumentsOption func(*UpdateDocumentsOptions)
type DeleteDocumentsOption func(*DeleteDocumentsOptions)
type DecrementDocumentAttributeOption func(*DecrementDocumentAttributeOptions)
type IncrementDocumentAttributeOption func(*IncrementDocumentAttributeOptions)

type Databases interface {
	List(optionalSetters ...ListOption) (*models.DatabaseList, error)
	Create(databaseId string, name string, optionalSetters ...CreateOption) (*models.Database, error)
	Get(databaseId string) (*models.Database, error)
	Update(databaseId string, name string, optionalSetters ...UpdateOption) (*models.Database, error)
	Delete(databaseId string) (*interface{}, error)

	ListCollections(databaseId string, optionalSetters ...ListCollectionsOption) (*models.CollectionList, error)
	CreateCollection(databaseId string, collectionId string, name string, optionalSetters ...CreateCollectionOption) (*models.Collection, error)
	GetCollection(databaseId string, collectionId string) (*models.Collection, error)
	UpdateCollection(databaseId string, collectionId string, name string, optionalSetters ...UpdateCollectionOption) (*models.Collection, error)
	DeleteCollection(databaseId string, collectionId string) (*interface{}, error)

	ListAttributes(databaseId string, collectionId string, optionalSetters ...ListAttributesOption) (*models.AttributeList, error)
	CreateBooleanAttribute(databaseId string, collectionId string, key string, required bool, optionalSetters ...CreateBooleanAttributeOption) (*models.AttributeBoolean, error)
	UpdateBooleanAttribute(databaseId string, collectionId string, key string, required bool, Default bool, optionalSetters ...UpdateBooleanAttributeOption) (*models.AttributeBoolean, error)
	CreateDatetimeAttribute(databaseId string, collectionId string, key string, required bool, optionalSetters ...CreateDatetimeAttributeOption) (*models.AttributeDatetime, error)
	UpdateDatetimeAttribute(databaseId string, collectionId string, key string, required bool, Default string, optionalSetters ...UpdateDatetimeAttributeOption) (*models.AttributeDatetime, error)
	CreateEmailAttribute(databaseId string, collectionId string, key string, required bool, optionalSetters ...CreateEmailAttributeOption) (*models.AttributeEmail, error)
	UpdateEmailAttribute(databaseId string, collectionId string, key string, required bool, Default string, optionalSetters ...UpdateEmailAttributeOption) (*models.AttributeEmail, error)
	CreateEnumAttribute(databaseId string, collectionId string, key string, elements []string, required bool, optionalSetters ...CreateEnumAttributeOption) (*models.AttributeEnum, error)
	UpdateEnumAttribute(databaseId string, collectionId string, key string, elements []string, required bool, Default string, optionalSetters ...UpdateEnumAttributeOption) (*models.AttributeEnum, error)
	CreateFloatAttribute(databaseId string, collectionId string, key string, required bool, optionalSetters ...CreateFloatAttributeOption) (*models.AttributeFloat, error)
	UpdateFloatAttribute(databaseId string, collectionId string, key string, required bool, Default float64, optionalSetters ...UpdateFloatAttributeOption) (*models.AttributeFloat, error)
	CreateIntegerAttribute(databaseId string, collectionId string, key string, required bool, optionalSetters ...CreateIntegerAttributeOption) (*models.AttributeInteger, error)
	UpdateIntegerAttribute(databaseId string, collectionId string, key string, required bool, Default int, optionalSetters ...UpdateIntegerAttributeOption) (*models.AttributeInteger, error)
	CreateIpAttribute(databaseId string, collectionId string, key string, required bool, optionalSetters ...CreateIpAttributeOption) (*models.AttributeIp, error)
	UpdateIpAttribute(databaseId string, collectionId string, key string, required bool, Default string, optionalSetters ...UpdateIpAttributeOption) (*models.AttributeIp, error)
	CreateRelationshipAttribute(databaseId string, collectionId string, relatedCollectionId string, Type string, optionalSetters ...CreateRelationshipAttributeOption) (*models.AttributeRelationship, error)
	UpdateRelationshipAttribute(databaseId string, collectionId string, key string, optionalSetters ...UpdateRelationshipAttributeOption) (*models.AttributeRelationship, error)
	CreateStringAttribute(databaseId string, collectionId string, key string, size int, required bool, optionalSetters ...CreateStringAttributeOption) (*models.AttributeString, error)
	UpdateStringAttribute(databaseId string, collectionId string, key string, required bool, Default string, optionalSetters ...UpdateStringAttributeOption) (*models.AttributeString, error)
	CreateUrlAttribute(databaseId string, collectionId string, key string, required bool, optionalSetters ...CreateUrlAttributeOption) (*models.AttributeUrl, error)
	UpdateUrlAttribute(databaseId string, collectionId string, key string, required bool, Default string, optionalSetters ...UpdateUrlAttributeOption) (*models.AttributeUrl, error)
	GetAttribute(databaseId string, collectionId string, key string) (*interface{}, error)
	DeleteAttribute(databaseId string, collectionId string, key string) (*interface{}, error)

	ListIndexes(databaseId string, collectionId string, optionalSetters ...ListIndexesOption) (*models.IndexList, error)
	CreateIndex(databaseId string, collectionId string, key string, Type string, attributes []string, optionalSetters ...CreateIndexOption) (*models.Index, error)
	GetIndex(databaseId string, collectionId string, key string) (*models.Index, error)
	DeleteIndex(databaseId string, collectionId string, key string) (*interface{}, error)

	ListDocuments(databaseId string, collectionId string, optionalSetters ...ListDocumentsOption) (*models.DocumentList, error)
	CreateDocument(databaseId string, collectionId string, documentId string, data interface{}, optionalSetters ...CreateDocumentOption) (*models.Document, error)
	CreateDocuments(databaseId string, collectionId string, documents []interface{}) (*models.DocumentList, error)
	UpsertDocuments(databaseId string, collectionId string, documents []interface{}) (*models.DocumentList, error)
	UpdateDocuments(databaseId string, collectionId string, optionalSetters ...UpdateDocumentsOption) (*models.DocumentList, error)
	DeleteDocuments(databaseId string, collectionId string, optionalSetters ...DeleteDocumentsOption) (*models.DocumentList, error)
	GetDocument(databaseId string, collectionId string, documentId string, optionalSetters ...GetDocumentOption) (*models.Document, error)
	UpsertDocument(databaseId string, collectionId string, documentId string, data interface{}, optionalSetters ...UpsertDocumentOption) (*models.Document, error)
	UpdateDocument(databaseId string, collectionId string, documentId string, optionalSetters ...UpdateDocumentOption) (*models.Document, error)
	DeleteDocument(databaseId string, collectionId string, documentId string) (*interface{}, error)
	DecrementDocumentAttribute(databaseId string, collectionId string, documentId string, attribute string, optionalSetters ...DecrementDocumentAttributeOption) (*models.Document, error)
	IncrementDocumentAttribute(databaseId string, collectionId string, documentId string, attribute string, optionalSetters ...IncrementDocumentAttributeOption) (*models.Document, error)

	WithListQueries(v []string) ListOption
	WithListSearch(v string) ListOption
	WithCreateEnabled(v bool) CreateOption
	WithUpdateEnabled(v bool) UpdateOption
	WithListCollectionsQueries(v []string) ListCollectionsOption
	WithListCollectionsSearch(v string) ListCollectionsOption
	WithCreateCollectionPermissions(v []string) CreateCollectionOption
	WithCreateCollectionDocumentSecurity(v bool) CreateCollectionOption
	WithCreateCollectionEnabled(v bool) CreateCollectionOption
	WithUpdateCollectionPermissions(v []string) UpdateCollectionOption
	WithUpdateCollectionDocumentSecurity(v bool) UpdateCollectionOption
	WithUpdateCollectionEnabled(v bool) UpdateCollectionOption
	WithListAttributesQueries(v []string) ListAttributesOption
	WithCreateBooleanAttributeDefault(v bool) CreateBooleanAttributeOption
	WithCreateBooleanAttributeArray(v bool) CreateBooleanAttributeOption
	WithUpdateBooleanAttributeNewKey(v string) UpdateBooleanAttributeOption
	WithCreateDatetimeAttributeDefault(v string) CreateDatetimeAttributeOption
	WithCreateDatetimeAttributeArray(v bool) CreateDatetimeAttributeOption
	WithUpdateDatetimeAttributeNewKey(v string) UpdateDatetimeAttributeOption
	WithCreateEmailAttributeDefault(v string) CreateEmailAttributeOption
	WithCreateEmailAttributeArray(v bool) CreateEmailAttributeOption
	WithUpdateEmailAttributeNewKey(v string) UpdateEmailAttributeOption
	WithCreateEnumAttributeDefault(v string) CreateEnumAttributeOption
	WithCreateEnumAttributeArray(v bool) CreateEnumAttributeOption
	WithUpdateEnumAttributeNewKey(v string) UpdateEnumAttributeOption
	WithCreateFloatAttributeMin(v float64) CreateFloatAttributeOption
	WithCreateFloatAttributeMax(v float64) CreateFloatAttributeOption
	WithCreateFloatAttributeDefault(v float64) CreateFloatAttributeOption
	WithCreateFloatAttributeArray(v bool) CreateFloatAttributeOption
	WithUpdateFloatAttributeMin(v float64) UpdateFloatAttributeOption
	WithUpdateFloatAttributeMax(v float64) UpdateFloatAttributeOption
	WithUpdateFloatAttributeNewKey(v string) UpdateFloatAttributeOption
	WithCreateIntegerAttributeMin(v int) CreateIntegerAttributeOption
	WithCreateIntegerAttributeMax(v int) CreateIntegerAttributeOption
	WithCreateIntegerAttributeDefault(v int) CreateIntegerAttributeOption
	WithCreateIntegerAttributeArray(v bool) CreateIntegerAttributeOption
	WithUpdateIntegerAttributeMin(v int) UpdateIntegerAttributeOption
	WithUpdateIntegerAttributeMax(v int) UpdateIntegerAttributeOption
	WithUpdateIntegerAttributeNewKey(v string) UpdateIntegerAttributeOption
	WithCreateIpAttributeDefault(v string) CreateIpAttributeOption
	WithCreateIpAttributeArray(v bool) CreateIpAttributeOption
	WithUpdateIpAttributeNewKey(v string) UpdateIpAttributeOption
	WithCreateRelationshipAttributeTwoWay(v bool) CreateRelationshipAttributeOption
	WithCreateRelationshipAttributeKey(v string) CreateRelationshipAttributeOption
	WithCreateRelationshipAttributeTwoWayKey(v string) CreateRelationshipAttributeOption
	WithCreateRelationshipAttributeOnDelete(v string) CreateRelationshipAttributeOption
	WithUpdateRelationshipAttributeOnDelete(v string) UpdateRelationshipAttributeOption
	WithUpdateRelationshipAttributeNewKey(v string) UpdateRelationshipAttributeOption
	WithCreateStringAttributeDefault(v string) CreateStringAttributeOption
	WithCreateStringAttributeArray(v bool) CreateStringAttributeOption
	WithCreateStringAttributeEncrypt(v bool) CreateStringAttributeOption
	WithUpdateStringAttributeSize(v int) UpdateStringAttributeOption
	WithUpdateStringAttributeNewKey(v string) UpdateStringAttributeOption
	WithCreateUrlAttributeDefault(v string) CreateUrlAttributeOption
	WithCreateUrlAttributeArray(v bool) CreateUrlAttributeOption
	WithUpdateUrlAttributeNewKey(v string) UpdateUrlAttributeOption
	WithListIndexesQueries(v []string) ListIndexesOption
	WithCreateIndexOrders(v []string) CreateIndexOption
	WithCreateIndexLengths(v []int) CreateIndexOption
	WithListDocumentsQueries(v []string) ListDocumentsOption
	WithCreateDocumentPermissions(v []string) CreateDocumentOption
	WithGetDocumentQueries(v []string) GetDocumentOption
	WithUpsertDocumentPermissions(v []string) UpsertDocumentOption
	WithUpdateDocumentData(v interface{}) UpdateDocumentOption
	WithUpdateDocumentPermissions(v []string) UpdateDocumentOption
	WithDeleteDocumentsQueries(v []string) DeleteDocumentsOption
	WithDecrementDocumentAttributeValue(v float64) DecrementDocumentAttributeOption
	WithDecrementDocumentAttributeMin(v float64) DecrementDocumentAttributeOption
	WithIncrementDocumentAttributeValue(v float64) IncrementDocumentAttributeOption
	WithIncrementDocumentAttributeMax(v float64) IncrementDocumentAttributeOption
	WithUpdateDocumentsData(v interface{}) UpdateDocumentsOption
}
