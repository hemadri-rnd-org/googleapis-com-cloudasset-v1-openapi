package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloud-asset-api/mcp-server/config"
	"github.com/cloud-asset-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Cloudasset_analyzeiampolicyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		scopeVal, ok := args["scope"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: scope"), nil
		}
		scope, ok := scopeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: scope"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["analysisQuery.accessSelector.permissions"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("analysisQuery.accessSelector.permissions=%v", val))
		}
		if val, ok := args["analysisQuery.accessSelector.roles"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("analysisQuery.accessSelector.roles=%v", val))
		}
		if val, ok := args["analysisQuery.conditionContext.accessTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("analysisQuery.conditionContext.accessTime=%v", val))
		}
		if val, ok := args["analysisQuery.identitySelector.identity"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("analysisQuery.identitySelector.identity=%v", val))
		}
		if val, ok := args["analysisQuery.options.analyzeServiceAccountImpersonation"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("analysisQuery.options.analyzeServiceAccountImpersonation=%v", val))
		}
		if val, ok := args["analysisQuery.options.expandGroups"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("analysisQuery.options.expandGroups=%v", val))
		}
		if val, ok := args["analysisQuery.options.expandResources"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("analysisQuery.options.expandResources=%v", val))
		}
		if val, ok := args["analysisQuery.options.expandRoles"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("analysisQuery.options.expandRoles=%v", val))
		}
		if val, ok := args["analysisQuery.options.outputGroupEdges"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("analysisQuery.options.outputGroupEdges=%v", val))
		}
		if val, ok := args["analysisQuery.options.outputResourceEdges"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("analysisQuery.options.outputResourceEdges=%v", val))
		}
		if val, ok := args["analysisQuery.resourceSelector.fullResourceName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("analysisQuery.resourceSelector.fullResourceName=%v", val))
		}
		if val, ok := args["executionTimeout"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("executionTimeout=%v", val))
		}
		if val, ok := args["savedAnalysisQuery"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("savedAnalysisQuery=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/%s:analyzeIamPolicy%s", cfg.BaseURL, scope, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.AnalyzeIamPolicyResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCloudasset_analyzeiampolicyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_scope_analyzeIamPolicy",
		mcp.WithDescription("Analyzes IAM policies to answer which identities have what accesses on which resources."),
		mcp.WithString("scope", mcp.Required(), mcp.Description("Required. The relative name of the root asset. Only resources and IAM policies within the scope will be analyzed. This can only be an organization number (such as \"organizations/123\"), a folder number (such as \"folders/123\"), a project ID (such as \"projects/my-project-id\"), or a project number (such as \"projects/12345\"). To know how to get organization ID, visit [here ](https://cloud.google.com/resource-manager/docs/creating-managing-organization#retrieving_your_organization_id). To know how to get folder or project ID, visit [here ](https://cloud.google.com/resource-manager/docs/creating-managing-folders#viewing_or_listing_folders_and_projects).")),
		mcp.WithArray("analysisQuery.accessSelector.permissions", mcp.Description("Optional. The permissions to appear in result.")),
		mcp.WithArray("analysisQuery.accessSelector.roles", mcp.Description("Optional. The roles to appear in result.")),
		mcp.WithString("analysisQuery.conditionContext.accessTime", mcp.Description("The hypothetical access timestamp to evaluate IAM conditions. Note that this value must not be earlier than the current time; otherwise, an INVALID_ARGUMENT error will be returned.")),
		mcp.WithString("analysisQuery.identitySelector.identity", mcp.Description("Required. The identity appear in the form of principals in [IAM policy binding](https://cloud.google.com/iam/reference/rest/v1/Binding). The examples of supported forms are: \"user:mike@example.com\", \"group:admins@example.com\", \"domain:google.com\", \"serviceAccount:my-project-id@appspot.gserviceaccount.com\". Notice that wildcard characters (such as * and ?) are not supported. You must give a specific identity.")),
		mcp.WithBoolean("analysisQuery.options.analyzeServiceAccountImpersonation", mcp.Description("Optional. If true, the response will include access analysis from identities to resources via service account impersonation. This is a very expensive operation, because many derived queries will be executed. We highly recommend you use AssetService.AnalyzeIamPolicyLongrunning RPC instead. For example, if the request analyzes for which resources user A has permission P, and there's an IAM policy states user A has iam.serviceAccounts.getAccessToken permission to a service account SA, and there's another IAM policy states service account SA has permission P to a Google Cloud folder F, then user A potentially has access to the Google Cloud folder F. And those advanced analysis results will be included in AnalyzeIamPolicyResponse.service_account_impersonation_analysis. Another example, if the request analyzes for who has permission P to a Google Cloud folder F, and there's an IAM policy states user A has iam.serviceAccounts.actAs permission to a service account SA, and there's another IAM policy states service account SA has permission P to the Google Cloud folder F, then user A potentially has access to the Google Cloud folder F. And those advanced analysis results will be included in AnalyzeIamPolicyResponse.service_account_impersonation_analysis. Only the following permissions are considered in this analysis: * `iam.serviceAccounts.actAs` * `iam.serviceAccounts.signBlob` * `iam.serviceAccounts.signJwt` * `iam.serviceAccounts.getAccessToken` * `iam.serviceAccounts.getOpenIdToken` * `iam.serviceAccounts.implicitDelegation` Default is false.")),
		mcp.WithBoolean("analysisQuery.options.expandGroups", mcp.Description("Optional. If true, the identities section of the result will expand any Google groups appearing in an IAM policy binding. If IamPolicyAnalysisQuery.identity_selector is specified, the identity in the result will be determined by the selector, and this flag is not allowed to set. If true, the default max expansion per group is 1000 for AssetService.AnalyzeIamPolicy][]. Default is false.")),
		mcp.WithBoolean("analysisQuery.options.expandResources", mcp.Description("Optional. If true and IamPolicyAnalysisQuery.resource_selector is not specified, the resource section of the result will expand any resource attached to an IAM policy to include resources lower in the resource hierarchy. For example, if the request analyzes for which resources user A has permission P, and the results include an IAM policy with P on a Google Cloud folder, the results will also include resources in that folder with permission P. If true and IamPolicyAnalysisQuery.resource_selector is specified, the resource section of the result will expand the specified resource to include resources lower in the resource hierarchy. Only project or lower resources are supported. Folder and organization resources cannot be used together with this option. For example, if the request analyzes for which users have permission P on a Google Cloud project with this option enabled, the results will include all users who have permission P on that project or any lower resource. If true, the default max expansion per resource is 1000 for AssetService.AnalyzeIamPolicy][] and 100000 for AssetService.AnalyzeIamPolicyLongrunning][]. Default is false.")),
		mcp.WithBoolean("analysisQuery.options.expandRoles", mcp.Description("Optional. If true, the access section of result will expand any roles appearing in IAM policy bindings to include their permissions. If IamPolicyAnalysisQuery.access_selector is specified, the access section of the result will be determined by the selector, and this flag is not allowed to set. Default is false.")),
		mcp.WithBoolean("analysisQuery.options.outputGroupEdges", mcp.Description("Optional. If true, the result will output the relevant membership relationships between groups and other groups, and between groups and principals. Default is false.")),
		mcp.WithBoolean("analysisQuery.options.outputResourceEdges", mcp.Description("Optional. If true, the result will output the relevant parent/child relationships between resources. Default is false.")),
		mcp.WithString("analysisQuery.resourceSelector.fullResourceName", mcp.Description("Required. The [full resource name] (https://cloud.google.com/asset-inventory/docs/resource-name-format) of a resource of [supported resource types](https://cloud.google.com/asset-inventory/docs/supported-asset-types#analyzable_asset_types).")),
		mcp.WithString("executionTimeout", mcp.Description("Optional. Amount of time executable has to complete. See JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json). If this field is set with a value less than the RPC deadline, and the execution of your query hasn't finished in the specified execution timeout, you will get a response with partial result. Otherwise, your query's execution will continue until the RPC deadline. If it's not finished until then, you will get a DEADLINE_EXCEEDED error. Default is empty.")),
		mcp.WithString("savedAnalysisQuery", mcp.Description("Optional. The name of a saved query, which must be in the format of: * projects/project_number/savedQueries/saved_query_id * folders/folder_number/savedQueries/saved_query_id * organizations/organization_number/savedQueries/saved_query_id If both `analysis_query` and `saved_analysis_query` are provided, they will be merged together with the `saved_analysis_query` as base and the `analysis_query` as overrides. For more details of the merge behavior, refer to the [MergeFrom](https://developers.google.com/protocol-buffers/docs/reference/cpp/google.protobuf.message#Message.MergeFrom.details) page. Note that you cannot override primitive fields with default value, such as 0 or empty string, etc., because we use proto3, which doesn't support field presence yet.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Cloudasset_analyzeiampolicyHandler(cfg),
	}
}
