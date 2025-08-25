package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GoogleCloudAssetV1p7beta1RelatedAssets represents the GoogleCloudAssetV1p7beta1RelatedAssets schema from the OpenAPI specification
type GoogleCloudAssetV1p7beta1RelatedAssets struct {
	Assets []GoogleCloudAssetV1p7beta1RelatedAsset `json:"assets,omitempty"` // The peer resources of the relationship.
	Relationshipattributes GoogleCloudAssetV1p7beta1RelationshipAttributes `json:"relationshipAttributes,omitempty"` // The relationship attributes which include `type`, `source_resource_type`, `target_resource_type` and `action`.
}

// Resource represents the Resource schema from the OpenAPI specification
type Resource struct {
	Parent string `json:"parent,omitempty"` // The full name of the immediate parent of this resource. See [Resource Names](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more information. For Google Cloud assets, this value is the parent resource defined in the [IAM policy hierarchy](https://cloud.google.com/iam/docs/overview#policy_hierarchy). Example: `//cloudresourcemanager.googleapis.com/projects/my_project_123`
	Resourceurl string `json:"resourceUrl,omitempty"` // The REST URL for accessing the resource. An HTTP `GET` request using this URL returns the resource itself. Example: `https://cloudresourcemanager.googleapis.com/v1/projects/my-project-123` This value is unspecified for resources without a REST API.
	Version string `json:"version,omitempty"` // The API version. Example: `v1`
	Data map[string]interface{} `json:"data,omitempty"` // The content of the resource, in which some sensitive fields are removed and may not be present.
	Discoverydocumenturi string `json:"discoveryDocumentUri,omitempty"` // The URL of the discovery document containing the resource's JSON schema. Example: `https://www.googleapis.com/discovery/v1/apis/compute/v1/rest` This value is unspecified for resources that do not have an API based on a discovery document, such as Cloud Bigtable.
	Discoveryname string `json:"discoveryName,omitempty"` // The JSON schema name listed in the discovery document. Example: `Project` This value is unspecified for resources that do not have an API based on a discovery document, such as Cloud Bigtable.
	Location string `json:"location,omitempty"` // The location of the resource in Google Cloud, such as its zone and region. For more information, see https://cloud.google.com/about/locations/.
}

// QueryAssetsOutputConfig represents the QueryAssetsOutputConfig schema from the OpenAPI specification
type QueryAssetsOutputConfig struct {
	Bigquerydestination GoogleCloudAssetV1QueryAssetsOutputConfigBigQueryDestination `json:"bigqueryDestination,omitempty"` // BigQuery destination.
}

// AuditLogConfig represents the AuditLogConfig schema from the OpenAPI specification
type AuditLogConfig struct {
	Exemptedmembers []string `json:"exemptedMembers,omitempty"` // Specifies the identities that do not cause logging for this type of permission. Follows the same format of Binding.members.
	Logtype string `json:"logType,omitempty"` // The log type that this config enables.
}

// GoogleIdentityAccesscontextmanagerV1DevicePolicy represents the GoogleIdentityAccesscontextmanagerV1DevicePolicy schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1DevicePolicy struct {
	Requirecorpowned bool `json:"requireCorpOwned,omitempty"` // Whether the device needs to be corp owned.
	Requirescreenlock bool `json:"requireScreenlock,omitempty"` // Whether or not screenlock is required for the DevicePolicy to be true. Defaults to `false`.
	Alloweddevicemanagementlevels []string `json:"allowedDeviceManagementLevels,omitempty"` // Allowed device management levels, an empty list allows all management levels.
	Allowedencryptionstatuses []string `json:"allowedEncryptionStatuses,omitempty"` // Allowed encryptions statuses, an empty list allows all statuses.
	Osconstraints []GoogleIdentityAccesscontextmanagerV1OsConstraint `json:"osConstraints,omitempty"` // Allowed OS versions, an empty list allows all types and all versions.
	Requireadminapproval bool `json:"requireAdminApproval,omitempty"` // Whether the device needs to be approved by the customer admin.
}

// ExportAssetsRequest represents the ExportAssetsRequest schema from the OpenAPI specification
type ExportAssetsRequest struct {
	Relationshiptypes []string `json:"relationshipTypes,omitempty"` // A list of relationship types to export, for example: `INSTANCE_TO_INSTANCEGROUP`. This field should only be specified if content_type=RELATIONSHIP. * If specified: it snapshots specified relationships. It returns an error if any of the [relationship_types] doesn't belong to the supported relationship types of the [asset_types] or if any of the [asset_types] doesn't belong to the source types of the [relationship_types]. * Otherwise: it snapshots the supported relationships for all [asset_types] or returns an error if any of the [asset_types] has no relationship support. An unspecified asset types field means all supported asset_types. See [Introduction to Cloud Asset Inventory](https://cloud.google.com/asset-inventory/docs/overview) for all supported asset types and relationship types.
	Assettypes []string `json:"assetTypes,omitempty"` // A list of asset types to take a snapshot for. For example: "compute.googleapis.com/Disk". Regular expressions are also supported. For example: * "compute.googleapis.com.*" snapshots resources whose asset type starts with "compute.googleapis.com". * ".*Instance" snapshots resources whose asset type ends with "Instance". * ".*Instance.*" snapshots resources whose asset type contains "Instance". See [RE2](https://github.com/google/re2/wiki/Syntax) for all supported regular expression syntax. If the regular expression does not match any supported asset type, an INVALID_ARGUMENT error will be returned. If specified, only matching assets will be returned, otherwise, it will snapshot all asset types. See [Introduction to Cloud Asset Inventory](https://cloud.google.com/asset-inventory/docs/overview) for all supported asset types.
	Contenttype string `json:"contentType,omitempty"` // Asset content type. If not specified, no content but the asset name will be returned.
	Outputconfig OutputConfig `json:"outputConfig,omitempty"` // Output configuration for export assets destination.
	Readtime string `json:"readTime,omitempty"` // Timestamp to take an asset snapshot. This can only be set to a timestamp between the current time and the current time minus 35 days (inclusive). If not specified, the current time will be used. Due to delays in resource data collection and indexing, there is a volatile window during which running the same query may get different results.
}

// WindowsUpdatePackage represents the WindowsUpdatePackage schema from the OpenAPI specification
type WindowsUpdatePackage struct {
	Description string `json:"description,omitempty"` // The localized description of the update package.
	Moreinfourls []string `json:"moreInfoUrls,omitempty"` // A collection of URLs that provide more information about the update package.
	Updateid string `json:"updateId,omitempty"` // Gets the identifier of an update package. Stays the same across revisions.
	Title string `json:"title,omitempty"` // The localized title of the update package.
	Kbarticleids []string `json:"kbArticleIds,omitempty"` // A collection of Microsoft Knowledge Base article IDs that are associated with the update package.
	Lastdeploymentchangetime string `json:"lastDeploymentChangeTime,omitempty"` // The last published date of the update, in (UTC) date and time.
	Revisionnumber int `json:"revisionNumber,omitempty"` // The revision number of this update package.
	Supporturl string `json:"supportUrl,omitempty"` // A hyperlink to the language-specific support information for the update.
	Categories []WindowsUpdateCategory `json:"categories,omitempty"` // The categories that are associated with this update package.
}

// SoftwarePackage represents the SoftwarePackage schema from the OpenAPI specification
type SoftwarePackage struct {
	Aptpackage VersionedPackage `json:"aptPackage,omitempty"` // Information related to the a standard versioned package. This includes package info for APT, Yum, Zypper, and Googet package managers.
	Cospackage VersionedPackage `json:"cosPackage,omitempty"` // Information related to the a standard versioned package. This includes package info for APT, Yum, Zypper, and Googet package managers.
	Googetpackage VersionedPackage `json:"googetPackage,omitempty"` // Information related to the a standard versioned package. This includes package info for APT, Yum, Zypper, and Googet package managers.
	Wuapackage WindowsUpdatePackage `json:"wuaPackage,omitempty"` // Details related to a Windows Update package. Field data and names are taken from Windows Update API IUpdate Interface: https://docs.microsoft.com/en-us/windows/win32/api/_wua/ Descriptive fields like title, and description are localized based on the locale of the VM being updated.
	Qfepackage WindowsQuickFixEngineeringPackage `json:"qfePackage,omitempty"` // Information related to a Quick Fix Engineering package. Fields are taken from Windows QuickFixEngineering Interface and match the source names: https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-quickfixengineering
	Windowsapplication WindowsApplication `json:"windowsApplication,omitempty"` // Contains information about a Windows application that is retrieved from the Windows Registry. For more information about these fields, see: https://docs.microsoft.com/en-us/windows/win32/msi/uninstall-registry-key
	Zypperpatch ZypperPatch `json:"zypperPatch,omitempty"` // Details related to a Zypper Patch.
	Yumpackage VersionedPackage `json:"yumPackage,omitempty"` // Information related to the a standard versioned package. This includes package info for APT, Yum, Zypper, and Googet package managers.
	Zypperpackage VersionedPackage `json:"zypperPackage,omitempty"` // Information related to the a standard versioned package. This includes package info for APT, Yum, Zypper, and Googet package managers.
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Code int `json:"code,omitempty"` // The status code, which should be an enum value of google.rpc.Code.
	Details []map[string]interface{} `json:"details,omitempty"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.
	Message string `json:"message,omitempty"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
}

// AnalyzeMoveResponse represents the AnalyzeMoveResponse schema from the OpenAPI specification
type AnalyzeMoveResponse struct {
	Moveanalysis []MoveAnalysis `json:"moveAnalysis,omitempty"` // The list of analyses returned from performing the intended resource move analysis. The analysis is grouped by different Google Cloud services.
}

// CreateFeedRequest represents the CreateFeedRequest schema from the OpenAPI specification
type CreateFeedRequest struct {
	Feedid string `json:"feedId,omitempty"` // Required. This is the client-assigned asset feed identifier and it needs to be unique under a specific parent project/folder/organization.
	Feed Feed `json:"feed,omitempty"` // An asset feed used to export asset updates to a destinations. An asset feed filter controls what updates are exported. The asset feed must be created within a project, organization, or folder. Supported destinations are: Pub/Sub topics.
}

// TimeWindow represents the TimeWindow schema from the OpenAPI specification
type TimeWindow struct {
	Endtime string `json:"endTime,omitempty"` // End time of the time window (inclusive). If not specified, the current timestamp is used instead.
	Starttime string `json:"startTime,omitempty"` // Start time of the time window (exclusive).
}

// Empty represents the Empty schema from the OpenAPI specification
type Empty struct {
}

// Asset represents the Asset schema from the OpenAPI specification
type Asset struct {
	Assettype string `json:"assetType,omitempty"` // The type of the asset. Example: `compute.googleapis.com/Disk` See [Supported asset types](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for more information.
	Iampolicy Policy `json:"iamPolicy,omitempty"` // An Identity and Access Management (IAM) policy, which specifies access controls for Google Cloud resources. A `Policy` is a collection of `bindings`. A `binding` binds one or more `members`, or principals, to a single `role`. Principals can be user accounts, service accounts, Google groups, and domains (such as G Suite). A `role` is a named list of permissions; each `role` can be an IAM predefined role or a user-created custom role. For some types of Google Cloud resources, a `binding` can also specify a `condition`, which is a logical expression that allows access to a resource only if the expression evaluates to `true`. A condition can add constraints based on attributes of the request, the resource, or both. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies). **JSON example:** ``` { "bindings": [ { "role": "roles/resourcemanager.organizationAdmin", "members": [ "user:mike@example.com", "group:admins@example.com", "domain:google.com", "serviceAccount:my-project-id@appspot.gserviceaccount.com" ] }, { "role": "roles/resourcemanager.organizationViewer", "members": [ "user:eve@example.com" ], "condition": { "title": "expirable access", "description": "Does not grant access after Sep 2020", "expression": "request.time < timestamp('2020-10-01T00:00:00.000Z')", } } ], "etag": "BwWWja0YfJA=", "version": 3 } ``` **YAML example:** ``` bindings: - members: - user:mike@example.com - group:admins@example.com - domain:google.com - serviceAccount:my-project-id@appspot.gserviceaccount.com role: roles/resourcemanager.organizationAdmin - members: - user:eve@example.com role: roles/resourcemanager.organizationViewer condition: title: expirable access description: Does not grant access after Sep 2020 expression: request.time < timestamp('2020-10-01T00:00:00.000Z') etag: BwWWja0YfJA= version: 3 ``` For a description of IAM and its features, see the [IAM documentation](https://cloud.google.com/iam/docs/).
	Orgpolicy []GoogleCloudOrgpolicyV1Policy `json:"orgPolicy,omitempty"` // A representation of an [organization policy](https://cloud.google.com/resource-manager/docs/organization-policy/overview#organization_policy). There can be more than one organization policy with different constraints set on a given resource.
	Osinventory Inventory `json:"osInventory,omitempty"` // This API resource represents the available inventory data for a Compute Engine virtual machine (VM) instance at a given point in time. You can use this API resource to determine the inventory data of your VM. For more information, see [Information provided by OS inventory management](https://cloud.google.com/compute/docs/instances/os-inventory-management#data-collected).
	Serviceperimeter GoogleIdentityAccesscontextmanagerV1ServicePerimeter `json:"servicePerimeter,omitempty"` // `ServicePerimeter` describes a set of Google Cloud resources which can freely import and export data amongst themselves, but not export outside of the `ServicePerimeter`. If a request with a source within this `ServicePerimeter` has a target outside of the `ServicePerimeter`, the request will be blocked. Otherwise the request is allowed. There are two types of Service Perimeter - Regular and Bridge. Regular Service Perimeters cannot overlap, a single Google Cloud project or VPC network can only belong to a single regular Service Perimeter. Service Perimeter Bridges can contain only Google Cloud projects as members, a single Google Cloud project may belong to multiple Service Perimeter Bridges.
	Ancestors []string `json:"ancestors,omitempty"` // The ancestry path of an asset in Google Cloud [resource hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy), represented as a list of relative resource names. An ancestry path starts with the closest ancestor in the hierarchy and ends at root. If the asset is a project, folder, or organization, the ancestry path starts from the asset itself. Example: `["projects/123456789", "folders/5432", "organizations/1234"]`
	Name string `json:"name,omitempty"` // The full name of the asset. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1` See [Resource names](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more information.
	Relatedasset RelatedAsset `json:"relatedAsset,omitempty"` // An asset identifier in Google Cloud which contains its name, type and ancestors. An asset can be any resource in the Google Cloud [resource hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy), a resource outside the Google Cloud resource hierarchy (such as Google Kubernetes Engine clusters and objects), or a policy (e.g. IAM policy). See [Supported asset types](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for more information.
	Relatedassets RelatedAssets `json:"relatedAssets,omitempty"` // DEPRECATED. This message only presents for the purpose of backward-compatibility. The server will never populate this message in responses. The detailed related assets with the `relationship_type`.
	Accesspolicy GoogleIdentityAccesscontextmanagerV1AccessPolicy `json:"accessPolicy,omitempty"` // `AccessPolicy` is a container for `AccessLevels` (which define the necessary attributes to use Google Cloud services) and `ServicePerimeters` (which define regions of services able to freely pass data within a perimeter). An access policy is globally visible within an organization, and the restrictions it specifies apply to all projects within an organization.
	Resource Resource `json:"resource,omitempty"` // A representation of a Google Cloud resource.
	Updatetime string `json:"updateTime,omitempty"` // The last update timestamp of an asset. update_time is updated when create/update/delete operation is performed.
	Accesslevel GoogleIdentityAccesscontextmanagerV1AccessLevel `json:"accessLevel,omitempty"` // An `AccessLevel` is a label that can be applied to requests to Google Cloud services, along with a list of requirements necessary for the label to be applied.
}

// GoogleCloudAssetV1p7beta1RelationshipAttributes represents the GoogleCloudAssetV1p7beta1RelationshipAttributes schema from the OpenAPI specification
type GoogleCloudAssetV1p7beta1RelationshipAttributes struct {
	Sourceresourcetype string `json:"sourceResourceType,omitempty"` // The source asset type. Example: `compute.googleapis.com/Instance`
	Targetresourcetype string `json:"targetResourceType,omitempty"` // The target asset type. Example: `compute.googleapis.com/Disk`
	TypeField string `json:"type,omitempty"` // The unique identifier of the relationship type. Example: `INSTANCE_TO_INSTANCEGROUP`
	Action string `json:"action,omitempty"` // The detail of the relationship, e.g. `contains`, `attaches`
}

// SearchAllIamPoliciesResponse represents the SearchAllIamPoliciesResponse schema from the OpenAPI specification
type SearchAllIamPoliciesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Set if there are more results than those appearing in this response; to get the next set of results, call this method again, using this value as the `page_token`.
	Results []IamPolicySearchResult `json:"results,omitempty"` // A list of IAM policies that match the search query. Related information such as the associated resource is returned along with the policy.
}

// SearchAllResourcesResponse represents the SearchAllResourcesResponse schema from the OpenAPI specification
type SearchAllResourcesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // If there are more results than those appearing in this response, then `next_page_token` is included. To get the next set of results, call this method again using the value of `next_page_token` as `page_token`.
	Results []ResourceSearchResult `json:"results,omitempty"` // A list of Resources that match the search query. It contains the resource standard metadata information.
}

// IamPolicyAnalysisOutputConfig represents the IamPolicyAnalysisOutputConfig schema from the OpenAPI specification
type IamPolicyAnalysisOutputConfig struct {
	Bigquerydestination GoogleCloudAssetV1BigQueryDestination `json:"bigqueryDestination,omitempty"` // A BigQuery destination.
	Gcsdestination GoogleCloudAssetV1GcsDestination `json:"gcsDestination,omitempty"` // A Cloud Storage location.
}

// GoogleIdentityAccesscontextmanagerV1BasicLevel represents the GoogleIdentityAccesscontextmanagerV1BasicLevel schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1BasicLevel struct {
	Conditions []GoogleIdentityAccesscontextmanagerV1Condition `json:"conditions,omitempty"` // Required. A list of requirements for the `AccessLevel` to be granted.
	Combiningfunction string `json:"combiningFunction,omitempty"` // How the `conditions` list should be combined to determine if a request is granted this `AccessLevel`. If AND is used, each `Condition` in `conditions` must be satisfied for the `AccessLevel` to be applied. If OR is used, at least one `Condition` in `conditions` must be satisfied for the `AccessLevel` to be applied. Default behavior is AND.
}

// IamPolicyAnalysisQuery represents the IamPolicyAnalysisQuery schema from the OpenAPI specification
type IamPolicyAnalysisQuery struct {
	Accessselector AccessSelector `json:"accessSelector,omitempty"` // Specifies roles and/or permissions to analyze, to determine both the identities possessing them and the resources they control. If multiple values are specified, results will include roles or permissions matching any of them. The total number of roles and permissions should be equal or less than 10.
	Conditioncontext ConditionContext `json:"conditionContext,omitempty"` // The IAM conditions context.
	Identityselector IdentitySelector `json:"identitySelector,omitempty"` // Specifies an identity for which to determine resource access, based on roles assigned either directly to them or to the groups they belong to, directly or indirectly.
	Options Options `json:"options,omitempty"` // Contains query options.
	Resourceselector ResourceSelector `json:"resourceSelector,omitempty"` // Specifies the resource to analyze for access policies, which may be set directly on the resource, or on ancestors such as organizations, folders or projects.
	Scope string `json:"scope,omitempty"` // Required. The relative name of the root asset. Only resources and IAM policies within the scope will be analyzed. This can only be an organization number (such as "organizations/123"), a folder number (such as "folders/123"), a project ID (such as "projects/my-project-id"), or a project number (such as "projects/12345"). To know how to get organization ID, visit [here ](https://cloud.google.com/resource-manager/docs/creating-managing-organization#retrieving_your_organization_id). To know how to get folder or project ID, visit [here ](https://cloud.google.com/resource-manager/docs/creating-managing-folders#viewing_or_listing_folders_and_projects).
}

// GoogleCloudAssetV1ListConstraint represents the GoogleCloudAssetV1ListConstraint schema from the OpenAPI specification
type GoogleCloudAssetV1ListConstraint struct {
	Supportsunder bool `json:"supportsUnder,omitempty"` // Indicates whether subtrees of Cloud Resource Manager resource hierarchy can be used in `Policy.allowed_values` and `Policy.denied_values`. For example, `"under:folders/123"` would match any resource under the 'folders/123' folder.
	Supportsin bool `json:"supportsIn,omitempty"` // Indicates whether values grouped into categories can be used in `Policy.allowed_values` and `Policy.denied_values`. For example, `"in:Python"` would match any value in the 'Python' group.
}

// Operation represents the Operation schema from the OpenAPI specification
type Operation struct {
	Name string `json:"name,omitempty"` // The server-assigned name, which is only unique within the same service that originally returns it. If you use the default HTTP mapping, the `name` should be a resource name ending with `operations/{unique_id}`.
	Response map[string]interface{} `json:"response,omitempty"` // The normal, successful response of the operation. If the original method returns no data on success, such as `Delete`, the response is `google.protobuf.Empty`. If the original method is standard `Get`/`Create`/`Update`, the response should be the resource. For other methods, the response should have the type `XxxResponse`, where `Xxx` is the original method name. For example, if the original method name is `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.
	Done bool `json:"done,omitempty"` // If the value is `false`, it means the operation is still in progress. If `true`, the operation is completed, and either `error` or `response` is available.
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Service-specific metadata associated with the operation. It typically contains progress information and common metadata such as create time. Some services might not provide such metadata. Any method that returns a long-running operation should document the metadata type, if any.
}

// GoogleIdentityAccesscontextmanagerV1CustomLevel represents the GoogleIdentityAccesscontextmanagerV1CustomLevel schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1CustomLevel struct {
	Expr Expr `json:"expr,omitempty"` // Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
}

// GoogleIdentityAccesscontextmanagerV1EgressFrom represents the GoogleIdentityAccesscontextmanagerV1EgressFrom schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1EgressFrom struct {
	Identitytype string `json:"identityType,omitempty"` // Specifies the type of identities that are allowed access to outside the perimeter. If left unspecified, then members of `identities` field will be allowed access.
	Sourcerestriction string `json:"sourceRestriction,omitempty"` // Whether to enforce traffic restrictions based on `sources` field. If the `sources` fields is non-empty, then this field must be set to `SOURCE_RESTRICTION_ENABLED`.
	Sources []GoogleIdentityAccesscontextmanagerV1EgressSource `json:"sources,omitempty"` // Sources that this EgressPolicy authorizes access from. If this field is not empty, then `source_restriction` must be set to `SOURCE_RESTRICTION_ENABLED`.
	Identities []string `json:"identities,omitempty"` // A list of identities that are allowed access through this [EgressPolicy], in the format of `user:{email_id}` or `serviceAccount:{email_id}`.
}

// GoogleCloudAssetV1Edge represents the GoogleCloudAssetV1Edge schema from the OpenAPI specification
type GoogleCloudAssetV1Edge struct {
	Sourcenode string `json:"sourceNode,omitempty"` // The source node of the edge. For example, it could be a full resource name for a resource node or an email of an identity.
	Targetnode string `json:"targetNode,omitempty"` // The target node of the edge. For example, it could be a full resource name for a resource node or an email of an identity.
}

// GoogleCloudAssetV1Identity represents the GoogleCloudAssetV1Identity schema from the OpenAPI specification
type GoogleCloudAssetV1Identity struct {
	Analysisstate IamPolicyAnalysisState `json:"analysisState,omitempty"` // Represents the detailed state of an entity under analysis, such as a resource, an identity or an access.
	Name string `json:"name,omitempty"` // The identity of members, formatted as appear in an [IAM policy binding](https://cloud.google.com/iam/reference/rest/v1/Binding). For example, they might be formatted like the following: - user:foo@google.com - group:group1@google.com - serviceAccount:s1@prj1.iam.gserviceaccount.com - projectOwner:some_project_id - domain:google.com - allUsers
}

// GoogleIdentityAccesscontextmanagerV1ApiOperation represents the GoogleIdentityAccesscontextmanagerV1ApiOperation schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1ApiOperation struct {
	Methodselectors []GoogleIdentityAccesscontextmanagerV1MethodSelector `json:"methodSelectors,omitempty"` // API methods or permissions to allow. Method or permission must belong to the service specified by `service_name` field. A single MethodSelector entry with `*` specified for the `method` field will allow all methods AND permissions for the service specified in `service_name`.
	Servicename string `json:"serviceName,omitempty"` // The name of the API whose methods or permissions the IngressPolicy or EgressPolicy want to allow. A single ApiOperation with `service_name` field set to `*` will allow all methods AND permissions for all services.
}

// UpdateFeedRequest represents the UpdateFeedRequest schema from the OpenAPI specification
type UpdateFeedRequest struct {
	Feed Feed `json:"feed,omitempty"` // An asset feed used to export asset updates to a destinations. An asset feed filter controls what updates are exported. The asset feed must be created within a project, organization, or folder. Supported destinations are: Pub/Sub topics.
	Updatemask string `json:"updateMask,omitempty"` // Required. Only updates the `feed` fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server.
}

// RelationshipAttributes represents the RelationshipAttributes schema from the OpenAPI specification
type RelationshipAttributes struct {
	Action string `json:"action,omitempty"` // The detail of the relationship, e.g. `contains`, `attaches`
	Sourceresourcetype string `json:"sourceResourceType,omitempty"` // The source asset type. Example: `compute.googleapis.com/Instance`
	Targetresourcetype string `json:"targetResourceType,omitempty"` // The target asset type. Example: `compute.googleapis.com/Disk`
	TypeField string `json:"type,omitempty"` // The unique identifier of the relationship type. Example: `INSTANCE_TO_INSTANCEGROUP`
}

// GoogleIdentityAccesscontextmanagerV1VpcNetworkSource represents the GoogleIdentityAccesscontextmanagerV1VpcNetworkSource schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1VpcNetworkSource struct {
	Vpcsubnetwork GoogleIdentityAccesscontextmanagerV1VpcSubNetwork `json:"vpcSubnetwork,omitempty"` // Sub-segment ranges inside of a VPC Network.
}

// GoogleIdentityAccesscontextmanagerV1EgressPolicy represents the GoogleIdentityAccesscontextmanagerV1EgressPolicy schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1EgressPolicy struct {
	Egressfrom GoogleIdentityAccesscontextmanagerV1EgressFrom `json:"egressFrom,omitempty"` // Defines the conditions under which an EgressPolicy matches a request. Conditions based on information about the source of the request. Note that if the destination of the request is also protected by a ServicePerimeter, then that ServicePerimeter must have an IngressPolicy which allows access in order for this request to succeed.
	Egressto GoogleIdentityAccesscontextmanagerV1EgressTo `json:"egressTo,omitempty"` // Defines the conditions under which an EgressPolicy matches a request. Conditions are based on information about the ApiOperation intended to be performed on the `resources` specified. Note that if the destination of the request is also protected by a ServicePerimeter, then that ServicePerimeter must have an IngressPolicy which allows access in order for this request to succeed. The request must match `operations` AND `resources` fields in order to be allowed egress out of the perimeter.
}

// GoogleIdentityAccesscontextmanagerV1EgressTo represents the GoogleIdentityAccesscontextmanagerV1EgressTo schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1EgressTo struct {
	Externalresources []string `json:"externalResources,omitempty"` // A list of external resources that are allowed to be accessed. Only AWS and Azure resources are supported. For Amazon S3, the supported format is s3://BUCKET_NAME. For Azure Storage, the supported format is azure://myaccount.blob.core.windows.net/CONTAINER_NAME. A request matches if it contains an external resource in this list (Example: s3://bucket/path). Currently '*' is not allowed.
	Operations []GoogleIdentityAccesscontextmanagerV1ApiOperation `json:"operations,omitempty"` // A list of ApiOperations allowed to be performed by the sources specified in the corresponding EgressFrom. A request matches if it uses an operation/service in this list.
	Resources []string `json:"resources,omitempty"` // A list of resources, currently only projects in the form `projects/`, that are allowed to be accessed by sources defined in the corresponding EgressFrom. A request matches if it contains a resource in this list. If `*` is specified for `resources`, then this EgressTo rule will authorize access to all resources outside the perimeter.
}

// GoogleCloudAssetV1AnalyzeOrgPolicyGovernedAssetsResponseGovernedAsset represents the GoogleCloudAssetV1AnalyzeOrgPolicyGovernedAssetsResponseGovernedAsset schema from the OpenAPI specification
type GoogleCloudAssetV1AnalyzeOrgPolicyGovernedAssetsResponseGovernedAsset struct {
	Policybundle []AnalyzerOrgPolicy `json:"policyBundle,omitempty"` // The ordered list of all organization policies from the AnalyzeOrgPoliciesResponse.OrgPolicyResult.consolidated_policy.attached_resource to the scope specified in the request. If the constraint is defined with default policy, it will also appear in the list.
	Consolidatedpolicy AnalyzerOrgPolicy `json:"consolidatedPolicy,omitempty"` // This organization policy message is a modified version of the one defined in the Organization Policy system. This message contains several fields defined in the original organization policy with some new fields for analysis purpose.
	Governediampolicy GoogleCloudAssetV1AnalyzeOrgPolicyGovernedAssetsResponseGovernedIamPolicy `json:"governedIamPolicy,omitempty"` // The IAM policies governed by the organization policies of the AnalyzeOrgPolicyGovernedAssetsRequest.constraint.
	Governedresource GoogleCloudAssetV1AnalyzeOrgPolicyGovernedAssetsResponseGovernedResource `json:"governedResource,omitempty"` // The Google Cloud resources governed by the organization policies of the AnalyzeOrgPolicyGovernedAssetsRequest.constraint.
}

// GoogleIdentityAccesscontextmanagerV1MethodSelector represents the GoogleIdentityAccesscontextmanagerV1MethodSelector schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1MethodSelector struct {
	Method string `json:"method,omitempty"` // A valid method name for the corresponding `service_name` in ApiOperation. If `*` is used as the value for the `method`, then ALL methods and permissions are allowed.
	Permission string `json:"permission,omitempty"` // A valid Cloud IAM permission for the corresponding `service_name` in ApiOperation.
}

// Date represents the Date schema from the OpenAPI specification
type Date struct {
	Year int `json:"year,omitempty"` // Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Day int `json:"day,omitempty"` // Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Month int `json:"month,omitempty"` // Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
}

// ConditionEvaluation represents the ConditionEvaluation schema from the OpenAPI specification
type ConditionEvaluation struct {
	Evaluationvalue string `json:"evaluationValue,omitempty"` // The evaluation result.
}

// OrgPolicyResult represents the OrgPolicyResult schema from the OpenAPI specification
type OrgPolicyResult struct {
	Consolidatedpolicy AnalyzerOrgPolicy `json:"consolidatedPolicy,omitempty"` // This organization policy message is a modified version of the one defined in the Organization Policy system. This message contains several fields defined in the original organization policy with some new fields for analysis purpose.
	Folders []string `json:"folders,omitempty"` // The folder(s) that this consolidated policy belongs to, in the format of folders/{FOLDER_NUMBER}. This field is available when the consolidated policy belongs (directly or cascadingly) to one or more folders.
	Organization string `json:"organization,omitempty"` // The organization that this consolidated policy belongs to, in the format of organizations/{ORGANIZATION_NUMBER}. This field is available when the consolidated policy belongs (directly or cascadingly) to an organization.
	Policybundle []AnalyzerOrgPolicy `json:"policyBundle,omitempty"` // The ordered list of all organization policies from the AnalyzeOrgPoliciesResponse.OrgPolicyResult.consolidated_policy.attached_resource. to the scope specified in the request. If the constraint is defined with default policy, it will also appear in the list.
	Project string `json:"project,omitempty"` // The project that this consolidated policy belongs to, in the format of projects/{PROJECT_NUMBER}. This field is available when the consolidated policy belongs to a project.
}

// AnalyzeIamPolicyLongrunningMetadata represents the AnalyzeIamPolicyLongrunningMetadata schema from the OpenAPI specification
type AnalyzeIamPolicyLongrunningMetadata struct {
	Createtime string `json:"createTime,omitempty"` // Output only. The time the operation was created.
}

// VersionedPackage represents the VersionedPackage schema from the OpenAPI specification
type VersionedPackage struct {
	Version string `json:"version,omitempty"` // The version of the package.
	Architecture string `json:"architecture,omitempty"` // The system architecture this package is intended for.
	Packagename string `json:"packageName,omitempty"` // The name of the package.
}

// EffectiveTagDetails represents the EffectiveTagDetails schema from the OpenAPI specification
type EffectiveTagDetails struct {
	Attachedresource string `json:"attachedResource,omitempty"` // The [full resource name](https://cloud.google.com/asset-inventory/docs/resource-name-format) of the ancestor from which an effective_tag is inherited, according to [tag inheritance](https://cloud.google.com/resource-manager/docs/tags/tags-overview#inheritance).
	Effectivetags []Tag `json:"effectiveTags,omitempty"` // The effective tags inherited from the attached_resource. Note that tags with the same key but different values may attach to resources at a different hierarchy levels. The lower hierarchy tag value will overwrite the higher hierarchy tag value of the same tag key. In this case, the tag value at the higher hierarchy level will be removed. For more information, see [tag inheritance](https://cloud.google.com/resource-manager/docs/tags/tags-overview#inheritance).
}

// MoveAnalysisResult represents the MoveAnalysisResult schema from the OpenAPI specification
type MoveAnalysisResult struct {
	Blockers []MoveImpact `json:"blockers,omitempty"` // Blocking information that would prevent the target resource from moving to the specified destination at runtime.
	Warnings []MoveImpact `json:"warnings,omitempty"` // Warning information indicating that moving the target resource to the specified destination might be unsafe. This can include important policy information and configuration changes, but will not block moves at runtime.
}

// GoogleCloudAssetV1AnalyzeOrgPolicyGovernedAssetsResponseGovernedIamPolicy represents the GoogleCloudAssetV1AnalyzeOrgPolicyGovernedAssetsResponseGovernedIamPolicy schema from the OpenAPI specification
type GoogleCloudAssetV1AnalyzeOrgPolicyGovernedAssetsResponseGovernedIamPolicy struct {
	Attachedresource string `json:"attachedResource,omitempty"` // The full resource name of the resource on which this IAM policy is set. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. See [Cloud Asset Inventory Resource Name Format](https://cloud.google.com/asset-inventory/docs/resource-name-format) for more information.
	Folders []string `json:"folders,omitempty"` // The folder(s) that this IAM policy belongs to, in the format of folders/{FOLDER_NUMBER}. This field is available when the IAM policy belongs (directly or cascadingly) to one or more folders.
	Organization string `json:"organization,omitempty"` // The organization that this IAM policy belongs to, in the format of organizations/{ORGANIZATION_NUMBER}. This field is available when the IAM policy belongs (directly or cascadingly) to an organization.
	Policy Policy `json:"policy,omitempty"` // An Identity and Access Management (IAM) policy, which specifies access controls for Google Cloud resources. A `Policy` is a collection of `bindings`. A `binding` binds one or more `members`, or principals, to a single `role`. Principals can be user accounts, service accounts, Google groups, and domains (such as G Suite). A `role` is a named list of permissions; each `role` can be an IAM predefined role or a user-created custom role. For some types of Google Cloud resources, a `binding` can also specify a `condition`, which is a logical expression that allows access to a resource only if the expression evaluates to `true`. A condition can add constraints based on attributes of the request, the resource, or both. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies). **JSON example:** ``` { "bindings": [ { "role": "roles/resourcemanager.organizationAdmin", "members": [ "user:mike@example.com", "group:admins@example.com", "domain:google.com", "serviceAccount:my-project-id@appspot.gserviceaccount.com" ] }, { "role": "roles/resourcemanager.organizationViewer", "members": [ "user:eve@example.com" ], "condition": { "title": "expirable access", "description": "Does not grant access after Sep 2020", "expression": "request.time < timestamp('2020-10-01T00:00:00.000Z')", } } ], "etag": "BwWWja0YfJA=", "version": 3 } ``` **YAML example:** ``` bindings: - members: - user:mike@example.com - group:admins@example.com - domain:google.com - serviceAccount:my-project-id@appspot.gserviceaccount.com role: roles/resourcemanager.organizationAdmin - members: - user:eve@example.com role: roles/resourcemanager.organizationViewer condition: title: expirable access description: Does not grant access after Sep 2020 expression: request.time < timestamp('2020-10-01T00:00:00.000Z') etag: BwWWja0YfJA= version: 3 ``` For a description of IAM and its features, see the [IAM documentation](https://cloud.google.com/iam/docs/).
	Project string `json:"project,omitempty"` // The project that this IAM policy belongs to, in the format of projects/{PROJECT_NUMBER}. This field is available when the IAM policy belongs to a project.
	Assettype string `json:"assetType,omitempty"` // The asset type of the AnalyzeOrgPolicyGovernedAssetsResponse.GovernedIamPolicy.attached_resource. Example: `cloudresourcemanager.googleapis.com/Project` See [Cloud Asset Inventory Supported Asset Types](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for all supported asset types.
}

// GoogleCloudAssetV1Access represents the GoogleCloudAssetV1Access schema from the OpenAPI specification
type GoogleCloudAssetV1Access struct {
	Analysisstate IamPolicyAnalysisState `json:"analysisState,omitempty"` // Represents the detailed state of an entity under analysis, such as a resource, an identity or an access.
	Permission string `json:"permission,omitempty"` // The permission.
	Role string `json:"role,omitempty"` // The role.
}

// ZypperPatch represents the ZypperPatch schema from the OpenAPI specification
type ZypperPatch struct {
	Summary string `json:"summary,omitempty"` // Any summary information provided about this patch.
	Category string `json:"category,omitempty"` // The category of the patch.
	Patchname string `json:"patchName,omitempty"` // The name of the patch.
	Severity string `json:"severity,omitempty"` // The severity specified for this patch
}

// ResourceSelector represents the ResourceSelector schema from the OpenAPI specification
type ResourceSelector struct {
	Fullresourcename string `json:"fullResourceName,omitempty"` // Required. The [full resource name] (https://cloud.google.com/asset-inventory/docs/resource-name-format) of a resource of [supported resource types](https://cloud.google.com/asset-inventory/docs/supported-asset-types#analyzable_asset_types).
}

// AnalyzeIamPolicyResponse represents the AnalyzeIamPolicyResponse schema from the OpenAPI specification
type AnalyzeIamPolicyResponse struct {
	Mainanalysis IamPolicyAnalysis `json:"mainAnalysis,omitempty"` // An analysis message to group the query and results.
	Serviceaccountimpersonationanalysis []IamPolicyAnalysis `json:"serviceAccountImpersonationAnalysis,omitempty"` // The service account impersonation analysis if AnalyzeIamPolicyRequest.analyze_service_account_impersonation is enabled.
	Fullyexplored bool `json:"fullyExplored,omitempty"` // Represents whether all entries in the main_analysis and service_account_impersonation_analysis have been fully explored to answer the query in the request.
}

// IamPolicyAnalysisResult represents the IamPolicyAnalysisResult schema from the OpenAPI specification
type IamPolicyAnalysisResult struct {
	Accesscontrollists []GoogleCloudAssetV1AccessControlList `json:"accessControlLists,omitempty"` // The access control lists derived from the iam_binding that match or potentially match resource and access selectors specified in the request.
	Attachedresourcefullname string `json:"attachedResourceFullName,omitempty"` // The [full resource name](https://cloud.google.com/asset-inventory/docs/resource-name-format) of the resource to which the iam_binding policy attaches.
	Fullyexplored bool `json:"fullyExplored,omitempty"` // Represents whether all analyses on the iam_binding have successfully finished.
	Iambinding Binding `json:"iamBinding,omitempty"` // Associates `members`, or principals, with a `role`.
	Identitylist GoogleCloudAssetV1IdentityList `json:"identityList,omitempty"` // The identities and group edges.
}

// IdentitySelector represents the IdentitySelector schema from the OpenAPI specification
type IdentitySelector struct {
	Identity string `json:"identity,omitempty"` // Required. The identity appear in the form of principals in [IAM policy binding](https://cloud.google.com/iam/reference/rest/v1/Binding). The examples of supported forms are: "user:mike@example.com", "group:admins@example.com", "domain:google.com", "serviceAccount:my-project-id@appspot.gserviceaccount.com". Notice that wildcard characters (such as * and ?) are not supported. You must give a specific identity.
}

// OsInfo represents the OsInfo schema from the OpenAPI specification
type OsInfo struct {
	Longname string `json:"longName,omitempty"` // The operating system long name. For example 'Debian GNU/Linux 9' or 'Microsoft Window Server 2019 Datacenter'.
	Osconfigagentversion string `json:"osconfigAgentVersion,omitempty"` // The current version of the OS Config agent running on the VM.
	Shortname string `json:"shortName,omitempty"` // The operating system short name. For example, 'windows' or 'debian'.
	Version string `json:"version,omitempty"` // The version of the operating system.
	Architecture string `json:"architecture,omitempty"` // The system architecture of the operating system.
	Hostname string `json:"hostname,omitempty"` // The VM hostname.
	Kernelrelease string `json:"kernelRelease,omitempty"` // The kernel release of the operating system.
	Kernelversion string `json:"kernelVersion,omitempty"` // The kernel version of the operating system.
}

// RelatedAssets represents the RelatedAssets schema from the OpenAPI specification
type RelatedAssets struct {
	Relationshipattributes RelationshipAttributes `json:"relationshipAttributes,omitempty"` // DEPRECATED. This message only presents for the purpose of backward-compatibility. The server will never populate this message in responses. The relationship attributes which include `type`, `source_resource_type`, `target_resource_type` and `action`.
	Assets []RelatedAsset `json:"assets,omitempty"` // The peer resources of the relationship.
}

// ListSavedQueriesResponse represents the ListSavedQueriesResponse schema from the OpenAPI specification
type ListSavedQueriesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
	Savedqueries []SavedQuery `json:"savedQueries,omitempty"` // A list of savedQueries.
}

// GoogleCloudAssetV1GovernedContainer represents the GoogleCloudAssetV1GovernedContainer schema from the OpenAPI specification
type GoogleCloudAssetV1GovernedContainer struct {
	Policybundle []AnalyzerOrgPolicy `json:"policyBundle,omitempty"` // The ordered list of all organization policies from the AnalyzeOrgPoliciesResponse.OrgPolicyResult.consolidated_policy.attached_resource. to the scope specified in the request. If the constraint is defined with default policy, it will also appear in the list.
	Project string `json:"project,omitempty"` // The project that this resource belongs to, in the format of projects/{PROJECT_NUMBER}. This field is available when the resource belongs to a project.
	Consolidatedpolicy AnalyzerOrgPolicy `json:"consolidatedPolicy,omitempty"` // This organization policy message is a modified version of the one defined in the Organization Policy system. This message contains several fields defined in the original organization policy with some new fields for analysis purpose.
	Effectivetags []EffectiveTagDetails `json:"effectiveTags,omitempty"` // The effective tags on this resource.
	Folders []string `json:"folders,omitempty"` // The folder(s) that this resource belongs to, in the format of folders/{FOLDER_NUMBER}. This field is available when the resource belongs (directly or cascadingly) to one or more folders.
	Fullresourcename string `json:"fullResourceName,omitempty"` // The [full resource name] (https://cloud.google.com/asset-inventory/docs/resource-name-format) of an organization/folder/project resource.
	Organization string `json:"organization,omitempty"` // The organization that this resource belongs to, in the format of organizations/{ORGANIZATION_NUMBER}. This field is available when the resource belongs (directly or cascadingly) to an organization.
	Parent string `json:"parent,omitempty"` // The [full resource name] (https://cloud.google.com/asset-inventory/docs/resource-name-format) of the parent of AnalyzeOrgPolicyGovernedContainersResponse.GovernedContainer.full_resource_name.
}

// Explanation represents the Explanation schema from the OpenAPI specification
type Explanation struct {
	Matchedpermissions map[string]interface{} `json:"matchedPermissions,omitempty"` // The map from roles to their included permissions that match the permission query (i.e., a query containing `policy.role.permissions:`). Example: if query `policy.role.permissions:compute.disk.get` matches a policy binding that contains owner role, the matched_permissions will be `{"roles/owner": ["compute.disk.get"]}`. The roles can also be found in the returned `policy` bindings. Note that the map is populated only for requests with permission queries.
}

// ListFeedsResponse represents the ListFeedsResponse schema from the OpenAPI specification
type ListFeedsResponse struct {
	Feeds []Feed `json:"feeds,omitempty"` // A list of feeds.
}

// GoogleCloudOrgpolicyV1ListPolicy represents the GoogleCloudOrgpolicyV1ListPolicy schema from the OpenAPI specification
type GoogleCloudOrgpolicyV1ListPolicy struct {
	Allvalues string `json:"allValues,omitempty"` // The policy all_values state.
	Allowedvalues []string `json:"allowedValues,omitempty"` // List of values allowed at this resource. Can only be set if `all_values` is set to `ALL_VALUES_UNSPECIFIED`.
	Deniedvalues []string `json:"deniedValues,omitempty"` // List of values denied at this resource. Can only be set if `all_values` is set to `ALL_VALUES_UNSPECIFIED`.
	Inheritfromparent bool `json:"inheritFromParent,omitempty"` // Determines the inheritance behavior for this `Policy`. By default, a `ListPolicy` set at a resource supersedes any `Policy` set anywhere up the resource hierarchy. However, if `inherit_from_parent` is set to `true`, then the values from the effective `Policy` of the parent resource are inherited, meaning the values set in this `Policy` are added to the values inherited up the hierarchy. Setting `Policy` hierarchies that inherit both allowed values and denied values isn't recommended in most circumstances to keep the configuration simple and understandable. However, it is possible to set a `Policy` with `allowed_values` set that inherits a `Policy` with `denied_values` set. In this case, the values that are allowed must be in `allowed_values` and not present in `denied_values`. For example, suppose you have a `Constraint` `constraints/serviceuser.services`, which has a `constraint_type` of `list_constraint`, and with `constraint_default` set to `ALLOW`. Suppose that at the Organization level, a `Policy` is applied that restricts the allowed API activations to {`E1`, `E2`}. Then, if a `Policy` is applied to a project below the Organization that has `inherit_from_parent` set to `false` and field all_values set to DENY, then an attempt to activate any API will be denied. The following examples demonstrate different possible layerings for `projects/bar` parented by `organizations/foo`: Example 1 (no inherited values): `organizations/foo` has a `Policy` with values: {allowed_values: "E1" allowed_values:"E2"} `projects/bar` has `inherit_from_parent` `false` and values: {allowed_values: "E3" allowed_values: "E4"} The accepted values at `organizations/foo` are `E1`, `E2`. The accepted values at `projects/bar` are `E3`, and `E4`. Example 2 (inherited values): `organizations/foo` has a `Policy` with values: {allowed_values: "E1" allowed_values:"E2"} `projects/bar` has a `Policy` with values: {value: "E3" value: "E4" inherit_from_parent: true} The accepted values at `organizations/foo` are `E1`, `E2`. The accepted values at `projects/bar` are `E1`, `E2`, `E3`, and `E4`. Example 3 (inheriting both allowed and denied values): `organizations/foo` has a `Policy` with values: {allowed_values: "E1" allowed_values: "E2"} `projects/bar` has a `Policy` with: {denied_values: "E1"} The accepted values at `organizations/foo` are `E1`, `E2`. The value accepted at `projects/bar` is `E2`. Example 4 (RestoreDefault): `organizations/foo` has a `Policy` with values: {allowed_values: "E1" allowed_values:"E2"} `projects/bar` has a `Policy` with values: {RestoreDefault: {}} The accepted values at `organizations/foo` are `E1`, `E2`. The accepted values at `projects/bar` are either all or none depending on the value of `constraint_default` (if `ALLOW`, all; if `DENY`, none). Example 5 (no policy inherits parent policy): `organizations/foo` has no `Policy` set. `projects/bar` has no `Policy` set. The accepted values at both levels are either all or none depending on the value of `constraint_default` (if `ALLOW`, all; if `DENY`, none). Example 6 (ListConstraint allowing all): `organizations/foo` has a `Policy` with values: {allowed_values: "E1" allowed_values: "E2"} `projects/bar` has a `Policy` with: {all: ALLOW} The accepted values at `organizations/foo` are `E1`, E2`. Any value is accepted at `projects/bar`. Example 7 (ListConstraint allowing none): `organizations/foo` has a `Policy` with values: {allowed_values: "E1" allowed_values: "E2"} `projects/bar` has a `Policy` with: {all: DENY} The accepted values at `organizations/foo` are `E1`, E2`. No value is accepted at `projects/bar`. Example 10 (allowed and denied subtrees of Resource Manager hierarchy): Given the following resource hierarchy O1->{F1, F2}; F1->{P1}; F2->{P2, P3}, `organizations/foo` has a `Policy` with values: {allowed_values: "under:organizations/O1"} `projects/bar` has a `Policy` with: {allowed_values: "under:projects/P3"} {denied_values: "under:folders/F2"} The accepted values at `organizations/foo` are `organizations/O1`, `folders/F1`, `folders/F2`, `projects/P1`, `projects/P2`, `projects/P3`. The accepted values at `projects/bar` are `organizations/O1`, `folders/F1`, `projects/P1`.
	Suggestedvalue string `json:"suggestedValue,omitempty"` // Optional. The Google Cloud Console will try to default to a configuration that matches the value specified in this `Policy`. If `suggested_value` is not set, it will inherit the value specified higher in the hierarchy, unless `inherit_from_parent` is `false`.
}

// GoogleCloudAssetV1AnalyzeOrgPolicyGovernedAssetsResponseGovernedResource represents the GoogleCloudAssetV1AnalyzeOrgPolicyGovernedAssetsResponseGovernedResource schema from the OpenAPI specification
type GoogleCloudAssetV1AnalyzeOrgPolicyGovernedAssetsResponseGovernedResource struct {
	Folders []string `json:"folders,omitempty"` // The folder(s) that this resource belongs to, in the format of folders/{FOLDER_NUMBER}. This field is available when the resource belongs (directly or cascadingly) to one or more folders.
	Fullresourcename string `json:"fullResourceName,omitempty"` // The [full resource name] (https://cloud.google.com/asset-inventory/docs/resource-name-format) of the Google Cloud resource.
	Organization string `json:"organization,omitempty"` // The organization that this resource belongs to, in the format of organizations/{ORGANIZATION_NUMBER}. This field is available when the resource belongs (directly or cascadingly) to an organization.
	Parent string `json:"parent,omitempty"` // The [full resource name] (https://cloud.google.com/asset-inventory/docs/resource-name-format) of the parent of AnalyzeOrgPolicyGovernedAssetsResponse.GovernedResource.full_resource_name.
	Project string `json:"project,omitempty"` // The project that this resource belongs to, in the format of projects/{PROJECT_NUMBER}. This field is available when the resource belongs to a project.
	Assettype string `json:"assetType,omitempty"` // The asset type of the AnalyzeOrgPolicyGovernedAssetsResponse.GovernedResource.full_resource_name Example: `cloudresourcemanager.googleapis.com/Project` See [Cloud Asset Inventory Supported Asset Types](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for all supported asset types.
	Effectivetags []EffectiveTagDetails `json:"effectiveTags,omitempty"` // The effective tags on this resource.
}

// GoogleCloudAssetV1BigQueryDestination represents the GoogleCloudAssetV1BigQueryDestination schema from the OpenAPI specification
type GoogleCloudAssetV1BigQueryDestination struct {
	Partitionkey string `json:"partitionKey,omitempty"` // The partition key for BigQuery partitioned table.
	Tableprefix string `json:"tablePrefix,omitempty"` // Required. The prefix of the BigQuery tables to which the analysis results will be written. Tables will be created based on this table_prefix if not exist: * _analysis table will contain export operation's metadata. * _analysis_result will contain all the IamPolicyAnalysisResult. When [partition_key] is specified, both tables will be partitioned based on the [partition_key].
	Writedisposition string `json:"writeDisposition,omitempty"` // Optional. Specifies the action that occurs if the destination table or partition already exists. The following values are supported: * WRITE_TRUNCATE: If the table or partition already exists, BigQuery overwrites the entire table or all the partitions data. * WRITE_APPEND: If the table or partition already exists, BigQuery appends the data to the table or the latest partition. * WRITE_EMPTY: If the table already exists and contains data, an error is returned. The default value is WRITE_APPEND. Each action is atomic and only occurs if BigQuery is able to complete the job successfully. Details are at https://cloud.google.com/bigquery/docs/loading-data-local#appending_to_or_overwriting_a_table_using_a_local_file.
	Dataset string `json:"dataset,omitempty"` // Required. The BigQuery dataset in format "projects/projectId/datasets/datasetId", to which the analysis results should be exported. If this dataset does not exist, the export call will return an INVALID_ARGUMENT error.
}

// Binding represents the Binding schema from the OpenAPI specification
type Binding struct {
	Role string `json:"role,omitempty"` // Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`. For an overview of the IAM roles and permissions, see the [IAM documentation](https://cloud.google.com/iam/docs/roles-overview). For a list of the available pre-defined roles, see [here](https://cloud.google.com/iam/docs/understanding-roles).
	Condition Expr `json:"condition,omitempty"` // Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
	Members []string `json:"members,omitempty"` // Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `principal://iam.googleapis.com/locations/global/workforcePools/{pool_id}/subject/{subject_attribute_value}`: A single identity in a workforce identity pool. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/group/{group_id}`: All workforce identities in a group. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/attribute.{attribute_name}/{attribute_value}`: All workforce identities with a specific attribute value. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/*`: All identities in a workforce identity pool. * `principal://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/subject/{subject_attribute_value}`: A single identity in a workload identity pool. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/group/{group_id}`: A workload identity pool group. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/attribute.{attribute_name}/{attribute_value}`: All identities in a workload identity pool with a certain attribute. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/*`: All identities in a workload identity pool. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `deleted:principal://iam.googleapis.com/locations/global/workforcePools/{pool_id}/subject/{subject_attribute_value}`: Deleted single identity in a workforce identity pool. For example, `deleted:principal://iam.googleapis.com/locations/global/workforcePools/my-pool-id/subject/my-subject-attribute-value`.
}

// PartitionSpec represents the PartitionSpec schema from the OpenAPI specification
type PartitionSpec struct {
	Partitionkey string `json:"partitionKey,omitempty"` // The partition key for BigQuery partitioned table.
}

// AnalyzeOrgPolicyGovernedContainersResponse represents the AnalyzeOrgPolicyGovernedContainersResponse schema from the OpenAPI specification
type AnalyzeOrgPolicyGovernedContainersResponse struct {
	Constraint AnalyzerOrgPolicyConstraint `json:"constraint,omitempty"` // The organization policy constraint definition.
	Governedcontainers []GoogleCloudAssetV1GovernedContainer `json:"governedContainers,omitempty"` // The list of the analyzed governed containers.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The page token to fetch the next page for AnalyzeOrgPolicyGovernedContainersResponse.governed_containers.
}

// IamPolicySearchResult represents the IamPolicySearchResult schema from the OpenAPI specification
type IamPolicySearchResult struct {
	Folders []string `json:"folders,omitempty"` // The folder(s) that the IAM policy belongs to, in the form of folders/{FOLDER_NUMBER}. This field is available when the IAM policy belongs to one or more folders. To search against `folders`: * use a field query. Example: `folders:(123 OR 456)` * use a free text query. Example: `123` * specify the `scope` field as this folder in your search request.
	Organization string `json:"organization,omitempty"` // The organization that the IAM policy belongs to, in the form of organizations/{ORGANIZATION_NUMBER}. This field is available when the IAM policy belongs to an organization. To search against `organization`: * use a field query. Example: `organization:123` * use a free text query. Example: `123` * specify the `scope` field as this organization in your search request.
	Policy Policy `json:"policy,omitempty"` // An Identity and Access Management (IAM) policy, which specifies access controls for Google Cloud resources. A `Policy` is a collection of `bindings`. A `binding` binds one or more `members`, or principals, to a single `role`. Principals can be user accounts, service accounts, Google groups, and domains (such as G Suite). A `role` is a named list of permissions; each `role` can be an IAM predefined role or a user-created custom role. For some types of Google Cloud resources, a `binding` can also specify a `condition`, which is a logical expression that allows access to a resource only if the expression evaluates to `true`. A condition can add constraints based on attributes of the request, the resource, or both. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies). **JSON example:** ``` { "bindings": [ { "role": "roles/resourcemanager.organizationAdmin", "members": [ "user:mike@example.com", "group:admins@example.com", "domain:google.com", "serviceAccount:my-project-id@appspot.gserviceaccount.com" ] }, { "role": "roles/resourcemanager.organizationViewer", "members": [ "user:eve@example.com" ], "condition": { "title": "expirable access", "description": "Does not grant access after Sep 2020", "expression": "request.time < timestamp('2020-10-01T00:00:00.000Z')", } } ], "etag": "BwWWja0YfJA=", "version": 3 } ``` **YAML example:** ``` bindings: - members: - user:mike@example.com - group:admins@example.com - domain:google.com - serviceAccount:my-project-id@appspot.gserviceaccount.com role: roles/resourcemanager.organizationAdmin - members: - user:eve@example.com role: roles/resourcemanager.organizationViewer condition: title: expirable access description: Does not grant access after Sep 2020 expression: request.time < timestamp('2020-10-01T00:00:00.000Z') etag: BwWWja0YfJA= version: 3 ``` For a description of IAM and its features, see the [IAM documentation](https://cloud.google.com/iam/docs/).
	Project string `json:"project,omitempty"` // The project that the associated Google Cloud resource belongs to, in the form of projects/{PROJECT_NUMBER}. If an IAM policy is set on a resource (like VM instance, Cloud Storage bucket), the project field will indicate the project that contains the resource. If an IAM policy is set on a folder or orgnization, this field will be empty. To search against the `project`: * specify the `scope` field as this project in your search request.
	Resource string `json:"resource,omitempty"` // The full resource name of the resource associated with this IAM policy. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. See [Cloud Asset Inventory Resource Name Format](https://cloud.google.com/asset-inventory/docs/resource-name-format) for more information. To search against the `resource`: * use a field query. Example: `resource:organizations/123`
	Assettype string `json:"assetType,omitempty"` // The type of the resource associated with this IAM policy. Example: `compute.googleapis.com/Disk`. To search against the `asset_type`: * specify the `asset_types` field in your search request.
	Explanation Explanation `json:"explanation,omitempty"` // Explanation about the IAM policy search result.
}

// GoogleCloudAssetV1BooleanConstraint represents the GoogleCloudAssetV1BooleanConstraint schema from the OpenAPI specification
type GoogleCloudAssetV1BooleanConstraint struct {
}

// MoveAnalysis represents the MoveAnalysis schema from the OpenAPI specification
type MoveAnalysis struct {
	Analysis MoveAnalysisResult `json:"analysis,omitempty"` // An analysis result including blockers and warnings.
	Displayname string `json:"displayName,omitempty"` // The user friendly display name of the analysis. E.g. IAM, organization policy etc.
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
}

// GcsDestination represents the GcsDestination schema from the OpenAPI specification
type GcsDestination struct {
	Uri string `json:"uri,omitempty"` // The URI of the Cloud Storage object. It's the same URI that is used by gsutil. Example: "gs://bucket_name/object_name". See [Viewing and Editing Object Metadata](https://cloud.google.com/storage/docs/viewing-editing-metadata) for more information. If the specified Cloud Storage object already exists and there is no [hold](https://cloud.google.com/storage/docs/object-holds), it will be overwritten with the exported result.
	Uriprefix string `json:"uriPrefix,omitempty"` // The URI prefix of all generated Cloud Storage objects. Example: "gs://bucket_name/object_name_prefix". Each object URI is in format: "gs://bucket_name/object_name_prefix// and only contains assets for that type. starts from 0. Example: "gs://bucket_name/object_name_prefix/compute.googleapis.com/Disk/0" is the first shard of output objects containing all compute.googleapis.com/Disk assets. An INVALID_ARGUMENT error will be returned if file with the same name "gs://bucket_name/object_name_prefix" already exists.
}

// GoogleIdentityAccesscontextmanagerV1Condition represents the GoogleIdentityAccesscontextmanagerV1Condition schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1Condition struct {
	Ipsubnetworks []string `json:"ipSubnetworks,omitempty"` // CIDR block IP subnetwork specification. May be IPv4 or IPv6. Note that for a CIDR IP address block, the specified IP address portion must be properly truncated (i.e. all the host bits must be zero) or the input is considered malformed. For example, "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly, for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32" is not. The originating IP of a request must be in one of the listed subnets in order for this Condition to be true. If empty, all IP addresses are allowed.
	Members []string `json:"members,omitempty"` // The request must be made by one of the provided user or service accounts. Groups are not supported. Syntax: `user:{emailid}` `serviceAccount:{emailid}` If not specified, a request may come from any user.
	Negate bool `json:"negate,omitempty"` // Whether to negate the Condition. If true, the Condition becomes a NAND over its non-empty fields. Any non-empty field criteria evaluating to false will result in the Condition to be satisfied. Defaults to false.
	Regions []string `json:"regions,omitempty"` // The request must originate from one of the provided countries/regions. Must be valid ISO 3166-1 alpha-2 codes.
	Requiredaccesslevels []string `json:"requiredAccessLevels,omitempty"` // A list of other access levels defined in the same `Policy`, referenced by resource name. Referencing an `AccessLevel` which does not exist is an error. All access levels listed must be granted for the Condition to be true. Example: "`accessPolicies/MY_POLICY/accessLevels/LEVEL_NAME"`
	Vpcnetworksources []GoogleIdentityAccesscontextmanagerV1VpcNetworkSource `json:"vpcNetworkSources,omitempty"` // The request must originate from one of the provided VPC networks in Google Cloud. Cannot specify this field together with `ip_subnetworks`.
	Devicepolicy GoogleIdentityAccesscontextmanagerV1DevicePolicy `json:"devicePolicy,omitempty"` // `DevicePolicy` specifies device specific restrictions necessary to acquire a given access level. A `DevicePolicy` specifies requirements for requests from devices to be granted access levels, it does not do any enforcement on the device. `DevicePolicy` acts as an AND over all specified fields, and each repeated field is an OR over its elements. Any unset fields are ignored. For example, if the proto is { os_type : DESKTOP_WINDOWS, os_type : DESKTOP_LINUX, encryption_status: ENCRYPTED}, then the DevicePolicy will be true for requests originating from encrypted Linux desktops and encrypted Windows desktops.
}

// MoveImpact represents the MoveImpact schema from the OpenAPI specification
type MoveImpact struct {
	Detail string `json:"detail,omitempty"` // User friendly impact detail in a free form message.
}

// ConditionContext represents the ConditionContext schema from the OpenAPI specification
type ConditionContext struct {
	Accesstime string `json:"accessTime,omitempty"` // The hypothetical access timestamp to evaluate IAM conditions. Note that this value must not be earlier than the current time; otherwise, an INVALID_ARGUMENT error will be returned.
}

// AnalyzeOrgPolicyGovernedAssetsResponse represents the AnalyzeOrgPolicyGovernedAssetsResponse schema from the OpenAPI specification
type AnalyzeOrgPolicyGovernedAssetsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The page token to fetch the next page for AnalyzeOrgPolicyGovernedAssetsResponse.governed_assets.
	Constraint AnalyzerOrgPolicyConstraint `json:"constraint,omitempty"` // The organization policy constraint definition.
	Governedassets []GoogleCloudAssetV1AnalyzeOrgPolicyGovernedAssetsResponseGovernedAsset `json:"governedAssets,omitempty"` // The list of the analyzed governed assets.
}

// GoogleCloudAssetV1Resource represents the GoogleCloudAssetV1Resource schema from the OpenAPI specification
type GoogleCloudAssetV1Resource struct {
	Analysisstate IamPolicyAnalysisState `json:"analysisState,omitempty"` // Represents the detailed state of an entity under analysis, such as a resource, an identity or an access.
	Fullresourcename string `json:"fullResourceName,omitempty"` // The [full resource name](https://cloud.google.com/asset-inventory/docs/resource-name-format)
}

// AnalyzeIamPolicyLongrunningResponse represents the AnalyzeIamPolicyLongrunningResponse schema from the OpenAPI specification
type AnalyzeIamPolicyLongrunningResponse struct {
}

// GoogleCloudAssetV1AccessControlList represents the GoogleCloudAssetV1AccessControlList schema from the OpenAPI specification
type GoogleCloudAssetV1AccessControlList struct {
	Resources []GoogleCloudAssetV1Resource `json:"resources,omitempty"` // The resources that match one of the following conditions: - The resource_selector, if it is specified in request; - Otherwise, resources reachable from the policy attached resource.
	Accesses []GoogleCloudAssetV1Access `json:"accesses,omitempty"` // The accesses that match one of the following conditions: - The access_selector, if it is specified in request; - Otherwise, access specifiers reachable from the policy binding's role.
	Conditionevaluation ConditionEvaluation `json:"conditionEvaluation,omitempty"` // The condition evaluation.
	Resourceedges []GoogleCloudAssetV1Edge `json:"resourceEdges,omitempty"` // Resource edges of the graph starting from the policy attached resource to any descendant resources. The Edge.source_node contains the full resource name of a parent resource and Edge.target_node contains the full resource name of a child resource. This field is present only if the output_resource_edges option is enabled in request.
}

// ResourceSearchResult represents the ResourceSearchResult schema from the OpenAPI specification
type ResourceSearchResult struct {
	Relationships map[string]interface{} `json:"relationships,omitempty"` // A map of related resources of this resource, keyed by the relationship type. A relationship type is in the format of {SourceType}_{ACTION}_{DestType}. Example: `DISK_TO_INSTANCE`, `DISK_TO_NETWORK`, `INSTANCE_TO_INSTANCEGROUP`. See [supported relationship types](https://cloud.google.com/asset-inventory/docs/supported-asset-types#supported_relationship_types).
	Parentassettype string `json:"parentAssetType,omitempty"` // The type of this resource's immediate parent, if there is one. To search against the `parent_asset_type`: * Use a field query. Example: `parentAssetType:"cloudresourcemanager.googleapis.com/Project"` * Use a free text query. Example: `cloudresourcemanager.googleapis.com/Project`
	Attachedresources []AttachedResource `json:"attachedResources,omitempty"` // Attached resources of this resource. For example, an OSConfig Inventory is an attached resource of a Compute Instance. This field is repeated because a resource could have multiple attached resources. This `attached_resources` field is not searchable. Some attributes of the attached resources are exposed in `additional_attributes` field, so as to allow users to search on them.
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels associated with this resource. See [Labelling and grouping Google Cloud resources](https://cloud.google.com/blog/products/gcp/labelling-and-grouping-your-google-cloud-platform-resources) for more information. This field is available only when the resource's Protobuf contains it. To search against the `labels`: * Use a field query: - query on any label's key or value. Example: `labels:prod` - query by a given label. Example: `labels.env:prod` - query by a given label's existence. Example: `labels.env:*` * Use a free text query. Example: `prod`
	Tagvalueids []string `json:"tagValueIds,omitempty"` // This field is only present for the purpose of backward compatibility. Use the `tags` field instead. TagValue IDs, in the format of tagValues/{TAG_VALUE_ID}. To search against the `tagValueIds`: * Use a field query. Example: - `tagValueIds="tagValues/456"` * Use a free text query. Example: - `456`
	Createtime string `json:"createTime,omitempty"` // The create timestamp of this resource, at which the resource was created. The granularity is in seconds. Timestamp.nanos will always be 0. This field is available only when the resource's Protobuf contains it. To search against `create_time`: * Use a field query. - value in seconds since unix epoch. Example: `createTime > 1609459200` - value in date string. Example: `createTime > 2021-01-01` - value in date-time string (must be quoted). Example: `createTime > "2021-01-01T00:00:00"`
	Name string `json:"name,omitempty"` // The full resource name of this resource. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. See [Cloud Asset Inventory Resource Name Format](https://cloud.google.com/asset-inventory/docs/resource-name-format) for more information. To search against the `name`: * Use a field query. Example: `name:instance1` * Use a free text query. Example: `instance1`
	Networktags []string `json:"networkTags,omitempty"` // Network tags associated with this resource. Like labels, network tags are a type of annotations used to group Google Cloud resources. See [Labelling Google Cloud resources](https://cloud.google.com/blog/products/gcp/labelling-and-grouping-your-google-cloud-platform-resources) for more information. This field is available only when the resource's Protobuf contains it. To search against the `network_tags`: * Use a field query. Example: `networkTags:internal` * Use a free text query. Example: `internal`
	Assettype string `json:"assetType,omitempty"` // The type of this resource. Example: `compute.googleapis.com/Disk`. To search against the `asset_type`: * Specify the `asset_type` field in your search request.
	Description string `json:"description,omitempty"` // One or more paragraphs of text description of this resource. Maximum length could be up to 1M bytes. This field is available only when the resource's Protobuf contains it. To search against the `description`: * Use a field query. Example: `description:"important instance"` * Use a free text query. Example: `"important instance"`
	Effectivetags []EffectiveTagDetails `json:"effectiveTags,omitempty"` // The effective tags on this resource. All of the tags that are both attached to and inherited by a resource are collectively called the effective tags. For more information, see [tag inheritance](https://cloud.google.com/resource-manager/docs/tags/tags-overview#inheritance). To search against the `effective_tags`: * Use a field query. Example: - `effectiveTagKeys:"123456789/env*"` - `effectiveTagKeys="123456789/env"` - `effectiveTagKeys:"env"` - `effectiveTagValues:"env"` - `effectiveTagValues:"env/prod"` - `effectiveTagValues:"123456789/env/prod*"` - `effectiveTagValues="123456789/env/prod"` - `effectiveTagValueIds="tagValues/456"`
	Folders []string `json:"folders,omitempty"` // The folder(s) that this resource belongs to, in the form of folders/{FOLDER_NUMBER}. This field is available when the resource belongs to one or more folders. To search against `folders`: * Use a field query. Example: `folders:(123 OR 456)` * Use a free text query. Example: `123` * Specify the `scope` field as this folder in your search request.
	Location string `json:"location,omitempty"` // Location can be `global`, regional like `us-east1`, or zonal like `us-west1-b`. This field is available only when the resource's Protobuf contains it. To search against the `location`: * Use a field query. Example: `location:us-west*` * Use a free text query. Example: `us-west*`
	State string `json:"state,omitempty"` // The state of this resource. Different resources types have different state definitions that are mapped from various fields of different resource types. This field is available only when the resource's Protobuf contains it. Example: If the resource is an instance provided by Compute Engine, its state will include PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. See `status` definition in [API Reference](https://cloud.google.com/compute/docs/reference/rest/v1/instances). If the resource is a project provided by Resource Manager, its state will include LIFECYCLE_STATE_UNSPECIFIED, ACTIVE, DELETE_REQUESTED and DELETE_IN_PROGRESS. See `lifecycleState` definition in [API Reference](https://cloud.google.com/resource-manager/reference/rest/v1/projects). To search against the `state`: * Use a field query. Example: `state:RUNNING` * Use a free text query. Example: `RUNNING`
	Tagkeys []string `json:"tagKeys,omitempty"` // This field is only present for the purpose of backward compatibility. Use the `tags` field instead. TagKey namespaced names, in the format of {ORG_ID}/{TAG_KEY_SHORT_NAME}. To search against the `tagKeys`: * Use a field query. Example: - `tagKeys:"123456789/env*"` - `tagKeys="123456789/env"` - `tagKeys:"env"` * Use a free text query. Example: - `env`
	Tags []Tag `json:"tags,omitempty"` // The tags directly attached to this resource. To search against the `tags`: * Use a field query. Example: - `tagKeys:"123456789/env*"` - `tagKeys="123456789/env"` - `tagKeys:"env"` - `tagValues:"env"` - `tagValues:"env/prod"` - `tagValues:"123456789/env/prod*"` - `tagValues="123456789/env/prod"` - `tagValueIds="tagValues/456"` * Use a free text query. Example: - `env/prod`
	Additionalattributes map[string]interface{} `json:"additionalAttributes,omitempty"` // The additional searchable attributes of this resource. The attributes may vary from one resource type to another. Examples: `projectId` for Project, `dnsName` for DNS ManagedZone. This field contains a subset of the resource metadata fields that are returned by the List or Get APIs provided by the corresponding Google Cloud service (e.g., Compute Engine). see [API references and supported searchable attributes](https://cloud.google.com/asset-inventory/docs/supported-asset-types) to see which fields are included. You can search values of these fields through free text search. However, you should not consume the field programically as the field names and values may change as the Google Cloud service updates to a new incompatible API version. To search against the `additional_attributes`: * Use a free text query to match the attributes values. Example: to search `additional_attributes = { dnsName: "foobar" }`, you can issue a query `foobar`.
	Organization string `json:"organization,omitempty"` // The organization that this resource belongs to, in the form of organizations/{ORGANIZATION_NUMBER}. This field is available when the resource belongs to an organization. To search against `organization`: * Use a field query. Example: `organization:123` * Use a free text query. Example: `123` * Specify the `scope` field as this organization in your search request.
	Updatetime string `json:"updateTime,omitempty"` // The last update timestamp of this resource, at which the resource was last modified or deleted. The granularity is in seconds. Timestamp.nanos will always be 0. This field is available only when the resource's Protobuf contains it. To search against `update_time`: * Use a field query. - value in seconds since unix epoch. Example: `updateTime < 1609459200` - value in date string. Example: `updateTime < 2021-01-01` - value in date-time string (must be quoted). Example: `updateTime < "2021-01-01T00:00:00"`
	Kmskeys []string `json:"kmsKeys,omitempty"` // The Cloud KMS [CryptoKey](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys) names or [CryptoKeyVersion](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions) names. This field is available only when the resource's Protobuf contains it. To search against the `kms_keys`: * Use a field query. Example: `kmsKeys:key` * Use a free text query. Example: `key`
	Tagvalues []string `json:"tagValues,omitempty"` // This field is only present for the purpose of backward compatibility. Use the `tags` field instead. TagValue namespaced names, in the format of {ORG_ID}/{TAG_KEY_SHORT_NAME}/{TAG_VALUE_SHORT_NAME}. To search against the `tagValues`: * Use a field query. Example: - `tagValues:"env"` - `tagValues:"env/prod"` - `tagValues:"123456789/env/prod*"` - `tagValues="123456789/env/prod"` * Use a free text query. Example: - `prod`
	Versionedresources []VersionedResource `json:"versionedResources,omitempty"` // Versioned resource representations of this resource. This is repeated because there could be multiple versions of resource representations during version migration. This `versioned_resources` field is not searchable. Some attributes of the resource representations are exposed in `additional_attributes` field, so as to allow users to search on them.
	Kmskey string `json:"kmsKey,omitempty"` // The Cloud KMS [CryptoKey](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys) name or [CryptoKeyVersion](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions) name. This field only presents for the purpose of backward compatibility. Use the `kms_keys` field to retrieve Cloud KMS key information. This field is available only when the resource's Protobuf contains it and will only be populated for [these resource types](https://cloud.google.com/asset-inventory/docs/legacy-field-names#resource_types_with_the_to_be_deprecated_kmskey_field) for backward compatible purposes. To search against the `kms_key`: * Use a field query. Example: `kmsKey:key` * Use a free text query. Example: `key`
	Sccsecuritymarks map[string]interface{} `json:"sccSecurityMarks,omitempty"` // The actual content of Security Command Center security marks associated with the asset. To search against SCC SecurityMarks field: * Use a field query: - query by a given key value pair. Example: `sccSecurityMarks.foo=bar` - query by a given key's existence. Example: `sccSecurityMarks.foo:*`
	Displayname string `json:"displayName,omitempty"` // The display name of this resource. This field is available only when the resource's Protobuf contains it. To search against the `display_name`: * Use a field query. Example: `displayName:"My Instance"` * Use a free text query. Example: `"My Instance"`
	Parentfullresourcename string `json:"parentFullResourceName,omitempty"` // The full resource name of this resource's parent, if it has one. To search against the `parent_full_resource_name`: * Use a field query. Example: `parentFullResourceName:"project-name"` * Use a free text query. Example: `project-name`
	Project string `json:"project,omitempty"` // The project that this resource belongs to, in the form of projects/{PROJECT_NUMBER}. This field is available when the resource belongs to a project. To search against `project`: * Use a field query. Example: `project:12345` * Use a free text query. Example: `12345` * Specify the `scope` field as this project in your search request.
}

// QueryResult represents the QueryResult schema from the OpenAPI specification
type QueryResult struct {
	Rows []map[string]interface{} `json:"rows,omitempty"` // Each row hold a query result in the format of `Struct`.
	Schema TableSchema `json:"schema,omitempty"` // BigQuery Compatible table schema.
	Totalrows string `json:"totalRows,omitempty"` // Total rows of the whole query results.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to retrieve the next page of the results.
}

// IamPolicyAnalysis represents the IamPolicyAnalysis schema from the OpenAPI specification
type IamPolicyAnalysis struct {
	Analysisquery IamPolicyAnalysisQuery `json:"analysisQuery,omitempty"` // IAM policy analysis query message.
	Analysisresults []IamPolicyAnalysisResult `json:"analysisResults,omitempty"` // A list of IamPolicyAnalysisResult that matches the analysis query, or empty if no result is found.
	Fullyexplored bool `json:"fullyExplored,omitempty"` // Represents whether all entries in the analysis_results have been fully explored to answer the query.
	Noncriticalerrors []IamPolicyAnalysisState `json:"nonCriticalErrors,omitempty"` // A list of non-critical errors happened during the query handling.
}

// BatchGetAssetsHistoryResponse represents the BatchGetAssetsHistoryResponse schema from the OpenAPI specification
type BatchGetAssetsHistoryResponse struct {
	Assets []TemporalAsset `json:"assets,omitempty"` // A list of assets with valid time windows.
}

// OutputConfig represents the OutputConfig schema from the OpenAPI specification
type OutputConfig struct {
	Gcsdestination GcsDestination `json:"gcsDestination,omitempty"` // A Cloud Storage location.
	Bigquerydestination BigQueryDestination `json:"bigqueryDestination,omitempty"` // A BigQuery destination for exporting assets to.
}

// GoogleCloudOrgpolicyV1Policy represents the GoogleCloudOrgpolicyV1Policy schema from the OpenAPI specification
type GoogleCloudOrgpolicyV1Policy struct {
	Etag string `json:"etag,omitempty"` // An opaque tag indicating the current version of the `Policy`, used for concurrency control. When the `Policy` is returned from either a `GetPolicy` or a `ListOrgPolicy` request, this `etag` indicates the version of the current `Policy` to use when executing a read-modify-write loop. When the `Policy` is returned from a `GetEffectivePolicy` request, the `etag` will be unset. When the `Policy` is used in a `SetOrgPolicy` method, use the `etag` value that was returned from a `GetOrgPolicy` request as part of a read-modify-write loop for concurrency control. Not setting the `etag`in a `SetOrgPolicy` request will result in an unconditional write of the `Policy`.
	Listpolicy GoogleCloudOrgpolicyV1ListPolicy `json:"listPolicy,omitempty"` // Used in `policy_type` to specify how `list_policy` behaves at this resource. `ListPolicy` can define specific values and subtrees of Cloud Resource Manager resource hierarchy (`Organizations`, `Folders`, `Projects`) that are allowed or denied by setting the `allowed_values` and `denied_values` fields. This is achieved by using the `under:` and optional `is:` prefixes. The `under:` prefix is used to denote resource subtree values. The `is:` prefix is used to denote specific values, and is required only if the value contains a ":". Values prefixed with "is:" are treated the same as values with no prefix. Ancestry subtrees must be in one of the following formats: - "projects/", e.g. "projects/tokyo-rain-123" - "folders/", e.g. "folders/1234" - "organizations/", e.g. "organizations/1234" The `supports_under` field of the associated `Constraint` defines whether ancestry prefixes can be used. You can set `allowed_values` and `denied_values` in the same `Policy` if `all_values` is `ALL_VALUES_UNSPECIFIED`. `ALLOW` or `DENY` are used to allow or deny all values. If `all_values` is set to either `ALLOW` or `DENY`, `allowed_values` and `denied_values` must be unset.
	Restoredefault GoogleCloudOrgpolicyV1RestoreDefault `json:"restoreDefault,omitempty"` // Ignores policies set above this resource and restores the `constraint_default` enforcement behavior of the specific `Constraint` at this resource. Suppose that `constraint_default` is set to `ALLOW` for the `Constraint` `constraints/serviceuser.services`. Suppose that organization foo.com sets a `Policy` at their Organization resource node that restricts the allowed service activations to deny all service activations. They could then set a `Policy` with the `policy_type` `restore_default` on several experimental projects, restoring the `constraint_default` enforcement of the `Constraint` for only those projects, allowing those projects to have all services activated.
	Updatetime string `json:"updateTime,omitempty"` // The time stamp the `Policy` was previously updated. This is set by the server, not specified by the caller, and represents the last time a call to `SetOrgPolicy` was made for that `Policy`. Any value set by the client will be ignored.
	Version int `json:"version,omitempty"` // Version of the `Policy`. Default version is 0;
	Booleanpolicy GoogleCloudOrgpolicyV1BooleanPolicy `json:"booleanPolicy,omitempty"` // Used in `policy_type` to specify how `boolean_policy` will behave at this resource.
	Constraint string `json:"constraint,omitempty"` // The name of the `Constraint` the `Policy` is configuring, for example, `constraints/serviceuser.services`. A [list of available constraints](/resource-manager/docs/organization-policy/org-policy-constraints) is available. Immutable after creation.
}

// ListAssetsResponse represents the ListAssetsResponse schema from the OpenAPI specification
type ListAssetsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to retrieve the next page of results. It expires 72 hours after the page token for the first page is generated. Set to empty if there are no remaining results.
	Readtime string `json:"readTime,omitempty"` // Time the snapshot was taken.
	Assets []Asset `json:"assets,omitempty"` // Assets.
}

// Policy represents the Policy schema from the OpenAPI specification
type Policy struct {
	Bindings []Binding `json:"bindings,omitempty"` // Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Etag string `json:"etag,omitempty"` // `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Version int `json:"version,omitempty"` // Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Auditconfigs []AuditConfig `json:"auditConfigs,omitempty"` // Specifies cloud audit logging configuration for this policy.
}

// GoogleCloudAssetV1GcsDestination represents the GoogleCloudAssetV1GcsDestination schema from the OpenAPI specification
type GoogleCloudAssetV1GcsDestination struct {
	Uri string `json:"uri,omitempty"` // Required. The URI of the Cloud Storage object. It's the same URI that is used by gsutil. Example: "gs://bucket_name/object_name". See [Viewing and Editing Object Metadata](https://cloud.google.com/storage/docs/viewing-editing-metadata) for more information. If the specified Cloud Storage object already exists and there is no [hold](https://cloud.google.com/storage/docs/object-holds), it will be overwritten with the analysis result.
}

// GoogleIdentityAccesscontextmanagerV1ServicePerimeter represents the GoogleIdentityAccesscontextmanagerV1ServicePerimeter schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1ServicePerimeter struct {
	Description string `json:"description,omitempty"` // Description of the `ServicePerimeter` and its use. Does not affect behavior.
	Name string `json:"name,omitempty"` // Resource name for the `ServicePerimeter`. Format: `accessPolicies/{access_policy}/servicePerimeters/{service_perimeter}`. The `service_perimeter` component must begin with a letter, followed by alphanumeric characters or `_`. After you create a `ServicePerimeter`, you cannot change its `name`.
	Perimetertype string `json:"perimeterType,omitempty"` // Perimeter type indicator. A single project or VPC network is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, the restricted service list as well as access level lists must be empty.
	Spec GoogleIdentityAccesscontextmanagerV1ServicePerimeterConfig `json:"spec,omitempty"` // `ServicePerimeterConfig` specifies a set of Google Cloud resources that describe specific Service Perimeter configuration.
	Status GoogleIdentityAccesscontextmanagerV1ServicePerimeterConfig `json:"status,omitempty"` // `ServicePerimeterConfig` specifies a set of Google Cloud resources that describe specific Service Perimeter configuration.
	Title string `json:"title,omitempty"` // Human readable title. Must be unique within the Policy.
	Useexplicitdryrunspec bool `json:"useExplicitDryRunSpec,omitempty"` // Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists for all Service Perimeters, and that spec is identical to the status for those Service Perimeters. When this flag is set, it inhibits the generation of the implicit spec, thereby allowing the user to explicitly provide a configuration ("spec") to use in a dry-run version of the Service Perimeter. This allows the user to test changes to the enforced config ("status") without actually enforcing them. This testing is done through analyzing the differences between currently enforced and suggested restrictions. use_explicit_dry_run_spec must bet set to True if any of the fields in the spec are set to non-default values.
}

// AnalyzeOrgPoliciesResponse represents the AnalyzeOrgPoliciesResponse schema from the OpenAPI specification
type AnalyzeOrgPoliciesResponse struct {
	Orgpolicyresults []OrgPolicyResult `json:"orgPolicyResults,omitempty"` // The organization policies under the AnalyzeOrgPoliciesRequest.scope with the AnalyzeOrgPoliciesRequest.constraint.
	Constraint AnalyzerOrgPolicyConstraint `json:"constraint,omitempty"` // The organization policy constraint definition.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The page token to fetch the next page for AnalyzeOrgPoliciesResponse.org_policy_results.
}

// QueryAssetsRequest represents the QueryAssetsRequest schema from the OpenAPI specification
type QueryAssetsRequest struct {
	Timeout string `json:"timeout,omitempty"` // Optional. Specifies the maximum amount of time that the client is willing to wait for the query to complete. By default, this limit is 5 min for the first query, and 1 minute for the following queries. If the query is complete, the `done` field in the `QueryAssetsResponse` is true, otherwise false. Like BigQuery [jobs.query API](https://cloud.google.com/bigquery/docs/reference/rest/v2/jobs/query#queryrequest) The call is not guaranteed to wait for the specified timeout; it typically returns after around 200 seconds (200,000 milliseconds), even if the query is not complete. The field will be ignored when [output_config] is specified.
	Jobreference string `json:"jobReference,omitempty"` // Optional. Reference to the query job, which is from the `QueryAssetsResponse` of previous `QueryAssets` call.
	Outputconfig QueryAssetsOutputConfig `json:"outputConfig,omitempty"` // Output configuration query assets.
	Pagesize int `json:"pageSize,omitempty"` // Optional. The maximum number of rows to return in the results. Responses are limited to 10 MB and 1000 rows. By default, the maximum row count is 1000. When the byte or row count limit is reached, the rest of the query results will be paginated. The field will be ignored when [output_config] is specified.
	Pagetoken string `json:"pageToken,omitempty"` // Optional. A page token received from previous `QueryAssets`. The field will be ignored when [output_config] is specified.
	Readtime string `json:"readTime,omitempty"` // Optional. Queries cloud assets as they appeared at the specified point in time.
	Readtimewindow TimeWindow `json:"readTimeWindow,omitempty"` // A time window specified by its `start_time` and `end_time`.
	Statement string `json:"statement,omitempty"` // Optional. A SQL statement that's compatible with [BigQuery SQL](https://cloud.google.com/bigquery/docs/introduction-sql).
}

// GoogleIdentityAccesscontextmanagerV1ServicePerimeterConfig represents the GoogleIdentityAccesscontextmanagerV1ServicePerimeterConfig schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1ServicePerimeterConfig struct {
	Ingresspolicies []GoogleIdentityAccesscontextmanagerV1IngressPolicy `json:"ingressPolicies,omitempty"` // List of IngressPolicies to apply to the perimeter. A perimeter may have multiple IngressPolicies, each of which is evaluated separately. Access is granted if any Ingress Policy grants it. Must be empty for a perimeter bridge.
	Resources []string `json:"resources,omitempty"` // A list of Google Cloud resources that are inside of the service perimeter. Currently only projects and VPCs are allowed. Project format: `projects/{project_number}` VPC network format: `//compute.googleapis.com/projects/{PROJECT_ID}/global/networks/{NAME}`.
	Restrictedservices []string `json:"restrictedServices,omitempty"` // Google Cloud services that are subject to the Service Perimeter restrictions. For example, if `storage.googleapis.com` is specified, access to the storage buckets inside the perimeter must meet the perimeter's access restrictions.
	Vpcaccessibleservices GoogleIdentityAccesscontextmanagerV1VpcAccessibleServices `json:"vpcAccessibleServices,omitempty"` // Specifies how APIs are allowed to communicate within the Service Perimeter.
	Accesslevels []string `json:"accessLevels,omitempty"` // A list of `AccessLevel` resource names that allow resources within the `ServicePerimeter` to be accessed from the internet. `AccessLevels` listed must be in the same policy as this `ServicePerimeter`. Referencing a nonexistent `AccessLevel` is a syntax error. If no `AccessLevel` names are listed, resources within the perimeter can only be accessed via Google Cloud calls with request origins within the perimeter. Example: `"accessPolicies/MY_POLICY/accessLevels/MY_LEVEL"`. For Service Perimeter Bridge, must be empty.
	Egresspolicies []GoogleIdentityAccesscontextmanagerV1EgressPolicy `json:"egressPolicies,omitempty"` // List of EgressPolicies to apply to the perimeter. A perimeter may have multiple EgressPolicies, each of which is evaluated separately. Access is granted if any EgressPolicy grants it. Must be empty for a perimeter bridge.
}

// GoogleCloudAssetV1p7beta1RelatedAsset represents the GoogleCloudAssetV1p7beta1RelatedAsset schema from the OpenAPI specification
type GoogleCloudAssetV1p7beta1RelatedAsset struct {
	Asset string `json:"asset,omitempty"` // The full name of the asset. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1` See [Resource names](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more information.
	Assettype string `json:"assetType,omitempty"` // The type of the asset. Example: `compute.googleapis.com/Disk` See [Supported asset types](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for more information.
	Ancestors []string `json:"ancestors,omitempty"` // The ancestors of an asset in Google Cloud [resource hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy), represented as a list of relative resource names. An ancestry path starts with the closest ancestor in the hierarchy and ends at root. Example: `["projects/123456789", "folders/5432", "organizations/1234"]`
}

// PolicyInfo represents the PolicyInfo schema from the OpenAPI specification
type PolicyInfo struct {
	Attachedresource string `json:"attachedResource,omitempty"` // The full resource name the policy is directly attached to.
	Policy Policy `json:"policy,omitempty"` // An Identity and Access Management (IAM) policy, which specifies access controls for Google Cloud resources. A `Policy` is a collection of `bindings`. A `binding` binds one or more `members`, or principals, to a single `role`. Principals can be user accounts, service accounts, Google groups, and domains (such as G Suite). A `role` is a named list of permissions; each `role` can be an IAM predefined role or a user-created custom role. For some types of Google Cloud resources, a `binding` can also specify a `condition`, which is a logical expression that allows access to a resource only if the expression evaluates to `true`. A condition can add constraints based on attributes of the request, the resource, or both. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies). **JSON example:** ``` { "bindings": [ { "role": "roles/resourcemanager.organizationAdmin", "members": [ "user:mike@example.com", "group:admins@example.com", "domain:google.com", "serviceAccount:my-project-id@appspot.gserviceaccount.com" ] }, { "role": "roles/resourcemanager.organizationViewer", "members": [ "user:eve@example.com" ], "condition": { "title": "expirable access", "description": "Does not grant access after Sep 2020", "expression": "request.time < timestamp('2020-10-01T00:00:00.000Z')", } } ], "etag": "BwWWja0YfJA=", "version": 3 } ``` **YAML example:** ``` bindings: - members: - user:mike@example.com - group:admins@example.com - domain:google.com - serviceAccount:my-project-id@appspot.gserviceaccount.com role: roles/resourcemanager.organizationAdmin - members: - user:eve@example.com role: roles/resourcemanager.organizationViewer condition: title: expirable access description: Does not grant access after Sep 2020 expression: request.time < timestamp('2020-10-01T00:00:00.000Z') etag: BwWWja0YfJA= version: 3 ``` For a description of IAM and its features, see the [IAM documentation](https://cloud.google.com/iam/docs/).
}

// GoogleIdentityAccesscontextmanagerV1OsConstraint represents the GoogleIdentityAccesscontextmanagerV1OsConstraint schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1OsConstraint struct {
	Minimumversion string `json:"minimumVersion,omitempty"` // The minimum allowed OS version. If not set, any version of this OS satisfies the constraint. Format: `"major.minor.patch"`. Examples: `"10.5.301"`, `"9.2.1"`.
	Ostype string `json:"osType,omitempty"` // Required. The allowed OS type.
	Requireverifiedchromeos bool `json:"requireVerifiedChromeOs,omitempty"` // Only allows requests from devices with a verified Chrome OS. Verifications includes requirements that the device is enterprise-managed, conformant to domain policies, and the caller has permission to call the API targeted by the request.
}

// GoogleIdentityAccesscontextmanagerV1AccessPolicy represents the GoogleIdentityAccesscontextmanagerV1AccessPolicy schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1AccessPolicy struct {
	Title string `json:"title,omitempty"` // Required. Human readable title. Does not affect behavior.
	Etag string `json:"etag,omitempty"` // Output only. An opaque identifier for the current version of the `AccessPolicy`. This will always be a strongly validated etag, meaning that two Access Polices will be identical if and only if their etags are identical. Clients should not expect this to be in any specific format.
	Name string `json:"name,omitempty"` // Output only. Resource name of the `AccessPolicy`. Format: `accessPolicies/{access_policy}`
	Parent string `json:"parent,omitempty"` // Required. The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
	Scopes []string `json:"scopes,omitempty"` // The scopes of the AccessPolicy. Scopes define which resources a policy can restrict and where its resources can be referenced. For example, policy A with `scopes=["folders/123"]` has the following behavior: - ServicePerimeter can only restrict projects within `folders/123`. - ServicePerimeter within policy A can only reference access levels defined within policy A. - Only one policy can include a given scope; thus, attempting to create a second policy which includes `folders/123` will result in an error. If no scopes are provided, then any resource within the organization can be restricted. Scopes cannot be modified after a policy is created. Policies can only have a single scope. Format: list of `folders/{folder_number}` or `projects/{project_number}`
}

// GoogleIdentityAccesscontextmanagerV1IngressFrom represents the GoogleIdentityAccesscontextmanagerV1IngressFrom schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1IngressFrom struct {
	Identities []string `json:"identities,omitempty"` // A list of identities that are allowed access through this ingress policy, in the format of `user:{email_id}` or `serviceAccount:{email_id}`.
	Identitytype string `json:"identityType,omitempty"` // Specifies the type of identities that are allowed access from outside the perimeter. If left unspecified, then members of `identities` field will be allowed access.
	Sources []GoogleIdentityAccesscontextmanagerV1IngressSource `json:"sources,omitempty"` // Sources that this IngressPolicy authorizes access from.
}

// AnalyzerOrgPolicy represents the AnalyzerOrgPolicy schema from the OpenAPI specification
type AnalyzerOrgPolicy struct {
	Rules []GoogleCloudAssetV1Rule `json:"rules,omitempty"` // List of rules for this organization policy.
	Appliedresource string `json:"appliedResource,omitempty"` // The [full resource name] (https://cloud.google.com/asset-inventory/docs/resource-name-format) of an organization/folder/project resource where this organization policy applies to. For any user defined org policies, this field has the same value as the [attached_resource] field. Only for default policy, this field has the different value.
	Attachedresource string `json:"attachedResource,omitempty"` // The [full resource name] (https://cloud.google.com/asset-inventory/docs/resource-name-format) of an organization/folder/project resource where this organization policy is set. Notice that some type of constraints are defined with default policy. This field will be empty for them.
	Inheritfromparent bool `json:"inheritFromParent,omitempty"` // If `inherit_from_parent` is true, Rules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this policy becomes the effective root for evaluation.
	Reset bool `json:"reset,omitempty"` // Ignores policies set above this resource and restores the default behavior of the constraint at this resource. This field can be set in policies for either list or boolean constraints. If set, `rules` must be empty and `inherit_from_parent` must be set to false.
}

// RelatedResource represents the RelatedResource schema from the OpenAPI specification
type RelatedResource struct {
	Assettype string `json:"assetType,omitempty"` // The type of the asset. Example: `compute.googleapis.com/Instance`
	Fullresourcename string `json:"fullResourceName,omitempty"` // The full resource name of the related resource. Example: `//compute.googleapis.com/projects/my_proj_123/zones/instance/instance123`
}

// PubsubDestination represents the PubsubDestination schema from the OpenAPI specification
type PubsubDestination struct {
	Topic string `json:"topic,omitempty"` // The name of the Pub/Sub topic to publish to. Example: `projects/PROJECT_ID/topics/TOPIC_ID`.
}

// TableFieldSchema represents the TableFieldSchema schema from the OpenAPI specification
type TableFieldSchema struct {
	Field string `json:"field,omitempty"` // The field name. The name must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_), and must start with a letter or underscore. The maximum length is 128 characters.
	Fields []TableFieldSchema `json:"fields,omitempty"` // Describes the nested schema fields if the type property is set to RECORD.
	Mode string `json:"mode,omitempty"` // The field mode. Possible values include NULLABLE, REQUIRED and REPEATED. The default value is NULLABLE.
	TypeField string `json:"type,omitempty"` // The field data type. Possible values include * STRING * BYTES * INTEGER * FLOAT * BOOLEAN * TIMESTAMP * DATE * TIME * DATETIME * GEOGRAPHY, * NUMERIC, * BIGNUMERIC, * RECORD (where RECORD indicates that the field contains a nested schema).
}

// GoogleCloudOrgpolicyV1RestoreDefault represents the GoogleCloudOrgpolicyV1RestoreDefault schema from the OpenAPI specification
type GoogleCloudOrgpolicyV1RestoreDefault struct {
}

// TableSchema represents the TableSchema schema from the OpenAPI specification
type TableSchema struct {
	Fields []TableFieldSchema `json:"fields,omitempty"` // Describes the fields in a table.
}

// GoogleIdentityAccesscontextmanagerV1IngressPolicy represents the GoogleIdentityAccesscontextmanagerV1IngressPolicy schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1IngressPolicy struct {
	Ingressfrom GoogleIdentityAccesscontextmanagerV1IngressFrom `json:"ingressFrom,omitempty"` // Defines the conditions under which an IngressPolicy matches a request. Conditions are based on information about the source of the request. The request must satisfy what is defined in `sources` AND identity related fields in order to match.
	Ingressto GoogleIdentityAccesscontextmanagerV1IngressTo `json:"ingressTo,omitempty"` // Defines the conditions under which an IngressPolicy matches a request. Conditions are based on information about the ApiOperation intended to be performed on the target resource of the request. The request must satisfy what is defined in `operations` AND `resources` in order to match.
}

// AnalyzerOrgPolicyConstraint represents the AnalyzerOrgPolicyConstraint schema from the OpenAPI specification
type AnalyzerOrgPolicyConstraint struct {
	Customconstraint GoogleCloudAssetV1CustomConstraint `json:"customConstraint,omitempty"` // The definition of a custom constraint.
	Googledefinedconstraint GoogleCloudAssetV1Constraint `json:"googleDefinedConstraint,omitempty"` // The definition of a constraint.
}

// GoogleCloudAssetV1CustomConstraint represents the GoogleCloudAssetV1CustomConstraint schema from the OpenAPI specification
type GoogleCloudAssetV1CustomConstraint struct {
	Methodtypes []string `json:"methodTypes,omitempty"` // All the operations being applied for this constraint.
	Name string `json:"name,omitempty"` // Name of the constraint. This is unique within the organization. Format of the name should be * `organizations/{organization_id}/customConstraints/{custom_constraint_id}` Example : "organizations/123/customConstraints/custom.createOnlyE2TypeVms"
	Resourcetypes []string `json:"resourceTypes,omitempty"` // The Resource Instance type on which this policy applies to. Format will be of the form : "/" Example: * `compute.googleapis.com/Instance`.
	Actiontype string `json:"actionType,omitempty"` // Allow or deny type.
	Condition string `json:"condition,omitempty"` // Organization Policy condition/expression. For example: `resource.instanceName.matches("[production|test]_.*_(\d)+")'` or, `resource.management.auto_upgrade == true`
	Description string `json:"description,omitempty"` // Detailed information about this custom policy constraint.
	Displayname string `json:"displayName,omitempty"` // One line display name for the UI.
}

// AttachedResource represents the AttachedResource schema from the OpenAPI specification
type AttachedResource struct {
	Versionedresources []VersionedResource `json:"versionedResources,omitempty"` // Versioned resource representations of this attached resource. This is repeated because there could be multiple versions of the attached resource representations during version migration.
	Assettype string `json:"assetType,omitempty"` // The type of this attached resource. Example: `osconfig.googleapis.com/Inventory` You can find the supported attached asset types of each resource in this table: `https://cloud.google.com/asset-inventory/docs/supported-asset-types`
}

// GoogleIdentityAccesscontextmanagerV1VpcAccessibleServices represents the GoogleIdentityAccesscontextmanagerV1VpcAccessibleServices schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1VpcAccessibleServices struct {
	Allowedservices []string `json:"allowedServices,omitempty"` // The list of APIs usable within the Service Perimeter. Must be empty unless 'enable_restriction' is True. You can specify a list of individual services, as well as include the 'RESTRICTED-SERVICES' value, which automatically includes all of the services protected by the perimeter.
	Enablerestriction bool `json:"enableRestriction,omitempty"` // Whether to restrict API calls within the Service Perimeter to the list of APIs specified in 'allowed_services'.
}

// QueryAssetsResponse represents the QueryAssetsResponse schema from the OpenAPI specification
type QueryAssetsResponse struct {
	Done bool `json:"done,omitempty"` // The query response, which can be either an `error` or a valid `response`. If `done` == `false` and the query result is being saved in a output, the output_config field will be set. If `done` == `true`, exactly one of `error`, `query_result` or `output_config` will be set.
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
	Jobreference string `json:"jobReference,omitempty"` // Reference to a query job.
	Outputconfig QueryAssetsOutputConfig `json:"outputConfig,omitempty"` // Output configuration query assets.
	Queryresult QueryResult `json:"queryResult,omitempty"` // Execution results of the query. The result is formatted as rows represented by BigQuery compatible [schema]. When pagination is necessary, it will contains the page token to retrieve the results of following pages.
}

// BatchGetEffectiveIamPoliciesResponse represents the BatchGetEffectiveIamPoliciesResponse schema from the OpenAPI specification
type BatchGetEffectiveIamPoliciesResponse struct {
	Policyresults []EffectiveIamPolicy `json:"policyResults,omitempty"` // The effective policies for a batch of resources. Note that the results order is the same as the order of BatchGetEffectiveIamPoliciesRequest.names. When a resource does not have any effective IAM policies, its corresponding policy_result will contain empty EffectiveIamPolicy.policies.
}

// FeedOutputConfig represents the FeedOutputConfig schema from the OpenAPI specification
type FeedOutputConfig struct {
	Pubsubdestination PubsubDestination `json:"pubsubDestination,omitempty"` // A Pub/Sub destination.
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Tagvalue string `json:"tagValue,omitempty"` // TagValue namespaced name, in the format of {ORG_ID}/{TAG_KEY_SHORT_NAME}/{TAG_VALUE_SHORT_NAME}.
	Tagvalueid string `json:"tagValueId,omitempty"` // TagValue ID, in the format of tagValues/{TAG_VALUE_ID}.
	Tagkey string `json:"tagKey,omitempty"` // TagKey namespaced name, in the format of {ORG_ID}/{TAG_KEY_SHORT_NAME}.
}

// AuditConfig represents the AuditConfig schema from the OpenAPI specification
type AuditConfig struct {
	Auditlogconfigs []AuditLogConfig `json:"auditLogConfigs,omitempty"` // The configuration for logging of each type of permission.
	Service string `json:"service,omitempty"` // Specifies a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
}

// GoogleIdentityAccesscontextmanagerV1IngressSource represents the GoogleIdentityAccesscontextmanagerV1IngressSource schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1IngressSource struct {
	Resource string `json:"resource,omitempty"` // A Google Cloud resource that is allowed to ingress the perimeter. Requests from these resources will be allowed to access perimeter data. Currently only projects and VPCs are allowed. Project format: `projects/{project_number}` VPC network format: `//compute.googleapis.com/projects/{PROJECT_ID}/global/networks/{NAME}`. The project may be in any Google Cloud organization, not just the organization that the perimeter is defined in. `*` is not allowed, the case of allowing all Google Cloud resources only is not supported.
	Accesslevel string `json:"accessLevel,omitempty"` // An AccessLevel resource name that allow resources within the ServicePerimeters to be accessed from the internet. AccessLevels listed must be in the same policy as this ServicePerimeter. Referencing a nonexistent AccessLevel will cause an error. If no AccessLevel names are listed, resources within the perimeter can only be accessed via Google Cloud calls with request origins within the perimeter. Example: `accessPolicies/MY_POLICY/accessLevels/MY_LEVEL`. If a single `*` is specified for `access_level`, then all IngressSources will be allowed.
}

// AccessSelector represents the AccessSelector schema from the OpenAPI specification
type AccessSelector struct {
	Permissions []string `json:"permissions,omitempty"` // Optional. The permissions to appear in result.
	Roles []string `json:"roles,omitempty"` // Optional. The roles to appear in result.
}

// RelatedAsset represents the RelatedAsset schema from the OpenAPI specification
type RelatedAsset struct {
	Ancestors []string `json:"ancestors,omitempty"` // The ancestors of an asset in Google Cloud [resource hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy), represented as a list of relative resource names. An ancestry path starts with the closest ancestor in the hierarchy and ends at root. Example: `["projects/123456789", "folders/5432", "organizations/1234"]`
	Asset string `json:"asset,omitempty"` // The full name of the asset. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1` See [Resource names](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more information.
	Assettype string `json:"assetType,omitempty"` // The type of the asset. Example: `compute.googleapis.com/Disk` See [Supported asset types](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for more information.
	Relationshiptype string `json:"relationshipType,omitempty"` // The unique identifier of the relationship type. Example: `INSTANCE_TO_INSTANCEGROUP`
}

// GoogleCloudAssetV1p7beta1Resource represents the GoogleCloudAssetV1p7beta1Resource schema from the OpenAPI specification
type GoogleCloudAssetV1p7beta1Resource struct {
	Parent string `json:"parent,omitempty"` // The full name of the immediate parent of this resource. See [Resource Names](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more information. For Google Cloud assets, this value is the parent resource defined in the [IAM policy hierarchy](https://cloud.google.com/iam/docs/overview#policy_hierarchy). Example: `//cloudresourcemanager.googleapis.com/projects/my_project_123` For third-party assets, this field may be set differently.
	Resourceurl string `json:"resourceUrl,omitempty"` // The REST URL for accessing the resource. An HTTP `GET` request using this URL returns the resource itself. Example: `https://cloudresourcemanager.googleapis.com/v1/projects/my-project-123` This value is unspecified for resources without a REST API.
	Version string `json:"version,omitempty"` // The API version. Example: `v1`
	Data map[string]interface{} `json:"data,omitempty"` // The content of the resource, in which some sensitive fields are removed and may not be present.
	Discoverydocumenturi string `json:"discoveryDocumentUri,omitempty"` // The URL of the discovery document containing the resource's JSON schema. Example: `https://www.googleapis.com/discovery/v1/apis/compute/v1/rest` This value is unspecified for resources that do not have an API based on a discovery document, such as Cloud Bigtable.
	Discoveryname string `json:"discoveryName,omitempty"` // The JSON schema name listed in the discovery document. Example: `Project` This value is unspecified for resources that do not have an API based on a discovery document, such as Cloud Bigtable.
	Location string `json:"location,omitempty"` // The location of the resource in Google Cloud, such as its zone and region. For more information, see https://cloud.google.com/about/locations/.
}

// GoogleCloudAssetV1p7beta1Asset represents the GoogleCloudAssetV1p7beta1Asset schema from the OpenAPI specification
type GoogleCloudAssetV1p7beta1Asset struct {
	Assettype string `json:"assetType,omitempty"` // The type of the asset. Example: `compute.googleapis.com/Disk` See [Supported asset types](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for more information.
	Name string `json:"name,omitempty"` // The full name of the asset. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1` See [Resource names](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more information.
	Accesspolicy GoogleIdentityAccesscontextmanagerV1AccessPolicy `json:"accessPolicy,omitempty"` // `AccessPolicy` is a container for `AccessLevels` (which define the necessary attributes to use Google Cloud services) and `ServicePerimeters` (which define regions of services able to freely pass data within a perimeter). An access policy is globally visible within an organization, and the restrictions it specifies apply to all projects within an organization.
	Ancestors []string `json:"ancestors,omitempty"` // The ancestry path of an asset in Google Cloud [resource hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy), represented as a list of relative resource names. An ancestry path starts with the closest ancestor in the hierarchy and ends at root. If the asset is a project, folder, or organization, the ancestry path starts from the asset itself. Example: `["projects/123456789", "folders/5432", "organizations/1234"]`
	Orgpolicy []GoogleCloudOrgpolicyV1Policy `json:"orgPolicy,omitempty"` // A representation of an [organization policy](https://cloud.google.com/resource-manager/docs/organization-policy/overview#organization_policy). There can be more than one organization policy with different constraints set on a given resource.
	Relatedassets GoogleCloudAssetV1p7beta1RelatedAssets `json:"relatedAssets,omitempty"` // The detailed related assets with the `relationship_type`.
	Resource GoogleCloudAssetV1p7beta1Resource `json:"resource,omitempty"` // A representation of a Google Cloud resource.
	Iampolicy Policy `json:"iamPolicy,omitempty"` // An Identity and Access Management (IAM) policy, which specifies access controls for Google Cloud resources. A `Policy` is a collection of `bindings`. A `binding` binds one or more `members`, or principals, to a single `role`. Principals can be user accounts, service accounts, Google groups, and domains (such as G Suite). A `role` is a named list of permissions; each `role` can be an IAM predefined role or a user-created custom role. For some types of Google Cloud resources, a `binding` can also specify a `condition`, which is a logical expression that allows access to a resource only if the expression evaluates to `true`. A condition can add constraints based on attributes of the request, the resource, or both. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies). **JSON example:** ``` { "bindings": [ { "role": "roles/resourcemanager.organizationAdmin", "members": [ "user:mike@example.com", "group:admins@example.com", "domain:google.com", "serviceAccount:my-project-id@appspot.gserviceaccount.com" ] }, { "role": "roles/resourcemanager.organizationViewer", "members": [ "user:eve@example.com" ], "condition": { "title": "expirable access", "description": "Does not grant access after Sep 2020", "expression": "request.time < timestamp('2020-10-01T00:00:00.000Z')", } } ], "etag": "BwWWja0YfJA=", "version": 3 } ``` **YAML example:** ``` bindings: - members: - user:mike@example.com - group:admins@example.com - domain:google.com - serviceAccount:my-project-id@appspot.gserviceaccount.com role: roles/resourcemanager.organizationAdmin - members: - user:eve@example.com role: roles/resourcemanager.organizationViewer condition: title: expirable access description: Does not grant access after Sep 2020 expression: request.time < timestamp('2020-10-01T00:00:00.000Z') etag: BwWWja0YfJA= version: 3 ``` For a description of IAM and its features, see the [IAM documentation](https://cloud.google.com/iam/docs/).
	Serviceperimeter GoogleIdentityAccesscontextmanagerV1ServicePerimeter `json:"servicePerimeter,omitempty"` // `ServicePerimeter` describes a set of Google Cloud resources which can freely import and export data amongst themselves, but not export outside of the `ServicePerimeter`. If a request with a source within this `ServicePerimeter` has a target outside of the `ServicePerimeter`, the request will be blocked. Otherwise the request is allowed. There are two types of Service Perimeter - Regular and Bridge. Regular Service Perimeters cannot overlap, a single Google Cloud project or VPC network can only belong to a single regular Service Perimeter. Service Perimeter Bridges can contain only Google Cloud projects as members, a single Google Cloud project may belong to multiple Service Perimeter Bridges.
	Updatetime string `json:"updateTime,omitempty"` // The last update timestamp of an asset. update_time is updated when create/update/delete operation is performed.
	Accesslevel GoogleIdentityAccesscontextmanagerV1AccessLevel `json:"accessLevel,omitempty"` // An `AccessLevel` is a label that can be applied to requests to Google Cloud services, along with a list of requirements necessary for the label to be applied.
}

// GoogleIdentityAccesscontextmanagerV1IngressTo represents the GoogleIdentityAccesscontextmanagerV1IngressTo schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1IngressTo struct {
	Resources []string `json:"resources,omitempty"` // A list of resources, currently only projects in the form `projects/`, protected by this ServicePerimeter that are allowed to be accessed by sources defined in the corresponding IngressFrom. If a single `*` is specified, then access to all resources inside the perimeter are allowed.
	Operations []GoogleIdentityAccesscontextmanagerV1ApiOperation `json:"operations,omitempty"` // A list of ApiOperations allowed to be performed by the sources specified in corresponding IngressFrom in this ServicePerimeter.
}

// IamPolicyAnalysisState represents the IamPolicyAnalysisState schema from the OpenAPI specification
type IamPolicyAnalysisState struct {
	Cause string `json:"cause,omitempty"` // The human-readable description of the cause of failure.
	Code string `json:"code,omitempty"` // The Google standard error code that best describes the state. For example: - OK means the analysis on this entity has been successfully finished; - PERMISSION_DENIED means an access denied error is encountered; - DEADLINE_EXCEEDED means the analysis on this entity hasn't been started in time;
}

// WindowsQuickFixEngineeringPackage represents the WindowsQuickFixEngineeringPackage schema from the OpenAPI specification
type WindowsQuickFixEngineeringPackage struct {
	Caption string `json:"caption,omitempty"` // A short textual description of the QFE update.
	Description string `json:"description,omitempty"` // A textual description of the QFE update.
	Hotfixid string `json:"hotFixId,omitempty"` // Unique identifier associated with a particular QFE update.
	Installtime string `json:"installTime,omitempty"` // Date that the QFE update was installed. Mapped from installed_on field.
}

// QueryContent represents the QueryContent schema from the OpenAPI specification
type QueryContent struct {
	Iampolicyanalysisquery IamPolicyAnalysisQuery `json:"iamPolicyAnalysisQuery,omitempty"` // IAM policy analysis query message.
}

// VersionedResource represents the VersionedResource schema from the OpenAPI specification
type VersionedResource struct {
	Resource map[string]interface{} `json:"resource,omitempty"` // JSON representation of the resource as defined by the corresponding service providing this resource. Example: If the resource is an instance provided by Compute Engine, this field will contain the JSON representation of the instance as defined by Compute Engine: `https://cloud.google.com/compute/docs/reference/rest/v1/instances`. You can find the resource definition for each supported resource type in this table: `https://cloud.google.com/asset-inventory/docs/supported-asset-types`
	Version string `json:"version,omitempty"` // API version of the resource. Example: If the resource is an instance provided by Compute Engine v1 API as defined in `https://cloud.google.com/compute/docs/reference/rest/v1/instances`, version will be "v1".
}

// Inventory represents the Inventory schema from the OpenAPI specification
type Inventory struct {
	Osinfo OsInfo `json:"osInfo,omitempty"` // Operating system information for the VM.
	Updatetime string `json:"updateTime,omitempty"` // Output only. Timestamp of the last reported inventory for the VM.
	Items map[string]interface{} `json:"items,omitempty"` // Inventory items related to the VM keyed by an opaque unique identifier for each inventory item. The identifier is unique to each distinct and addressable inventory item and will change, when there is a new package version.
	Name string `json:"name,omitempty"` // Output only. The `Inventory` API resource name. Format: `projects/{project_number}/locations/{location}/instances/{instance_id}/inventory`
}

// Item represents the Item schema from the OpenAPI specification
type Item struct {
	Updatetime string `json:"updateTime,omitempty"` // When this inventory item was last modified.
	Availablepackage SoftwarePackage `json:"availablePackage,omitempty"` // Software package information of the operating system.
	Createtime string `json:"createTime,omitempty"` // When this inventory item was first detected.
	Id string `json:"id,omitempty"` // Identifier for this item, unique across items for this VM.
	Installedpackage SoftwarePackage `json:"installedPackage,omitempty"` // Software package information of the operating system.
	Origintype string `json:"originType,omitempty"` // The origin of this inventory item.
	TypeField string `json:"type,omitempty"` // The specific type of inventory, correlating to its specific details.
}

// Feed represents the Feed schema from the OpenAPI specification
type Feed struct {
	Condition Expr `json:"condition,omitempty"` // Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
	Contenttype string `json:"contentType,omitempty"` // Asset content type. If not specified, no content but the asset name and type will be returned.
	Feedoutputconfig FeedOutputConfig `json:"feedOutputConfig,omitempty"` // Output configuration for asset feed destination.
	Name string `json:"name,omitempty"` // Required. The format will be projects/{project_number}/feeds/{client-assigned_feed_identifier} or folders/{folder_number}/feeds/{client-assigned_feed_identifier} or organizations/{organization_number}/feeds/{client-assigned_feed_identifier} The client-assigned feed identifier must be unique within the parent project/folder/organization.
	Relationshiptypes []string `json:"relationshipTypes,omitempty"` // A list of relationship types to output, for example: `INSTANCE_TO_INSTANCEGROUP`. This field should only be specified if content_type=RELATIONSHIP. * If specified: it outputs specified relationship updates on the [asset_names] or the [asset_types]. It returns an error if any of the [relationship_types] doesn't belong to the supported relationship types of the [asset_names] or [asset_types], or any of the [asset_names] or the [asset_types] doesn't belong to the source types of the [relationship_types]. * Otherwise: it outputs the supported relationships of the types of [asset_names] and [asset_types] or returns an error if any of the [asset_names] or the [asset_types] has no replationship support. See [Introduction to Cloud Asset Inventory](https://cloud.google.com/asset-inventory/docs/overview) for all supported asset types and relationship types.
	Assetnames []string `json:"assetNames,omitempty"` // A list of the full names of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. For a list of the full names for supported asset types, see [Resource name format](/asset-inventory/docs/resource-name-format).
	Assettypes []string `json:"assetTypes,omitempty"` // A list of types of the assets to receive updates. You must specify either or both of asset_names and asset_types. Only asset updates matching specified asset_names or asset_types are exported to the feed. Example: `"compute.googleapis.com/Disk"` For a list of all supported asset types, see [Supported asset types](/asset-inventory/docs/supported-asset-types).
}

// GoogleCloudAssetV1IdentityList represents the GoogleCloudAssetV1IdentityList schema from the OpenAPI specification
type GoogleCloudAssetV1IdentityList struct {
	Identities []GoogleCloudAssetV1Identity `json:"identities,omitempty"` // Only the identities that match one of the following conditions will be presented: - The identity_selector, if it is specified in request; - Otherwise, identities reachable from the policy binding's members.
	Groupedges []GoogleCloudAssetV1Edge `json:"groupEdges,omitempty"` // Group identity edges of the graph starting from the binding's group members to any node of the identities. The Edge.source_node contains a group, such as `group:parent@google.com`. The Edge.target_node contains a member of the group, such as `group:child@google.com` or `user:foo@google.com`. This field is present only if the output_group_edges option is enabled in request.
}

// GoogleCloudOrgpolicyV1BooleanPolicy represents the GoogleCloudOrgpolicyV1BooleanPolicy schema from the OpenAPI specification
type GoogleCloudOrgpolicyV1BooleanPolicy struct {
	Enforced bool `json:"enforced,omitempty"` // If `true`, then the `Policy` is enforced. If `false`, then any configuration is acceptable. Suppose you have a `Constraint` `constraints/compute.disableSerialPortAccess` with `constraint_default` set to `ALLOW`. A `Policy` for that `Constraint` exhibits the following behavior: - If the `Policy` at this resource has enforced set to `false`, serial port connection attempts will be allowed. - If the `Policy` at this resource has enforced set to `true`, serial port connection attempts will be refused. - If the `Policy` at this resource is `RestoreDefault`, serial port connection attempts will be allowed. - If no `Policy` is set at this resource or anywhere higher in the resource hierarchy, serial port connection attempts will be allowed. - If no `Policy` is set at this resource, but one exists higher in the resource hierarchy, the behavior is as if the`Policy` were set at this resource. The following examples demonstrate the different possible layerings: Example 1 (nearest `Constraint` wins): `organizations/foo` has a `Policy` with: {enforced: false} `projects/bar` has no `Policy` set. The constraint at `projects/bar` and `organizations/foo` will not be enforced. Example 2 (enforcement gets replaced): `organizations/foo` has a `Policy` with: {enforced: false} `projects/bar` has a `Policy` with: {enforced: true} The constraint at `organizations/foo` is not enforced. The constraint at `projects/bar` is enforced. Example 3 (RestoreDefault): `organizations/foo` has a `Policy` with: {enforced: true} `projects/bar` has a `Policy` with: {RestoreDefault: {}} The constraint at `organizations/foo` is enforced. The constraint at `projects/bar` is not enforced, because `constraint_default` for the `Constraint` is `ALLOW`.
}

// Permissions represents the Permissions schema from the OpenAPI specification
type Permissions struct {
	Permissions []string `json:"permissions,omitempty"` // A list of permissions. A sample permission string: `compute.disk.get`.
}

// WindowsApplication represents the WindowsApplication schema from the OpenAPI specification
type WindowsApplication struct {
	Publisher string `json:"publisher,omitempty"` // The name of the manufacturer for the product or application.
	Displayname string `json:"displayName,omitempty"` // The name of the application or product.
	Displayversion string `json:"displayVersion,omitempty"` // The version of the product or application in string format.
	Helplink string `json:"helpLink,omitempty"` // The internet address for technical support.
	Installdate Date `json:"installDate,omitempty"` // Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
}

// GoogleCloudAssetV1Constraint represents the GoogleCloudAssetV1Constraint schema from the OpenAPI specification
type GoogleCloudAssetV1Constraint struct {
	Constraintdefault string `json:"constraintDefault,omitempty"` // The evaluation behavior of this constraint in the absence of 'Policy'.
	Description string `json:"description,omitempty"` // Detailed description of what this `Constraint` controls as well as how and where it is enforced.
	Displayname string `json:"displayName,omitempty"` // The human readable name of the constraint.
	Listconstraint GoogleCloudAssetV1ListConstraint `json:"listConstraint,omitempty"` // A `Constraint` that allows or disallows a list of string values, which are configured by an organization's policy administrator with a `Policy`.
	Name string `json:"name,omitempty"` // The unique name of the constraint. Format of the name should be * `constraints/{constraint_name}` For example, `constraints/compute.disableSerialPortAccess`.
	Booleanconstraint GoogleCloudAssetV1BooleanConstraint `json:"booleanConstraint,omitempty"` // A `Constraint` that is either enforced or not. For example a constraint `constraints/compute.disableSerialPortAccess`. If it is enforced on a VM instance, serial port connections will not be opened to that instance.
}

// GoogleCloudAssetV1QueryAssetsOutputConfigBigQueryDestination represents the GoogleCloudAssetV1QueryAssetsOutputConfigBigQueryDestination schema from the OpenAPI specification
type GoogleCloudAssetV1QueryAssetsOutputConfigBigQueryDestination struct {
	Dataset string `json:"dataset,omitempty"` // Required. The BigQuery dataset where the query results will be saved. It has the format of "projects/{projectId}/datasets/{datasetId}".
	Table string `json:"table,omitempty"` // Required. The BigQuery table where the query results will be saved. If this table does not exist, a new table with the given name will be created.
	Writedisposition string `json:"writeDisposition,omitempty"` // Specifies the action that occurs if the destination table or partition already exists. The following values are supported: * WRITE_TRUNCATE: If the table or partition already exists, BigQuery overwrites the entire table or all the partitions data. * WRITE_APPEND: If the table or partition already exists, BigQuery appends the data to the table or the latest partition. * WRITE_EMPTY: If the table already exists and contains data, a 'duplicate' error is returned in the job result. The default value is WRITE_EMPTY.
}

// RelatedResources represents the RelatedResources schema from the OpenAPI specification
type RelatedResources struct {
	Relatedresources []RelatedResource `json:"relatedResources,omitempty"` // The detailed related resources of the primary resource.
}

// GoogleIdentityAccesscontextmanagerV1AccessLevel represents the GoogleIdentityAccesscontextmanagerV1AccessLevel schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1AccessLevel struct {
	Custom GoogleIdentityAccesscontextmanagerV1CustomLevel `json:"custom,omitempty"` // `CustomLevel` is an `AccessLevel` using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request. See CEL spec at: https://github.com/google/cel-spec
	Description string `json:"description,omitempty"` // Description of the `AccessLevel` and its use. Does not affect behavior.
	Name string `json:"name,omitempty"` // Resource name for the `AccessLevel`. Format: `accessPolicies/{access_policy}/accessLevels/{access_level}`. The `access_level` component must begin with a letter, followed by alphanumeric characters or `_`. Its maximum length is 50 characters. After you create an `AccessLevel`, you cannot change its `name`.
	Title string `json:"title,omitempty"` // Human readable title. Must be unique within the Policy.
	Basic GoogleIdentityAccesscontextmanagerV1BasicLevel `json:"basic,omitempty"` // `BasicLevel` is an `AccessLevel` using a set of recommended features.
}

// WindowsUpdateCategory represents the WindowsUpdateCategory schema from the OpenAPI specification
type WindowsUpdateCategory struct {
	Id string `json:"id,omitempty"` // The identifier of the windows update category.
	Name string `json:"name,omitempty"` // The name of the windows update category.
}

// TemporalAsset represents the TemporalAsset schema from the OpenAPI specification
type TemporalAsset struct {
	Window TimeWindow `json:"window,omitempty"` // A time window specified by its `start_time` and `end_time`.
	Asset Asset `json:"asset,omitempty"` // An asset in Google Cloud. An asset can be any resource in the Google Cloud [resource hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy), a resource outside the Google Cloud resource hierarchy (such as Google Kubernetes Engine clusters and objects), or a policy (e.g. IAM policy), or a relationship (e.g. an INSTANCE_TO_INSTANCEGROUP relationship). See [Supported asset types](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for more information.
	Deleted bool `json:"deleted,omitempty"` // Whether the asset has been deleted or not.
	Priorasset Asset `json:"priorAsset,omitempty"` // An asset in Google Cloud. An asset can be any resource in the Google Cloud [resource hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy), a resource outside the Google Cloud resource hierarchy (such as Google Kubernetes Engine clusters and objects), or a policy (e.g. IAM policy), or a relationship (e.g. an INSTANCE_TO_INSTANCEGROUP relationship). See [Supported asset types](https://cloud.google.com/asset-inventory/docs/supported-asset-types) for more information.
	Priorassetstate string `json:"priorAssetState,omitempty"` // State of prior_asset.
}

// GoogleCloudAssetV1StringValues represents the GoogleCloudAssetV1StringValues schema from the OpenAPI specification
type GoogleCloudAssetV1StringValues struct {
	Allowedvalues []string `json:"allowedValues,omitempty"` // List of values allowed at this resource.
	Deniedvalues []string `json:"deniedValues,omitempty"` // List of values denied at this resource.
}

// GoogleCloudAssetV1Rule represents the GoogleCloudAssetV1Rule schema from the OpenAPI specification
type GoogleCloudAssetV1Rule struct {
	Allowall bool `json:"allowAll,omitempty"` // Setting this to true means that all values are allowed. This field can be set only in Policies for list constraints.
	Condition Expr `json:"condition,omitempty"` // Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
	Conditionevaluation ConditionEvaluation `json:"conditionEvaluation,omitempty"` // The condition evaluation.
	Denyall bool `json:"denyAll,omitempty"` // Setting this to true means that all values are denied. This field can be set only in Policies for list constraints.
	Enforce bool `json:"enforce,omitempty"` // If `true`, then the `Policy` is enforced. If `false`, then any configuration is acceptable. This field can be set only in Policies for boolean constraints.
	Values GoogleCloudAssetV1StringValues `json:"values,omitempty"` // The string values for the list constraints.
}

// Options represents the Options schema from the OpenAPI specification
type Options struct {
	Outputresourceedges bool `json:"outputResourceEdges,omitempty"` // Optional. If true, the result will output the relevant parent/child relationships between resources. Default is false.
	Analyzeserviceaccountimpersonation bool `json:"analyzeServiceAccountImpersonation,omitempty"` // Optional. If true, the response will include access analysis from identities to resources via service account impersonation. This is a very expensive operation, because many derived queries will be executed. We highly recommend you use AssetService.AnalyzeIamPolicyLongrunning RPC instead. For example, if the request analyzes for which resources user A has permission P, and there's an IAM policy states user A has iam.serviceAccounts.getAccessToken permission to a service account SA, and there's another IAM policy states service account SA has permission P to a Google Cloud folder F, then user A potentially has access to the Google Cloud folder F. And those advanced analysis results will be included in AnalyzeIamPolicyResponse.service_account_impersonation_analysis. Another example, if the request analyzes for who has permission P to a Google Cloud folder F, and there's an IAM policy states user A has iam.serviceAccounts.actAs permission to a service account SA, and there's another IAM policy states service account SA has permission P to the Google Cloud folder F, then user A potentially has access to the Google Cloud folder F. And those advanced analysis results will be included in AnalyzeIamPolicyResponse.service_account_impersonation_analysis. Only the following permissions are considered in this analysis: * `iam.serviceAccounts.actAs` * `iam.serviceAccounts.signBlob` * `iam.serviceAccounts.signJwt` * `iam.serviceAccounts.getAccessToken` * `iam.serviceAccounts.getOpenIdToken` * `iam.serviceAccounts.implicitDelegation` Default is false.
	Expandgroups bool `json:"expandGroups,omitempty"` // Optional. If true, the identities section of the result will expand any Google groups appearing in an IAM policy binding. If IamPolicyAnalysisQuery.identity_selector is specified, the identity in the result will be determined by the selector, and this flag is not allowed to set. If true, the default max expansion per group is 1000 for AssetService.AnalyzeIamPolicy][]. Default is false.
	Expandresources bool `json:"expandResources,omitempty"` // Optional. If true and IamPolicyAnalysisQuery.resource_selector is not specified, the resource section of the result will expand any resource attached to an IAM policy to include resources lower in the resource hierarchy. For example, if the request analyzes for which resources user A has permission P, and the results include an IAM policy with P on a Google Cloud folder, the results will also include resources in that folder with permission P. If true and IamPolicyAnalysisQuery.resource_selector is specified, the resource section of the result will expand the specified resource to include resources lower in the resource hierarchy. Only project or lower resources are supported. Folder and organization resources cannot be used together with this option. For example, if the request analyzes for which users have permission P on a Google Cloud project with this option enabled, the results will include all users who have permission P on that project or any lower resource. If true, the default max expansion per resource is 1000 for AssetService.AnalyzeIamPolicy][] and 100000 for AssetService.AnalyzeIamPolicyLongrunning][]. Default is false.
	Expandroles bool `json:"expandRoles,omitempty"` // Optional. If true, the access section of result will expand any roles appearing in IAM policy bindings to include their permissions. If IamPolicyAnalysisQuery.access_selector is specified, the access section of the result will be determined by the selector, and this flag is not allowed to set. Default is false.
	Outputgroupedges bool `json:"outputGroupEdges,omitempty"` // Optional. If true, the result will output the relevant membership relationships between groups and other groups, and between groups and principals. Default is false.
}

// EffectiveIamPolicy represents the EffectiveIamPolicy schema from the OpenAPI specification
type EffectiveIamPolicy struct {
	Fullresourcename string `json:"fullResourceName,omitempty"` // The [full_resource_name] (https://cloud.google.com/asset-inventory/docs/resource-name-format) for which the policies are computed. This is one of the BatchGetEffectiveIamPoliciesRequest.names the caller provides in the request.
	Policies []PolicyInfo `json:"policies,omitempty"` // The effective policies for the full_resource_name. These policies include the policy set on the full_resource_name and those set on its parents and ancestors up to the BatchGetEffectiveIamPoliciesRequest.scope. Note that these policies are not filtered according to the resource type of the full_resource_name. These policies are hierarchically ordered by PolicyInfo.attached_resource starting from full_resource_name itself to its parents and ancestors, such that policies[i]'s PolicyInfo.attached_resource is the child of policies[i+1]'s PolicyInfo.attached_resource, if policies[i+1] exists.
}

// GoogleIdentityAccesscontextmanagerV1VpcSubNetwork represents the GoogleIdentityAccesscontextmanagerV1VpcSubNetwork schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1VpcSubNetwork struct {
	Vpcipsubnetworks []string `json:"vpcIpSubnetworks,omitempty"` // CIDR block IP subnetwork specification. The IP address must be an IPv4 address and can be a public or private IP address. Note that for a CIDR IP address block, the specified IP address portion must be properly truncated (i.e. all the host bits must be zero) or the input is considered malformed. For example, "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. If empty, all IP addresses are allowed.
	Network string `json:"network,omitempty"` // Required. Network name. If the network is not part of the organization, the `compute.network.get` permission must be granted to the caller. Format: `//compute.googleapis.com/projects/{PROJECT_ID}/global/networks/{NETWORK_NAME}` Example: `//compute.googleapis.com/projects/my-project/global/networks/network-1`
}

// BigQueryDestination represents the BigQueryDestination schema from the OpenAPI specification
type BigQueryDestination struct {
	Dataset string `json:"dataset,omitempty"` // Required. The BigQuery dataset in format "projects/projectId/datasets/datasetId", to which the snapshot result should be exported. If this dataset does not exist, the export call returns an INVALID_ARGUMENT error. Setting the `contentType` for `exportAssets` determines the [schema](/asset-inventory/docs/exporting-to-bigquery#bigquery-schema) of the BigQuery table. Setting `separateTablesPerAssetType` to `TRUE` also influences the schema.
	Force bool `json:"force,omitempty"` // If the destination table already exists and this flag is `TRUE`, the table will be overwritten by the contents of assets snapshot. If the flag is `FALSE` or unset and the destination table already exists, the export call returns an INVALID_ARGUMEMT error.
	Partitionspec PartitionSpec `json:"partitionSpec,omitempty"` // Specifications of BigQuery partitioned table as export destination.
	Separatetablesperassettype bool `json:"separateTablesPerAssetType,omitempty"` // If this flag is `TRUE`, the snapshot results will be written to one or multiple tables, each of which contains results of one asset type. The [force] and [partition_spec] fields will apply to each of them. Field [table] will be concatenated with "_" and the asset type names (see https://cloud.google.com/asset-inventory/docs/supported-asset-types for supported asset types) to construct per-asset-type table names, in which all non-alphanumeric characters like "." and "/" will be substituted by "_". Example: if field [table] is "mytable" and snapshot results contain "storage.googleapis.com/Bucket" assets, the corresponding table name will be "mytable_storage_googleapis_com_Bucket". If any of these tables does not exist, a new table with the concatenated name will be created. When [content_type] in the ExportAssetsRequest is `RESOURCE`, the schema of each table will include RECORD-type columns mapped to the nested fields in the Asset.resource.data field of that asset type (up to the 15 nested level BigQuery supports (https://cloud.google.com/bigquery/docs/nested-repeated#limitations)). The fields in >15 nested levels will be stored in JSON format string as a child column of its parent RECORD column. If error occurs when exporting to any table, the whole export call will return an error but the export results that already succeed will persist. Example: if exporting to table_type_A succeeds when exporting to table_type_B fails during one export call, the results in table_type_A will persist and there will not be partial results persisting in a table.
	Table string `json:"table,omitempty"` // Required. The BigQuery table to which the snapshot result should be written. If this table does not exist, a new table with the given name will be created.
}

// Expr represents the Expr schema from the OpenAPI specification
type Expr struct {
	Title string `json:"title,omitempty"` // Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Description string `json:"description,omitempty"` // Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Expression string `json:"expression,omitempty"` // Textual representation of an expression in Common Expression Language syntax.
	Location string `json:"location,omitempty"` // Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
}

// AnalyzeIamPolicyLongrunningRequest represents the AnalyzeIamPolicyLongrunningRequest schema from the OpenAPI specification
type AnalyzeIamPolicyLongrunningRequest struct {
	Savedanalysisquery string `json:"savedAnalysisQuery,omitempty"` // Optional. The name of a saved query, which must be in the format of: * projects/project_number/savedQueries/saved_query_id * folders/folder_number/savedQueries/saved_query_id * organizations/organization_number/savedQueries/saved_query_id If both `analysis_query` and `saved_analysis_query` are provided, they will be merged together with the `saved_analysis_query` as base and the `analysis_query` as overrides. For more details of the merge behavior, refer to the [MergeFrom](https://developers.google.com/protocol-buffers/docs/reference/cpp/google.protobuf.message#Message.MergeFrom.details) doc. Note that you cannot override primitive fields with default value, such as 0 or empty string, etc., because we use proto3, which doesn't support field presence yet.
	Analysisquery IamPolicyAnalysisQuery `json:"analysisQuery,omitempty"` // IAM policy analysis query message.
	Outputconfig IamPolicyAnalysisOutputConfig `json:"outputConfig,omitempty"` // Output configuration for export IAM policy analysis destination.
}

// SavedQuery represents the SavedQuery schema from the OpenAPI specification
type SavedQuery struct {
	Content QueryContent `json:"content,omitempty"` // The query content.
	Createtime string `json:"createTime,omitempty"` // Output only. The create time of this saved query.
	Creator string `json:"creator,omitempty"` // Output only. The account's email address who has created this saved query.
	Description string `json:"description,omitempty"` // The description of this saved query. This value should be fewer than 255 characters.
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels applied on the resource. This value should not contain more than 10 entries. The key and value of each entry must be non-empty and fewer than 64 characters.
	Lastupdatetime string `json:"lastUpdateTime,omitempty"` // Output only. The last update time of this saved query.
	Lastupdater string `json:"lastUpdater,omitempty"` // Output only. The account's email address who has updated this saved query most recently.
	Name string `json:"name,omitempty"` // The resource name of the saved query. The format must be: * projects/project_number/savedQueries/saved_query_id * folders/folder_number/savedQueries/saved_query_id * organizations/organization_number/savedQueries/saved_query_id
}

// GoogleIdentityAccesscontextmanagerV1EgressSource represents the GoogleIdentityAccesscontextmanagerV1EgressSource schema from the OpenAPI specification
type GoogleIdentityAccesscontextmanagerV1EgressSource struct {
	Accesslevel string `json:"accessLevel,omitempty"` // An AccessLevel resource name that allows protected resources inside the ServicePerimeters to access outside the ServicePerimeter boundaries. AccessLevels listed must be in the same policy as this ServicePerimeter. Referencing a nonexistent AccessLevel will cause an error. If an AccessLevel name is not specified, only resources within the perimeter can be accessed through Google Cloud calls with request origins within the perimeter. Example: `accessPolicies/MY_POLICY/accessLevels/MY_LEVEL`. If a single `*` is specified for `access_level`, then all EgressSources will be allowed.
}
