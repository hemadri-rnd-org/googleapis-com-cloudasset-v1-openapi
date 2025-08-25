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

func Cloudasset_analyzeorgpolicygovernedcontainersHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["constraint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("constraint=%v", val))
		}
		if val, ok := args["filter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter=%v", val))
		}
		if val, ok := args["pageSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageSize=%v", val))
		}
		if val, ok := args["pageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageToken=%v", val))
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
		url := fmt.Sprintf("%s/v1/%s:analyzeOrgPolicyGovernedContainers%s", cfg.BaseURL, scope, queryString)
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
		var result models.AnalyzeOrgPolicyGovernedContainersResponse
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

func CreateCloudasset_analyzeorgpolicygovernedcontainersTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_scope_analyzeOrgPolicyGovernedContainers",
		mcp.WithDescription("Analyzes organization policies governed containers (projects, folders or organization) under a scope."),
		mcp.WithString("scope", mcp.Required(), mcp.Description("Required. The organization to scope the request. Only organization policies within the scope will be analyzed. The output containers will also be limited to the ones governed by those in-scope organization policies. * organizations/{ORGANIZATION_NUMBER} (e.g., \"organizations/123456\")")),
		mcp.WithString("constraint", mcp.Description("Required. The name of the constraint to analyze governed containers for. The analysis only contains organization policies for the provided constraint.")),
		mcp.WithString("filter", mcp.Description("The expression to filter AnalyzeOrgPolicyGovernedContainersResponse.governed_containers. Filtering is currently available for bare literal values and the following fields: * parent * consolidated_policy.rules.enforce When filtering by a specific field, the only supported operator is `=`. For example, filtering by parent=\"//cloudresourcemanager.googleapis.com/folders/001\" will return all the containers under \"folders/001\".")),
		mcp.WithNumber("pageSize", mcp.Description("The maximum number of items to return per page. If unspecified, AnalyzeOrgPolicyGovernedContainersResponse.governed_containers will contain 100 items with a maximum of 200.")),
		mcp.WithString("pageToken", mcp.Description("The pagination token to retrieve the next page.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Cloudasset_analyzeorgpolicygovernedcontainersHandler(cfg),
	}
}
