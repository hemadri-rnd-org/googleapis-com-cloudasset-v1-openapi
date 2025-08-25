package main

import (
	"github.com/cloud-asset-api/mcp-server/config"
	"github.com/cloud-asset-api/mcp-server/models"
	tools_v1 "github.com/cloud-asset-api/mcp-server/tools/v1"
	tools_assets "github.com/cloud-asset-api/mcp-server/tools/assets"
	tools_feeds "github.com/cloud-asset-api/mcp-server/tools/feeds"
	tools_effectiveiampolicies "github.com/cloud-asset-api/mcp-server/tools/effectiveiampolicies"
	tools_savedqueries "github.com/cloud-asset-api/mcp-server/tools/savedqueries"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_v1.CreateCloudasset_analyzeiampolicylongrunningTool(cfg),
		tools_v1.CreateCloudasset_analyzeiampolicyTool(cfg),
		tools_v1.CreateCloudasset_analyzeorgpolicygovernedassetsTool(cfg),
		tools_assets.CreateCloudasset_assets_listTool(cfg),
		tools_feeds.CreateCloudasset_feeds_listTool(cfg),
		tools_feeds.CreateCloudasset_feeds_createTool(cfg),
		tools_v1.CreateCloudasset_batchgetassetshistoryTool(cfg),
		tools_v1.CreateCloudasset_exportassetsTool(cfg),
		tools_v1.CreateCloudasset_analyzemoveTool(cfg),
		tools_effectiveiampolicies.CreateCloudasset_effectiveiampolicies_batchgetTool(cfg),
		tools_savedqueries.CreateCloudasset_savedqueries_listTool(cfg),
		tools_savedqueries.CreateCloudasset_savedqueries_createTool(cfg),
		tools_v1.CreateCloudasset_searchallresourcesTool(cfg),
		tools_v1.CreateCloudasset_analyzeorgpoliciesTool(cfg),
		tools_savedqueries.CreateCloudasset_savedqueries_deleteTool(cfg),
		tools_savedqueries.CreateCloudasset_savedqueries_getTool(cfg),
		tools_savedqueries.CreateCloudasset_savedqueries_patchTool(cfg),
		tools_v1.CreateCloudasset_queryassetsTool(cfg),
		tools_v1.CreateCloudasset_searchalliampoliciesTool(cfg),
		tools_v1.CreateCloudasset_analyzeorgpolicygovernedcontainersTool(cfg),
	}
}
