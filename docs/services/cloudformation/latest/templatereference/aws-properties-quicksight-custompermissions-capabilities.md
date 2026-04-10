This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::CustomPermissions Capabilities

A set of actions that correspond to Amazon Quick Sight permissions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "AddOrRunAnomalyDetectionForAnalyses" : String,
  "AmazonBedrockARSAction" : String,
  "AmazonBedrockFSAction" : String,
  "AmazonBedrockKRSAction" : String,
  "AmazonSThreeAction" : String,
  "Analysis" : String,
  "ApproveFlowShareRequests" : String,
  "AsanaAction" : String,
  "Automate" : String,
  "BambooHRAction" : String,
  "BoxAgentAction" : String,
  "BuildCalculatedFieldWithQ" : String,
  "CanvaAgentAction" : String,
  "ChatAgent" : String,
  "ComprehendAction" : String,
  "ComprehendMedicalAction" : String,
  "ConfluenceAction" : String,
  "CreateAndUpdateAmazonBedrockARSAction" : String,
  "CreateAndUpdateAmazonBedrockFSAction" : String,
  "CreateAndUpdateAmazonBedrockKRSAction" : String,
  "CreateAndUpdateAmazonSThreeAction" : String,
  "CreateAndUpdateAsanaAction" : String,
  "CreateAndUpdateBambooHRAction" : String,
  "CreateAndUpdateBoxAgentAction" : String,
  "CreateAndUpdateCanvaAgentAction" : String,
  "CreateAndUpdateComprehendAction" : String,
  "CreateAndUpdateComprehendMedicalAction" : String,
  "CreateAndUpdateConfluenceAction" : String,
  "CreateAndUpdateDashboardEmailReports" : String,
  "CreateAndUpdateDatasets" : String,
  "CreateAndUpdateDataSources" : String,
  "CreateAndUpdateFactSetAction" : String,
  "CreateAndUpdateGenericHTTPAction" : String,
  "CreateAndUpdateGithubAction" : String,
  "CreateAndUpdateGoogleCalendarAction" : String,
  "CreateAndUpdateHubspotAction" : String,
  "CreateAndUpdateHuggingFaceAction" : String,
  "CreateAndUpdateIntercomAction" : String,
  "CreateAndUpdateJiraAction" : String,
  "CreateAndUpdateKnowledgeBases" : String,
  "CreateAndUpdateLinearAction" : String,
  "CreateAndUpdateMCPAction" : String,
  "CreateAndUpdateMondayAction" : String,
  "CreateAndUpdateMSExchangeAction" : String,
  "CreateAndUpdateMSTeamsAction" : String,
  "CreateAndUpdateNewRelicAction" : String,
  "CreateAndUpdateNotionAction" : String,
  "CreateAndUpdateOneDriveAction" : String,
  "CreateAndUpdateOpenAPIAction" : String,
  "CreateAndUpdatePagerDutyAction" : String,
  "CreateAndUpdateSalesforceAction" : String,
  "CreateAndUpdateSandPGlobalEnergyAction" : String,
  "CreateAndUpdateSandPGMIAction" : String,
  "CreateAndUpdateSAPBillOfMaterialAction" : String,
  "CreateAndUpdateSAPBusinessPartnerAction" : String,
  "CreateAndUpdateSAPMaterialStockAction" : String,
  "CreateAndUpdateSAPPhysicalInventoryAction" : String,
  "CreateAndUpdateSAPProductMasterDataAction" : String,
  "CreateAndUpdateServiceNowAction" : String,
  "CreateAndUpdateSharePointAction" : String,
  "CreateAndUpdateSlackAction" : String,
  "CreateAndUpdateSmartsheetAction" : String,
  "CreateAndUpdateTextractAction" : String,
  "CreateAndUpdateThemes" : String,
  "CreateAndUpdateThresholdAlerts" : String,
  "CreateAndUpdateZendeskAction" : String,
  "CreateChatAgents" : String,
  "CreateDashboardExecutiveSummaryWithQ" : String,
  "CreateSharedFolders" : String,
  "CreateSPICEDataset" : String,
  "Dashboard" : String,
  "EditVisualWithQ" : String,
  "ExportToCsv" : String,
  "ExportToCsvInScheduledReports" : String,
  "ExportToExcel" : String,
  "ExportToExcelInScheduledReports" : String,
  "ExportToPdf" : String,
  "ExportToPdfInScheduledReports" : String,
  "Extension" : String,
  "FactSetAction" : String,
  "Flow" : String,
  "GenericHTTPAction" : String,
  "GithubAction" : String,
  "GoogleCalendarAction" : String,
  "HubspotAction" : String,
  "HuggingFaceAction" : String,
  "IncludeContentInScheduledReportsEmail" : String,
  "IntercomAction" : String,
  "JiraAction" : String,
  "KnowledgeBase" : String,
  "LinearAction" : String,
  "ManageSharedFolders" : String,
  "MCPAction" : String,
  "MondayAction" : String,
  "MSExchangeAction" : String,
  "MSTeamsAction" : String,
  "NewRelicAction" : String,
  "NotionAction" : String,
  "OneDriveAction" : String,
  "OpenAPIAction" : String,
  "PagerDutyAction" : String,
  "PerformFlowUiTask" : String,
  "PrintReports" : String,
  "PublishWithoutApproval" : String,
  "RenameSharedFolders" : String,
  "Research" : String,
  "SalesforceAction" : String,
  "SandPGlobalEnergyAction" : String,
  "SandPGMIAction" : String,
  "SAPBillOfMaterialAction" : String,
  "SAPBusinessPartnerAction" : String,
  "SAPMaterialStockAction" : String,
  "SAPPhysicalInventoryAction" : String,
  "SAPProductMasterDataAction" : String,
  "ServiceNowAction" : String,
  "ShareAmazonBedrockARSAction" : String,
  "ShareAmazonBedrockFSAction" : String,
  "ShareAmazonBedrockKRSAction" : String,
  "ShareAmazonSThreeAction" : String,
  "ShareAnalyses" : String,
  "ShareAsanaAction" : String,
  "ShareBambooHRAction" : String,
  "ShareBoxAgentAction" : String,
  "ShareCanvaAgentAction" : String,
  "ShareComprehendAction" : String,
  "ShareComprehendMedicalAction" : String,
  "ShareConfluenceAction" : String,
  "ShareDashboards" : String,
  "ShareDatasets" : String,
  "ShareDataSources" : String,
  "ShareFactSetAction" : String,
  "ShareGenericHTTPAction" : String,
  "ShareGithubAction" : String,
  "ShareGoogleCalendarAction" : String,
  "ShareHubspotAction" : String,
  "ShareHuggingFaceAction" : String,
  "ShareIntercomAction" : String,
  "ShareJiraAction" : String,
  "ShareKnowledgeBases" : String,
  "ShareLinearAction" : String,
  "ShareMCPAction" : String,
  "ShareMondayAction" : String,
  "ShareMSExchangeAction" : String,
  "ShareMSTeamsAction" : String,
  "ShareNewRelicAction" : String,
  "ShareNotionAction" : String,
  "ShareOneDriveAction" : String,
  "ShareOpenAPIAction" : String,
  "SharePagerDutyAction" : String,
  "SharePointAction" : String,
  "ShareSalesforceAction" : String,
  "ShareSandPGlobalEnergyAction" : String,
  "ShareSandPGMIAction" : String,
  "ShareSAPBillOfMaterialAction" : String,
  "ShareSAPBusinessPartnerAction" : String,
  "ShareSAPMaterialStockAction" : String,
  "ShareSAPPhysicalInventoryAction" : String,
  "ShareSAPProductMasterDataAction" : String,
  "ShareServiceNowAction" : String,
  "ShareSharePointAction" : String,
  "ShareSlackAction" : String,
  "ShareSmartsheetAction" : String,
  "ShareTextractAction" : String,
  "ShareZendeskAction" : String,
  "SlackAction" : String,
  "SmartsheetAction" : String,
  "Space" : String,
  "SubscribeDashboardEmailReports" : String,
  "TextractAction" : String,
  "Topic" : String,
  "UseAgentWebSearch" : String,
  "UseAmazonBedrockARSAction" : String,
  "UseAmazonBedrockFSAction" : String,
  "UseAmazonBedrockKRSAction" : String,
  "UseAmazonSThreeAction" : String,
  "UseAsanaAction" : String,
  "UseBambooHRAction" : String,
  "UseBedrockModels" : String,
  "UseBoxAgentAction" : String,
  "UseCanvaAgentAction" : String,
  "UseComprehendAction" : String,
  "UseComprehendMedicalAction" : String,
  "UseConfluenceAction" : String,
  "UseFactSetAction" : String,
  "UseGenericHTTPAction" : String,
  "UseGithubAction" : String,
  "UseGoogleCalendarAction" : String,
  "UseHubspotAction" : String,
  "UseHuggingFaceAction" : String,
  "UseIntercomAction" : String,
  "UseJiraAction" : String,
  "UseLinearAction" : String,
  "UseMCPAction" : String,
  "UseMondayAction" : String,
  "UseMSExchangeAction" : String,
  "UseMSTeamsAction" : String,
  "UseNewRelicAction" : String,
  "UseNotionAction" : String,
  "UseOneDriveAction" : String,
  "UseOpenAPIAction" : String,
  "UsePagerDutyAction" : String,
  "UseSalesforceAction" : String,
  "UseSandPGlobalEnergyAction" : String,
  "UseSandPGMIAction" : String,
  "UseSAPBillOfMaterialAction" : String,
  "UseSAPBusinessPartnerAction" : String,
  "UseSAPMaterialStockAction" : String,
  "UseSAPPhysicalInventoryAction" : String,
  "UseSAPProductMasterDataAction" : String,
  "UseServiceNowAction" : String,
  "UseSharePointAction" : String,
  "UseSlackAction" : String,
  "UseSmartsheetAction" : String,
  "UseTextractAction" : String,
  "UseZendeskAction" : String,
  "ViewAccountSPICECapacity" : String,
  "ZendeskAction" : String
}

```

### YAML

```yaml

  Action: String
  AddOrRunAnomalyDetectionForAnalyses: String
  AmazonBedrockARSAction: String
  AmazonBedrockFSAction: String
  AmazonBedrockKRSAction: String
  AmazonSThreeAction: String
  Analysis: String
  ApproveFlowShareRequests: String
  AsanaAction: String
  Automate: String
  BambooHRAction: String
  BoxAgentAction: String
  BuildCalculatedFieldWithQ: String
  CanvaAgentAction: String
  ChatAgent: String
  ComprehendAction: String
  ComprehendMedicalAction: String
  ConfluenceAction: String
  CreateAndUpdateAmazonBedrockARSAction: String
  CreateAndUpdateAmazonBedrockFSAction: String
  CreateAndUpdateAmazonBedrockKRSAction: String
  CreateAndUpdateAmazonSThreeAction: String
  CreateAndUpdateAsanaAction: String
  CreateAndUpdateBambooHRAction: String
  CreateAndUpdateBoxAgentAction: String
  CreateAndUpdateCanvaAgentAction: String
  CreateAndUpdateComprehendAction: String
  CreateAndUpdateComprehendMedicalAction: String
  CreateAndUpdateConfluenceAction: String
  CreateAndUpdateDashboardEmailReports: String
  CreateAndUpdateDatasets: String
  CreateAndUpdateDataSources: String
  CreateAndUpdateFactSetAction: String
  CreateAndUpdateGenericHTTPAction: String
  CreateAndUpdateGithubAction: String
  CreateAndUpdateGoogleCalendarAction: String
  CreateAndUpdateHubspotAction: String
  CreateAndUpdateHuggingFaceAction: String
  CreateAndUpdateIntercomAction: String
  CreateAndUpdateJiraAction: String
  CreateAndUpdateKnowledgeBases: String
  CreateAndUpdateLinearAction: String
  CreateAndUpdateMCPAction: String
  CreateAndUpdateMondayAction: String
  CreateAndUpdateMSExchangeAction: String
  CreateAndUpdateMSTeamsAction: String
  CreateAndUpdateNewRelicAction: String
  CreateAndUpdateNotionAction: String
  CreateAndUpdateOneDriveAction: String
  CreateAndUpdateOpenAPIAction: String
  CreateAndUpdatePagerDutyAction: String
  CreateAndUpdateSalesforceAction: String
  CreateAndUpdateSandPGlobalEnergyAction: String
  CreateAndUpdateSandPGMIAction: String
  CreateAndUpdateSAPBillOfMaterialAction: String
  CreateAndUpdateSAPBusinessPartnerAction: String
  CreateAndUpdateSAPMaterialStockAction: String
  CreateAndUpdateSAPPhysicalInventoryAction: String
  CreateAndUpdateSAPProductMasterDataAction: String
  CreateAndUpdateServiceNowAction: String
  CreateAndUpdateSharePointAction: String
  CreateAndUpdateSlackAction: String
  CreateAndUpdateSmartsheetAction: String
  CreateAndUpdateTextractAction: String
  CreateAndUpdateThemes: String
  CreateAndUpdateThresholdAlerts: String
  CreateAndUpdateZendeskAction: String
  CreateChatAgents: String
  CreateDashboardExecutiveSummaryWithQ: String
  CreateSharedFolders: String
  CreateSPICEDataset: String
  Dashboard: String
  EditVisualWithQ: String
  ExportToCsv: String
  ExportToCsvInScheduledReports: String
  ExportToExcel: String
  ExportToExcelInScheduledReports: String
  ExportToPdf: String
  ExportToPdfInScheduledReports: String
  Extension: String
  FactSetAction: String
  Flow: String
  GenericHTTPAction: String
  GithubAction: String
  GoogleCalendarAction: String
  HubspotAction: String
  HuggingFaceAction: String
  IncludeContentInScheduledReportsEmail: String
  IntercomAction: String
  JiraAction: String
  KnowledgeBase: String
  LinearAction: String
  ManageSharedFolders: String
  MCPAction: String
  MondayAction: String
  MSExchangeAction: String
  MSTeamsAction: String
  NewRelicAction: String
  NotionAction: String
  OneDriveAction: String
  OpenAPIAction: String
  PagerDutyAction: String
  PerformFlowUiTask: String
  PrintReports: String
  PublishWithoutApproval: String
  RenameSharedFolders: String
  Research: String
  SalesforceAction: String
  SandPGlobalEnergyAction: String
  SandPGMIAction: String
  SAPBillOfMaterialAction: String
  SAPBusinessPartnerAction: String
  SAPMaterialStockAction: String
  SAPPhysicalInventoryAction: String
  SAPProductMasterDataAction: String
  ServiceNowAction: String
  ShareAmazonBedrockARSAction: String
  ShareAmazonBedrockFSAction: String
  ShareAmazonBedrockKRSAction: String
  ShareAmazonSThreeAction: String
  ShareAnalyses: String
  ShareAsanaAction: String
  ShareBambooHRAction: String
  ShareBoxAgentAction: String
  ShareCanvaAgentAction: String
  ShareComprehendAction: String
  ShareComprehendMedicalAction: String
  ShareConfluenceAction: String
  ShareDashboards: String
  ShareDatasets: String
  ShareDataSources: String
  ShareFactSetAction: String
  ShareGenericHTTPAction: String
  ShareGithubAction: String
  ShareGoogleCalendarAction: String
  ShareHubspotAction: String
  ShareHuggingFaceAction: String
  ShareIntercomAction: String
  ShareJiraAction: String
  ShareKnowledgeBases: String
  ShareLinearAction: String
  ShareMCPAction: String
  ShareMondayAction: String
  ShareMSExchangeAction: String
  ShareMSTeamsAction: String
  ShareNewRelicAction: String
  ShareNotionAction: String
  ShareOneDriveAction: String
  ShareOpenAPIAction: String
  SharePagerDutyAction: String
  SharePointAction: String
  ShareSalesforceAction: String
  ShareSandPGlobalEnergyAction: String
  ShareSandPGMIAction: String
  ShareSAPBillOfMaterialAction: String
  ShareSAPBusinessPartnerAction: String
  ShareSAPMaterialStockAction: String
  ShareSAPPhysicalInventoryAction: String
  ShareSAPProductMasterDataAction: String
  ShareServiceNowAction: String
  ShareSharePointAction: String
  ShareSlackAction: String
  ShareSmartsheetAction: String
  ShareTextractAction: String
  ShareZendeskAction: String
  SlackAction: String
  SmartsheetAction: String
  Space: String
  SubscribeDashboardEmailReports: String
  TextractAction: String
  Topic: String
  UseAgentWebSearch: String
  UseAmazonBedrockARSAction: String
  UseAmazonBedrockFSAction: String
  UseAmazonBedrockKRSAction: String
  UseAmazonSThreeAction: String
  UseAsanaAction: String
  UseBambooHRAction: String
  UseBedrockModels: String
  UseBoxAgentAction: String
  UseCanvaAgentAction: String
  UseComprehendAction: String
  UseComprehendMedicalAction: String
  UseConfluenceAction: String
  UseFactSetAction: String
  UseGenericHTTPAction: String
  UseGithubAction: String
  UseGoogleCalendarAction: String
  UseHubspotAction: String
  UseHuggingFaceAction: String
  UseIntercomAction: String
  UseJiraAction: String
  UseLinearAction: String
  UseMCPAction: String
  UseMondayAction: String
  UseMSExchangeAction: String
  UseMSTeamsAction: String
  UseNewRelicAction: String
  UseNotionAction: String
  UseOneDriveAction: String
  UseOpenAPIAction: String
  UsePagerDutyAction: String
  UseSalesforceAction: String
  UseSandPGlobalEnergyAction: String
  UseSandPGMIAction: String
  UseSAPBillOfMaterialAction: String
  UseSAPBusinessPartnerAction: String
  UseSAPMaterialStockAction: String
  UseSAPPhysicalInventoryAction: String
  UseSAPProductMasterDataAction: String
  UseServiceNowAction: String
  UseSharePointAction: String
  UseSlackAction: String
  UseSmartsheetAction: String
  UseTextractAction: String
  UseZendeskAction: String
  ViewAccountSPICECapacity: String
  ZendeskAction: String

```

## Properties

`Action`

The ability to perform actions in external services through Action connectors. Actions allow users to interact with third-party systems.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AddOrRunAnomalyDetectionForAnalyses`

The ability to add or run anomaly detection.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AmazonBedrockARSAction`

The ability to perform actions using Bedrock Agent connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AmazonBedrockFSAction`

The ability to perform actions using Bedrock Runtime connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AmazonBedrockKRSAction`

The ability to perform actions using Bedrock Data Automation Runtime connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AmazonSThreeAction`

The ability to perform actions using Amazon S3 connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Analysis`

The ability to perform analysis-related actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApproveFlowShareRequests`

The ability to review and approve sharing requests of Flows.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AsanaAction`

The ability to perform actions using Asana connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Automate`

The ability to perform automate-related actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BambooHRAction`

The ability to perform actions using BambooHR connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BoxAgentAction`

The ability to perform actions using Box Agent connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BuildCalculatedFieldWithQ`

The ability to Build Calculation with AI

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CanvaAgentAction`

The ability to perform actions using Canva Agent connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChatAgent`

The ability to perform chat-related actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComprehendAction`

The ability to perform actions using Comprehend connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComprehendMedicalAction`

The ability to perform actions using Comprehend Medical connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfluenceAction`

The ability to perform actions using Atlassian Confluence Cloud connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateAmazonBedrockARSAction`

The ability to create and update Bedrock Agent actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateAmazonBedrockFSAction`

The ability to create and update Bedrock Runtime actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateAmazonBedrockKRSAction`

The ability to create and update Bedrock Data Automation Runtime actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateAmazonSThreeAction`

The ability to create and update Amazon S3 actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateAsanaAction`

The ability to create and update Asana actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateBambooHRAction`

The ability to create and update BambooHR actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateBoxAgentAction`

The ability to create and update Box Agent actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateCanvaAgentAction`

The ability to create and update Canva Agent actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateComprehendAction`

The ability to create and update Comprehend actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateComprehendMedicalAction`

The ability to create and update Comprehend Medical actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateConfluenceAction`

The ability to create and update Atlassian Confluence Cloud actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateDashboardEmailReports`

The ability to create and update email reports.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateDatasets`

The ability to create and update datasets.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateDataSources`

The ability to create and update data sources.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateFactSetAction`

The ability to create and update FactSet actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateGenericHTTPAction`

The ability to create and update REST API connection actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateGithubAction`

The ability to create and update GitHub actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateGoogleCalendarAction`

The ability to create and update Google Calendar actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateHubspotAction`

The ability to create and update Hubspot actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateHuggingFaceAction`

The ability to create and update HuggingFace actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateIntercomAction`

The ability to create and update Intercom actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateJiraAction`

The ability to create and update Jira actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateKnowledgeBases`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateLinearAction`

The ability to create and update Linear actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateMCPAction`

The ability to create and update Model Context Protocol actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateMondayAction`

The ability to create and update Monday actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateMSExchangeAction`

The ability to create and update Microsoft Outlook actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateMSTeamsAction`

The ability to create and update Microsoft Teams actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateNewRelicAction`

The ability to create and update New Relic actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateNotionAction`

The ability to create and update Notion actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateOneDriveAction`

The ability to create and update Microsoft OneDrive actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateOpenAPIAction`

The ability to create and update OpenAPI Specification actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdatePagerDutyAction`

The ability to create and update PagerDuty Advance actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateSalesforceAction`

The ability to create and update Salesforce actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateSandPGlobalEnergyAction`

The ability to create and update S&P Global Energy actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateSandPGMIAction`

The ability to create and update S&P Global Market Intelligence actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateSAPBillOfMaterialAction`

The ability to create and update SAP Bill of Materials actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateSAPBusinessPartnerAction`

The ability to create and update SAP Business Partner actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateSAPMaterialStockAction`

The ability to create and update SAP Material Stock actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateSAPPhysicalInventoryAction`

The ability to create and update SAP Physical Inventory actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateSAPProductMasterDataAction`

The ability to create and update SAP Product Master actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateServiceNowAction`

The ability to create and update ServiceNow actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateSharePointAction`

The ability to create and update Microsoft SharePoint Online actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateSlackAction`

The ability to create and update Slack actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateSmartsheetAction`

The ability to create and update Smartsheet actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateTextractAction`

The ability to create and update Textract actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateThemes`

The ability to export to Create and Update themes.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateThresholdAlerts`

The ability to create and update threshold alerts.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateAndUpdateZendeskAction`

The ability to create and update Zendesk actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateChatAgents`

The ability to create chat agents.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateDashboardExecutiveSummaryWithQ`

The ability to Create Executive Summary

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateSharedFolders`

The ability to create shared folders.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateSPICEDataset`

The ability to create a SPICE dataset.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dashboard`

The ability to perform dashboard-related actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EditVisualWithQ`

The ability to Edit Visual with AI

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportToCsv`

The ability to export to CSV files from the UI.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportToCsvInScheduledReports`

The ability to export to CSV files in scheduled email reports.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportToExcel`

The ability to export to Excel files from the UI.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportToExcelInScheduledReports`

The ability to export to Excel files in scheduled email reports.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportToPdf`

The ability to export to PDF files from the UI.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportToPdfInScheduledReports`

The ability to export to PDF files in scheduled email reports.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Extension`

The ability to perform Extension-related actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FactSetAction`

The ability to perform actions using FactSet connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Flow`

The ability to perform flow-related actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GenericHTTPAction`

The ability to perform actions using REST API connection connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GithubAction`

The ability to perform actions using GitHub connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GoogleCalendarAction`

The ability to perform actions using Google Calendar connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HubspotAction`

The ability to perform actions using Hubspot connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HuggingFaceAction`

The ability to perform actions using HuggingFace connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeContentInScheduledReportsEmail`

The ability to include content in scheduled email reports.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntercomAction`

The ability to perform actions using Intercom connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JiraAction`

The ability to perform actions using Jira connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KnowledgeBase`

The ability to use knowledge bases to specify content from external applications.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LinearAction`

The ability to perform actions using Linear connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManageSharedFolders`

The ability to create, update, delete and view shared folders (both restricted and unrestricted), ability to add any asset to shared folders, and ability to share the folders.

**Note:** This does _not_ prevent inheriting access to assets that others share with them through folder membership.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MCPAction`

The ability to perform actions using Model Context Protocol connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MondayAction`

The ability to perform actions using Monday connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MSExchangeAction`

The ability to perform actions using Microsoft Outlook connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MSTeamsAction`

The ability to perform actions using Microsoft Teams connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NewRelicAction`

The ability to perform actions using New Relic connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotionAction`

The ability to perform actions using Notion connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OneDriveAction`

The ability to perform actions using Microsoft OneDrive connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenAPIAction`

The ability to perform actions using OpenAPI Specification connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PagerDutyAction`

The ability to perform actions using PagerDuty Advance connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerformFlowUiTask`

The ability to use UI Agent step to perform tasks on public websites.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrintReports`

The ability to print reports.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublishWithoutApproval`

The ability to enable approvals for flow share.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RenameSharedFolders`

The ability to rename shared folders.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Research`

The ability to perform research-related actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SalesforceAction`

The ability to perform actions using Salesforce connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SandPGlobalEnergyAction`

The ability to perform actions using S&P Global Energy connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SandPGMIAction`

The ability to perform actions using S&P Global Market Intelligence connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPBillOfMaterialAction`

The ability to perform actions using SAP Bill of Materials connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPBusinessPartnerAction`

The ability to perform actions using SAP Business Partner connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPMaterialStockAction`

The ability to perform actions using SAP Material Stock connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPPhysicalInventoryAction`

The ability to perform actions using SAP Physical Inventory connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPProductMasterDataAction`

The ability to perform actions using SAP Product Master connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNowAction`

The ability to perform actions using ServiceNow connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareAmazonBedrockARSAction`

The ability to share Bedrock Agent actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareAmazonBedrockFSAction`

The ability to share Bedrock Runtime actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareAmazonBedrockKRSAction`

The ability to share Bedrock Data Automation Runtime actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareAmazonSThreeAction`

The ability to share Amazon S3 actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareAnalyses`

The ability to share analyses.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareAsanaAction`

The ability to share Asana actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareBambooHRAction`

The ability to share BambooHR actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareBoxAgentAction`

The ability to share Box Agent actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareCanvaAgentAction`

The ability to share Canva Agent actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareComprehendAction`

The ability to share Comprehend actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareComprehendMedicalAction`

The ability to share Comprehend Medical actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareConfluenceAction`

The ability to share Atlassian Confluence Cloud actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareDashboards`

The ability to share dashboards.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareDatasets`

The ability to share datasets.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareDataSources`

The ability to share data sources.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareFactSetAction`

The ability to share FactSet actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareGenericHTTPAction`

The ability to share REST API connection actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareGithubAction`

The ability to share GitHub actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareGoogleCalendarAction`

The ability to share Google Calendar actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareHubspotAction`

The ability to share Hubspot actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareHuggingFaceAction`

The ability to share HuggingFace actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareIntercomAction`

The ability to share Intercom actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareJiraAction`

The ability to share Jira actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareKnowledgeBases`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareLinearAction`

The ability to share Linear actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareMCPAction`

The ability to share Model Context Protocol actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareMondayAction`

The ability to share Monday actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareMSExchangeAction`

The ability to share Microsoft Outlook actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareMSTeamsAction`

The ability to share Microsoft Teams actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareNewRelicAction`

The ability to share New Relic actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareNotionAction`

The ability to share Notion actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareOneDriveAction`

The ability to share Microsoft OneDrive actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareOpenAPIAction`

The ability to share OpenAPI Specification actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SharePagerDutyAction`

The ability to share PagerDuty Advance actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SharePointAction`

The ability to perform actions using Microsoft SharePoint Online connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareSalesforceAction`

The ability to share Salesforce actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareSandPGlobalEnergyAction`

The ability to share S&P Global Energy actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareSandPGMIAction`

The ability to share S&P Global Market Intelligence actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareSAPBillOfMaterialAction`

The ability to share SAP Bill of Materials actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareSAPBusinessPartnerAction`

The ability to share SAP Business Partner actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareSAPMaterialStockAction`

The ability to share SAP Material Stock actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareSAPPhysicalInventoryAction`

The ability to share SAP Physical Inventory actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareSAPProductMasterDataAction`

The ability to share SAP Product Master actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareServiceNowAction`

The ability to share ServiceNow actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareSharePointAction`

The ability to share Microsoft SharePoint Online actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareSlackAction`

The ability to share Slack actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareSmartsheetAction`

The ability to share Smartsheet actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareTextractAction`

The ability to share Textract actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareZendeskAction`

The ability to share Zendesk actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlackAction`

The ability to perform actions using Slack connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmartsheetAction`

The ability to perform actions using Smartsheet connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Space`

The ability to perform space-related actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubscribeDashboardEmailReports`

The ability to subscribe to email reports.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextractAction`

The ability to perform actions using Textract connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Topic`

The ability to perform Topic-related actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseAgentWebSearch`

The ability to use internet to enhance results in Chat Agents, Flows, and Quick Research.
Web search queries will be processed securely in an AWS region `us-east-1`.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseAmazonBedrockARSAction`

The ability to use Bedrock Agent actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseAmazonBedrockFSAction`

The ability to use Bedrock Runtime actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseAmazonBedrockKRSAction`

The ability to use Bedrock Data Automation Runtime actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseAmazonSThreeAction`

The ability to use Amazon S3 actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseAsanaAction`

The ability to use Asana actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseBambooHRAction`

The ability to use BambooHR actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseBedrockModels`

The ability to use Bedrock models for general knowledge step in flows.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseBoxAgentAction`

The ability to use Box Agent actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseCanvaAgentAction`

The ability to use Canva Agent actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseComprehendAction`

The ability to use Comprehend actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseComprehendMedicalAction`

The ability to use Comprehend Medical actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseConfluenceAction`

The ability to use Atlassian Confluence Cloud actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseFactSetAction`

The ability to use FactSet actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseGenericHTTPAction`

The ability to use REST API connection actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseGithubAction`

The ability to use GitHub actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseGoogleCalendarAction`

The ability to use Google Calendar actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseHubspotAction`

The ability to use Hubspot actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseHuggingFaceAction`

The ability to use HuggingFace actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseIntercomAction`

The ability to use Intercom actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseJiraAction`

The ability to use Jira actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseLinearAction`

The ability to use Linear actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseMCPAction`

The ability to use Model Context Protocol actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseMondayAction`

The ability to use Monday actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseMSExchangeAction`

The ability to use Microsoft Outlook actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseMSTeamsAction`

The ability to use Microsoft Teams actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseNewRelicAction`

The ability to use New Relic actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseNotionAction`

The ability to use Notion actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseOneDriveAction`

The ability to use Microsoft OneDrive actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseOpenAPIAction`

The ability to use OpenAPI Specification actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsePagerDutyAction`

The ability to use PagerDuty Advance actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseSalesforceAction`

The ability to use Salesforce actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseSandPGlobalEnergyAction`

The ability to use S&P Global Energy actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseSandPGMIAction`

The ability to use S&P Global Market Intelligence actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseSAPBillOfMaterialAction`

The ability to use SAP Bill of Materials actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseSAPBusinessPartnerAction`

The ability to use SAP Business Partner actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseSAPMaterialStockAction`

The ability to use SAP Material Stock actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseSAPPhysicalInventoryAction`

The ability to use SAP Physical Inventory actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseSAPProductMasterDataAction`

The ability to use SAP Product Master actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseServiceNowAction`

The ability to use ServiceNow actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseSharePointAction`

The ability to use Microsoft SharePoint Online actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseSlackAction`

The ability to use Slack actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseSmartsheetAction`

The ability to use Smartsheet actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseTextractAction`

The ability to use Textract actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseZendeskAction`

The ability to use Zendesk actions.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ViewAccountSPICECapacity`

The ability to view account SPICE capacity.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ZendeskAction`

The ability to perform actions using Zendesk connectors.

_Required_: No

_Type_: String

_Allowed values_: `DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::QuickSight::CustomPermissions

Tag

All content copied from https://docs.aws.amazon.com/.
