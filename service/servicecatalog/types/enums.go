// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessLevelFilterKey string

// Enum values for AccessLevelFilterKey
const (
	AccessLevelFilterKeyAccount AccessLevelFilterKey = "Account"
	AccessLevelFilterKeyRole    AccessLevelFilterKey = "Role"
	AccessLevelFilterKeyUser    AccessLevelFilterKey = "User"
)

type AccessStatus string

// Enum values for AccessStatus
const (
	AccessStatusEnabled      AccessStatus = "ENABLED"
	AccessStatusUnder_change AccessStatus = "UNDER_CHANGE"
	AccessStatusDisabled     AccessStatus = "DISABLED"
)

type ChangeAction string

// Enum values for ChangeAction
const (
	ChangeActionAdd    ChangeAction = "ADD"
	ChangeActionModify ChangeAction = "MODIFY"
	ChangeActionRemove ChangeAction = "REMOVE"
)

type CopyOption string

// Enum values for CopyOption
const (
	CopyOptionCopytags CopyOption = "CopyTags"
)

type CopyProductStatus string

// Enum values for CopyProductStatus
const (
	CopyProductStatusSucceeded   CopyProductStatus = "SUCCEEDED"
	CopyProductStatusIn_progress CopyProductStatus = "IN_PROGRESS"
	CopyProductStatusFailed      CopyProductStatus = "FAILED"
)

type EvaluationType string

// Enum values for EvaluationType
const (
	EvaluationTypeStatic  EvaluationType = "STATIC"
	EvaluationTypeDynamic EvaluationType = "DYNAMIC"
)

type OrganizationNodeType string

// Enum values for OrganizationNodeType
const (
	OrganizationNodeTypeOrganization        OrganizationNodeType = "ORGANIZATION"
	OrganizationNodeTypeOrganizational_unit OrganizationNodeType = "ORGANIZATIONAL_UNIT"
	OrganizationNodeTypeAccount             OrganizationNodeType = "ACCOUNT"
)

type PortfolioShareType string

// Enum values for PortfolioShareType
const (
	PortfolioShareTypeImported           PortfolioShareType = "IMPORTED"
	PortfolioShareTypeAws_servicecatalog PortfolioShareType = "AWS_SERVICECATALOG"
	PortfolioShareTypeAws_organizations  PortfolioShareType = "AWS_ORGANIZATIONS"
)

type PrincipalType string

// Enum values for PrincipalType
const (
	PrincipalTypeIam PrincipalType = "IAM"
)

type ProductSource string

// Enum values for ProductSource
const (
	ProductSourceAccount ProductSource = "ACCOUNT"
)

type ProductType string

// Enum values for ProductType
const (
	ProductTypeCloud_formation_template ProductType = "CLOUD_FORMATION_TEMPLATE"
	ProductTypeMarketplace              ProductType = "MARKETPLACE"
)

type ProductViewFilterBy string

// Enum values for ProductViewFilterBy
const (
	ProductViewFilterByFulltextsearch  ProductViewFilterBy = "FullTextSearch"
	ProductViewFilterByOwner           ProductViewFilterBy = "Owner"
	ProductViewFilterByProducttype     ProductViewFilterBy = "ProductType"
	ProductViewFilterBySourceproductid ProductViewFilterBy = "SourceProductId"
)

type ProductViewSortBy string

// Enum values for ProductViewSortBy
const (
	ProductViewSortByTitle        ProductViewSortBy = "Title"
	ProductViewSortByVersioncount ProductViewSortBy = "VersionCount"
	ProductViewSortByCreationdate ProductViewSortBy = "CreationDate"
)

type PropertyKey string

// Enum values for PropertyKey
const (
	PropertyKeyOwner PropertyKey = "OWNER"
)

type ProvisionedProductPlanStatus string

// Enum values for ProvisionedProductPlanStatus
const (
	ProvisionedProductPlanStatusCreate_in_progress  ProvisionedProductPlanStatus = "CREATE_IN_PROGRESS"
	ProvisionedProductPlanStatusCreate_success      ProvisionedProductPlanStatus = "CREATE_SUCCESS"
	ProvisionedProductPlanStatusCreate_failed       ProvisionedProductPlanStatus = "CREATE_FAILED"
	ProvisionedProductPlanStatusExecute_in_progress ProvisionedProductPlanStatus = "EXECUTE_IN_PROGRESS"
	ProvisionedProductPlanStatusExecute_success     ProvisionedProductPlanStatus = "EXECUTE_SUCCESS"
	ProvisionedProductPlanStatusExecute_failed      ProvisionedProductPlanStatus = "EXECUTE_FAILED"
)

type ProvisionedProductPlanType string

// Enum values for ProvisionedProductPlanType
const (
	ProvisionedProductPlanTypeCloudformation ProvisionedProductPlanType = "CLOUDFORMATION"
)

type ProvisionedProductStatus string

// Enum values for ProvisionedProductStatus
const (
	ProvisionedProductStatusAvailable        ProvisionedProductStatus = "AVAILABLE"
	ProvisionedProductStatusUnder_change     ProvisionedProductStatus = "UNDER_CHANGE"
	ProvisionedProductStatusTainted          ProvisionedProductStatus = "TAINTED"
	ProvisionedProductStatusError            ProvisionedProductStatus = "ERROR"
	ProvisionedProductStatusPlan_in_progress ProvisionedProductStatus = "PLAN_IN_PROGRESS"
)

type ProvisionedProductViewFilterBy string

// Enum values for ProvisionedProductViewFilterBy
const (
	ProvisionedProductViewFilterBySearchquery ProvisionedProductViewFilterBy = "SearchQuery"
)

type ProvisioningArtifactGuidance string

// Enum values for ProvisioningArtifactGuidance
const (
	ProvisioningArtifactGuidanceDefault    ProvisioningArtifactGuidance = "DEFAULT"
	ProvisioningArtifactGuidanceDeprecated ProvisioningArtifactGuidance = "DEPRECATED"
)

type ProvisioningArtifactPropertyName string

// Enum values for ProvisioningArtifactPropertyName
const (
	ProvisioningArtifactPropertyNameId ProvisioningArtifactPropertyName = "Id"
)

type ProvisioningArtifactType string

// Enum values for ProvisioningArtifactType
const (
	ProvisioningArtifactTypeCloud_formation_template ProvisioningArtifactType = "CLOUD_FORMATION_TEMPLATE"
	ProvisioningArtifactTypeMarketplace_ami          ProvisioningArtifactType = "MARKETPLACE_AMI"
	ProvisioningArtifactTypeMarketplace_car          ProvisioningArtifactType = "MARKETPLACE_CAR"
)

type RecordStatus string

// Enum values for RecordStatus
const (
	RecordStatusCreated              RecordStatus = "CREATED"
	RecordStatusIn_progress          RecordStatus = "IN_PROGRESS"
	RecordStatusIn_progress_in_error RecordStatus = "IN_PROGRESS_IN_ERROR"
	RecordStatusSucceeded            RecordStatus = "SUCCEEDED"
	RecordStatusFailed               RecordStatus = "FAILED"
)

type Replacement string

// Enum values for Replacement
const (
	ReplacementTrue        Replacement = "TRUE"
	ReplacementFalse       Replacement = "FALSE"
	ReplacementConditional Replacement = "CONDITIONAL"
)

type RequiresRecreation string

// Enum values for RequiresRecreation
const (
	RequiresRecreationNever         RequiresRecreation = "NEVER"
	RequiresRecreationConditionally RequiresRecreation = "CONDITIONALLY"
	RequiresRecreationAlways        RequiresRecreation = "ALWAYS"
)

type ResourceAttribute string

// Enum values for ResourceAttribute
const (
	ResourceAttributeProperties     ResourceAttribute = "PROPERTIES"
	ResourceAttributeMetadata       ResourceAttribute = "METADATA"
	ResourceAttributeCreationpolicy ResourceAttribute = "CREATIONPOLICY"
	ResourceAttributeUpdatepolicy   ResourceAttribute = "UPDATEPOLICY"
	ResourceAttributeDeletionpolicy ResourceAttribute = "DELETIONPOLICY"
	ResourceAttributeTags           ResourceAttribute = "TAGS"
)

type ServiceActionAssociationErrorCode string

// Enum values for ServiceActionAssociationErrorCode
const (
	ServiceActionAssociationErrorCodeDuplicateresourceexception ServiceActionAssociationErrorCode = "DUPLICATE_RESOURCE"
	ServiceActionAssociationErrorCodeInternalfailure            ServiceActionAssociationErrorCode = "INTERNAL_FAILURE"
	ServiceActionAssociationErrorCodeLimitexceededexception     ServiceActionAssociationErrorCode = "LIMIT_EXCEEDED"
	ServiceActionAssociationErrorCodeResourcenotfoundexception  ServiceActionAssociationErrorCode = "RESOURCE_NOT_FOUND"
	ServiceActionAssociationErrorCodeThrottlingexception        ServiceActionAssociationErrorCode = "THROTTLING"
)

type ServiceActionDefinitionKey string

// Enum values for ServiceActionDefinitionKey
const (
	ServiceActionDefinitionKeyName       ServiceActionDefinitionKey = "Name"
	ServiceActionDefinitionKeyVersion    ServiceActionDefinitionKey = "Version"
	ServiceActionDefinitionKeyAssumerole ServiceActionDefinitionKey = "AssumeRole"
	ServiceActionDefinitionKeyParameters ServiceActionDefinitionKey = "Parameters"
)

type ServiceActionDefinitionType string

// Enum values for ServiceActionDefinitionType
const (
	ServiceActionDefinitionTypeSsmautomation ServiceActionDefinitionType = "SSM_AUTOMATION"
)

type ShareStatus string

// Enum values for ShareStatus
const (
	ShareStatusNot_started           ShareStatus = "NOT_STARTED"
	ShareStatusIn_progress           ShareStatus = "IN_PROGRESS"
	ShareStatusCompleted             ShareStatus = "COMPLETED"
	ShareStatusCompleted_with_errors ShareStatus = "COMPLETED_WITH_ERRORS"
	ShareStatusError                 ShareStatus = "ERROR"
)

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

type StackInstanceStatus string

// Enum values for StackInstanceStatus
const (
	StackInstanceStatusCurrent    StackInstanceStatus = "CURRENT"
	StackInstanceStatusOutdated   StackInstanceStatus = "OUTDATED"
	StackInstanceStatusInoperable StackInstanceStatus = "INOPERABLE"
)

type StackSetOperationType string

// Enum values for StackSetOperationType
const (
	StackSetOperationTypeCreate StackSetOperationType = "CREATE"
	StackSetOperationTypeUpdate StackSetOperationType = "UPDATE"
	StackSetOperationTypeDelete StackSetOperationType = "DELETE"
)

type Status string

// Enum values for Status
const (
	StatusAvailable Status = "AVAILABLE"
	StatusCreating  Status = "CREATING"
	StatusFailed    Status = "FAILED"
)