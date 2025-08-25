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

func Cloudasset_batchgetassetshistoryHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		parentVal, ok := args["parent"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: parent"), nil
		}
		parent, ok := parentVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: parent"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["assetNames"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("assetNames=%v", val))
		}
		if val, ok := args["contentType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("contentType=%v", val))
		}
		if val, ok := args["readTimeWindow.endTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("readTimeWindow.endTime=%v", val))
		}
		if val, ok := args["readTimeWindow.startTime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("readTimeWindow.startTime=%v", val))
		}
		if val, ok := args["relationshipTypes"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("relationshipTypes=%v", val))
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
		url := fmt.Sprintf("%s/v1/%s:batchGetAssetsHistory%s", cfg.BaseURL, parent, queryString)
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
		var result models.BatchGetAssetsHistoryResponse
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

func CreateCloudasset_batchgetassetshistoryTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_parent_batchGetAssetsHistory",
		mcp.WithDescription("Batch gets the update history of assets that overlap a time window. For IAM_POLICY content, this API outputs history when the asset and its attached IAM POLICY both exist. This can create gaps in the output history. Otherwise, this API outputs history with asset in both non-delete or deleted status. If a specified asset does not exist, this API returns an INVALID_ARGUMENT error."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. The relative name of the root asset. It can only be an organization number (such as \"organizations/123\"), a project ID (such as \"projects/my-project-id\")\", or a project number (such as \"projects/12345\").")),
		mcp.WithArray("assetNames", mcp.Description("A list of the full names of the assets. See: https://cloud.google.com/asset-inventory/docs/resource-name-format Example: `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`. The request becomes a no-op if the asset name list is empty, and the max size of the asset name list is 100 in one request.")),
		mcp.WithString("contentType", mcp.Description("Optional. The content type.")),
		mcp.WithString("readTimeWindow.endTime", mcp.Description("End time of the time window (inclusive). If not specified, the current timestamp is used instead.")),
		mcp.WithString("readTimeWindow.startTime", mcp.Description("Start time of the time window (exclusive).")),
		mcp.WithArray("relationshipTypes", mcp.Description("Optional. A list of relationship types to output, for example: `INSTANCE_TO_INSTANCEGROUP`. This field should only be specified if content_type=RELATIONSHIP. * If specified: it outputs specified relationships' history on the [asset_names]. It returns an error if any of the [relationship_types] doesn't belong to the supported relationship types of the [asset_names] or if any of the [asset_names]'s types doesn't belong to the source types of the [relationship_types]. * Otherwise: it outputs the supported relationships' history on the [asset_names] or returns an error if any of the [asset_names]'s types has no relationship support. See [Introduction to Cloud Asset Inventory](https://cloud.google.com/asset-inventory/docs/overview) for all supported asset types and relationship types.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Cloudasset_batchgetassetshistoryHandler(cfg),
	}
}
