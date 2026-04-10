# Resource type support

The following table lists AWS resource types that currently support import, drift
detection, and infrastructure as code (IaC) generator operations. Each resource type name links
to its corresponding reference topic in the [CloudFormation Template Reference Guide](../templatereference/introduction.md).

A resource that supports resource import could also support auto-import. For more information, see [Import AWS resources into a CloudFormation stack](import-resources.md).

This is not an exhaustive list of AWS resources. If a specific resource type isn't listed,
it likely means that it's not accessible through the AWS Cloud Control API. For more information, see
[Resource types that support Cloud Control API](../../../cloudcontrolapi/latest/userguide/supported-resources.md) in the _Cloud Control API User Guide_.
Each AWS service independently determines which resource types to make accessible through
Cloud Control API.

CloudFormation also supports import and drift detection operations for private resource types
that are provisionable (those with provisioning types of either `FULLY_MUTABLE` or
`IMMUTABLE`). To import or perform drift detection on a private resource type, you
must first register the default version of that resource type in your account, and ensure it's
provisionable. For more information, see [Use third-party private extensions that have been shared with you](registry-private.md).

Note that IaC generator only supports AWS resources that are compatible with Cloud Control API in
your Region.

To programmatically access information about public and private provisionable resource
types, you can use the AWS Cloud Control API. For more information, see [Determining if a resource type supports Cloud Control API](../../../cloudcontrolapi/latest/userguide/resource-types.md#resource-types-determine-support) in the _Cloud Control API User_
_Guide_.

To get started with import, drift detection, or IaC generator, here are some useful
topics to review:

- [Import AWS resources into a CloudFormation stack](import-resources.md)

- [Detect unmanaged configuration changes to stacks and resources with drift detection](using-cfn-stack-drift.md)

- [Generate templates from existing resources with IaC generator](generate-iac.md)

ResourceImportDrift detectionIaC generator

[`AWS::ACMPCA::Certificate`](../templatereference/aws-resource-acmpca-certificate.md)

Yes

No

No

[`AWS::ACMPCA::CertificateAuthority`](../templatereference/aws-resource-acmpca-certificateauthority.md)

Yes

Yes

Yes

[`AWS::ACMPCA::CertificateAuthorityActivation`](../templatereference/aws-resource-acmpca-certificateauthorityactivation.md)

Yes

No

No

[`AWS::ACMPCA::Permission`](../templatereference/aws-resource-acmpca-permission.md)

Yes

Yes

No

[`AWS::AIOps::InvestigationGroup`](../templatereference/aws-resource-aiops-investigationgroup.md)

Yes

Yes

No

[`AWS::APS::AnomalyDetector`](../templatereference/aws-resource-aps-anomalydetector.md)

Yes

Yes

No

[`AWS::APS::ResourcePolicy`](../templatereference/aws-resource-aps-resourcepolicy.md)

Yes

Yes

No

[`AWS::APS::RuleGroupsNamespace`](../templatereference/aws-resource-aps-rulegroupsnamespace.md)

Yes

Yes

Yes

[`AWS::APS::Scraper`](../templatereference/aws-resource-aps-scraper.md)

Yes

Yes

No

[`AWS::APS::Workspace`](../templatereference/aws-resource-aps-workspace.md)

Yes

Yes

Yes

[`AWS::ARCRegionSwitch::Plan`](../templatereference/aws-resource-arcregionswitch-plan.md)

Yes

Yes

No

[`AWS::ARCZonalShift::AutoshiftObserverNotificationStatus`](../templatereference/aws-resource-arczonalshift-autoshiftobservernotificationstatus.md)

Yes

Yes

No

[`AWS::ARCZonalShift::ZonalAutoshiftConfiguration`](../templatereference/aws-resource-arczonalshift-zonalautoshiftconfiguration.md)

Yes

Yes

No

[`AWS::AccessAnalyzer::Analyzer`](../templatereference/aws-resource-accessanalyzer-analyzer.md)

Yes

Yes

Yes

[`AWS::AmazonMQ::Broker`](../templatereference/aws-resource-amazonmq-broker.md)

Yes

Yes

No

[`AWS::AmazonMQ::Configuration`](../templatereference/aws-resource-amazonmq-configuration.md)

Yes

Yes

No

[`AWS::Amplify::App`](../templatereference/aws-resource-amplify-app.md)

Yes

Yes

Yes

[`AWS::Amplify::Branch`](../templatereference/aws-resource-amplify-branch.md)

Yes

Yes

Yes

[`AWS::Amplify::Domain`](../templatereference/aws-resource-amplify-domain.md)

Yes

Yes

Yes

[`AWS::AmplifyUIBuilder::Component`](../templatereference/aws-resource-amplifyuibuilder-component.md)

Yes

Yes

Yes

[`AWS::AmplifyUIBuilder::Form`](../templatereference/aws-resource-amplifyuibuilder-form.md)

Yes

Yes

Yes

[`AWS::AmplifyUIBuilder::Theme`](../templatereference/aws-resource-amplifyuibuilder-theme.md)

Yes

Yes

Yes

[`AWS::ApiGateway::Account`](../templatereference/aws-resource-apigateway-account.md)

Yes

Yes

No

[`AWS::ApiGateway::ApiKey`](../templatereference/aws-resource-apigateway-apikey.md)

Yes

Yes

Yes

[`AWS::ApiGateway::Authorizer`](../templatereference/aws-resource-apigateway-authorizer.md)

Yes

Yes

No

[`AWS::ApiGateway::BasePathMapping`](../templatereference/aws-resource-apigateway-basepathmapping.md)

Yes

Yes

No

[`AWS::ApiGateway::BasePathMappingV2`](../templatereference/aws-resource-apigateway-basepathmappingv2.md)

Yes

Yes

No

[`AWS::ApiGateway::ClientCertificate`](../templatereference/aws-resource-apigateway-clientcertificate.md)

Yes

Yes

Yes

[`AWS::ApiGateway::Deployment`](../templatereference/aws-resource-apigateway-deployment.md)

Yes

Yes

Yes

[`AWS::ApiGateway::DocumentationPart`](../templatereference/aws-resource-apigateway-documentationpart.md)

Yes

Yes

No

[`AWS::ApiGateway::DocumentationVersion`](../templatereference/aws-resource-apigateway-documentationversion.md)

Yes

Yes

Yes

[`AWS::ApiGateway::DomainName`](../templatereference/aws-resource-apigateway-domainname.md)

Yes

No

No

[`AWS::ApiGateway::DomainNameAccessAssociation`](../templatereference/aws-resource-apigateway-domainnameaccessassociation.md)

Yes

Yes

No

[`AWS::ApiGateway::DomainNameV2`](../templatereference/aws-resource-apigateway-domainnamev2.md)

Yes

Yes

No

[`AWS::ApiGateway::GatewayResponse`](../templatereference/aws-resource-apigateway-gatewayresponse.md)

Yes

No

No

[`AWS::ApiGateway::Method`](../templatereference/aws-resource-apigateway-method.md)

Yes

Yes

No

[`AWS::ApiGateway::Model`](../templatereference/aws-resource-apigateway-model.md)

Yes

Yes

No

[`AWS::ApiGateway::RequestValidator`](../templatereference/aws-resource-apigateway-requestvalidator.md)

Yes

Yes

No

[`AWS::ApiGateway::Resource`](../templatereference/aws-resource-apigateway-resource.md)

Yes

Yes

No

[`AWS::ApiGateway::RestApi`](../templatereference/aws-resource-apigateway-restapi.md)

Yes

Yes

Yes

[`AWS::ApiGateway::Stage`](../templatereference/aws-resource-apigateway-stage.md)

Yes

Yes

Yes

[`AWS::ApiGateway::UsagePlan`](../templatereference/aws-resource-apigateway-usageplan.md)

Yes

Yes

Yes

[`AWS::ApiGateway::UsagePlanKey`](../templatereference/aws-resource-apigateway-usageplankey.md)

Yes

No

Yes

[`AWS::ApiGateway::VpcLink`](../templatereference/aws-resource-apigateway-vpclink.md)

Yes

Yes

No

[`AWS::ApiGatewayV2::Api`](../templatereference/aws-resource-apigatewayv2-api.md)

Yes

Yes

Yes

[`AWS::ApiGatewayV2::ApiMapping`](../templatereference/aws-resource-apigatewayv2-apimapping.md)

Yes

Yes

Yes

[`AWS::ApiGatewayV2::Authorizer`](../templatereference/aws-resource-apigatewayv2-authorizer.md)

Yes

Yes

No

[`AWS::ApiGatewayV2::Deployment`](../templatereference/aws-resource-apigatewayv2-deployment.md)

Yes

Yes

Yes

[`AWS::ApiGatewayV2::DomainName`](../templatereference/aws-resource-apigatewayv2-domainname.md)

Yes

Yes

Yes

[`AWS::ApiGatewayV2::Integration`](../templatereference/aws-resource-apigatewayv2-integration.md)

Yes

No

No

[`AWS::ApiGatewayV2::IntegrationResponse`](../templatereference/aws-resource-apigatewayv2-integrationresponse.md)

Yes

Yes

No

[`AWS::ApiGatewayV2::Model`](../templatereference/aws-resource-apigatewayv2-model.md)

Yes

Yes

No

[`AWS::ApiGatewayV2::Route`](../templatereference/aws-resource-apigatewayv2-route.md)

Yes

Yes

No

[`AWS::ApiGatewayV2::RouteResponse`](../templatereference/aws-resource-apigatewayv2-routeresponse.md)

Yes

Yes

No

[`AWS::ApiGatewayV2::RoutingRule`](../templatereference/aws-resource-apigatewayv2-routingrule.md)

Yes

Yes

No

[`AWS::ApiGatewayV2::VpcLink`](../templatereference/aws-resource-apigatewayv2-vpclink.md)

Yes

Yes

Yes

[`AWS::AppConfig::Application`](../templatereference/aws-resource-appconfig-application.md)

Yes

Yes

Yes

[`AWS::AppConfig::ConfigurationProfile`](../templatereference/aws-resource-appconfig-configurationprofile.md)

Yes

Yes

No

[`AWS::AppConfig::Deployment`](../templatereference/aws-resource-appconfig-deployment.md)

Yes

Yes

No

[`AWS::AppConfig::DeploymentStrategy`](../templatereference/aws-resource-appconfig-deploymentstrategy.md)

Yes

Yes

No

[`AWS::AppConfig::Environment`](../templatereference/aws-resource-appconfig-environment.md)

Yes

Yes

No

[`AWS::AppConfig::Extension`](../templatereference/aws-resource-appconfig-extension.md)

Yes

Yes

No

[`AWS::AppConfig::ExtensionAssociation`](../templatereference/aws-resource-appconfig-extensionassociation.md)

Yes

Yes

No

[`AWS::AppConfig::HostedConfigurationVersion`](../templatereference/aws-resource-appconfig-hostedconfigurationversion.md)

Yes

Yes

No

[`AWS::AppFlow::Connector`](../templatereference/aws-resource-appflow-connector.md)

Yes

Yes

Yes

[`AWS::AppFlow::ConnectorProfile`](../templatereference/aws-resource-appflow-connectorprofile.md)

Yes

Yes

Yes

[`AWS::AppFlow::Flow`](../templatereference/aws-resource-appflow-flow.md)

Yes

Yes

Yes

[`AWS::AppIntegrations::Application`](../templatereference/aws-resource-appintegrations-application.md)

Yes

Yes

No

[`AWS::AppIntegrations::DataIntegration`](../templatereference/aws-resource-appintegrations-dataintegration.md)

Yes

Yes

Yes

[`AWS::AppIntegrations::EventIntegration`](../templatereference/aws-resource-appintegrations-eventintegration.md)

Yes

Yes

Yes

[`AWS::AppRunner::AutoScalingConfiguration`](../templatereference/aws-resource-apprunner-autoscalingconfiguration.md)

Yes

Yes

No

[`AWS::AppRunner::ObservabilityConfiguration`](../templatereference/aws-resource-apprunner-observabilityconfiguration.md)

Yes

Yes

No

[`AWS::AppRunner::Service`](../templatereference/aws-resource-apprunner-service.md)

Yes

Yes

No

[`AWS::AppRunner::VpcConnector`](../templatereference/aws-resource-apprunner-vpcconnector.md)

Yes

Yes

No

[`AWS::AppRunner::VpcIngressConnection`](../templatereference/aws-resource-apprunner-vpcingressconnection.md)

Yes

Yes

No

[`AWS::AppStream::AppBlock`](../templatereference/aws-resource-appstream-appblock.md)

Yes

Yes

No

[`AWS::AppStream::AppBlockBuilder`](../templatereference/aws-resource-appstream-appblockbuilder.md)

Yes

Yes

No

[`AWS::AppStream::Application`](../templatereference/aws-resource-appstream-application.md)

Yes

Yes

No

[`AWS::AppStream::ApplicationEntitlementAssociation`](../templatereference/aws-resource-appstream-applicationentitlementassociation.md)

Yes

Yes

No

[`AWS::AppStream::ApplicationFleetAssociation`](../templatereference/aws-resource-appstream-applicationfleetassociation.md)

Yes

Yes

No

[`AWS::AppStream::DirectoryConfig`](../templatereference/aws-resource-appstream-directoryconfig.md)

Yes

Yes

Yes

[`AWS::AppStream::Entitlement`](../templatereference/aws-resource-appstream-entitlement.md)

Yes

Yes

No

[`AWS::AppStream::ImageBuilder`](../templatereference/aws-resource-appstream-imagebuilder.md)

Yes

Yes

Yes

[`AWS::AppSync::Api`](../templatereference/aws-resource-appsync-api.md)

Yes

Yes

No

[`AWS::AppSync::ChannelNamespace`](../templatereference/aws-resource-appsync-channelnamespace.md)

Yes

Yes

No

[`AWS::AppSync::DataSource`](../templatereference/aws-resource-appsync-datasource.md)

Yes

Yes

No

[`AWS::AppSync::DomainName`](../templatereference/aws-resource-appsync-domainname.md)

Yes

Yes

Yes

[`AWS::AppSync::DomainNameApiAssociation`](../templatereference/aws-resource-appsync-domainnameapiassociation.md)

Yes

Yes

No

[`AWS::AppSync::FunctionConfiguration`](../templatereference/aws-resource-appsync-functionconfiguration.md)

Yes

Yes

No

[`AWS::AppSync::GraphQLApi`](../templatereference/aws-resource-appsync-graphqlapi.md)

No

Yes

No

[`AWS::AppSync::Resolver`](../templatereference/aws-resource-appsync-resolver.md)

Yes

Yes

No

[`AWS::AppSync::SourceApiAssociation`](../templatereference/aws-resource-appsync-sourceapiassociation.md)

Yes

Yes

Yes

[`AWS::AppTest::TestCase`](../templatereference/aws-resource-apptest-testcase.md)

Yes

Yes

No

[`AWS::ApplicationAutoScaling::ScalableTarget`](../templatereference/aws-resource-applicationautoscaling-scalabletarget.md)

Yes

Yes

Yes

[`AWS::ApplicationAutoScaling::ScalingPolicy`](../templatereference/aws-resource-applicationautoscaling-scalingpolicy.md)

Yes

No

No

[`AWS::ApplicationInsights::Application`](../templatereference/aws-resource-applicationinsights-application.md)

Yes

Yes

No

[`AWS::ApplicationSignals::Discovery`](../templatereference/aws-resource-applicationsignals-discovery.md)

Yes

Yes

No

[`AWS::ApplicationSignals::GroupingConfiguration`](../templatereference/aws-resource-applicationsignals-groupingconfiguration.md)

Yes

Yes

No

[`AWS::ApplicationSignals::ServiceLevelObjective`](../templatereference/aws-resource-applicationsignals-servicelevelobjective.md)

Yes

Yes

No

[`AWS::Athena::CapacityReservation`](../templatereference/aws-resource-athena-capacityreservation.md)

Yes

Yes

Yes

[`AWS::Athena::DataCatalog`](../templatereference/aws-resource-athena-datacatalog.md)

Yes

Yes

Yes

[`AWS::Athena::NamedQuery`](../templatereference/aws-resource-athena-namedquery.md)

Yes

Yes

Yes

[`AWS::Athena::PreparedStatement`](../templatereference/aws-resource-athena-preparedstatement.md)

Yes

Yes

Yes

[`AWS::Athena::WorkGroup`](../templatereference/aws-resource-athena-workgroup.md)

Yes

No

Yes

[`AWS::AuditManager::Assessment`](../templatereference/aws-resource-auditmanager-assessment.md)

Yes

No

No

[`AWS::AutoScaling::AutoScalingGroup`](../templatereference/aws-resource-autoscaling-autoscalinggroup.md)

Yes

Yes

Yes

[`AWS::AutoScaling::LaunchConfiguration`](../templatereference/aws-resource-autoscaling-launchconfiguration.md)

Yes

Yes

Yes

[`AWS::AutoScaling::LifecycleHook`](../templatereference/aws-resource-autoscaling-lifecyclehook.md)

Yes

Yes

Yes

[`AWS::AutoScaling::ScalingPolicy`](../templatereference/aws-resource-autoscaling-scalingpolicy.md)

Yes

Yes

Yes

[`AWS::AutoScaling::ScheduledAction`](../templatereference/aws-resource-autoscaling-scheduledaction.md)

Yes

Yes

Yes

[`AWS::AutoScaling::WarmPool`](../templatereference/aws-resource-autoscaling-warmpool.md)

Yes

Yes

No

[`AWS::B2BI::Capability`](../templatereference/aws-resource-b2bi-capability.md)

Yes

Yes

No

[`AWS::B2BI::Partnership`](../templatereference/aws-resource-b2bi-partnership.md)

Yes

Yes

No

[`AWS::B2BI::Profile`](../templatereference/aws-resource-b2bi-profile.md)

Yes

Yes

No

[`AWS::B2BI::Transformer`](../templatereference/aws-resource-b2bi-transformer.md)

Yes

Yes

No

[`AWS::BCMDataExports::Export`](../templatereference/aws-resource-bcmdataexports-export.md)

Yes

Yes

No

[`AWS::Backup::BackupPlan`](../templatereference/aws-resource-backup-backupplan.md)

Yes

Yes

No

[`AWS::Backup::BackupSelection`](../templatereference/aws-resource-backup-backupselection.md)

Yes

No

Yes

[`AWS::Backup::BackupVault`](../templatereference/aws-resource-backup-backupvault.md)

Yes

Yes

No

[`AWS::Backup::Framework`](../templatereference/aws-resource-backup-framework.md)

Yes

Yes

Yes

[`AWS::Backup::LogicallyAirGappedBackupVault`](../templatereference/aws-resource-backup-logicallyairgappedbackupvault.md)

Yes

Yes

No

[`AWS::Backup::ReportPlan`](../templatereference/aws-resource-backup-reportplan.md)

Yes

Yes

Yes

[`AWS::Backup::RestoreTestingPlan`](../templatereference/aws-resource-backup-restoretestingplan.md)

Yes

Yes

No

[`AWS::Backup::RestoreTestingSelection`](../templatereference/aws-resource-backup-restoretestingselection.md)

Yes

Yes

No

[`AWS::Backup::TieringConfiguration`](../templatereference/aws-resource-backup-tieringconfiguration.md)

Yes

Yes

No

[`AWS::BackupGateway::Hypervisor`](../templatereference/aws-resource-backupgateway-hypervisor.md)

Yes

Yes

No

[`AWS::Batch::ComputeEnvironment`](../templatereference/aws-resource-batch-computeenvironment.md)

Yes

Yes

Yes

[`AWS::Batch::ConsumableResource`](../templatereference/aws-resource-batch-consumableresource.md)

Yes

Yes

No

[`AWS::Batch::JobDefinition`](../templatereference/aws-resource-batch-jobdefinition.md)

Yes

Yes

No

[`AWS::Batch::JobQueue`](../templatereference/aws-resource-batch-jobqueue.md)

Yes

Yes

Yes

[`AWS::Batch::SchedulingPolicy`](../templatereference/aws-resource-batch-schedulingpolicy.md)

Yes

Yes

Yes

[`AWS::Batch::ServiceEnvironment`](../templatereference/aws-resource-batch-serviceenvironment.md)

Yes

Yes

No

[`AWS::Bedrock::Agent`](../templatereference/aws-resource-bedrock-agent.md)

Yes

Yes

Yes

[`AWS::Bedrock::AgentAlias`](../templatereference/aws-resource-bedrock-agentalias.md)

Yes

Yes

Yes

[`AWS::Bedrock::ApplicationInferenceProfile`](../templatereference/aws-resource-bedrock-applicationinferenceprofile.md)

Yes

Yes

Yes

[`AWS::Bedrock::AutomatedReasoningPolicy`](../templatereference/aws-resource-bedrock-automatedreasoningpolicy.md)

Yes

Yes

No

[`AWS::Bedrock::AutomatedReasoningPolicyVersion`](../templatereference/aws-resource-bedrock-automatedreasoningpolicyversion.md)

Yes

Yes

No

[`AWS::Bedrock::Blueprint`](../templatereference/aws-resource-bedrock-blueprint.md)

Yes

Yes

Yes

[`AWS::Bedrock::DataAutomationProject`](../templatereference/aws-resource-bedrock-dataautomationproject.md)

Yes

Yes

Yes

[`AWS::Bedrock::DataSource`](../templatereference/aws-resource-bedrock-datasource.md)

Yes

Yes

Yes

[`AWS::Bedrock::Flow`](../templatereference/aws-resource-bedrock-flow.md)

Yes

Yes

Yes

[`AWS::Bedrock::FlowAlias`](../templatereference/aws-resource-bedrock-flowalias.md)

Yes

Yes

Yes

[`AWS::Bedrock::FlowVersion`](../templatereference/aws-resource-bedrock-flowversion.md)

Yes

Yes

Yes

[`AWS::Bedrock::Guardrail`](../templatereference/aws-resource-bedrock-guardrail.md)

Yes

Yes

Yes

[`AWS::Bedrock::GuardrailVersion`](../templatereference/aws-resource-bedrock-guardrailversion.md)

Yes

Yes

No

[`AWS::Bedrock::IntelligentPromptRouter`](../templatereference/aws-resource-bedrock-intelligentpromptrouter.md)

Yes

Yes

Yes

[`AWS::Bedrock::KnowledgeBase`](../templatereference/aws-resource-bedrock-knowledgebase.md)

Yes

No

Yes

[`AWS::Bedrock::Prompt`](../templatereference/aws-resource-bedrock-prompt.md)

Yes

Yes

Yes

[`AWS::Bedrock::PromptVersion`](../templatereference/aws-resource-bedrock-promptversion.md)

Yes

Yes

Yes

[`AWS::BedrockAgentCore::BrowserCustom`](../templatereference/aws-resource-bedrockagentcore-browsercustom.md)

Yes

Yes

No

[`AWS::BedrockAgentCore::BrowserProfile`](../templatereference/aws-resource-bedrockagentcore-browserprofile.md)

Yes

Yes

No

[`AWS::BedrockAgentCore::CodeInterpreterCustom`](../templatereference/aws-resource-bedrockagentcore-codeinterpretercustom.md)

Yes

Yes

No

[`AWS::BedrockAgentCore::Evaluator`](../templatereference/aws-resource-bedrockagentcore-evaluator.md)

Yes

Yes

No

[`AWS::BedrockAgentCore::Gateway`](../templatereference/aws-resource-bedrockagentcore-gateway.md)

Yes

Yes

No

[`AWS::BedrockAgentCore::GatewayTarget`](../templatereference/aws-resource-bedrockagentcore-gatewaytarget.md)

Yes

Yes

No

[`AWS::BedrockAgentCore::Memory`](../templatereference/aws-resource-bedrockagentcore-memory.md)

Yes

Yes

No

[`AWS::BedrockAgentCore::OnlineEvaluationConfig`](../templatereference/aws-resource-bedrockagentcore-onlineevaluationconfig.md)

Yes

Yes

No

[`AWS::BedrockAgentCore::Policy`](../templatereference/aws-resource-bedrockagentcore-policy.md)

Yes

Yes

No

[`AWS::BedrockAgentCore::PolicyEngine`](../templatereference/aws-resource-bedrockagentcore-policyengine.md)

Yes

Yes

No

[`AWS::BedrockAgentCore::Runtime`](../templatereference/aws-resource-bedrockagentcore-runtime.md)

Yes

Yes

No

[`AWS::BedrockAgentCore::RuntimeEndpoint`](../templatereference/aws-resource-bedrockagentcore-runtimeendpoint.md)

Yes

Yes

No

[`AWS::BedrockAgentCore::WorkloadIdentity`](../templatereference/aws-resource-bedrockagentcore-workloadidentity.md)

Yes

Yes

No

[`AWS::BedrockMantle::Project`](../templatereference/aws-resource-bedrockmantle-project.md)

Yes

Yes

No

[`AWS::Billing::BillingView`](../templatereference/aws-resource-billing-billingview.md)

Yes

Yes

No

[`AWS::BillingConductor::BillingGroup`](../templatereference/aws-resource-billingconductor-billinggroup.md)

Yes

Yes

Yes

[`AWS::BillingConductor::CustomLineItem`](../templatereference/aws-resource-billingconductor-customlineitem.md)

Yes

Yes

Yes

[`AWS::BillingConductor::PricingPlan`](../templatereference/aws-resource-billingconductor-pricingplan.md)

Yes

Yes

Yes

[`AWS::BillingConductor::PricingRule`](../templatereference/aws-resource-billingconductor-pricingrule.md)

Yes

Yes

Yes

[`AWS::Budgets::BudgetsAction`](../templatereference/aws-resource-budgets-budgetsaction.md)

Yes

Yes

Yes

[`AWS::CE::AnomalyMonitor`](../templatereference/aws-resource-ce-anomalymonitor.md)

Yes

Yes

No

[`AWS::CE::AnomalySubscription`](../templatereference/aws-resource-ce-anomalysubscription.md)

Yes

Yes

No

[`AWS::CE::CostCategory`](../templatereference/aws-resource-ce-costcategory.md)

Yes

Yes

Yes

[`AWS::CUR::ReportDefinition`](../templatereference/aws-resource-cur-reportdefinition.md)

Yes

Yes

No

[`AWS::Cases::CaseRule`](../templatereference/aws-resource-cases-caserule.md)

Yes

Yes

No

[`AWS::Cases::Domain`](../templatereference/aws-resource-cases-domain.md)

Yes

Yes

No

[`AWS::Cases::Field`](../templatereference/aws-resource-cases-field.md)

Yes

Yes

No

[`AWS::Cases::Layout`](../templatereference/aws-resource-cases-layout.md)

Yes

Yes

No

[`AWS::Cases::Template`](../templatereference/aws-resource-cases-template.md)

Yes

Yes

No

[`AWS::Cassandra::Keyspace`](../templatereference/aws-resource-cassandra-keyspace.md)

Yes

Yes

No

[`AWS::Cassandra::Table`](../templatereference/aws-resource-cassandra-table.md)

Yes

Yes

Yes

[`AWS::Cassandra::Type`](../templatereference/aws-resource-cassandra-type.md)

Yes

Yes

No

[`AWS::CertificateManager::Account`](../templatereference/aws-resource-certificatemanager-account.md)

Yes

Yes

No

[`AWS::Chatbot::CustomAction`](../templatereference/aws-resource-chatbot-customaction.md)

Yes

Yes

No

[`AWS::Chatbot::MicrosoftTeamsChannelConfiguration`](../templatereference/aws-resource-chatbot-microsoftteamschannelconfiguration.md)

Yes

Yes

Yes

[`AWS::Chatbot::SlackChannelConfiguration`](../templatereference/aws-resource-chatbot-slackchannelconfiguration.md)

Yes

Yes

Yes

[`AWS::CleanRooms::AnalysisTemplate`](../templatereference/aws-resource-cleanrooms-analysistemplate.md)

Yes

Yes

Yes

[`AWS::CleanRooms::Collaboration`](../templatereference/aws-resource-cleanrooms-collaboration.md)

Yes

Yes

Yes

[`AWS::CleanRooms::ConfiguredTable`](../templatereference/aws-resource-cleanrooms-configuredtable.md)

Yes

Yes

Yes

[`AWS::CleanRooms::ConfiguredTableAssociation`](../templatereference/aws-resource-cleanrooms-configuredtableassociation.md)

Yes

Yes

Yes

[`AWS::CleanRooms::IdMappingTable`](../templatereference/aws-resource-cleanrooms-idmappingtable.md)

Yes

Yes

No

[`AWS::CleanRooms::IdNamespaceAssociation`](../templatereference/aws-resource-cleanrooms-idnamespaceassociation.md)

Yes

Yes

No

[`AWS::CleanRooms::Membership`](../templatereference/aws-resource-cleanrooms-membership.md)

Yes

Yes

Yes

[`AWS::CleanRooms::PrivacyBudgetTemplate`](../templatereference/aws-resource-cleanrooms-privacybudgettemplate.md)

Yes

Yes

No

[`AWS::CleanRoomsML::ConfiguredModelAlgorithm`](../templatereference/aws-resource-cleanroomsml-configuredmodelalgorithm.md)

Yes

Yes

No

[`AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation`](../templatereference/aws-resource-cleanroomsml-configuredmodelalgorithmassociation.md)

Yes

Yes

No

[`AWS::CleanRoomsML::TrainingDataset`](../templatereference/aws-resource-cleanroomsml-trainingdataset.md)

Yes

Yes

No

[`AWS::CloudFormation::GuardHook`](../templatereference/aws-resource-cloudformation-guardhook.md)

Yes

Yes

No

[`AWS::CloudFormation::HookDefaultVersion`](../templatereference/aws-resource-cloudformation-hookdefaultversion.md)

Yes

Yes

Yes

[`AWS::CloudFormation::HookTypeConfig`](../templatereference/aws-resource-cloudformation-hooktypeconfig.md)

Yes

Yes

Yes

[`AWS::CloudFormation::HookVersion`](../templatereference/aws-resource-cloudformation-hookversion.md)

Yes

Yes

Yes

[`AWS::CloudFormation::LambdaHook`](../templatereference/aws-resource-cloudformation-lambdahook.md)

Yes

Yes

No

[`AWS::CloudFormation::ModuleDefaultVersion`](../templatereference/aws-resource-cloudformation-moduledefaultversion.md)

Yes

Yes

No

[`AWS::CloudFormation::ModuleVersion`](../templatereference/aws-resource-cloudformation-moduleversion.md)

Yes

Yes

No

[`AWS::CloudFormation::PublicTypeVersion`](../templatereference/aws-resource-cloudformation-publictypeversion.md)

Yes

Yes

No

[`AWS::CloudFormation::Publisher`](../templatereference/aws-resource-cloudformation-publisher.md)

Yes

Yes

No

[`AWS::CloudFormation::ResourceDefaultVersion`](../templatereference/aws-resource-cloudformation-resourcedefaultversion.md)

Yes

Yes

No

[`AWS::CloudFormation::ResourceVersion`](../templatereference/aws-resource-cloudformation-resourceversion.md)

Yes

Yes

No

[`AWS::CloudFormation::Stack`](../templatereference/aws-resource-cloudformation-stack.md)

Yes

No

No

[`AWS::CloudFormation::StackSet`](../templatereference/aws-resource-cloudformation-stackset.md)

Yes

Yes

No

[`AWS::CloudFormation::TypeActivation`](../templatereference/aws-resource-cloudformation-typeactivation.md)

Yes

Yes

No

[`AWS::CloudFront::AnycastIpList`](../templatereference/aws-resource-cloudfront-anycastiplist.md)

Yes

Yes

No

[`AWS::CloudFront::CachePolicy`](../templatereference/aws-resource-cloudfront-cachepolicy.md)

Yes

Yes

Yes

[`AWS::CloudFront::CloudFrontOriginAccessIdentity`](../templatereference/aws-resource-cloudfront-cloudfrontoriginaccessidentity.md)

Yes

Yes

Yes

[`AWS::CloudFront::ConnectionFunction`](../templatereference/aws-resource-cloudfront-connectionfunction.md)

Yes

Yes

No

[`AWS::CloudFront::ConnectionGroup`](../templatereference/aws-resource-cloudfront-connectiongroup.md)

Yes

Yes

No

[`AWS::CloudFront::ContinuousDeploymentPolicy`](../templatereference/aws-resource-cloudfront-continuousdeploymentpolicy.md)

Yes

Yes

Yes

[`AWS::CloudFront::Distribution`](../templatereference/aws-resource-cloudfront-distribution.md)

Yes

Yes

Yes

[`AWS::CloudFront::DistributionTenant`](../templatereference/aws-resource-cloudfront-distributiontenant.md)

Yes

Yes

No

[`AWS::CloudFront::Function`](../templatereference/aws-resource-cloudfront-function.md)

Yes

Yes

No

[`AWS::CloudFront::KeyGroup`](../templatereference/aws-resource-cloudfront-keygroup.md)

Yes

Yes

Yes

[`AWS::CloudFront::KeyValueStore`](../templatereference/aws-resource-cloudfront-keyvaluestore.md)

Yes

Yes

No

[`AWS::CloudFront::MonitoringSubscription`](../templatereference/aws-resource-cloudfront-monitoringsubscription.md)

Yes

Yes

No

[`AWS::CloudFront::OriginAccessControl`](../templatereference/aws-resource-cloudfront-originaccesscontrol.md)

Yes

Yes

Yes

[`AWS::CloudFront::OriginRequestPolicy`](../templatereference/aws-resource-cloudfront-originrequestpolicy.md)

Yes

Yes

Yes

[`AWS::CloudFront::PublicKey`](../templatereference/aws-resource-cloudfront-publickey.md)

Yes

Yes

Yes

[`AWS::CloudFront::RealtimeLogConfig`](../templatereference/aws-resource-cloudfront-realtimelogconfig.md)

Yes

Yes

Yes

[`AWS::CloudFront::ResponseHeadersPolicy`](../templatereference/aws-resource-cloudfront-responseheaderspolicy.md)

Yes

Yes

No

[`AWS::CloudFront::TrustStore`](../templatereference/aws-resource-cloudfront-truststore.md)

Yes

Yes

No

[`AWS::CloudFront::VpcOrigin`](../templatereference/aws-resource-cloudfront-vpcorigin.md)

Yes

Yes

No

[`AWS::CloudTrail::Channel`](../templatereference/aws-resource-cloudtrail-channel.md)

Yes

Yes

No

[`AWS::CloudTrail::Dashboard`](../templatereference/aws-resource-cloudtrail-dashboard.md)

Yes

Yes

No

[`AWS::CloudTrail::EventDataStore`](../templatereference/aws-resource-cloudtrail-eventdatastore.md)

Yes

Yes

Yes

[`AWS::CloudTrail::ResourcePolicy`](../templatereference/aws-resource-cloudtrail-resourcepolicy.md)

Yes

Yes

No

[`AWS::CloudTrail::Trail`](../templatereference/aws-resource-cloudtrail-trail.md)

Yes

Yes

Yes

[`AWS::CloudWatch::Alarm`](../templatereference/aws-resource-cloudwatch-alarm.md)

Yes

Yes

Yes

[`AWS::CloudWatch::AlarmMuteRule`](../templatereference/aws-resource-cloudwatch-alarmmuterule.md)

Yes

Yes

No

[`AWS::CloudWatch::CompositeAlarm`](../templatereference/aws-resource-cloudwatch-compositealarm.md)

Yes

Yes

Yes

[`AWS::CloudWatch::Dashboard`](../templatereference/aws-resource-cloudwatch-dashboard.md)

Yes

Yes

No

[`AWS::CloudWatch::MetricStream`](../templatereference/aws-resource-cloudwatch-metricstream.md)

Yes

Yes

No

[`AWS::CodeArtifact::Domain`](../templatereference/aws-resource-codeartifact-domain.md)

Yes

Yes

Yes

[`AWS::CodeArtifact::PackageGroup`](../templatereference/aws-resource-codeartifact-packagegroup.md)

Yes

Yes

No

[`AWS::CodeArtifact::Repository`](../templatereference/aws-resource-codeartifact-repository.md)

Yes

Yes

Yes

[`AWS::CodeBuild::Fleet`](../templatereference/aws-resource-codebuild-fleet.md)

Yes

Yes

No

[`AWS::CodeConnections::Connection`](../templatereference/aws-resource-codeconnections-connection.md)

Yes

Yes

No

[`AWS::CodeDeploy::Application`](../templatereference/aws-resource-codedeploy-application.md)

Yes

Yes

Yes

[`AWS::CodeDeploy::DeploymentConfig`](../templatereference/aws-resource-codedeploy-deploymentconfig.md)

Yes

Yes

No

[`AWS::CodeDeploy::DeploymentGroup`](../templatereference/aws-resource-codedeploy-deploymentgroup.md)

Yes

Yes

No

[`AWS::CodeGuruProfiler::ProfilingGroup`](../templatereference/aws-resource-codeguruprofiler-profilinggroup.md)

Yes

Yes

Yes

[`AWS::CodeGuruReviewer::RepositoryAssociation`](../templatereference/aws-resource-codegurureviewer-repositoryassociation.md)

Yes

Yes

No

[`AWS::CodePipeline::CustomActionType`](../templatereference/aws-resource-codepipeline-customactiontype.md)

Yes

Yes

No

[`AWS::CodePipeline::Pipeline`](../templatereference/aws-resource-codepipeline-pipeline.md)

Yes

Yes

No

[`AWS::CodePipeline::Webhook`](../templatereference/aws-resource-codepipeline-webhook.md)

Yes

Yes

No

[`AWS::CodeStarConnections::Connection`](../templatereference/aws-resource-codestarconnections-connection.md)

Yes

Yes

Yes

[`AWS::CodeStarConnections::RepositoryLink`](../templatereference/aws-resource-codestarconnections-repositorylink.md)

Yes

Yes

No

[`AWS::CodeStarConnections::SyncConfiguration`](../templatereference/aws-resource-codestarconnections-syncconfiguration.md)

Yes

Yes

No

[`AWS::CodeStarNotifications::NotificationRule`](../templatereference/aws-resource-codestarnotifications-notificationrule.md)

Yes

Yes

Yes

[`AWS::Cognito::IdentityPool`](../templatereference/aws-resource-cognito-identitypool.md)

Yes

Yes

No

[`AWS::Cognito::IdentityPoolPrincipalTag`](../templatereference/aws-resource-cognito-identitypoolprincipaltag.md)

Yes

Yes

Yes

[`AWS::Cognito::IdentityPoolRoleAttachment`](../templatereference/aws-resource-cognito-identitypoolroleattachment.md)

Yes

No

No

[`AWS::Cognito::LogDeliveryConfiguration`](../templatereference/aws-resource-cognito-logdeliveryconfiguration.md)

Yes

Yes

No

[`AWS::Cognito::ManagedLoginBranding`](../templatereference/aws-resource-cognito-managedloginbranding.md)

Yes

Yes

No

[`AWS::Cognito::Terms`](../templatereference/aws-resource-cognito-terms.md)

Yes

Yes

No

[`AWS::Cognito::UserPool`](../templatereference/aws-resource-cognito-userpool.md)

Yes

No

No

[`AWS::Cognito::UserPoolClient`](../templatereference/aws-resource-cognito-userpoolclient.md)

Yes

Yes

No

[`AWS::Cognito::UserPoolDomain`](../templatereference/aws-resource-cognito-userpooldomain.md)

Yes

Yes

No

[`AWS::Cognito::UserPoolGroup`](../templatereference/aws-resource-cognito-userpoolgroup.md)

Yes

Yes

Yes

[`AWS::Cognito::UserPoolIdentityProvider`](../templatereference/aws-resource-cognito-userpoolidentityprovider.md)

Yes

Yes

No

[`AWS::Cognito::UserPoolResourceServer`](../templatereference/aws-resource-cognito-userpoolresourceserver.md)

Yes

Yes

No

[`AWS::Cognito::UserPoolRiskConfigurationAttachment`](../templatereference/aws-resource-cognito-userpoolriskconfigurationattachment.md)

Yes

Yes

No

[`AWS::Cognito::UserPoolUICustomizationAttachment`](../templatereference/aws-resource-cognito-userpooluicustomizationattachment.md)

Yes

Yes

No

[`AWS::Cognito::UserPoolUser`](../templatereference/aws-resource-cognito-userpooluser.md)

Yes

No

No

[`AWS::Cognito::UserPoolUserToGroupAttachment`](../templatereference/aws-resource-cognito-userpoolusertogroupattachment.md)

Yes

Yes

No

[`AWS::Comprehend::DocumentClassifier`](../templatereference/aws-resource-comprehend-documentclassifier.md)

Yes

Yes

Yes

[`AWS::Comprehend::Flywheel`](../templatereference/aws-resource-comprehend-flywheel.md)

Yes

Yes

Yes

[`AWS::ComputeOptimizer::AutomationRule`](../templatereference/aws-resource-computeoptimizer-automationrule.md)

Yes

Yes

No

[`AWS::Config::AggregationAuthorization`](../templatereference/aws-resource-config-aggregationauthorization.md)

Yes

Yes

Yes

[`AWS::Config::ConfigRule`](../templatereference/aws-resource-config-configrule.md)

Yes

Yes

Yes

[`AWS::Config::ConfigurationAggregator`](../templatereference/aws-resource-config-configurationaggregator.md)

Yes

Yes

Yes

[`AWS::Config::ConformancePack`](../templatereference/aws-resource-config-conformancepack.md)

Yes

Yes

Yes

[`AWS::Config::OrganizationConformancePack`](../templatereference/aws-resource-config-organizationconformancepack.md)

Yes

Yes

Yes

[`AWS::Config::StoredQuery`](../templatereference/aws-resource-config-storedquery.md)

Yes

Yes

Yes

[`AWS::Connect::AgentStatus`](../templatereference/aws-resource-connect-agentstatus.md)

Yes

Yes

Yes

[`AWS::Connect::ApprovedOrigin`](../templatereference/aws-resource-connect-approvedorigin.md)

Yes

Yes

Yes

[`AWS::Connect::ContactFlow`](../templatereference/aws-resource-connect-contactflow.md)

Yes

Yes

Yes

[`AWS::Connect::ContactFlowModule`](../templatereference/aws-resource-connect-contactflowmodule.md)

Yes

Yes

Yes

[`AWS::Connect::ContactFlowModuleAlias`](../templatereference/aws-resource-connect-contactflowmodulealias.md)

Yes

Yes

No

[`AWS::Connect::ContactFlowModuleVersion`](../templatereference/aws-resource-connect-contactflowmoduleversion.md)

Yes

Yes

No

[`AWS::Connect::ContactFlowVersion`](../templatereference/aws-resource-connect-contactflowversion.md)

Yes

Yes

Yes

[`AWS::Connect::DataTable`](../templatereference/aws-resource-connect-datatable.md)

Yes

Yes

No

[`AWS::Connect::DataTableAttribute`](../templatereference/aws-resource-connect-datatableattribute.md)

Yes

Yes

No

[`AWS::Connect::DataTableRecord`](../templatereference/aws-resource-connect-datatablerecord.md)

Yes

Yes

No

[`AWS::Connect::EmailAddress`](../templatereference/aws-resource-connect-emailaddress.md)

Yes

Yes

Yes

[`AWS::Connect::EvaluationForm`](../templatereference/aws-resource-connect-evaluationform.md)

Yes

Yes

Yes

[`AWS::Connect::HoursOfOperation`](../templatereference/aws-resource-connect-hoursofoperation.md)

Yes

Yes

Yes

[`AWS::Connect::Instance`](../templatereference/aws-resource-connect-instance.md)

Yes

Yes

Yes

[`AWS::Connect::InstanceStorageConfig`](../templatereference/aws-resource-connect-instancestorageconfig.md)

Yes

Yes

Yes

[`AWS::Connect::IntegrationAssociation`](../templatereference/aws-resource-connect-integrationassociation.md)

Yes

Yes

Yes

[`AWS::Connect::Notification`](../templatereference/aws-resource-connect-notification.md)

Yes

Yes

No

[`AWS::Connect::PhoneNumber`](../templatereference/aws-resource-connect-phonenumber.md)

Yes

Yes

Yes

[`AWS::Connect::PredefinedAttribute`](../templatereference/aws-resource-connect-predefinedattribute.md)

Yes

Yes

Yes

[`AWS::Connect::Prompt`](../templatereference/aws-resource-connect-prompt.md)

Yes

Yes

Yes

[`AWS::Connect::Queue`](../templatereference/aws-resource-connect-queue.md)

Yes

Yes

Yes

[`AWS::Connect::QuickConnect`](../templatereference/aws-resource-connect-quickconnect.md)

Yes

Yes

Yes

[`AWS::Connect::RoutingProfile`](../templatereference/aws-resource-connect-routingprofile.md)

Yes

Yes

Yes

[`AWS::Connect::Rule`](../templatereference/aws-resource-connect-rule.md)

Yes

Yes

No

[`AWS::Connect::SecurityKey`](../templatereference/aws-resource-connect-securitykey.md)

Yes

Yes

Yes

[`AWS::Connect::SecurityProfile`](../templatereference/aws-resource-connect-securityprofile.md)

Yes

Yes

Yes

[`AWS::Connect::TaskTemplate`](../templatereference/aws-resource-connect-tasktemplate.md)

Yes

Yes

Yes

[`AWS::Connect::TrafficDistributionGroup`](../templatereference/aws-resource-connect-trafficdistributiongroup.md)

Yes

Yes

Yes

[`AWS::Connect::User`](../templatereference/aws-resource-connect-user.md)

Yes

Yes

Yes

[`AWS::Connect::UserHierarchyGroup`](../templatereference/aws-resource-connect-userhierarchygroup.md)

Yes

Yes

Yes

[`AWS::Connect::UserHierarchyStructure`](../templatereference/aws-resource-connect-userhierarchystructure.md)

Yes

Yes

No

[`AWS::Connect::View`](../templatereference/aws-resource-connect-view.md)

Yes

Yes

Yes

[`AWS::Connect::ViewVersion`](../templatereference/aws-resource-connect-viewversion.md)

Yes

Yes

Yes

[`AWS::Connect::Workspace`](../templatereference/aws-resource-connect-workspace.md)

Yes

Yes

No

[`AWS::ConnectCampaigns::Campaign`](../templatereference/aws-resource-connectcampaigns-campaign.md)

Yes

No

Yes

[`AWS::ConnectCampaignsV2::Campaign`](../templatereference/aws-resource-connectcampaignsv2-campaign.md)

Yes

Yes

Yes

[`AWS::ControlTower::EnabledBaseline`](../templatereference/aws-resource-controltower-enabledbaseline.md)

Yes

Yes

No

[`AWS::ControlTower::EnabledControl`](../templatereference/aws-resource-controltower-enabledcontrol.md)

Yes

Yes

Yes

[`AWS::ControlTower::LandingZone`](../templatereference/aws-resource-controltower-landingzone.md)

Yes

Yes

No

[`AWS::CustomerProfiles::CalculatedAttributeDefinition`](../templatereference/aws-resource-customerprofiles-calculatedattributedefinition.md)

Yes

Yes

Yes

[`AWS::CustomerProfiles::Domain`](../templatereference/aws-resource-customerprofiles-domain.md)

Yes

Yes

Yes

[`AWS::CustomerProfiles::EventStream`](../templatereference/aws-resource-customerprofiles-eventstream.md)

Yes

Yes

Yes

[`AWS::CustomerProfiles::EventTrigger`](../templatereference/aws-resource-customerprofiles-eventtrigger.md)

Yes

Yes

No

[`AWS::CustomerProfiles::Integration`](../templatereference/aws-resource-customerprofiles-integration.md)

Yes

Yes

Yes

[`AWS::CustomerProfiles::ObjectType`](../templatereference/aws-resource-customerprofiles-objecttype.md)

Yes

Yes

Yes

[`AWS::CustomerProfiles::SegmentDefinition`](../templatereference/aws-resource-customerprofiles-segmentdefinition.md)

Yes

Yes

No

[`AWS::DMS::DataMigration`](../templatereference/aws-resource-dms-datamigration.md)

Yes

Yes

No

[`AWS::DMS::DataProvider`](../templatereference/aws-resource-dms-dataprovider.md)

Yes

Yes

No

[`AWS::DMS::InstanceProfile`](../templatereference/aws-resource-dms-instanceprofile.md)

Yes

Yes

No

[`AWS::DMS::MigrationProject`](../templatereference/aws-resource-dms-migrationproject.md)

Yes

Yes

No

[`AWS::DMS::ReplicationConfig`](../templatereference/aws-resource-dms-replicationconfig.md)

Yes

Yes

Yes

[`AWS::DSQL::Cluster`](../templatereference/aws-resource-dsql-cluster.md)

Yes

Yes

No

[`AWS::DataBrew::Dataset`](../templatereference/aws-resource-databrew-dataset.md)

Yes

Yes

Yes

[`AWS::DataBrew::Job`](../templatereference/aws-resource-databrew-job.md)

Yes

Yes

Yes

[`AWS::DataBrew::Project`](../templatereference/aws-resource-databrew-project.md)

Yes

Yes

Yes

[`AWS::DataBrew::Recipe`](../templatereference/aws-resource-databrew-recipe.md)

Yes

Yes

Yes

[`AWS::DataBrew::Ruleset`](../templatereference/aws-resource-databrew-ruleset.md)

Yes

Yes

Yes

[`AWS::DataBrew::Schedule`](../templatereference/aws-resource-databrew-schedule.md)

Yes

Yes

Yes

[`AWS::DataPipeline::Pipeline`](../templatereference/aws-resource-datapipeline-pipeline.md)

Yes

Yes

Yes

[`AWS::DataSync::Agent`](../templatereference/aws-resource-datasync-agent.md)

Yes

Yes

Yes

[`AWS::DataSync::LocationAzureBlob`](../templatereference/aws-resource-datasync-locationazureblob.md)

Yes

Yes

Yes

[`AWS::DataSync::LocationEFS`](../templatereference/aws-resource-datasync-locationefs.md)

Yes

Yes

Yes

[`AWS::DataSync::LocationFSxLustre`](../templatereference/aws-resource-datasync-locationfsxlustre.md)

Yes

Yes

Yes

[`AWS::DataSync::LocationFSxONTAP`](../templatereference/aws-resource-datasync-locationfsxontap.md)

Yes

Yes

Yes

[`AWS::DataSync::LocationFSxOpenZFS`](../templatereference/aws-resource-datasync-locationfsxopenzfs.md)

Yes

Yes

Yes

[`AWS::DataSync::LocationFSxWindows`](../templatereference/aws-resource-datasync-locationfsxwindows.md)

Yes

Yes

Yes

[`AWS::DataSync::LocationHDFS`](../templatereference/aws-resource-datasync-locationhdfs.md)

Yes

Yes

Yes

[`AWS::DataSync::LocationNFS`](../templatereference/aws-resource-datasync-locationnfs.md)

Yes

Yes

Yes

[`AWS::DataSync::LocationObjectStorage`](../templatereference/aws-resource-datasync-locationobjectstorage.md)

Yes

Yes

Yes

[`AWS::DataSync::LocationS3`](../templatereference/aws-resource-datasync-locations3.md)

Yes

Yes

Yes

[`AWS::DataSync::LocationSMB`](../templatereference/aws-resource-datasync-locationsmb.md)

Yes

Yes

Yes

[`AWS::DataSync::Task`](../templatereference/aws-resource-datasync-task.md)

Yes

Yes

Yes

[`AWS::DataZone::Connection`](../templatereference/aws-resource-datazone-connection.md)

Yes

Yes

No

[`AWS::DataZone::DataSource`](../templatereference/aws-resource-datazone-datasource.md)

Yes

Yes

No

[`AWS::DataZone::Domain`](../templatereference/aws-resource-datazone-domain.md)

Yes

Yes

No

[`AWS::DataZone::DomainUnit`](../templatereference/aws-resource-datazone-domainunit.md)

Yes

Yes

No

[`AWS::DataZone::Environment`](../templatereference/aws-resource-datazone-environment.md)

Yes

Yes

No

[`AWS::DataZone::EnvironmentActions`](../templatereference/aws-resource-datazone-environmentactions.md)

Yes

Yes

No

[`AWS::DataZone::EnvironmentBlueprintConfiguration`](../templatereference/aws-resource-datazone-environmentblueprintconfiguration.md)

Yes

Yes

No

[`AWS::DataZone::EnvironmentProfile`](../templatereference/aws-resource-datazone-environmentprofile.md)

Yes

Yes

No

[`AWS::DataZone::FormType`](../templatereference/aws-resource-datazone-formtype.md)

Yes

Yes

No

[`AWS::DataZone::GroupProfile`](../templatereference/aws-resource-datazone-groupprofile.md)

Yes

Yes

No

[`AWS::DataZone::Owner`](../templatereference/aws-resource-datazone-owner.md)

Yes

Yes

No

[`AWS::DataZone::PolicyGrant`](../templatereference/aws-resource-datazone-policygrant.md)

Yes

Yes

No

[`AWS::DataZone::Project`](../templatereference/aws-resource-datazone-project.md)

Yes

Yes

No

[`AWS::DataZone::ProjectMembership`](../templatereference/aws-resource-datazone-projectmembership.md)

Yes

Yes

No

[`AWS::DataZone::ProjectProfile`](../templatereference/aws-resource-datazone-projectprofile.md)

Yes

Yes

No

[`AWS::DataZone::SubscriptionTarget`](../templatereference/aws-resource-datazone-subscriptiontarget.md)

Yes

Yes

No

[`AWS::DataZone::UserProfile`](../templatereference/aws-resource-datazone-userprofile.md)

Yes

Yes

No

[`AWS::Deadline::Farm`](../templatereference/aws-resource-deadline-farm.md)

Yes

Yes

Yes

[`AWS::Deadline::Fleet`](../templatereference/aws-resource-deadline-fleet.md)

Yes

Yes

Yes

[`AWS::Deadline::LicenseEndpoint`](../templatereference/aws-resource-deadline-licenseendpoint.md)

Yes

Yes

Yes

[`AWS::Deadline::Limit`](../templatereference/aws-resource-deadline-limit.md)

Yes

Yes

Yes

[`AWS::Deadline::MeteredProduct`](../templatereference/aws-resource-deadline-meteredproduct.md)

Yes

Yes

No

[`AWS::Deadline::Monitor`](../templatereference/aws-resource-deadline-monitor.md)

Yes

Yes

Yes

[`AWS::Deadline::Queue`](../templatereference/aws-resource-deadline-queue.md)

Yes

Yes

Yes

[`AWS::Deadline::QueueEnvironment`](../templatereference/aws-resource-deadline-queueenvironment.md)

Yes

Yes

No

[`AWS::Deadline::QueueFleetAssociation`](../templatereference/aws-resource-deadline-queuefleetassociation.md)

Yes

Yes

No

[`AWS::Deadline::QueueLimitAssociation`](../templatereference/aws-resource-deadline-queuelimitassociation.md)

Yes

Yes

No

[`AWS::Deadline::StorageProfile`](../templatereference/aws-resource-deadline-storageprofile.md)

Yes

Yes

Yes

[`AWS::Detective::Graph`](../templatereference/aws-resource-detective-graph.md)

Yes

Yes

Yes

[`AWS::Detective::MemberInvitation`](../templatereference/aws-resource-detective-memberinvitation.md)

Yes

Yes

Yes

[`AWS::Detective::OrganizationAdmin`](../templatereference/aws-resource-detective-organizationadmin.md)

Yes

Yes

Yes

[`AWS::DevOpsAgent::AgentSpace`](../templatereference/aws-resource-devopsagent-agentspace.md)

Yes

Yes

No

[`AWS::DevOpsAgent::Association`](../templatereference/aws-resource-devopsagent-association.md)

Yes

Yes

No

[`AWS::DevOpsAgent::Service`](../templatereference/aws-resource-devopsagent-service.md)

Yes

Yes

No

[`AWS::DevOpsGuru::LogAnomalyDetectionIntegration`](../templatereference/aws-resource-devopsguru-loganomalydetectionintegration.md)

Yes

Yes

Yes

[`AWS::DevOpsGuru::NotificationChannel`](../templatereference/aws-resource-devopsguru-notificationchannel.md)

Yes

Yes

Yes

[`AWS::DevOpsGuru::ResourceCollection`](../templatereference/aws-resource-devopsguru-resourcecollection.md)

Yes

Yes

Yes

[`AWS::DeviceFarm::DevicePool`](../templatereference/aws-resource-devicefarm-devicepool.md)

Yes

Yes

No

[`AWS::DeviceFarm::InstanceProfile`](../templatereference/aws-resource-devicefarm-instanceprofile.md)

Yes

Yes

No

[`AWS::DeviceFarm::NetworkProfile`](../templatereference/aws-resource-devicefarm-networkprofile.md)

Yes

Yes

No

[`AWS::DeviceFarm::Project`](../templatereference/aws-resource-devicefarm-project.md)

Yes

Yes

No

[`AWS::DeviceFarm::TestGridProject`](../templatereference/aws-resource-devicefarm-testgridproject.md)

Yes

Yes

No

[`AWS::DeviceFarm::VPCEConfiguration`](../templatereference/aws-resource-devicefarm-vpceconfiguration.md)

Yes

Yes

No

[`AWS::DirectConnect::Connection`](../templatereference/aws-resource-directconnect-connection.md)

Yes

Yes

No

[`AWS::DirectConnect::DirectConnectGateway`](../templatereference/aws-resource-directconnect-directconnectgateway.md)

Yes

Yes

No

[`AWS::DirectConnect::DirectConnectGatewayAssociation`](../templatereference/aws-resource-directconnect-directconnectgatewayassociation.md)

Yes

Yes

No

[`AWS::DirectConnect::Lag`](../templatereference/aws-resource-directconnect-lag.md)

Yes

Yes

No

[`AWS::DirectConnect::PrivateVirtualInterface`](../templatereference/aws-resource-directconnect-privatevirtualinterface.md)

Yes

Yes

No

[`AWS::DirectConnect::PublicVirtualInterface`](../templatereference/aws-resource-directconnect-publicvirtualinterface.md)

Yes

Yes

No

[`AWS::DirectConnect::TransitVirtualInterface`](../templatereference/aws-resource-directconnect-transitvirtualinterface.md)

Yes

Yes

No

[`AWS::DirectoryService::SimpleAD`](../templatereference/aws-resource-directoryservice-simplead.md)

Yes

Yes

Yes

[`AWS::DocDB::GlobalCluster`](../templatereference/aws-resource-docdb-globalcluster.md)

Yes

Yes

No

[`AWS::DocDBElastic::Cluster`](../templatereference/aws-resource-docdbelastic-cluster.md)

Yes

Yes

Yes

[`AWS::DynamoDB::GlobalTable`](../templatereference/aws-resource-dynamodb-globaltable.md)

Yes

Yes

No

[`AWS::DynamoDB::Table`](../templatereference/aws-resource-dynamodb-table.md)

Yes

Yes

Yes

[`AWS::EC2::CapacityManagerDataExport`](../templatereference/aws-resource-ec2-capacitymanagerdataexport.md)

Yes

Yes

No

[`AWS::EC2::CapacityReservation`](../templatereference/aws-resource-ec2-capacityreservation.md)

Yes

Yes

Yes

[`AWS::EC2::CapacityReservationFleet`](../templatereference/aws-resource-ec2-capacityreservationfleet.md)

Yes

Yes

Yes

[`AWS::EC2::CarrierGateway`](../templatereference/aws-resource-ec2-carriergateway.md)

Yes

Yes

Yes

[`AWS::EC2::CustomerGateway`](../templatereference/aws-resource-ec2-customergateway.md)

Yes

Yes

Yes

[`AWS::EC2::DHCPOptions`](../templatereference/aws-resource-ec2-dhcpoptions.md)

Yes

Yes

Yes

[`AWS::EC2::EC2Fleet`](../templatereference/aws-resource-ec2-ec2fleet.md)

Yes

Yes

No

[`AWS::EC2::EIP`](../templatereference/aws-resource-ec2-eip.md)

Yes

Yes

Yes

[`AWS::EC2::EIPAssociation`](../templatereference/aws-resource-ec2-eipassociation.md)

Yes

No

Yes

[`AWS::EC2::EgressOnlyInternetGateway`](../templatereference/aws-resource-ec2-egressonlyinternetgateway.md)

Yes

Yes

Yes

[`AWS::EC2::EnclaveCertificateIamRoleAssociation`](../templatereference/aws-resource-ec2-enclavecertificateiamroleassociation.md)

Yes

Yes

Yes

[`AWS::EC2::FlowLog`](../templatereference/aws-resource-ec2-flowlog.md)

Yes

Yes

Yes

[`AWS::EC2::GatewayRouteTableAssociation`](../templatereference/aws-resource-ec2-gatewayroutetableassociation.md)

Yes

Yes

No

[`AWS::EC2::Host`](../templatereference/aws-resource-ec2-host.md)

Yes

Yes

Yes

[`AWS::EC2::IPAM`](../templatereference/aws-resource-ec2-ipam.md)

Yes

Yes

Yes

[`AWS::EC2::IPAMAllocation`](../templatereference/aws-resource-ec2-ipamallocation.md)

Yes

Yes

Yes

[`AWS::EC2::IPAMPool`](../templatereference/aws-resource-ec2-ipampool.md)

Yes

Yes

Yes

[`AWS::EC2::IPAMPoolCidr`](../templatereference/aws-resource-ec2-ipampoolcidr.md)

Yes

Yes

Yes

[`AWS::EC2::IPAMPrefixListResolver`](../templatereference/aws-resource-ec2-ipamprefixlistresolver.md)

Yes

Yes

No

[`AWS::EC2::IPAMPrefixListResolverTarget`](../templatereference/aws-resource-ec2-ipamprefixlistresolvertarget.md)

Yes

Yes

No

[`AWS::EC2::IPAMResourceDiscovery`](../templatereference/aws-resource-ec2-ipamresourcediscovery.md)

Yes

Yes

Yes

[`AWS::EC2::IPAMResourceDiscoveryAssociation`](../templatereference/aws-resource-ec2-ipamresourcediscoveryassociation.md)

Yes

Yes

Yes

[`AWS::EC2::IPAMScope`](../templatereference/aws-resource-ec2-ipamscope.md)

Yes

Yes

Yes

[`AWS::EC2::Instance`](../templatereference/aws-resource-ec2-instance.md)

Yes

Yes

Yes

[`AWS::EC2::InstanceConnectEndpoint`](../templatereference/aws-resource-ec2-instanceconnectendpoint.md)

Yes

Yes

Yes

[`AWS::EC2::InternetGateway`](../templatereference/aws-resource-ec2-internetgateway.md)

Yes

Yes

Yes

[`AWS::EC2::IpPoolRouteTableAssociation`](../templatereference/aws-resource-ec2-ippoolroutetableassociation.md)

Yes

Yes

No

[`AWS::EC2::KeyPair`](../templatereference/aws-resource-ec2-keypair.md)

Yes

Yes

Yes

[`AWS::EC2::LaunchTemplate`](../templatereference/aws-resource-ec2-launchtemplate.md)

Yes

Yes

No

[`AWS::EC2::LocalGatewayRoute`](../templatereference/aws-resource-ec2-localgatewayroute.md)

Yes

Yes

Yes

[`AWS::EC2::LocalGatewayRouteTable`](../templatereference/aws-resource-ec2-localgatewayroutetable.md)

Yes

Yes

Yes

[`AWS::EC2::LocalGatewayRouteTableVPCAssociation`](../templatereference/aws-resource-ec2-localgatewayroutetablevpcassociation.md)

Yes

Yes

Yes

[`AWS::EC2::LocalGatewayRouteTableVirtualInterfaceGroupAssociation`](../templatereference/aws-resource-ec2-localgatewayroutetablevirtualinterfacegroupassociation.md)

Yes

Yes

Yes

[`AWS::EC2::LocalGatewayVirtualInterface`](../templatereference/aws-resource-ec2-localgatewayvirtualinterface.md)

Yes

Yes

No

[`AWS::EC2::LocalGatewayVirtualInterfaceGroup`](../templatereference/aws-resource-ec2-localgatewayvirtualinterfacegroup.md)

Yes

Yes

No

[`AWS::EC2::NatGateway`](../templatereference/aws-resource-ec2-natgateway.md)

Yes

Yes

Yes

[`AWS::EC2::NetworkAcl`](../templatereference/aws-resource-ec2-networkacl.md)

Yes

Yes

Yes

[`AWS::EC2::NetworkInsightsAccessScope`](../templatereference/aws-resource-ec2-networkinsightsaccessscope.md)

Yes

Yes

No

[`AWS::EC2::NetworkInsightsAccessScopeAnalysis`](../templatereference/aws-resource-ec2-networkinsightsaccessscopeanalysis.md)

Yes

Yes

Yes

[`AWS::EC2::NetworkInsightsAnalysis`](../templatereference/aws-resource-ec2-networkinsightsanalysis.md)

Yes

Yes

No

[`AWS::EC2::NetworkInsightsPath`](../templatereference/aws-resource-ec2-networkinsightspath.md)

Yes

Yes

Yes

[`AWS::EC2::NetworkInterface`](../templatereference/aws-resource-ec2-networkinterface.md)

Yes

Yes

Yes

[`AWS::EC2::NetworkInterfaceAttachment`](../templatereference/aws-resource-ec2-networkinterfaceattachment.md)

Yes

Yes

Yes

[`AWS::EC2::NetworkPerformanceMetricSubscription`](../templatereference/aws-resource-ec2-networkperformancemetricsubscription.md)

Yes

Yes

Yes

[`AWS::EC2::PlacementGroup`](../templatereference/aws-resource-ec2-placementgroup.md)

Yes

Yes

Yes

[`AWS::EC2::PrefixList`](../templatereference/aws-resource-ec2-prefixlist.md)

Yes

Yes

No

[`AWS::EC2::Route`](../templatereference/aws-resource-ec2-route.md)

Yes

Yes

Yes

[`AWS::EC2::RouteServer`](../templatereference/aws-resource-ec2-routeserver.md)

Yes

Yes

No

[`AWS::EC2::RouteServerAssociation`](../templatereference/aws-resource-ec2-routeserverassociation.md)

Yes

Yes

No

[`AWS::EC2::RouteServerEndpoint`](../templatereference/aws-resource-ec2-routeserverendpoint.md)

Yes

Yes

No

[`AWS::EC2::RouteServerPeer`](../templatereference/aws-resource-ec2-routeserverpeer.md)

Yes

Yes

No

[`AWS::EC2::RouteServerPropagation`](../templatereference/aws-resource-ec2-routeserverpropagation.md)

Yes

Yes

No

[`AWS::EC2::RouteTable`](../templatereference/aws-resource-ec2-routetable.md)

Yes

Yes

Yes

[`AWS::EC2::SecurityGroup`](../templatereference/aws-resource-ec2-securitygroup.md)

Yes

Yes

Yes

[`AWS::EC2::SecurityGroupEgress`](../templatereference/aws-resource-ec2-securitygroupegress.md)

Yes

No

No

[`AWS::EC2::SecurityGroupIngress`](../templatereference/aws-resource-ec2-securitygroupingress.md)

Yes

No

No

[`AWS::EC2::SecurityGroupVpcAssociation`](../templatereference/aws-resource-ec2-securitygroupvpcassociation.md)

Yes

Yes

No

[`AWS::EC2::SnapshotBlockPublicAccess`](../templatereference/aws-resource-ec2-snapshotblockpublicaccess.md)

Yes

Yes

No

[`AWS::EC2::SpotFleet`](../templatereference/aws-resource-ec2-spotfleet.md)

Yes

Yes

No

[`AWS::EC2::Subnet`](../templatereference/aws-resource-ec2-subnet.md)

Yes

Yes

Yes

[`AWS::EC2::SubnetCidrBlock`](../templatereference/aws-resource-ec2-subnetcidrblock.md)

Yes

No

Yes

[`AWS::EC2::SubnetNetworkAclAssociation`](../templatereference/aws-resource-ec2-subnetnetworkaclassociation.md)

Yes

Yes

Yes

[`AWS::EC2::SubnetRouteTableAssociation`](../templatereference/aws-resource-ec2-subnetroutetableassociation.md)

Yes

No

Yes

[`AWS::EC2::TrafficMirrorFilter`](../templatereference/aws-resource-ec2-trafficmirrorfilter.md)

Yes

Yes

No

[`AWS::EC2::TrafficMirrorFilterRule`](../templatereference/aws-resource-ec2-trafficmirrorfilterrule.md)

Yes

Yes

No

[`AWS::EC2::TrafficMirrorSession`](../templatereference/aws-resource-ec2-trafficmirrorsession.md)

Yes

Yes

No

[`AWS::EC2::TrafficMirrorTarget`](../templatereference/aws-resource-ec2-trafficmirrortarget.md)

Yes

Yes

No

[`AWS::EC2::TransitGateway`](../templatereference/aws-resource-ec2-transitgateway.md)

Yes

Yes

Yes

[`AWS::EC2::TransitGatewayAttachment`](../templatereference/aws-resource-ec2-transitgatewayattachment.md)

Yes

Yes

Yes

[`AWS::EC2::TransitGatewayConnect`](../templatereference/aws-resource-ec2-transitgatewayconnect.md)

Yes

Yes

Yes

[`AWS::EC2::TransitGatewayConnectPeer`](../templatereference/aws-resource-ec2-transitgatewayconnectpeer.md)

Yes

Yes

No

[`AWS::EC2::TransitGatewayMeteringPolicy`](../templatereference/aws-resource-ec2-transitgatewaymeteringpolicy.md)

Yes

Yes

No

[`AWS::EC2::TransitGatewayMeteringPolicyEntry`](../templatereference/aws-resource-ec2-transitgatewaymeteringpolicyentry.md)

Yes

Yes

No

[`AWS::EC2::TransitGatewayMulticastDomain`](../templatereference/aws-resource-ec2-transitgatewaymulticastdomain.md)

Yes

Yes

Yes

[`AWS::EC2::TransitGatewayMulticastDomainAssociation`](../templatereference/aws-resource-ec2-transitgatewaymulticastdomainassociation.md)

Yes

Yes

Yes

[`AWS::EC2::TransitGatewayMulticastGroupMember`](../templatereference/aws-resource-ec2-transitgatewaymulticastgroupmember.md)

Yes

Yes

Yes

[`AWS::EC2::TransitGatewayMulticastGroupSource`](../templatereference/aws-resource-ec2-transitgatewaymulticastgroupsource.md)

Yes

Yes

Yes

[`AWS::EC2::TransitGatewayPeeringAttachment`](../templatereference/aws-resource-ec2-transitgatewaypeeringattachment.md)

Yes

Yes

Yes

[`AWS::EC2::TransitGatewayRoute`](../templatereference/aws-resource-ec2-transitgatewayroute.md)

Yes

Yes

No

[`AWS::EC2::TransitGatewayRouteTable`](../templatereference/aws-resource-ec2-transitgatewayroutetable.md)

Yes

Yes

Yes

[`AWS::EC2::TransitGatewayRouteTableAssociation`](../templatereference/aws-resource-ec2-transitgatewayroutetableassociation.md)

Yes

Yes

No

[`AWS::EC2::TransitGatewayRouteTablePropagation`](../templatereference/aws-resource-ec2-transitgatewayroutetablepropagation.md)

Yes

Yes

No

[`AWS::EC2::TransitGatewayVpcAttachment`](../templatereference/aws-resource-ec2-transitgatewayvpcattachment.md)

Yes

Yes

Yes

[`AWS::EC2::VPC`](../templatereference/aws-resource-ec2-vpc.md)

Yes

Yes

Yes

[`AWS::EC2::VPCBlockPublicAccessExclusion`](../templatereference/aws-resource-ec2-vpcblockpublicaccessexclusion.md)

Yes

Yes

No

[`AWS::EC2::VPCBlockPublicAccessOptions`](../templatereference/aws-resource-ec2-vpcblockpublicaccessoptions.md)

Yes

Yes

No

[`AWS::EC2::VPCCidrBlock`](../templatereference/aws-resource-ec2-vpccidrblock.md)

Yes

Yes

No

[`AWS::EC2::VPCDHCPOptionsAssociation`](../templatereference/aws-resource-ec2-vpcdhcpoptionsassociation.md)

Yes

Yes

Yes

[`AWS::EC2::VPCEncryptionControl`](../templatereference/aws-resource-ec2-vpcencryptioncontrol.md)

Yes

Yes

No

[`AWS::EC2::VPCEndpoint`](../templatereference/aws-resource-ec2-vpcendpoint.md)

Yes

No

Yes

[`AWS::EC2::VPCEndpointConnectionNotification`](../templatereference/aws-resource-ec2-vpcendpointconnectionnotification.md)

Yes

Yes

Yes

[`AWS::EC2::VPCEndpointService`](../templatereference/aws-resource-ec2-vpcendpointservice.md)

Yes

Yes

Yes

[`AWS::EC2::VPCEndpointServicePermissions`](../templatereference/aws-resource-ec2-vpcendpointservicepermissions.md)

Yes

Yes

No

[`AWS::EC2::VPCGatewayAttachment`](../templatereference/aws-resource-ec2-vpcgatewayattachment.md)

Yes

No

No

[`AWS::EC2::VPCPeeringConnection`](../templatereference/aws-resource-ec2-vpcpeeringconnection.md)

Yes

Yes

Yes

[`AWS::EC2::VPNConcentrator`](../templatereference/aws-resource-ec2-vpnconcentrator.md)

Yes

Yes

No

[`AWS::EC2::VPNConnection`](../templatereference/aws-resource-ec2-vpnconnection.md)

Yes

Yes

Yes

[`AWS::EC2::VPNConnectionRoute`](../templatereference/aws-resource-ec2-vpnconnectionroute.md)

Yes

Yes

Yes

[`AWS::EC2::VPNGateway`](../templatereference/aws-resource-ec2-vpngateway.md)

Yes

Yes

Yes

[`AWS::EC2::VerifiedAccessEndpoint`](../templatereference/aws-resource-ec2-verifiedaccessendpoint.md)

Yes

Yes

Yes

[`AWS::EC2::VerifiedAccessGroup`](../templatereference/aws-resource-ec2-verifiedaccessgroup.md)

Yes

Yes

Yes

[`AWS::EC2::VerifiedAccessInstance`](../templatereference/aws-resource-ec2-verifiedaccessinstance.md)

Yes

Yes

Yes

[`AWS::EC2::VerifiedAccessTrustProvider`](../templatereference/aws-resource-ec2-verifiedaccesstrustprovider.md)

Yes

Yes

Yes

[`AWS::EC2::Volume`](../templatereference/aws-resource-ec2-volume.md)

Yes

Yes

Yes

[`AWS::EC2::VolumeAttachment`](../templatereference/aws-resource-ec2-volumeattachment.md)

Yes

Yes

Yes

[`AWS::ECR::PublicRepository`](../templatereference/aws-resource-ecr-publicrepository.md)

Yes

Yes

Yes

[`AWS::ECR::PullThroughCacheRule`](../templatereference/aws-resource-ecr-pullthroughcacherule.md)

Yes

Yes

Yes

[`AWS::ECR::PullTimeUpdateExclusion`](../templatereference/aws-resource-ecr-pulltimeupdateexclusion.md)

Yes

Yes

No

[`AWS::ECR::RegistryPolicy`](../templatereference/aws-resource-ecr-registrypolicy.md)

Yes

Yes

Yes

[`AWS::ECR::RegistryScanningConfiguration`](../templatereference/aws-resource-ecr-registryscanningconfiguration.md)

Yes

Yes

No

[`AWS::ECR::ReplicationConfiguration`](../templatereference/aws-resource-ecr-replicationconfiguration.md)

Yes

Yes

Yes

[`AWS::ECR::Repository`](../templatereference/aws-resource-ecr-repository.md)

Yes

Yes

Yes

[`AWS::ECR::RepositoryCreationTemplate`](../templatereference/aws-resource-ecr-repositorycreationtemplate.md)

Yes

Yes

No

[`AWS::ECR::SigningConfiguration`](../templatereference/aws-resource-ecr-signingconfiguration.md)

Yes

Yes

No

[`AWS::ECS::CapacityProvider`](../templatereference/aws-resource-ecs-capacityprovider.md)

Yes

Yes

No

[`AWS::ECS::Cluster`](../templatereference/aws-resource-ecs-cluster.md)

Yes

Yes

Yes

[`AWS::ECS::ClusterCapacityProviderAssociations`](../templatereference/aws-resource-ecs-clustercapacityproviderassociations.md)

Yes

No

Yes

[`AWS::ECS::ExpressGatewayService`](../templatereference/aws-resource-ecs-expressgatewayservice.md)

Yes

Yes

No

[`AWS::ECS::PrimaryTaskSet`](../templatereference/aws-resource-ecs-primarytaskset.md)

Yes

Yes

No

[`AWS::ECS::Service`](../templatereference/aws-resource-ecs-service.md)

Yes

Yes

Yes

[`AWS::ECS::TaskDefinition`](../templatereference/aws-resource-ecs-taskdefinition.md)

Yes

Yes

Yes

[`AWS::ECS::TaskSet`](../templatereference/aws-resource-ecs-taskset.md)

Yes

Yes

No

[`AWS::EFS::AccessPoint`](../templatereference/aws-resource-efs-accesspoint.md)

Yes

Yes

Yes

[`AWS::EFS::FileSystem`](../templatereference/aws-resource-efs-filesystem.md)

Yes

Yes

Yes

[`AWS::EFS::MountTarget`](../templatereference/aws-resource-efs-mounttarget.md)

Yes

Yes

Yes

[`AWS::EKS::AccessEntry`](../templatereference/aws-resource-eks-accessentry.md)

Yes

Yes

No

[`AWS::EKS::Addon`](../templatereference/aws-resource-eks-addon.md)

Yes

Yes

Yes

[`AWS::EKS::Capability`](../templatereference/aws-resource-eks-capability.md)

Yes

Yes

No

[`AWS::EKS::Cluster`](../templatereference/aws-resource-eks-cluster.md)

Yes

No

Yes

[`AWS::EKS::FargateProfile`](../templatereference/aws-resource-eks-fargateprofile.md)

Yes

Yes

Yes

[`AWS::EKS::IdentityProviderConfig`](../templatereference/aws-resource-eks-identityproviderconfig.md)

Yes

Yes

Yes

[`AWS::EKS::Nodegroup`](../templatereference/aws-resource-eks-nodegroup.md)

Yes

No

Yes

[`AWS::EKS::PodIdentityAssociation`](../templatereference/aws-resource-eks-podidentityassociation.md)

Yes

Yes

No

[`AWS::EMR::SecurityConfiguration`](../templatereference/aws-resource-emr-securityconfiguration.md)

Yes

Yes

Yes

[`AWS::EMR::Step`](../templatereference/aws-resource-emr-step.md)

Yes

Yes

No

[`AWS::EMR::Studio`](../templatereference/aws-resource-emr-studio.md)

Yes

Yes

Yes

[`AWS::EMR::StudioSessionMapping`](../templatereference/aws-resource-emr-studiosessionmapping.md)

Yes

Yes

Yes

[`AWS::EMR::WALWorkspace`](../templatereference/aws-resource-emr-walworkspace.md)

Yes

Yes

No

[`AWS::EMRContainers::Endpoint`](../templatereference/aws-resource-emrcontainers-endpoint.md)

Yes

Yes

No

[`AWS::EMRContainers::SecurityConfiguration`](../templatereference/aws-resource-emrcontainers-securityconfiguration.md)

Yes

Yes

No

[`AWS::EMRContainers::VirtualCluster`](../templatereference/aws-resource-emrcontainers-virtualcluster.md)

Yes

Yes

No

[`AWS::EMRServerless::Application`](../templatereference/aws-resource-emrserverless-application.md)

Yes

Yes

Yes

[`AWS::EVS::Environment`](../templatereference/aws-resource-evs-environment.md)

Yes

Yes

No

[`AWS::ElastiCache::GlobalReplicationGroup`](../templatereference/aws-resource-elasticache-globalreplicationgroup.md)

Yes

Yes

No

[`AWS::ElastiCache::ParameterGroup`](../templatereference/aws-resource-elasticache-parametergroup.md)

Yes

Yes

No

[`AWS::ElastiCache::ReplicationGroup`](../templatereference/aws-resource-elasticache-replicationgroup.md)

Yes

Yes

No

[`AWS::ElastiCache::ServerlessCache`](../templatereference/aws-resource-elasticache-serverlesscache.md)

Yes

Yes

No

[`AWS::ElastiCache::SubnetGroup`](../templatereference/aws-resource-elasticache-subnetgroup.md)

Yes

Yes

Yes

[`AWS::ElastiCache::User`](../templatereference/aws-resource-elasticache-user.md)

Yes

Yes

Yes

[`AWS::ElastiCache::UserGroup`](../templatereference/aws-resource-elasticache-usergroup.md)

Yes

Yes

Yes

[`AWS::ElasticBeanstalk::Application`](../templatereference/aws-resource-elasticbeanstalk-application.md)

Yes

Yes

Yes

[`AWS::ElasticBeanstalk::ApplicationVersion`](../templatereference/aws-resource-elasticbeanstalk-applicationversion.md)

Yes

Yes

Yes

[`AWS::ElasticBeanstalk::ConfigurationTemplate`](../templatereference/aws-resource-elasticbeanstalk-configurationtemplate.md)

Yes

Yes

Yes

[`AWS::ElasticBeanstalk::Environment`](../templatereference/aws-resource-elasticbeanstalk-environment.md)

Yes

Yes

No

[`AWS::ElasticLoadBalancing::LoadBalancer`](../templatereference/aws-resource-elasticloadbalancing-loadbalancer.md)

Yes

Yes

No

[`AWS::ElasticLoadBalancingV2::Listener`](../templatereference/aws-resource-elasticloadbalancingv2-listener.md)

Yes

Yes

Yes

[`AWS::ElasticLoadBalancingV2::ListenerRule`](../templatereference/aws-resource-elasticloadbalancingv2-listenerrule.md)

Yes

Yes

Yes

[`AWS::ElasticLoadBalancingV2::LoadBalancer`](../templatereference/aws-resource-elasticloadbalancingv2-loadbalancer.md)

Yes

Yes

Yes

[`AWS::ElasticLoadBalancingV2::TargetGroup`](../templatereference/aws-resource-elasticloadbalancingv2-targetgroup.md)

Yes

Yes

Yes

[`AWS::ElasticLoadBalancingV2::TrustStore`](../templatereference/aws-resource-elasticloadbalancingv2-truststore.md)

Yes

Yes

No

[`AWS::ElasticLoadBalancingV2::TrustStoreRevocation`](../templatereference/aws-resource-elasticloadbalancingv2-truststorerevocation.md)

Yes

Yes

No

[`AWS::ElementalInference::Feed`](../templatereference/aws-resource-elementalinference-feed.md)

Yes

Yes

No

[`AWS::EntityResolution::IdMappingWorkflow`](../templatereference/aws-resource-entityresolution-idmappingworkflow.md)

Yes

Yes

Yes

[`AWS::EntityResolution::IdNamespace`](../templatereference/aws-resource-entityresolution-idnamespace.md)

Yes

Yes

No

[`AWS::EntityResolution::MatchingWorkflow`](../templatereference/aws-resource-entityresolution-matchingworkflow.md)

Yes

Yes

Yes

[`AWS::EntityResolution::PolicyStatement`](../templatereference/aws-resource-entityresolution-policystatement.md)

Yes

Yes

No

[`AWS::EntityResolution::SchemaMapping`](../templatereference/aws-resource-entityresolution-schemamapping.md)

Yes

Yes

Yes

[`AWS::EventSchemas::Discoverer`](../templatereference/aws-resource-eventschemas-discoverer.md)

Yes

Yes

No

[`AWS::EventSchemas::Registry`](../templatereference/aws-resource-eventschemas-registry.md)

Yes

Yes

No

[`AWS::EventSchemas::RegistryPolicy`](../templatereference/aws-resource-eventschemas-registrypolicy.md)

Yes

Yes

No

[`AWS::EventSchemas::Schema`](../templatereference/aws-resource-eventschemas-schema.md)

Yes

Yes

No

[`AWS::Events::ApiDestination`](../templatereference/aws-resource-events-apidestination.md)

Yes

Yes

Yes

[`AWS::Events::Archive`](../templatereference/aws-resource-events-archive.md)

Yes

Yes

Yes

[`AWS::Events::Connection`](../templatereference/aws-resource-events-connection.md)

Yes

Yes

No

[`AWS::Events::Endpoint`](../templatereference/aws-resource-events-endpoint.md)

Yes

Yes

Yes

[`AWS::Events::EventBus`](../templatereference/aws-resource-events-eventbus.md)

Yes

Yes

Yes

[`AWS::Events::EventBusPolicy`](../templatereference/aws-resource-events-eventbuspolicy.md)

Yes

No

No

[`AWS::Events::Rule`](../templatereference/aws-resource-events-rule.md)

Yes

Yes

Yes

[`AWS::Evidently::Experiment`](../templatereference/aws-resource-evidently-experiment.md)

Yes

Yes

No

[`AWS::Evidently::Feature`](../templatereference/aws-resource-evidently-feature.md)

Yes

No

No

[`AWS::Evidently::Launch`](../templatereference/aws-resource-evidently-launch.md)

Yes

Yes

No

[`AWS::Evidently::Project`](../templatereference/aws-resource-evidently-project.md)

Yes

Yes

No

[`AWS::Evidently::Segment`](../templatereference/aws-resource-evidently-segment.md)

Yes

Yes

No

[`AWS::FIS::ExperimentTemplate`](../templatereference/aws-resource-fis-experimenttemplate.md)

Yes

Yes

Yes

[`AWS::FIS::TargetAccountConfiguration`](../templatereference/aws-resource-fis-targetaccountconfiguration.md)

Yes

Yes

No

[`AWS::FMS::NotificationChannel`](../templatereference/aws-resource-fms-notificationchannel.md)

Yes

Yes

No

[`AWS::FMS::Policy`](../templatereference/aws-resource-fms-policy.md)

Yes

No

No

[`AWS::FMS::ResourceSet`](../templatereference/aws-resource-fms-resourceset.md)

Yes

Yes

No

[`AWS::FSx::DataRepositoryAssociation`](../templatereference/aws-resource-fsx-datarepositoryassociation.md)

Yes

Yes

Yes

[`AWS::FSx::S3AccessPointAttachment`](../templatereference/aws-resource-fsx-s3accesspointattachment.md)

Yes

Yes

No

[`AWS::FinSpace::Environment`](../templatereference/aws-resource-finspace-environment.md)

Yes

Yes

Yes

[`AWS::Forecast::Dataset`](../templatereference/aws-resource-forecast-dataset.md)

Yes

Yes

No

[`AWS::Forecast::DatasetGroup`](../templatereference/aws-resource-forecast-datasetgroup.md)

Yes

Yes

No

[`AWS::FraudDetector::Detector`](../templatereference/aws-resource-frauddetector-detector.md)

Yes

No

Yes

[`AWS::FraudDetector::EntityType`](../templatereference/aws-resource-frauddetector-entitytype.md)

Yes

Yes

Yes

[`AWS::FraudDetector::EventType`](../templatereference/aws-resource-frauddetector-eventtype.md)

Yes

Yes

Yes

[`AWS::FraudDetector::Label`](../templatereference/aws-resource-frauddetector-label.md)

Yes

Yes

Yes

[`AWS::FraudDetector::List`](../templatereference/aws-resource-frauddetector-list.md)

Yes

Yes

Yes

[`AWS::FraudDetector::Outcome`](../templatereference/aws-resource-frauddetector-outcome.md)

Yes

Yes

Yes

[`AWS::FraudDetector::Variable`](../templatereference/aws-resource-frauddetector-variable.md)

Yes

Yes

Yes

[`AWS::GameLift::Alias`](../templatereference/aws-resource-gamelift-alias.md)

Yes

Yes

Yes

[`AWS::GameLift::Build`](../templatereference/aws-resource-gamelift-build.md)

Yes

Yes

Yes

[`AWS::GameLift::ContainerFleet`](../templatereference/aws-resource-gamelift-containerfleet.md)

Yes

Yes

No

[`AWS::GameLift::ContainerGroupDefinition`](../templatereference/aws-resource-gamelift-containergroupdefinition.md)

Yes

Yes

No

[`AWS::GameLift::Fleet`](../templatereference/aws-resource-gamelift-fleet.md)

Yes

Yes

Yes

[`AWS::GameLift::GameServerGroup`](../templatereference/aws-resource-gamelift-gameservergroup.md)

Yes

Yes

No

[`AWS::GameLift::GameSessionQueue`](../templatereference/aws-resource-gamelift-gamesessionqueue.md)

Yes

Yes

No

[`AWS::GameLift::Location`](../templatereference/aws-resource-gamelift-location.md)

Yes

Yes

No

[`AWS::GameLift::MatchmakingConfiguration`](../templatereference/aws-resource-gamelift-matchmakingconfiguration.md)

Yes

Yes

No

[`AWS::GameLift::MatchmakingRuleSet`](../templatereference/aws-resource-gamelift-matchmakingruleset.md)

Yes

Yes

No

[`AWS::GameLift::Script`](../templatereference/aws-resource-gamelift-script.md)

Yes

Yes

No

[`AWS::GameLiftStreams::Application`](../templatereference/aws-resource-gameliftstreams-application.md)

Yes

Yes

No

[`AWS::GameLiftStreams::StreamGroup`](../templatereference/aws-resource-gameliftstreams-streamgroup.md)

Yes

Yes

No

[`AWS::GlobalAccelerator::Accelerator`](../templatereference/aws-resource-globalaccelerator-accelerator.md)

Yes

Yes

Yes

[`AWS::GlobalAccelerator::CrossAccountAttachment`](../templatereference/aws-resource-globalaccelerator-crossaccountattachment.md)

Yes

Yes

No

[`AWS::GlobalAccelerator::EndpointGroup`](../templatereference/aws-resource-globalaccelerator-endpointgroup.md)

Yes

Yes

Yes

[`AWS::GlobalAccelerator::Listener`](../templatereference/aws-resource-globalaccelerator-listener.md)

Yes

Yes

Yes

[`AWS::Glue::Catalog`](../templatereference/aws-resource-glue-catalog.md)

Yes

Yes

No

[`AWS::Glue::Crawler`](../templatereference/aws-resource-glue-crawler.md)

Yes

No

No

[`AWS::Glue::Database`](../templatereference/aws-resource-glue-database.md)

Yes

Yes

No

[`AWS::Glue::IdentityCenterConfiguration`](../templatereference/aws-resource-glue-identitycenterconfiguration.md)

Yes

Yes

No

[`AWS::Glue::Integration`](../templatereference/aws-resource-glue-integration.md)

Yes

Yes

No

[`AWS::Glue::IntegrationResourceProperty`](../templatereference/aws-resource-glue-integrationresourceproperty.md)

Yes

Yes

No

[`AWS::Glue::Job`](../templatereference/aws-resource-glue-job.md)

Yes

Yes

No

[`AWS::Glue::Registry`](../templatereference/aws-resource-glue-registry.md)

Yes

Yes

Yes

[`AWS::Glue::Schema`](../templatereference/aws-resource-glue-schema.md)

Yes

Yes

Yes

[`AWS::Glue::SchemaVersion`](../templatereference/aws-resource-glue-schemaversion.md)

Yes

Yes

Yes

[`AWS::Glue::SchemaVersionMetadata`](../templatereference/aws-resource-glue-schemaversionmetadata.md)

Yes

Yes

Yes

[`AWS::Glue::Trigger`](../templatereference/aws-resource-glue-trigger.md)

Yes

Yes

No

[`AWS::Glue::UsageProfile`](../templatereference/aws-resource-glue-usageprofile.md)

Yes

Yes

No

[`AWS::Grafana::Workspace`](../templatereference/aws-resource-grafana-workspace.md)

Yes

Yes

Yes

[`AWS::GreengrassV2::ComponentVersion`](../templatereference/aws-resource-greengrassv2-componentversion.md)

Yes

Yes

No

[`AWS::GreengrassV2::Deployment`](../templatereference/aws-resource-greengrassv2-deployment.md)

Yes

Yes

Yes

[`AWS::GroundStation::Config`](../templatereference/aws-resource-groundstation-config.md)

Yes

Yes

Yes

[`AWS::GroundStation::DataflowEndpointGroup`](../templatereference/aws-resource-groundstation-dataflowendpointgroup.md)

Yes

Yes

Yes

[`AWS::GroundStation::DataflowEndpointGroupV2`](../templatereference/aws-resource-groundstation-dataflowendpointgroupv2.md)

Yes

Yes

No

[`AWS::GroundStation::MissionProfile`](../templatereference/aws-resource-groundstation-missionprofile.md)

Yes

Yes

Yes

[`AWS::GuardDuty::Detector`](../templatereference/aws-resource-guardduty-detector.md)

Yes

Yes

Yes

[`AWS::GuardDuty::Filter`](../templatereference/aws-resource-guardduty-filter.md)

Yes

Yes

No

[`AWS::GuardDuty::IPSet`](../templatereference/aws-resource-guardduty-ipset.md)

Yes

Yes

No

[`AWS::GuardDuty::MalwareProtectionPlan`](../templatereference/aws-resource-guardduty-malwareprotectionplan.md)

Yes

Yes

No

[`AWS::GuardDuty::Master`](../templatereference/aws-resource-guardduty-master.md)

Yes

Yes

No

[`AWS::GuardDuty::Member`](../templatereference/aws-resource-guardduty-member.md)

Yes

Yes

No

[`AWS::GuardDuty::PublishingDestination`](../templatereference/aws-resource-guardduty-publishingdestination.md)

Yes

Yes

No

[`AWS::GuardDuty::ThreatEntitySet`](../templatereference/aws-resource-guardduty-threatentityset.md)

Yes

Yes

No

[`AWS::GuardDuty::ThreatIntelSet`](../templatereference/aws-resource-guardduty-threatintelset.md)

Yes

Yes

No

[`AWS::GuardDuty::TrustedEntitySet`](../templatereference/aws-resource-guardduty-trustedentityset.md)

Yes

Yes

No

[`AWS::HealthImaging::Datastore`](../templatereference/aws-resource-healthimaging-datastore.md)

Yes

Yes

Yes

[`AWS::HealthLake::FHIRDatastore`](../templatereference/aws-resource-healthlake-fhirdatastore.md)

Yes

Yes

Yes

[`AWS::IAM::Group`](../templatereference/aws-resource-iam-group.md)

Yes

Yes

Yes

[`AWS::IAM::GroupPolicy`](../templatereference/aws-resource-iam-grouppolicy.md)

Yes

Yes

No

[`AWS::IAM::InstanceProfile`](../templatereference/aws-resource-iam-instanceprofile.md)

Yes

Yes

Yes

[`AWS::IAM::ManagedPolicy`](../templatereference/aws-resource-iam-managedpolicy.md)

Yes

Yes

Yes

[`AWS::IAM::OIDCProvider`](../templatereference/aws-resource-iam-oidcprovider.md)

Yes

Yes

Yes

[`AWS::IAM::Role`](../templatereference/aws-resource-iam-role.md)

Yes

Yes

Yes

[`AWS::IAM::RolePolicy`](../templatereference/aws-resource-iam-rolepolicy.md)

Yes

Yes

No

[`AWS::IAM::SAMLProvider`](../templatereference/aws-resource-iam-samlprovider.md)

Yes

Yes

Yes

[`AWS::IAM::ServerCertificate`](../templatereference/aws-resource-iam-servercertificate.md)

Yes

Yes

Yes

[`AWS::IAM::ServiceLinkedRole`](../templatereference/aws-resource-iam-servicelinkedrole.md)

Yes

Yes

No

[`AWS::IAM::User`](../templatereference/aws-resource-iam-user.md)

Yes

Yes

Yes

[`AWS::IAM::UserPolicy`](../templatereference/aws-resource-iam-userpolicy.md)

Yes

Yes

No

[`AWS::IAM::VirtualMFADevice`](../templatereference/aws-resource-iam-virtualmfadevice.md)

Yes

Yes

No

[`AWS::IVS::Channel`](../templatereference/aws-resource-ivs-channel.md)

Yes

Yes

Yes

[`AWS::IVS::EncoderConfiguration`](../templatereference/aws-resource-ivs-encoderconfiguration.md)

Yes

Yes

No

[`AWS::IVS::IngestConfiguration`](../templatereference/aws-resource-ivs-ingestconfiguration.md)

Yes

Yes

No

[`AWS::IVS::PlaybackKeyPair`](../templatereference/aws-resource-ivs-playbackkeypair.md)

Yes

Yes

Yes

[`AWS::IVS::PlaybackRestrictionPolicy`](../templatereference/aws-resource-ivs-playbackrestrictionpolicy.md)

Yes

Yes

No

[`AWS::IVS::PublicKey`](../templatereference/aws-resource-ivs-publickey.md)

Yes

Yes

No

[`AWS::IVS::RecordingConfiguration`](../templatereference/aws-resource-ivs-recordingconfiguration.md)

Yes

Yes

Yes

[`AWS::IVS::Stage`](../templatereference/aws-resource-ivs-stage.md)

Yes

Yes

No

[`AWS::IVS::StorageConfiguration`](../templatereference/aws-resource-ivs-storageconfiguration.md)

Yes

Yes

No

[`AWS::IVS::StreamKey`](../templatereference/aws-resource-ivs-streamkey.md)

Yes

Yes

No

[`AWS::IVSChat::LoggingConfiguration`](../templatereference/aws-resource-ivschat-loggingconfiguration.md)

Yes

Yes

Yes

[`AWS::IVSChat::Room`](../templatereference/aws-resource-ivschat-room.md)

Yes

Yes

Yes

[`AWS::IdentityStore::Group`](../templatereference/aws-resource-identitystore-group.md)

Yes

Yes

Yes

[`AWS::IdentityStore::GroupMembership`](../templatereference/aws-resource-identitystore-groupmembership.md)

Yes

Yes

Yes

[`AWS::ImageBuilder::Component`](../templatereference/aws-resource-imagebuilder-component.md)

Yes

Yes

Yes

[`AWS::ImageBuilder::ContainerRecipe`](../templatereference/aws-resource-imagebuilder-containerrecipe.md)

Yes

No

No

[`AWS::ImageBuilder::DistributionConfiguration`](../templatereference/aws-resource-imagebuilder-distributionconfiguration.md)

Yes

Yes

Yes

[`AWS::ImageBuilder::Image`](../templatereference/aws-resource-imagebuilder-image.md)

Yes

Yes

Yes

[`AWS::ImageBuilder::ImagePipeline`](../templatereference/aws-resource-imagebuilder-imagepipeline.md)

Yes

Yes

Yes

[`AWS::ImageBuilder::ImageRecipe`](../templatereference/aws-resource-imagebuilder-imagerecipe.md)

Yes

Yes

Yes

[`AWS::ImageBuilder::InfrastructureConfiguration`](../templatereference/aws-resource-imagebuilder-infrastructureconfiguration.md)

Yes

Yes

Yes

[`AWS::ImageBuilder::LifecyclePolicy`](../templatereference/aws-resource-imagebuilder-lifecyclepolicy.md)

Yes

Yes

No

[`AWS::ImageBuilder::Workflow`](../templatereference/aws-resource-imagebuilder-workflow.md)

Yes

Yes

No

[`AWS::Inspector::AssessmentTarget`](../templatereference/aws-resource-inspector-assessmenttarget.md)

Yes

Yes

Yes

[`AWS::Inspector::AssessmentTemplate`](../templatereference/aws-resource-inspector-assessmenttemplate.md)

Yes

Yes

Yes

[`AWS::Inspector::ResourceGroup`](../templatereference/aws-resource-inspector-resourcegroup.md)

Yes

Yes

No

[`AWS::InspectorV2::CisScanConfiguration`](../templatereference/aws-resource-inspectorv2-cisscanconfiguration.md)

Yes

Yes

No

[`AWS::InspectorV2::CodeSecurityIntegration`](../templatereference/aws-resource-inspectorv2-codesecurityintegration.md)

Yes

Yes

No

[`AWS::InspectorV2::CodeSecurityScanConfiguration`](../templatereference/aws-resource-inspectorv2-codesecurityscanconfiguration.md)

Yes

Yes

No

[`AWS::InspectorV2::Filter`](../templatereference/aws-resource-inspectorv2-filter.md)

Yes

Yes

No

[`AWS::InternetMonitor::Monitor`](../templatereference/aws-resource-internetmonitor-monitor.md)

Yes

Yes

Yes

[`AWS::Invoicing::InvoiceUnit`](../templatereference/aws-resource-invoicing-invoiceunit.md)

Yes

Yes

No

[`AWS::IoT::AccountAuditConfiguration`](../templatereference/aws-resource-iot-accountauditconfiguration.md)

Yes

Yes

Yes

[`AWS::IoT::Authorizer`](../templatereference/aws-resource-iot-authorizer.md)

Yes

Yes

Yes

[`AWS::IoT::BillingGroup`](../templatereference/aws-resource-iot-billinggroup.md)

Yes

Yes

Yes

[`AWS::IoT::CACertificate`](../templatereference/aws-resource-iot-cacertificate.md)

Yes

Yes

Yes

[`AWS::IoT::Certificate`](../templatereference/aws-resource-iot-certificate.md)

Yes

Yes

Yes

[`AWS::IoT::CertificateProvider`](../templatereference/aws-resource-iot-certificateprovider.md)

Yes

Yes

No

[`AWS::IoT::Command`](../templatereference/aws-resource-iot-command.md)

Yes

Yes

No

[`AWS::IoT::CustomMetric`](../templatereference/aws-resource-iot-custommetric.md)

Yes

Yes

Yes

[`AWS::IoT::Dimension`](../templatereference/aws-resource-iot-dimension.md)

Yes

Yes

Yes

[`AWS::IoT::DomainConfiguration`](../templatereference/aws-resource-iot-domainconfiguration.md)

Yes

Yes

No

[`AWS::IoT::EncryptionConfiguration`](../templatereference/aws-resource-iot-encryptionconfiguration.md)

Yes

Yes

No

[`AWS::IoT::FleetMetric`](../templatereference/aws-resource-iot-fleetmetric.md)

Yes

Yes

No

[`AWS::IoT::JobTemplate`](../templatereference/aws-resource-iot-jobtemplate.md)

Yes

Yes

No

[`AWS::IoT::Logging`](../templatereference/aws-resource-iot-logging.md)

Yes

Yes

Yes

[`AWS::IoT::MitigationAction`](../templatereference/aws-resource-iot-mitigationaction.md)

Yes

Yes

Yes

[`AWS::IoT::Policy`](../templatereference/aws-resource-iot-policy.md)

Yes

No

Yes

[`AWS::IoT::ProvisioningTemplate`](../templatereference/aws-resource-iot-provisioningtemplate.md)

Yes

Yes

Yes

[`AWS::IoT::ResourceSpecificLogging`](../templatereference/aws-resource-iot-resourcespecificlogging.md)

Yes

Yes

Yes

[`AWS::IoT::RoleAlias`](../templatereference/aws-resource-iot-rolealias.md)

Yes

Yes

Yes

[`AWS::IoT::ScheduledAudit`](../templatereference/aws-resource-iot-scheduledaudit.md)

Yes

Yes

Yes

[`AWS::IoT::SecurityProfile`](../templatereference/aws-resource-iot-securityprofile.md)

Yes

No

Yes

[`AWS::IoT::SoftwarePackage`](../templatereference/aws-resource-iot-softwarepackage.md)

Yes

Yes

Yes

[`AWS::IoT::SoftwarePackageVersion`](../templatereference/aws-resource-iot-softwarepackageversion.md)

Yes

Yes

No

[`AWS::IoT::Thing`](../templatereference/aws-resource-iot-thing.md)

Yes

Yes

Yes

[`AWS::IoT::ThingGroup`](../templatereference/aws-resource-iot-thinggroup.md)

Yes

Yes

Yes

[`AWS::IoT::ThingType`](../templatereference/aws-resource-iot-thingtype.md)

Yes

Yes

No

[`AWS::IoT::TopicRule`](../templatereference/aws-resource-iot-topicrule.md)

Yes

Yes

Yes

[`AWS::IoT::TopicRuleDestination`](../templatereference/aws-resource-iot-topicruledestination.md)

Yes

Yes

Yes

[`AWS::IoTAnalytics::Channel`](../templatereference/aws-resource-iotanalytics-channel.md)

Yes

Yes

Yes

[`AWS::IoTAnalytics::Dataset`](../templatereference/aws-resource-iotanalytics-dataset.md)

Yes

Yes

Yes

[`AWS::IoTAnalytics::Datastore`](../templatereference/aws-resource-iotanalytics-datastore.md)

Yes

Yes

Yes

[`AWS::IoTAnalytics::Pipeline`](../templatereference/aws-resource-iotanalytics-pipeline.md)

Yes

Yes

Yes

[`AWS::IoTCoreDeviceAdvisor::SuiteDefinition`](../templatereference/aws-resource-iotcoredeviceadvisor-suitedefinition.md)

Yes

Yes

Yes

[`AWS::IoTEvents::AlarmModel`](../templatereference/aws-resource-iotevents-alarmmodel.md)

Yes

Yes

Yes

[`AWS::IoTEvents::DetectorModel`](../templatereference/aws-resource-iotevents-detectormodel.md)

Yes

Yes

Yes

[`AWS::IoTEvents::Input`](../templatereference/aws-resource-iotevents-input.md)

Yes

Yes

Yes

[`AWS::IoTFleetWise::Campaign`](../templatereference/aws-resource-iotfleetwise-campaign.md)

Yes

Yes

Yes

[`AWS::IoTFleetWise::DecoderManifest`](../templatereference/aws-resource-iotfleetwise-decodermanifest.md)

Yes

Yes

Yes

[`AWS::IoTFleetWise::Fleet`](../templatereference/aws-resource-iotfleetwise-fleet.md)

Yes

Yes

Yes

[`AWS::IoTFleetWise::ModelManifest`](../templatereference/aws-resource-iotfleetwise-modelmanifest.md)

Yes

No

Yes

[`AWS::IoTFleetWise::SignalCatalog`](../templatereference/aws-resource-iotfleetwise-signalcatalog.md)

Yes

No

Yes

[`AWS::IoTFleetWise::StateTemplate`](../templatereference/aws-resource-iotfleetwise-statetemplate.md)

Yes

Yes

No

[`AWS::IoTFleetWise::Vehicle`](../templatereference/aws-resource-iotfleetwise-vehicle.md)

Yes

Yes

Yes

[`AWS::IoTSiteWise::AccessPolicy`](../templatereference/aws-resource-iotsitewise-accesspolicy.md)

Yes

Yes

No

[`AWS::IoTSiteWise::Asset`](../templatereference/aws-resource-iotsitewise-asset.md)

Yes

Yes

Yes

[`AWS::IoTSiteWise::AssetModel`](../templatereference/aws-resource-iotsitewise-assetmodel.md)

Yes

Yes

Yes

[`AWS::IoTSiteWise::ComputationModel`](../templatereference/aws-resource-iotsitewise-computationmodel.md)

Yes

Yes

No

[`AWS::IoTSiteWise::Dashboard`](../templatereference/aws-resource-iotsitewise-dashboard.md)

Yes

Yes

No

[`AWS::IoTSiteWise::Dataset`](../templatereference/aws-resource-iotsitewise-dataset.md)

Yes

Yes

No

[`AWS::IoTSiteWise::Gateway`](../templatereference/aws-resource-iotsitewise-gateway.md)

Yes

Yes

Yes

[`AWS::IoTSiteWise::Portal`](../templatereference/aws-resource-iotsitewise-portal.md)

Yes

Yes

No

[`AWS::IoTSiteWise::Project`](../templatereference/aws-resource-iotsitewise-project.md)

Yes

Yes

No

[`AWS::IoTTwinMaker::ComponentType`](../templatereference/aws-resource-iottwinmaker-componenttype.md)

Yes

Yes

Yes

[`AWS::IoTTwinMaker::Entity`](../templatereference/aws-resource-iottwinmaker-entity.md)

Yes

Yes

Yes

[`AWS::IoTTwinMaker::Scene`](../templatereference/aws-resource-iottwinmaker-scene.md)

Yes

Yes

Yes

[`AWS::IoTTwinMaker::SyncJob`](../templatereference/aws-resource-iottwinmaker-syncjob.md)

Yes

Yes

Yes

[`AWS::IoTTwinMaker::Workspace`](../templatereference/aws-resource-iottwinmaker-workspace.md)

Yes

Yes

Yes

[`AWS::IoTWireless::Destination`](../templatereference/aws-resource-iotwireless-destination.md)

Yes

Yes

Yes

[`AWS::IoTWireless::DeviceProfile`](../templatereference/aws-resource-iotwireless-deviceprofile.md)

Yes

Yes

Yes

[`AWS::IoTWireless::FuotaTask`](../templatereference/aws-resource-iotwireless-fuotatask.md)

Yes

Yes

Yes

[`AWS::IoTWireless::MulticastGroup`](../templatereference/aws-resource-iotwireless-multicastgroup.md)

Yes

Yes

Yes

[`AWS::IoTWireless::NetworkAnalyzerConfiguration`](../templatereference/aws-resource-iotwireless-networkanalyzerconfiguration.md)

Yes

Yes

Yes

[`AWS::IoTWireless::PartnerAccount`](../templatereference/aws-resource-iotwireless-partneraccount.md)

Yes

Yes

Yes

[`AWS::IoTWireless::ServiceProfile`](../templatereference/aws-resource-iotwireless-serviceprofile.md)

Yes

Yes

Yes

[`AWS::IoTWireless::TaskDefinition`](../templatereference/aws-resource-iotwireless-taskdefinition.md)

Yes

Yes

Yes

[`AWS::IoTWireless::WirelessDevice`](../templatereference/aws-resource-iotwireless-wirelessdevice.md)

Yes

Yes

Yes

[`AWS::IoTWireless::WirelessDeviceImportTask`](../templatereference/aws-resource-iotwireless-wirelessdeviceimporttask.md)

Yes

Yes

Yes

[`AWS::IoTWireless::WirelessGateway`](../templatereference/aws-resource-iotwireless-wirelessgateway.md)

Yes

Yes

Yes

[`AWS::KMS::Alias`](../templatereference/aws-resource-kms-alias.md)

Yes

No

Yes

[`AWS::KMS::Key`](../templatereference/aws-resource-kms-key.md)

Yes

Yes

Yes

[`AWS::KMS::ReplicaKey`](../templatereference/aws-resource-kms-replicakey.md)

Yes

No

Yes

[`AWS::KafkaConnect::Connector`](../templatereference/aws-resource-kafkaconnect-connector.md)

Yes

No

No

[`AWS::KafkaConnect::CustomPlugin`](../templatereference/aws-resource-kafkaconnect-customplugin.md)

Yes

Yes

No

[`AWS::KafkaConnect::WorkerConfiguration`](../templatereference/aws-resource-kafkaconnect-workerconfiguration.md)

Yes

Yes

No

[`AWS::Kendra::DataSource`](../templatereference/aws-resource-kendra-datasource.md)

Yes

Yes

Yes

[`AWS::Kendra::Faq`](../templatereference/aws-resource-kendra-faq.md)

Yes

Yes

Yes

[`AWS::Kendra::Index`](../templatereference/aws-resource-kendra-index.md)

Yes

Yes

Yes

[`AWS::KendraRanking::ExecutionPlan`](../templatereference/aws-resource-kendraranking-executionplan.md)

Yes

Yes

Yes

[`AWS::Kinesis::ResourcePolicy`](../templatereference/aws-resource-kinesis-resourcepolicy.md)

Yes

Yes

No

[`AWS::Kinesis::Stream`](../templatereference/aws-resource-kinesis-stream.md)

Yes

Yes

Yes

[`AWS::Kinesis::StreamConsumer`](../templatereference/aws-resource-kinesis-streamconsumer.md)

Yes

Yes

No

[`AWS::KinesisAnalyticsV2::Application`](../templatereference/aws-resource-kinesisanalyticsv2-application.md)

Yes

Yes

No

[`AWS::KinesisFirehose::DeliveryStream`](../templatereference/aws-resource-kinesisfirehose-deliverystream.md)

Yes

No

No

[`AWS::KinesisVideo::SignalingChannel`](../templatereference/aws-resource-kinesisvideo-signalingchannel.md)

Yes

Yes

No

[`AWS::KinesisVideo::Stream`](../templatereference/aws-resource-kinesisvideo-stream.md)

Yes

Yes

No

[`AWS::LakeFormation::DataCellsFilter`](../templatereference/aws-resource-lakeformation-datacellsfilter.md)

Yes

Yes

No

[`AWS::LakeFormation::PrincipalPermissions`](../templatereference/aws-resource-lakeformation-principalpermissions.md)

Yes

Yes

No

[`AWS::LakeFormation::Tag`](../templatereference/aws-resource-lakeformation-tag.md)

Yes

Yes

No

[`AWS::LakeFormation::TagAssociation`](../templatereference/aws-resource-lakeformation-tagassociation.md)

Yes

Yes

No

[`AWS::Lambda::Alias`](../templatereference/aws-resource-lambda-alias.md)

Yes

Yes

No

[`AWS::Lambda::CapacityProvider`](../templatereference/aws-resource-lambda-capacityprovider.md)

Yes

Yes

No

[`AWS::Lambda::CodeSigningConfig`](../templatereference/aws-resource-lambda-codesigningconfig.md)

Yes

Yes

Yes

[`AWS::Lambda::EventInvokeConfig`](../templatereference/aws-resource-lambda-eventinvokeconfig.md)

Yes

Yes

No

[`AWS::Lambda::EventSourceMapping`](../templatereference/aws-resource-lambda-eventsourcemapping.md)

Yes

No

Yes

[`AWS::Lambda::Function`](../templatereference/aws-resource-lambda-function.md)

Yes

Yes

Yes

[`AWS::Lambda::LayerVersion`](../templatereference/aws-resource-lambda-layerversion.md)

Yes

Yes

No

[`AWS::Lambda::LayerVersionPermission`](../templatereference/aws-resource-lambda-layerversionpermission.md)

Yes

Yes

Yes

[`AWS::Lambda::Permission`](../templatereference/aws-resource-lambda-permission.md)

Yes

Yes

Yes

[`AWS::Lambda::Url`](../templatereference/aws-resource-lambda-url.md)

Yes

Yes

Yes

[`AWS::Lambda::Version`](../templatereference/aws-resource-lambda-version.md)

Yes

Yes

Yes

[`AWS::LaunchWizard::Deployment`](../templatereference/aws-resource-launchwizard-deployment.md)

Yes

Yes

No

[`AWS::Lex::Bot`](../templatereference/aws-resource-lex-bot.md)

Yes

Yes

No

[`AWS::Lex::BotAlias`](../templatereference/aws-resource-lex-botalias.md)

Yes

Yes

No

[`AWS::Lex::BotVersion`](../templatereference/aws-resource-lex-botversion.md)

Yes

Yes

No

[`AWS::Lex::ResourcePolicy`](../templatereference/aws-resource-lex-resourcepolicy.md)

Yes

Yes

No

[`AWS::LicenseManager::Grant`](../templatereference/aws-resource-licensemanager-grant.md)

Yes

Yes

No

[`AWS::LicenseManager::License`](../templatereference/aws-resource-licensemanager-license.md)

Yes

Yes

No

[`AWS::Lightsail::Alarm`](../templatereference/aws-resource-lightsail-alarm.md)

Yes

Yes

Yes

[`AWS::Lightsail::Bucket`](../templatereference/aws-resource-lightsail-bucket.md)

Yes

Yes

Yes

[`AWS::Lightsail::Certificate`](../templatereference/aws-resource-lightsail-certificate.md)

Yes

Yes

Yes

[`AWS::Lightsail::Container`](../templatereference/aws-resource-lightsail-container.md)

Yes

Yes

Yes

[`AWS::Lightsail::Database`](../templatereference/aws-resource-lightsail-database.md)

Yes

Yes

No

[`AWS::Lightsail::DatabaseSnapshot`](../templatereference/aws-resource-lightsail-databasesnapshot.md)

Yes

Yes

No

[`AWS::Lightsail::Disk`](../templatereference/aws-resource-lightsail-disk.md)

Yes

Yes

Yes

[`AWS::Lightsail::DiskSnapshot`](../templatereference/aws-resource-lightsail-disksnapshot.md)

Yes

Yes

No

[`AWS::Lightsail::Distribution`](../templatereference/aws-resource-lightsail-distribution.md)

Yes

Yes

Yes

[`AWS::Lightsail::Domain`](../templatereference/aws-resource-lightsail-domain.md)

Yes

Yes

No

[`AWS::Lightsail::Instance`](../templatereference/aws-resource-lightsail-instance.md)

Yes

Yes

Yes

[`AWS::Lightsail::InstanceSnapshot`](../templatereference/aws-resource-lightsail-instancesnapshot.md)

Yes

Yes

No

[`AWS::Lightsail::LoadBalancer`](../templatereference/aws-resource-lightsail-loadbalancer.md)

Yes

Yes

Yes

[`AWS::Lightsail::LoadBalancerTlsCertificate`](../templatereference/aws-resource-lightsail-loadbalancertlscertificate.md)

Yes

Yes

Yes

[`AWS::Lightsail::StaticIp`](../templatereference/aws-resource-lightsail-staticip.md)

Yes

Yes

Yes

[`AWS::Location::APIKey`](../templatereference/aws-resource-location-apikey.md)

Yes

Yes

No

[`AWS::Location::GeofenceCollection`](../templatereference/aws-resource-location-geofencecollection.md)

Yes

Yes

Yes

[`AWS::Location::Map`](../templatereference/aws-resource-location-map.md)

Yes

Yes

Yes

[`AWS::Location::PlaceIndex`](../templatereference/aws-resource-location-placeindex.md)

Yes

Yes

Yes

[`AWS::Location::RouteCalculator`](../templatereference/aws-resource-location-routecalculator.md)

Yes

Yes

Yes

[`AWS::Location::Tracker`](../templatereference/aws-resource-location-tracker.md)

Yes

Yes

Yes

[`AWS::Location::TrackerConsumer`](../templatereference/aws-resource-location-trackerconsumer.md)

Yes

Yes

Yes

[`AWS::Logs::AccountPolicy`](../templatereference/aws-resource-logs-accountpolicy.md)

Yes

Yes

Yes

[`AWS::Logs::Delivery`](../templatereference/aws-resource-logs-delivery.md)

Yes

Yes

No

[`AWS::Logs::DeliveryDestination`](../templatereference/aws-resource-logs-deliverydestination.md)

Yes

Yes

No

[`AWS::Logs::DeliverySource`](../templatereference/aws-resource-logs-deliverysource.md)

Yes

Yes

No

[`AWS::Logs::Destination`](../templatereference/aws-resource-logs-destination.md)

Yes

Yes

Yes

[`AWS::Logs::Integration`](../templatereference/aws-resource-logs-integration.md)

Yes

Yes

No

[`AWS::Logs::LogAnomalyDetector`](../templatereference/aws-resource-logs-loganomalydetector.md)

Yes

Yes

No

[`AWS::Logs::LogGroup`](../templatereference/aws-resource-logs-loggroup.md)

Yes

Yes

Yes

[`AWS::Logs::LogStream`](../templatereference/aws-resource-logs-logstream.md)

Yes

Yes

Yes

[`AWS::Logs::MetricFilter`](../templatereference/aws-resource-logs-metricfilter.md)

Yes

Yes

Yes

[`AWS::Logs::QueryDefinition`](../templatereference/aws-resource-logs-querydefinition.md)

Yes

Yes

No

[`AWS::Logs::ResourcePolicy`](../templatereference/aws-resource-logs-resourcepolicy.md)

Yes

Yes

Yes

[`AWS::Logs::ScheduledQuery`](../templatereference/aws-resource-logs-scheduledquery.md)

Yes

Yes

No

[`AWS::Logs::SubscriptionFilter`](../templatereference/aws-resource-logs-subscriptionfilter.md)

Yes

Yes

Yes

[`AWS::Logs::Transformer`](../templatereference/aws-resource-logs-transformer.md)

Yes

Yes

No

[`AWS::LookoutEquipment::InferenceScheduler`](../templatereference/aws-resource-lookoutequipment-inferencescheduler.md)

Yes

Yes

No

[`AWS::LookoutVision::Project`](../templatereference/aws-resource-lookoutvision-project.md)

Yes

Yes

Yes

[`AWS::M2::Application`](../templatereference/aws-resource-m2-application.md)

Yes

Yes

No

[`AWS::M2::Deployment`](../templatereference/aws-resource-m2-deployment.md)

Yes

Yes

No

[`AWS::M2::Environment`](../templatereference/aws-resource-m2-environment.md)

Yes

Yes

Yes

[`AWS::MPA::ApprovalTeam`](../templatereference/aws-resource-mpa-approvalteam.md)

Yes

Yes

No

[`AWS::MPA::IdentitySource`](../templatereference/aws-resource-mpa-identitysource.md)

Yes

Yes

No

[`AWS::MSK::BatchScramSecret`](../templatereference/aws-resource-msk-batchscramsecret.md)

Yes

Yes

Yes

[`AWS::MSK::Cluster`](../templatereference/aws-resource-msk-cluster.md)

Yes

Yes

Yes

[`AWS::MSK::ClusterPolicy`](../templatereference/aws-resource-msk-clusterpolicy.md)

Yes

Yes

Yes

[`AWS::MSK::Configuration`](../templatereference/aws-resource-msk-configuration.md)

Yes

Yes

Yes

[`AWS::MSK::Replicator`](../templatereference/aws-resource-msk-replicator.md)

Yes

Yes

Yes

[`AWS::MSK::ServerlessCluster`](../templatereference/aws-resource-msk-serverlesscluster.md)

Yes

Yes

Yes

[`AWS::MSK::Topic`](../templatereference/aws-resource-msk-topic.md)

Yes

Yes

No

[`AWS::MSK::VpcConnection`](../templatereference/aws-resource-msk-vpcconnection.md)

Yes

Yes

Yes

[`AWS::MWAA::Environment`](../templatereference/aws-resource-mwaa-environment.md)

Yes

Yes

No

[`AWS::MWAAServerless::Workflow`](../templatereference/aws-resource-mwaaserverless-workflow.md)

Yes

Yes

No

[`AWS::Macie::AllowList`](../templatereference/aws-resource-macie-allowlist.md)

Yes

Yes

No

[`AWS::Macie::CustomDataIdentifier`](../templatereference/aws-resource-macie-customdataidentifier.md)

Yes

Yes

No

[`AWS::Macie::FindingsFilter`](../templatereference/aws-resource-macie-findingsfilter.md)

Yes

Yes

No

[`AWS::Macie::Session`](../templatereference/aws-resource-macie-session.md)

Yes

Yes

Yes

[`AWS::ManagedBlockchain::Accessor`](../templatereference/aws-resource-managedblockchain-accessor.md)

Yes

Yes

Yes

[`AWS::MediaConnect::Bridge`](../templatereference/aws-resource-mediaconnect-bridge.md)

Yes

Yes

Yes

[`AWS::MediaConnect::BridgeOutput`](../templatereference/aws-resource-mediaconnect-bridgeoutput.md)

Yes

Yes

No

[`AWS::MediaConnect::BridgeSource`](../templatereference/aws-resource-mediaconnect-bridgesource.md)

Yes

Yes

No

[`AWS::MediaConnect::Flow`](../templatereference/aws-resource-mediaconnect-flow.md)

Yes

Yes

No

[`AWS::MediaConnect::FlowEntitlement`](../templatereference/aws-resource-mediaconnect-flowentitlement.md)

Yes

Yes

No

[`AWS::MediaConnect::FlowOutput`](../templatereference/aws-resource-mediaconnect-flowoutput.md)

Yes

Yes

No

[`AWS::MediaConnect::FlowSource`](../templatereference/aws-resource-mediaconnect-flowsource.md)

Yes

Yes

No

[`AWS::MediaConnect::FlowVpcInterface`](../templatereference/aws-resource-mediaconnect-flowvpcinterface.md)

Yes

Yes

Yes

[`AWS::MediaConnect::Gateway`](../templatereference/aws-resource-mediaconnect-gateway.md)

Yes

Yes

Yes

[`AWS::MediaConnect::RouterInput`](../templatereference/aws-resource-mediaconnect-routerinput.md)

Yes

Yes

No

[`AWS::MediaConnect::RouterNetworkInterface`](../templatereference/aws-resource-mediaconnect-routernetworkinterface.md)

Yes

Yes

No

[`AWS::MediaConnect::RouterOutput`](../templatereference/aws-resource-mediaconnect-routeroutput.md)

Yes

Yes

No

[`AWS::MediaLive::ChannelPlacementGroup`](../templatereference/aws-resource-medialive-channelplacementgroup.md)

Yes

Yes

No

[`AWS::MediaLive::CloudWatchAlarmTemplate`](../templatereference/aws-resource-medialive-cloudwatchalarmtemplate.md)

Yes

Yes

No

[`AWS::MediaLive::CloudWatchAlarmTemplateGroup`](../templatereference/aws-resource-medialive-cloudwatchalarmtemplategroup.md)

Yes

Yes

No

[`AWS::MediaLive::Cluster`](../templatereference/aws-resource-medialive-cluster.md)

Yes

Yes

No

[`AWS::MediaLive::EventBridgeRuleTemplate`](../templatereference/aws-resource-medialive-eventbridgeruletemplate.md)

Yes

Yes

No

[`AWS::MediaLive::EventBridgeRuleTemplateGroup`](../templatereference/aws-resource-medialive-eventbridgeruletemplategroup.md)

Yes

Yes

No

[`AWS::MediaLive::Multiplex`](../templatereference/aws-resource-medialive-multiplex.md)

Yes

Yes

No

[`AWS::MediaLive::Multiplexprogram`](../templatereference/aws-resource-medialive-multiplexprogram.md)

Yes

Yes

No

[`AWS::MediaLive::Network`](../templatereference/aws-resource-medialive-network.md)

Yes

Yes

No

[`AWS::MediaLive::SdiSource`](../templatereference/aws-resource-medialive-sdisource.md)

Yes

Yes

No

[`AWS::MediaLive::SignalMap`](../templatereference/aws-resource-medialive-signalmap.md)

Yes

Yes

No

[`AWS::MediaPackage::Asset`](../templatereference/aws-resource-mediapackage-asset.md)

Yes

Yes

Yes

[`AWS::MediaPackage::Channel`](../templatereference/aws-resource-mediapackage-channel.md)

Yes

Yes

Yes

[`AWS::MediaPackage::OriginEndpoint`](../templatereference/aws-resource-mediapackage-originendpoint.md)

Yes

Yes

Yes

[`AWS::MediaPackage::PackagingConfiguration`](../templatereference/aws-resource-mediapackage-packagingconfiguration.md)

Yes

Yes

Yes

[`AWS::MediaPackage::PackagingGroup`](../templatereference/aws-resource-mediapackage-packaginggroup.md)

Yes

Yes

Yes

[`AWS::MediaPackageV2::Channel`](../templatereference/aws-resource-mediapackagev2-channel.md)

Yes

Yes

Yes

[`AWS::MediaPackageV2::ChannelGroup`](../templatereference/aws-resource-mediapackagev2-channelgroup.md)

Yes

Yes

Yes

[`AWS::MediaPackageV2::ChannelPolicy`](../templatereference/aws-resource-mediapackagev2-channelpolicy.md)

Yes

Yes

No

[`AWS::MediaPackageV2::OriginEndpoint`](../templatereference/aws-resource-mediapackagev2-originendpoint.md)

Yes

Yes

Yes

[`AWS::MediaPackageV2::OriginEndpointPolicy`](../templatereference/aws-resource-mediapackagev2-originendpointpolicy.md)

Yes

Yes

No

[`AWS::MediaTailor::Channel`](../templatereference/aws-resource-mediatailor-channel.md)

Yes

Yes

No

[`AWS::MediaTailor::ChannelPolicy`](../templatereference/aws-resource-mediatailor-channelpolicy.md)

Yes

Yes

No

[`AWS::MediaTailor::LiveSource`](../templatereference/aws-resource-mediatailor-livesource.md)

Yes

Yes

Yes

[`AWS::MediaTailor::PlaybackConfiguration`](../templatereference/aws-resource-mediatailor-playbackconfiguration.md)

Yes

Yes

No

[`AWS::MediaTailor::SourceLocation`](../templatereference/aws-resource-mediatailor-sourcelocation.md)

Yes

Yes

Yes

[`AWS::MediaTailor::VodSource`](../templatereference/aws-resource-mediatailor-vodsource.md)

Yes

Yes

Yes

[`AWS::MemoryDB::ACL`](../templatereference/aws-resource-memorydb-acl.md)

Yes

Yes

Yes

[`AWS::MemoryDB::Cluster`](../templatereference/aws-resource-memorydb-cluster.md)

Yes

Yes

Yes

[`AWS::MemoryDB::MultiRegionCluster`](../templatereference/aws-resource-memorydb-multiregioncluster.md)

Yes

Yes

No

[`AWS::MemoryDB::ParameterGroup`](../templatereference/aws-resource-memorydb-parametergroup.md)

Yes

Yes

Yes

[`AWS::MemoryDB::SubnetGroup`](../templatereference/aws-resource-memorydb-subnetgroup.md)

Yes

Yes

Yes

[`AWS::MemoryDB::User`](../templatereference/aws-resource-memorydb-user.md)

Yes

Yes

Yes

[`AWS::Neptune::DBCluster`](../templatereference/aws-resource-neptune-dbcluster.md)

Yes

Yes

Yes

[`AWS::Neptune::DBClusterParameterGroup`](../templatereference/aws-resource-neptune-dbclusterparametergroup.md)

Yes

Yes

No

[`AWS::Neptune::DBInstance`](../templatereference/aws-resource-neptune-dbinstance.md)

Yes

Yes

No

[`AWS::Neptune::DBParameterGroup`](../templatereference/aws-resource-neptune-dbparametergroup.md)

Yes

Yes

No

[`AWS::Neptune::DBSubnetGroup`](../templatereference/aws-resource-neptune-dbsubnetgroup.md)

Yes

Yes

No

[`AWS::Neptune::EventSubscription`](../templatereference/aws-resource-neptune-eventsubscription.md)

Yes

Yes

No

[`AWS::NeptuneGraph::Graph`](../templatereference/aws-resource-neptunegraph-graph.md)

Yes

Yes

No

[`AWS::NeptuneGraph::PrivateGraphEndpoint`](../templatereference/aws-resource-neptunegraph-privategraphendpoint.md)

Yes

Yes

No

[`AWS::NetworkFirewall::Firewall`](../templatereference/aws-resource-networkfirewall-firewall.md)

Yes

Yes

Yes

[`AWS::NetworkFirewall::FirewallPolicy`](../templatereference/aws-resource-networkfirewall-firewallpolicy.md)

Yes

Yes

Yes

[`AWS::NetworkFirewall::LoggingConfiguration`](../templatereference/aws-resource-networkfirewall-loggingconfiguration.md)

Yes

Yes

No

[`AWS::NetworkFirewall::RuleGroup`](../templatereference/aws-resource-networkfirewall-rulegroup.md)

Yes

Yes

Yes

[`AWS::NetworkFirewall::TLSInspectionConfiguration`](../templatereference/aws-resource-networkfirewall-tlsinspectionconfiguration.md)

Yes

Yes

No

[`AWS::NetworkFirewall::VpcEndpointAssociation`](../templatereference/aws-resource-networkfirewall-vpcendpointassociation.md)

Yes

Yes

No

[`AWS::NetworkManager::ConnectAttachment`](../templatereference/aws-resource-networkmanager-connectattachment.md)

Yes

Yes

Yes

[`AWS::NetworkManager::ConnectPeer`](../templatereference/aws-resource-networkmanager-connectpeer.md)

Yes

Yes

Yes

[`AWS::NetworkManager::CoreNetwork`](../templatereference/aws-resource-networkmanager-corenetwork.md)

Yes

Yes

Yes

[`AWS::NetworkManager::CoreNetworkPrefixListAssociation`](../templatereference/aws-resource-networkmanager-corenetworkprefixlistassociation.md)

Yes

Yes

No

[`AWS::NetworkManager::CustomerGatewayAssociation`](../templatereference/aws-resource-networkmanager-customergatewayassociation.md)

Yes

Yes

Yes

[`AWS::NetworkManager::Device`](../templatereference/aws-resource-networkmanager-device.md)

Yes

Yes

Yes

[`AWS::NetworkManager::DirectConnectGatewayAttachment`](../templatereference/aws-resource-networkmanager-directconnectgatewayattachment.md)

Yes

Yes

No

[`AWS::NetworkManager::GlobalNetwork`](../templatereference/aws-resource-networkmanager-globalnetwork.md)

Yes

Yes

Yes

[`AWS::NetworkManager::Link`](../templatereference/aws-resource-networkmanager-link.md)

Yes

Yes

Yes

[`AWS::NetworkManager::LinkAssociation`](../templatereference/aws-resource-networkmanager-linkassociation.md)

Yes

Yes

Yes

[`AWS::NetworkManager::Site`](../templatereference/aws-resource-networkmanager-site.md)

Yes

Yes

Yes

[`AWS::NetworkManager::SiteToSiteVpnAttachment`](../templatereference/aws-resource-networkmanager-sitetositevpnattachment.md)

Yes

Yes

Yes

[`AWS::NetworkManager::TransitGatewayPeering`](../templatereference/aws-resource-networkmanager-transitgatewaypeering.md)

Yes

Yes

Yes

[`AWS::NetworkManager::TransitGatewayRegistration`](../templatereference/aws-resource-networkmanager-transitgatewayregistration.md)

Yes

Yes

Yes

[`AWS::NetworkManager::TransitGatewayRouteTableAttachment`](../templatereference/aws-resource-networkmanager-transitgatewayroutetableattachment.md)

Yes

Yes

Yes

[`AWS::NetworkManager::VpcAttachment`](../templatereference/aws-resource-networkmanager-vpcattachment.md)

Yes

Yes

Yes

[`AWS::Notifications::ChannelAssociation`](../templatereference/aws-resource-notifications-channelassociation.md)

Yes

Yes

No

[`AWS::Notifications::EventRule`](../templatereference/aws-resource-notifications-eventrule.md)

Yes

Yes

No

[`AWS::Notifications::ManagedNotificationAccountContactAssociation`](../templatereference/aws-resource-notifications-managednotificationaccountcontactassociation.md)

Yes

Yes

No

[`AWS::Notifications::ManagedNotificationAdditionalChannelAssociation`](../templatereference/aws-resource-notifications-managednotificationadditionalchannelassociation.md)

Yes

Yes

No

[`AWS::Notifications::NotificationConfiguration`](../templatereference/aws-resource-notifications-notificationconfiguration.md)

Yes

Yes

No

[`AWS::Notifications::NotificationHub`](../templatereference/aws-resource-notifications-notificationhub.md)

Yes

Yes

No

[`AWS::Notifications::OrganizationalUnitAssociation`](../templatereference/aws-resource-notifications-organizationalunitassociation.md)

Yes

Yes

No

[`AWS::NotificationsContacts::EmailContact`](../templatereference/aws-resource-notificationscontacts-emailcontact.md)

Yes

Yes

No

[`AWS::ODB::CloudAutonomousVmCluster`](../templatereference/aws-resource-odb-cloudautonomousvmcluster.md)

Yes

Yes

No

[`AWS::ODB::CloudExadataInfrastructure`](../templatereference/aws-resource-odb-cloudexadatainfrastructure.md)

Yes

Yes

No

[`AWS::ODB::CloudVmCluster`](../templatereference/aws-resource-odb-cloudvmcluster.md)

Yes

Yes

No

[`AWS::ODB::OdbNetwork`](../templatereference/aws-resource-odb-odbnetwork.md)

Yes

Yes

No

[`AWS::ODB::OdbPeeringConnection`](../templatereference/aws-resource-odb-odbpeeringconnection.md)

Yes

Yes

No

[`AWS::OSIS::Pipeline`](../templatereference/aws-resource-osis-pipeline.md)

Yes

Yes

Yes

[`AWS::Oam::Link`](../templatereference/aws-resource-oam-link.md)

Yes

Yes

Yes

[`AWS::Oam::Sink`](../templatereference/aws-resource-oam-sink.md)

Yes

Yes

Yes

[`AWS::ObservabilityAdmin::OrganizationCentralizationRule`](../templatereference/aws-resource-observabilityadmin-organizationcentralizationrule.md)

Yes

Yes

No

[`AWS::ObservabilityAdmin::OrganizationTelemetryRule`](../templatereference/aws-resource-observabilityadmin-organizationtelemetryrule.md)

Yes

Yes

No

[`AWS::ObservabilityAdmin::S3TableIntegration`](../templatereference/aws-resource-observabilityadmin-s3tableintegration.md)

Yes

Yes

No

[`AWS::ObservabilityAdmin::TelemetryEnrichment`](../templatereference/aws-resource-observabilityadmin-telemetryenrichment.md)

Yes

Yes

No

[`AWS::ObservabilityAdmin::TelemetryPipelines`](../templatereference/aws-resource-observabilityadmin-telemetrypipelines.md)

Yes

Yes

No

[`AWS::ObservabilityAdmin::TelemetryRule`](../templatereference/aws-resource-observabilityadmin-telemetryrule.md)

Yes

Yes

No

[`AWS::Omics::AnnotationStore`](../templatereference/aws-resource-omics-annotationstore.md)

Yes

Yes

Yes

[`AWS::Omics::ReferenceStore`](../templatereference/aws-resource-omics-referencestore.md)

Yes

Yes

Yes

[`AWS::Omics::RunGroup`](../templatereference/aws-resource-omics-rungroup.md)

Yes

Yes

Yes

[`AWS::Omics::SequenceStore`](../templatereference/aws-resource-omics-sequencestore.md)

Yes

Yes

Yes

[`AWS::Omics::VariantStore`](../templatereference/aws-resource-omics-variantstore.md)

Yes

Yes

Yes

[`AWS::Omics::Workflow`](../templatereference/aws-resource-omics-workflow.md)

Yes

Yes

Yes

[`AWS::Omics::WorkflowVersion`](../templatereference/aws-resource-omics-workflowversion.md)

Yes

Yes

No

[`AWS::OpenSearchServerless::AccessPolicy`](../templatereference/aws-resource-opensearchserverless-accesspolicy.md)

Yes

Yes

Yes

[`AWS::OpenSearchServerless::Collection`](../templatereference/aws-resource-opensearchserverless-collection.md)

Yes

Yes

No

[`AWS::OpenSearchServerless::CollectionGroup`](../templatereference/aws-resource-opensearchserverless-collectiongroup.md)

Yes

Yes

No

[`AWS::OpenSearchServerless::Index`](../templatereference/aws-resource-opensearchserverless-index.md)

Yes

Yes

No

[`AWS::OpenSearchServerless::LifecyclePolicy`](../templatereference/aws-resource-opensearchserverless-lifecyclepolicy.md)

Yes

Yes

No

[`AWS::OpenSearchServerless::SecurityConfig`](../templatereference/aws-resource-opensearchserverless-securityconfig.md)

Yes

Yes

Yes

[`AWS::OpenSearchServerless::SecurityPolicy`](../templatereference/aws-resource-opensearchserverless-securitypolicy.md)

Yes

Yes

Yes

[`AWS::OpenSearchServerless::VpcEndpoint`](../templatereference/aws-resource-opensearchserverless-vpcendpoint.md)

Yes

Yes

Yes

[`AWS::OpenSearchService::Application`](../templatereference/aws-resource-opensearchservice-application.md)

Yes

Yes

No

[`AWS::OpenSearchService::Domain`](../templatereference/aws-resource-opensearchservice-domain.md)

Yes

Yes

No

[`AWS::Organizations::Account`](../templatereference/aws-resource-organizations-account.md)

Yes

Yes

Yes

[`AWS::Organizations::Organization`](../templatereference/aws-resource-organizations-organization.md)

Yes

Yes

Yes

[`AWS::Organizations::OrganizationalUnit`](../templatereference/aws-resource-organizations-organizationalunit.md)

Yes

Yes

Yes

[`AWS::Organizations::Policy`](../templatereference/aws-resource-organizations-policy.md)

Yes

Yes

Yes

[`AWS::Organizations::ResourcePolicy`](../templatereference/aws-resource-organizations-resourcepolicy.md)

Yes

Yes

Yes

[`AWS::PCAConnectorAD::Connector`](../templatereference/aws-resource-pcaconnectorad-connector.md)

Yes

Yes

Yes

[`AWS::PCAConnectorAD::DirectoryRegistration`](../templatereference/aws-resource-pcaconnectorad-directoryregistration.md)

Yes

Yes

Yes

[`AWS::PCAConnectorAD::ServicePrincipalName`](../templatereference/aws-resource-pcaconnectorad-serviceprincipalname.md)

Yes

Yes

Yes

[`AWS::PCAConnectorAD::Template`](../templatereference/aws-resource-pcaconnectorad-template.md)

Yes

Yes

Yes

[`AWS::PCAConnectorAD::TemplateGroupAccessControlEntry`](../templatereference/aws-resource-pcaconnectorad-templategroupaccesscontrolentry.md)

Yes

Yes

Yes

[`AWS::PCAConnectorSCEP::Challenge`](../templatereference/aws-resource-pcaconnectorscep-challenge.md)

Yes

Yes

No

[`AWS::PCAConnectorSCEP::Connector`](../templatereference/aws-resource-pcaconnectorscep-connector.md)

Yes

Yes

No

[`AWS::PCS::Cluster`](../templatereference/aws-resource-pcs-cluster.md)

Yes

Yes

No

[`AWS::PCS::ComputeNodeGroup`](../templatereference/aws-resource-pcs-computenodegroup.md)

Yes

Yes

No

[`AWS::PCS::Queue`](../templatereference/aws-resource-pcs-queue.md)

Yes

Yes

No

[`AWS::Panorama::ApplicationInstance`](../templatereference/aws-resource-panorama-applicationinstance.md)

Yes

Yes

Yes

[`AWS::Panorama::Package`](../templatereference/aws-resource-panorama-package.md)

Yes

Yes

Yes

[`AWS::Panorama::PackageVersion`](../templatereference/aws-resource-panorama-packageversion.md)

Yes

Yes

No

[`AWS::PaymentCryptography::Alias`](../templatereference/aws-resource-paymentcryptography-alias.md)

Yes

Yes

No

[`AWS::PaymentCryptography::Key`](../templatereference/aws-resource-paymentcryptography-key.md)

Yes

Yes

No

[`AWS::Personalize::Dataset`](../templatereference/aws-resource-personalize-dataset.md)

Yes

Yes

Yes

[`AWS::Personalize::DatasetGroup`](../templatereference/aws-resource-personalize-datasetgroup.md)

Yes

Yes

Yes

[`AWS::Personalize::Schema`](../templatereference/aws-resource-personalize-schema.md)

Yes

Yes

Yes

[`AWS::Personalize::Solution`](../templatereference/aws-resource-personalize-solution.md)

Yes

Yes

No

[`AWS::Pinpoint::InAppTemplate`](../templatereference/aws-resource-pinpoint-inapptemplate.md)

Yes

Yes

Yes

[`AWS::Pipes::Pipe`](../templatereference/aws-resource-pipes-pipe.md)

Yes

Yes

Yes

[`AWS::Proton::EnvironmentAccountConnection`](../templatereference/aws-resource-proton-environmentaccountconnection.md)

Yes

Yes

Yes

[`AWS::Proton::EnvironmentTemplate`](../templatereference/aws-resource-proton-environmenttemplate.md)

Yes

Yes

Yes

[`AWS::Proton::ServiceTemplate`](../templatereference/aws-resource-proton-servicetemplate.md)

Yes

Yes

Yes

[`AWS::QBusiness::Application`](../templatereference/aws-resource-qbusiness-application.md)

Yes

Yes

No

[`AWS::QBusiness::DataAccessor`](../templatereference/aws-resource-qbusiness-dataaccessor.md)

Yes

Yes

No

[`AWS::QBusiness::DataSource`](../templatereference/aws-resource-qbusiness-datasource.md)

Yes

Yes

No

[`AWS::QBusiness::Index`](../templatereference/aws-resource-qbusiness-index.md)

Yes

Yes

No

[`AWS::QBusiness::Permission`](../templatereference/aws-resource-qbusiness-permission.md)

Yes

Yes

No

[`AWS::QBusiness::Plugin`](../templatereference/aws-resource-qbusiness-plugin.md)

Yes

Yes

No

[`AWS::QBusiness::Retriever`](../templatereference/aws-resource-qbusiness-retriever.md)

Yes

Yes

No

[`AWS::QBusiness::WebExperience`](../templatereference/aws-resource-qbusiness-webexperience.md)

Yes

Yes

No

[`AWS::QLDB::Stream`](../templatereference/aws-resource-qldb-stream.md)

Yes

Yes

Yes

[`AWS::QuickSight::ActionConnector`](../templatereference/aws-resource-quicksight-actionconnector.md)

Yes

Yes

No

[`AWS::QuickSight::Analysis`](../templatereference/aws-resource-quicksight-analysis.md)

Yes

Yes

No

[`AWS::QuickSight::CustomPermissions`](../templatereference/aws-resource-quicksight-custompermissions.md)

Yes

Yes

No

[`AWS::QuickSight::Dashboard`](../templatereference/aws-resource-quicksight-dashboard.md)

Yes

Yes

No

[`AWS::QuickSight::DataSet`](../templatereference/aws-resource-quicksight-dataset.md)

Yes

Yes

Yes

[`AWS::QuickSight::DataSource`](../templatereference/aws-resource-quicksight-datasource.md)

Yes

Yes

Yes

[`AWS::QuickSight::Folder`](../templatereference/aws-resource-quicksight-folder.md)

Yes

Yes

No

[`AWS::QuickSight::RefreshSchedule`](../templatereference/aws-resource-quicksight-refreshschedule.md)

Yes

Yes

Yes

[`AWS::QuickSight::Template`](../templatereference/aws-resource-quicksight-template.md)

Yes

Yes

No

[`AWS::QuickSight::Theme`](../templatereference/aws-resource-quicksight-theme.md)

Yes

Yes

No

[`AWS::QuickSight::Topic`](../templatereference/aws-resource-quicksight-topic.md)

Yes

Yes

Yes

[`AWS::QuickSight::VPCConnection`](../templatereference/aws-resource-quicksight-vpcconnection.md)

Yes

Yes

Yes

[`AWS::RAM::Permission`](../templatereference/aws-resource-ram-permission.md)

Yes

Yes

No

[`AWS::RAM::ResourceShare`](../templatereference/aws-resource-ram-resourceshare.md)

Yes

Yes

No

[`AWS::RDS::CustomDBEngineVersion`](../templatereference/aws-resource-rds-customdbengineversion.md)

Yes

Yes

No

[`AWS::RDS::DBCluster`](../templatereference/aws-resource-rds-dbcluster.md)

Yes

Yes

Yes

[`AWS::RDS::DBClusterParameterGroup`](../templatereference/aws-resource-rds-dbclusterparametergroup.md)

Yes

Yes

No

[`AWS::RDS::DBInstance`](../templatereference/aws-resource-rds-dbinstance.md)

Yes

Yes

Yes

[`AWS::RDS::DBParameterGroup`](../templatereference/aws-resource-rds-dbparametergroup.md)

Yes

Yes

No

[`AWS::RDS::DBProxy`](../templatereference/aws-resource-rds-dbproxy.md)

Yes

Yes

Yes

[`AWS::RDS::DBProxyEndpoint`](../templatereference/aws-resource-rds-dbproxyendpoint.md)

Yes

Yes

Yes

[`AWS::RDS::DBProxyTargetGroup`](../templatereference/aws-resource-rds-dbproxytargetgroup.md)

Yes

Yes

Yes

[`AWS::RDS::DBShardGroup`](../templatereference/aws-resource-rds-dbshardgroup.md)

Yes

Yes

No

[`AWS::RDS::DBSubnetGroup`](../templatereference/aws-resource-rds-dbsubnetgroup.md)

Yes

Yes

Yes

[`AWS::RDS::EventSubscription`](../templatereference/aws-resource-rds-eventsubscription.md)

Yes

Yes

Yes

[`AWS::RDS::GlobalCluster`](../templatereference/aws-resource-rds-globalcluster.md)

Yes

Yes

Yes

[`AWS::RDS::Integration`](../templatereference/aws-resource-rds-integration.md)

Yes

Yes

No

[`AWS::RDS::OptionGroup`](../templatereference/aws-resource-rds-optiongroup.md)

Yes

Yes

Yes

[`AWS::RTBFabric::InboundExternalLink`](../templatereference/aws-resource-rtbfabric-inboundexternallink.md)

Yes

Yes

No

[`AWS::RTBFabric::Link`](../templatereference/aws-resource-rtbfabric-link.md)

Yes

Yes

No

[`AWS::RTBFabric::OutboundExternalLink`](../templatereference/aws-resource-rtbfabric-outboundexternallink.md)

Yes

Yes

No

[`AWS::RTBFabric::RequesterGateway`](../templatereference/aws-resource-rtbfabric-requestergateway.md)

Yes

Yes

No

[`AWS::RTBFabric::ResponderGateway`](../templatereference/aws-resource-rtbfabric-respondergateway.md)

Yes

Yes

No

[`AWS::RUM::AppMonitor`](../templatereference/aws-resource-rum-appmonitor.md)

Yes

Yes

Yes

[`AWS::Rbin::Rule`](../templatereference/aws-resource-rbin-rule.md)

Yes

Yes

No

[`AWS::Redshift::Cluster`](../templatereference/aws-resource-redshift-cluster.md)

Yes

Yes

Yes

[`AWS::Redshift::ClusterParameterGroup`](../templatereference/aws-resource-redshift-clusterparametergroup.md)

Yes

Yes

No

[`AWS::Redshift::ClusterSubnetGroup`](../templatereference/aws-resource-redshift-clustersubnetgroup.md)

Yes

Yes

No

[`AWS::Redshift::EndpointAccess`](../templatereference/aws-resource-redshift-endpointaccess.md)

Yes

Yes

Yes

[`AWS::Redshift::EndpointAuthorization`](../templatereference/aws-resource-redshift-endpointauthorization.md)

Yes

Yes

Yes

[`AWS::Redshift::EventSubscription`](../templatereference/aws-resource-redshift-eventsubscription.md)

Yes

Yes

No

[`AWS::Redshift::Integration`](../templatereference/aws-resource-redshift-integration.md)

Yes

Yes

No

[`AWS::Redshift::ScheduledAction`](../templatereference/aws-resource-redshift-scheduledaction.md)

Yes

Yes

Yes

[`AWS::RedshiftServerless::Namespace`](../templatereference/aws-resource-redshiftserverless-namespace.md)

Yes

Yes

No

[`AWS::RedshiftServerless::Snapshot`](../templatereference/aws-resource-redshiftserverless-snapshot.md)

Yes

Yes

No

[`AWS::RedshiftServerless::Workgroup`](../templatereference/aws-resource-redshiftserverless-workgroup.md)

Yes

Yes

No

[`AWS::RefactorSpaces::Application`](../templatereference/aws-resource-refactorspaces-application.md)

Yes

Yes

Yes

[`AWS::RefactorSpaces::Environment`](../templatereference/aws-resource-refactorspaces-environment.md)

Yes

Yes

Yes

[`AWS::RefactorSpaces::Route`](../templatereference/aws-resource-refactorspaces-route.md)

Yes

Yes

Yes

[`AWS::RefactorSpaces::Service`](../templatereference/aws-resource-refactorspaces-service.md)

Yes

Yes

Yes

[`AWS::Rekognition::Collection`](../templatereference/aws-resource-rekognition-collection.md)

Yes

Yes

Yes

[`AWS::Rekognition::Project`](../templatereference/aws-resource-rekognition-project.md)

Yes

Yes

Yes

[`AWS::Rekognition::StreamProcessor`](../templatereference/aws-resource-rekognition-streamprocessor.md)

Yes

Yes

Yes

[`AWS::ResilienceHub::App`](../templatereference/aws-resource-resiliencehub-app.md)

Yes

Yes

Yes

[`AWS::ResilienceHub::ResiliencyPolicy`](../templatereference/aws-resource-resiliencehub-resiliencypolicy.md)

Yes

Yes

Yes

[`AWS::ResourceExplorer2::DefaultViewAssociation`](../templatereference/aws-resource-resourceexplorer2-defaultviewassociation.md)

Yes

Yes

No

[`AWS::ResourceExplorer2::Index`](../templatereference/aws-resource-resourceexplorer2-index.md)

Yes

Yes

Yes

[`AWS::ResourceExplorer2::View`](../templatereference/aws-resource-resourceexplorer2-view.md)

Yes

Yes

Yes

[`AWS::ResourceGroups::Group`](../templatereference/aws-resource-resourcegroups-group.md)

Yes

Yes

Yes

[`AWS::ResourceGroups::TagSyncTask`](../templatereference/aws-resource-resourcegroups-tagsynctask.md)

Yes

Yes

No

[`AWS::RolesAnywhere::CRL`](../templatereference/aws-resource-rolesanywhere-crl.md)

Yes

Yes

Yes

[`AWS::RolesAnywhere::Profile`](../templatereference/aws-resource-rolesanywhere-profile.md)

Yes

Yes

Yes

[`AWS::RolesAnywhere::TrustAnchor`](../templatereference/aws-resource-rolesanywhere-trustanchor.md)

Yes

Yes

No

[`AWS::Route53::CidrCollection`](../templatereference/aws-resource-route53-cidrcollection.md)

Yes

Yes

Yes

[`AWS::Route53::DNSSEC`](../templatereference/aws-resource-route53-dnssec.md)

Yes

Yes

Yes

[`AWS::Route53::HealthCheck`](../templatereference/aws-resource-route53-healthcheck.md)

Yes

Yes

Yes

[`AWS::Route53::HostedZone`](../templatereference/aws-resource-route53-hostedzone.md)

Yes

Yes

Yes

[`AWS::Route53::KeySigningKey`](../templatereference/aws-resource-route53-keysigningkey.md)

Yes

Yes

Yes

[`AWS::Route53GlobalResolver::AccessSource`](../templatereference/aws-resource-route53globalresolver-accesssource.md)

Yes

Yes

No

[`AWS::Route53GlobalResolver::AccessToken`](../templatereference/aws-resource-route53globalresolver-accesstoken.md)

Yes

Yes

No

[`AWS::Route53GlobalResolver::DnsView`](../templatereference/aws-resource-route53globalresolver-dnsview.md)

Yes

Yes

No

[`AWS::Route53GlobalResolver::FirewallDomainList`](../templatereference/aws-resource-route53globalresolver-firewalldomainlist.md)

Yes

Yes

No

[`AWS::Route53GlobalResolver::FirewallRule`](../templatereference/aws-resource-route53globalresolver-firewallrule.md)

Yes

Yes

No

[`AWS::Route53GlobalResolver::GlobalResolver`](../templatereference/aws-resource-route53globalresolver-globalresolver.md)

Yes

Yes

No

[`AWS::Route53GlobalResolver::HostedZoneAssociation`](../templatereference/aws-resource-route53globalresolver-hostedzoneassociation.md)

Yes

Yes

No

[`AWS::Route53Profiles::Profile`](../templatereference/aws-resource-route53profiles-profile.md)

Yes

Yes

No

[`AWS::Route53Profiles::ProfileAssociation`](../templatereference/aws-resource-route53profiles-profileassociation.md)

Yes

Yes

No

[`AWS::Route53Profiles::ProfileResourceAssociation`](../templatereference/aws-resource-route53profiles-profileresourceassociation.md)

Yes

Yes

No

[`AWS::Route53RecoveryControl::Cluster`](../templatereference/aws-resource-route53recoverycontrol-cluster.md)

Yes

Yes

Yes

[`AWS::Route53RecoveryControl::ControlPanel`](../templatereference/aws-resource-route53recoverycontrol-controlpanel.md)

Yes

Yes

Yes

[`AWS::Route53RecoveryControl::RoutingControl`](../templatereference/aws-resource-route53recoverycontrol-routingcontrol.md)

Yes

Yes

Yes

[`AWS::Route53RecoveryControl::SafetyRule`](../templatereference/aws-resource-route53recoverycontrol-safetyrule.md)

Yes

Yes

Yes

[`AWS::Route53RecoveryReadiness::Cell`](../templatereference/aws-resource-route53recoveryreadiness-cell.md)

Yes

Yes

Yes

[`AWS::Route53RecoveryReadiness::ReadinessCheck`](../templatereference/aws-resource-route53recoveryreadiness-readinesscheck.md)

Yes

Yes

Yes

[`AWS::Route53RecoveryReadiness::RecoveryGroup`](../templatereference/aws-resource-route53recoveryreadiness-recoverygroup.md)

Yes

Yes

Yes

[`AWS::Route53RecoveryReadiness::ResourceSet`](../templatereference/aws-resource-route53recoveryreadiness-resourceset.md)

Yes

Yes

Yes

[`AWS::Route53Resolver::FirewallDomainList`](../templatereference/aws-resource-route53resolver-firewalldomainlist.md)

Yes

Yes

No

[`AWS::Route53Resolver::FirewallRuleGroup`](../templatereference/aws-resource-route53resolver-firewallrulegroup.md)

Yes

Yes

Yes

[`AWS::Route53Resolver::FirewallRuleGroupAssociation`](../templatereference/aws-resource-route53resolver-firewallrulegroupassociation.md)

Yes

Yes

Yes

[`AWS::Route53Resolver::OutpostResolver`](../templatereference/aws-resource-route53resolver-outpostresolver.md)

Yes

Yes

Yes

[`AWS::Route53Resolver::ResolverConfig`](../templatereference/aws-resource-route53resolver-resolverconfig.md)

Yes

Yes

Yes

[`AWS::Route53Resolver::ResolverDNSSECConfig`](../templatereference/aws-resource-route53resolver-resolverdnssecconfig.md)

Yes

Yes

Yes

[`AWS::Route53Resolver::ResolverEndpoint`](../templatereference/aws-resource-route53resolver-resolverendpoint.md)

Yes

Yes

No

[`AWS::Route53Resolver::ResolverQueryLoggingConfig`](../templatereference/aws-resource-route53resolver-resolverqueryloggingconfig.md)

Yes

Yes

Yes

[`AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation`](../templatereference/aws-resource-route53resolver-resolverqueryloggingconfigassociation.md)

Yes

Yes

Yes

[`AWS::Route53Resolver::ResolverRule`](../templatereference/aws-resource-route53resolver-resolverrule.md)

Yes

Yes

No

[`AWS::Route53Resolver::ResolverRuleAssociation`](../templatereference/aws-resource-route53resolver-resolverruleassociation.md)

Yes

Yes

Yes

[`AWS::S3::AccessGrant`](../templatereference/aws-resource-s3-accessgrant.md)

Yes

Yes

No

[`AWS::S3::AccessGrantsInstance`](../templatereference/aws-resource-s3-accessgrantsinstance.md)

Yes

Yes

No

[`AWS::S3::AccessGrantsLocation`](../templatereference/aws-resource-s3-accessgrantslocation.md)

Yes

Yes

No

[`AWS::S3::AccessPoint`](../templatereference/aws-resource-s3-accesspoint.md)

Yes

Yes

Yes

[`AWS::S3::Bucket`](../templatereference/aws-resource-s3-bucket.md)

Yes

Yes

Yes

[`AWS::S3::BucketPolicy`](../templatereference/aws-resource-s3-bucketpolicy.md)

Yes

No

Yes

[`AWS::S3::MultiRegionAccessPoint`](../templatereference/aws-resource-s3-multiregionaccesspoint.md)

Yes

Yes

Yes

[`AWS::S3::MultiRegionAccessPointPolicy`](../templatereference/aws-resource-s3-multiregionaccesspointpolicy.md)

Yes

No

No

[`AWS::S3::StorageLens`](../templatereference/aws-resource-s3-storagelens.md)

Yes

Yes

Yes

[`AWS::S3::StorageLensGroup`](../templatereference/aws-resource-s3-storagelensgroup.md)

Yes

Yes

No

[`AWS::S3Express::AccessPoint`](../templatereference/aws-resource-s3express-accesspoint.md)

Yes

Yes

No

[`AWS::S3Express::BucketPolicy`](../templatereference/aws-resource-s3express-bucketpolicy.md)

Yes

Yes

No

[`AWS::S3Express::DirectoryBucket`](../templatereference/aws-resource-s3express-directorybucket.md)

Yes

Yes

No

[`AWS::S3ObjectLambda::AccessPoint`](../templatereference/aws-resource-s3objectlambda-accesspoint.md)

Yes

Yes

Yes

[`AWS::S3ObjectLambda::AccessPointPolicy`](../templatereference/aws-resource-s3objectlambda-accesspointpolicy.md)

Yes

Yes

No

[`AWS::S3Outposts::AccessPoint`](../templatereference/aws-resource-s3outposts-accesspoint.md)

Yes

Yes

Yes

[`AWS::S3Outposts::Bucket`](../templatereference/aws-resource-s3outposts-bucket.md)

Yes

Yes

Yes

[`AWS::S3Outposts::BucketPolicy`](../templatereference/aws-resource-s3outposts-bucketpolicy.md)

Yes

Yes

No

[`AWS::S3Outposts::Endpoint`](../templatereference/aws-resource-s3outposts-endpoint.md)

Yes

Yes

Yes

[`AWS::S3Tables::Namespace`](../templatereference/aws-resource-s3tables-namespace.md)

Yes

Yes

No

[`AWS::S3Tables::Table`](../templatereference/aws-resource-s3tables-table.md)

Yes

Yes

No

[`AWS::S3Tables::TableBucket`](../templatereference/aws-resource-s3tables-tablebucket.md)

Yes

Yes

No

[`AWS::S3Tables::TableBucketPolicy`](../templatereference/aws-resource-s3tables-tablebucketpolicy.md)

Yes

Yes

No

[`AWS::S3Tables::TablePolicy`](../templatereference/aws-resource-s3tables-tablepolicy.md)

Yes

Yes

No

[`AWS::S3Vectors::Index`](../templatereference/aws-resource-s3vectors-index.md)

Yes

Yes

No

[`AWS::S3Vectors::VectorBucket`](../templatereference/aws-resource-s3vectors-vectorbucket.md)

Yes

Yes

No

[`AWS::S3Vectors::VectorBucketPolicy`](../templatereference/aws-resource-s3vectors-vectorbucketpolicy.md)

Yes

Yes

No

[`AWS::SES::ConfigurationSet`](../templatereference/aws-resource-ses-configurationset.md)

Yes

Yes

No

[`AWS::SES::ConfigurationSetEventDestination`](../templatereference/aws-resource-ses-configurationseteventdestination.md)

No

Yes

No

[`AWS::SES::ContactList`](../templatereference/aws-resource-ses-contactlist.md)

Yes

Yes

No

[`AWS::SES::CustomVerificationEmailTemplate`](../templatereference/aws-resource-ses-customverificationemailtemplate.md)

Yes

Yes

No

[`AWS::SES::DedicatedIpPool`](../templatereference/aws-resource-ses-dedicatedippool.md)

Yes

Yes

Yes

[`AWS::SES::EmailIdentity`](../templatereference/aws-resource-ses-emailidentity.md)

Yes

Yes

Yes

[`AWS::SES::MailManagerAddonInstance`](../templatereference/aws-resource-ses-mailmanageraddoninstance.md)

Yes

Yes

No

[`AWS::SES::MailManagerAddonSubscription`](../templatereference/aws-resource-ses-mailmanageraddonsubscription.md)

Yes

Yes

No

[`AWS::SES::MailManagerAddressList`](../templatereference/aws-resource-ses-mailmanageraddresslist.md)

Yes

Yes

No

[`AWS::SES::MailManagerArchive`](../templatereference/aws-resource-ses-mailmanagerarchive.md)

Yes

Yes

No

[`AWS::SES::MailManagerIngressPoint`](../templatereference/aws-resource-ses-mailmanageringresspoint.md)

Yes

Yes

No

[`AWS::SES::MailManagerRelay`](../templatereference/aws-resource-ses-mailmanagerrelay.md)

Yes

Yes

No

[`AWS::SES::MailManagerRuleSet`](../templatereference/aws-resource-ses-mailmanagerruleset.md)

Yes

Yes

No

[`AWS::SES::MailManagerTrafficPolicy`](../templatereference/aws-resource-ses-mailmanagertrafficpolicy.md)

Yes

Yes

No

[`AWS::SES::MultiRegionEndpoint`](../templatereference/aws-resource-ses-multiregionendpoint.md)

Yes

Yes

No

[`AWS::SES::Template`](../templatereference/aws-resource-ses-template.md)

Yes

Yes

No

[`AWS::SES::Tenant`](../templatereference/aws-resource-ses-tenant.md)

Yes

Yes

No

[`AWS::SES::VdmAttributes`](../templatereference/aws-resource-ses-vdmattributes.md)

Yes

Yes

No

[`AWS::SMSVOICE::ConfigurationSet`](../templatereference/aws-resource-smsvoice-configurationset.md)

Yes

Yes

No

[`AWS::SMSVOICE::OptOutList`](../templatereference/aws-resource-smsvoice-optoutlist.md)

Yes

Yes

No

[`AWS::SMSVOICE::PhoneNumber`](../templatereference/aws-resource-smsvoice-phonenumber.md)

Yes

Yes

No

[`AWS::SMSVOICE::Pool`](../templatereference/aws-resource-smsvoice-pool.md)

Yes

Yes

No

[`AWS::SMSVOICE::ProtectConfiguration`](../templatereference/aws-resource-smsvoice-protectconfiguration.md)

Yes

Yes

No

[`AWS::SMSVOICE::ResourcePolicy`](../templatereference/aws-resource-smsvoice-resourcepolicy.md)

Yes

Yes

No

[`AWS::SMSVOICE::SenderId`](../templatereference/aws-resource-smsvoice-senderid.md)

Yes

Yes

No

[`AWS::SNS::Subscription`](../templatereference/aws-resource-sns-subscription.md)

Yes

No

No

[`AWS::SNS::Topic`](../templatereference/aws-resource-sns-topic.md)

Yes

Yes

Yes

[`AWS::SNS::TopicInlinePolicy`](../templatereference/aws-resource-sns-topicinlinepolicy.md)

Yes

Yes

No

[`AWS::SQS::Queue`](../templatereference/aws-resource-sqs-queue.md)

Yes

Yes

Yes

[`AWS::SQS::QueueInlinePolicy`](../templatereference/aws-resource-sqs-queueinlinepolicy.md)

Yes

Yes

No

[`AWS::SSM::Association`](../templatereference/aws-resource-ssm-association.md)

Yes

Yes

Yes

[`AWS::SSM::Document`](../templatereference/aws-resource-ssm-document.md)

Yes

No

No

[`AWS::SSM::MaintenanceWindow`](../templatereference/aws-resource-ssm-maintenancewindow.md)

Yes

Yes

No

[`AWS::SSM::MaintenanceWindowTarget`](../templatereference/aws-resource-ssm-maintenancewindowtarget.md)

Yes

Yes

No

[`AWS::SSM::MaintenanceWindowTask`](../templatereference/aws-resource-ssm-maintenancewindowtask.md)

Yes

Yes

No

[`AWS::SSM::Parameter`](../templatereference/aws-resource-ssm-parameter.md)

Yes

Yes

No

[`AWS::SSM::PatchBaseline`](../templatereference/aws-resource-ssm-patchbaseline.md)

Yes

Yes

No

[`AWS::SSM::ResourceDataSync`](../templatereference/aws-resource-ssm-resourcedatasync.md)

Yes

Yes

No

[`AWS::SSM::ResourcePolicy`](../templatereference/aws-resource-ssm-resourcepolicy.md)

Yes

Yes

Yes

[`AWS::SSMContacts::Contact`](../templatereference/aws-resource-ssmcontacts-contact.md)

Yes

Yes

Yes

[`AWS::SSMContacts::ContactChannel`](../templatereference/aws-resource-ssmcontacts-contactchannel.md)

Yes

Yes

Yes

[`AWS::SSMContacts::Plan`](../templatereference/aws-resource-ssmcontacts-plan.md)

Yes

Yes

No

[`AWS::SSMContacts::Rotation`](../templatereference/aws-resource-ssmcontacts-rotation.md)

Yes

Yes

No

[`AWS::SSMGuiConnect::Preferences`](../templatereference/aws-resource-ssmguiconnect-preferences.md)

Yes

Yes

No

[`AWS::SSMIncidents::ReplicationSet`](../templatereference/aws-resource-ssmincidents-replicationset.md)

Yes

Yes

Yes

[`AWS::SSMIncidents::ResponsePlan`](../templatereference/aws-resource-ssmincidents-responseplan.md)

Yes

Yes

Yes

[`AWS::SSMQuickSetup::ConfigurationManager`](../templatereference/aws-resource-ssmquicksetup-configurationmanager.md)

Yes

Yes

No

[`AWS::SSMQuickSetup::LifecycleAutomation`](../templatereference/aws-resource-ssmquicksetup-lifecycleautomation.md)

Yes

Yes

No

[`AWS::SSO::Application`](../templatereference/aws-resource-sso-application.md)

Yes

Yes

No

[`AWS::SSO::ApplicationAssignment`](../templatereference/aws-resource-sso-applicationassignment.md)

Yes

Yes

No

[`AWS::SSO::Assignment`](../templatereference/aws-resource-sso-assignment.md)

Yes

Yes

No

[`AWS::SSO::Instance`](../templatereference/aws-resource-sso-instance.md)

Yes

Yes

No

[`AWS::SSO::InstanceAccessControlAttributeConfiguration`](../templatereference/aws-resource-sso-instanceaccesscontrolattributeconfiguration.md)

Yes

Yes

Yes

[`AWS::SageMaker::App`](../templatereference/aws-resource-sagemaker-app.md)

Yes

Yes

No

[`AWS::SageMaker::AppImageConfig`](../templatereference/aws-resource-sagemaker-appimageconfig.md)

Yes

Yes

No

[`AWS::SageMaker::Cluster`](../templatereference/aws-resource-sagemaker-cluster.md)

Yes

Yes

No

[`AWS::SageMaker::DataQualityJobDefinition`](../templatereference/aws-resource-sagemaker-dataqualityjobdefinition.md)

Yes

Yes

No

[`AWS::SageMaker::Device`](../templatereference/aws-resource-sagemaker-device.md)

Yes

Yes

No

[`AWS::SageMaker::DeviceFleet`](../templatereference/aws-resource-sagemaker-devicefleet.md)

Yes

Yes

No

[`AWS::SageMaker::Domain`](../templatereference/aws-resource-sagemaker-domain.md)

Yes

Yes

No

[`AWS::SageMaker::Endpoint`](../templatereference/aws-resource-sagemaker-endpoint.md)

No

Yes

No

[`AWS::SageMaker::FeatureGroup`](../templatereference/aws-resource-sagemaker-featuregroup.md)

Yes

Yes

Yes

[`AWS::SageMaker::Image`](../templatereference/aws-resource-sagemaker-image.md)

Yes

Yes

Yes

[`AWS::SageMaker::ImageVersion`](../templatereference/aws-resource-sagemaker-imageversion.md)

Yes

Yes

Yes

[`AWS::SageMaker::InferenceComponent`](../templatereference/aws-resource-sagemaker-inferencecomponent.md)

Yes

Yes

No

[`AWS::SageMaker::InferenceExperiment`](../templatereference/aws-resource-sagemaker-inferenceexperiment.md)

Yes

Yes

Yes

[`AWS::SageMaker::MlflowTrackingServer`](../templatereference/aws-resource-sagemaker-mlflowtrackingserver.md)

Yes

Yes

No

[`AWS::SageMaker::ModelBiasJobDefinition`](../templatereference/aws-resource-sagemaker-modelbiasjobdefinition.md)

Yes

Yes

No

[`AWS::SageMaker::ModelCard`](../templatereference/aws-resource-sagemaker-modelcard.md)

Yes

Yes

Yes

[`AWS::SageMaker::ModelExplainabilityJobDefinition`](../templatereference/aws-resource-sagemaker-modelexplainabilityjobdefinition.md)

Yes

Yes

No

[`AWS::SageMaker::ModelPackage`](../templatereference/aws-resource-sagemaker-modelpackage.md)

Yes

Yes

Yes

[`AWS::SageMaker::ModelPackageGroup`](../templatereference/aws-resource-sagemaker-modelpackagegroup.md)

Yes

Yes

Yes

[`AWS::SageMaker::ModelQualityJobDefinition`](../templatereference/aws-resource-sagemaker-modelqualityjobdefinition.md)

Yes

Yes

No

[`AWS::SageMaker::MonitoringSchedule`](../templatereference/aws-resource-sagemaker-monitoringschedule.md)

Yes

No

No

[`AWS::SageMaker::PartnerApp`](../templatereference/aws-resource-sagemaker-partnerapp.md)

Yes

Yes

No

[`AWS::SageMaker::Pipeline`](../templatereference/aws-resource-sagemaker-pipeline.md)

Yes

Yes

Yes

[`AWS::SageMaker::ProcessingJob`](../templatereference/aws-resource-sagemaker-processingjob.md)

Yes

Yes

No

[`AWS::SageMaker::Project`](../templatereference/aws-resource-sagemaker-project.md)

Yes

Yes

Yes

[`AWS::SageMaker::Space`](../templatereference/aws-resource-sagemaker-space.md)

Yes

Yes

No

[`AWS::SageMaker::StudioLifecycleConfig`](../templatereference/aws-resource-sagemaker-studiolifecycleconfig.md)

Yes

Yes

No

[`AWS::SageMaker::UserProfile`](../templatereference/aws-resource-sagemaker-userprofile.md)

Yes

Yes

No

[`AWS::Scheduler::Schedule`](../templatereference/aws-resource-scheduler-schedule.md)

Yes

Yes

Yes

[`AWS::Scheduler::ScheduleGroup`](../templatereference/aws-resource-scheduler-schedulegroup.md)

Yes

Yes

No

[`AWS::SecretsManager::ResourcePolicy`](../templatereference/aws-resource-secretsmanager-resourcepolicy.md)

Yes

No

No

[`AWS::SecretsManager::RotationSchedule`](../templatereference/aws-resource-secretsmanager-rotationschedule.md)

Yes

Yes

No

[`AWS::SecretsManager::Secret`](../templatereference/aws-resource-secretsmanager-secret.md)

Yes

Yes

Yes

[`AWS::SecretsManager::SecretTargetAttachment`](../templatereference/aws-resource-secretsmanager-secrettargetattachment.md)

Yes

No

No

[`AWS::SecurityHub::AggregatorV2`](../templatereference/aws-resource-securityhub-aggregatorv2.md)

Yes

Yes

No

[`AWS::SecurityHub::AutomationRule`](../templatereference/aws-resource-securityhub-automationrule.md)

Yes

Yes

No

[`AWS::SecurityHub::AutomationRuleV2`](../templatereference/aws-resource-securityhub-automationrulev2.md)

Yes

Yes

No

[`AWS::SecurityHub::ConfigurationPolicy`](../templatereference/aws-resource-securityhub-configurationpolicy.md)

Yes

Yes

No

[`AWS::SecurityHub::ConnectorV2`](../templatereference/aws-resource-securityhub-connectorv2.md)

Yes

Yes

No

[`AWS::SecurityHub::DelegatedAdmin`](../templatereference/aws-resource-securityhub-delegatedadmin.md)

Yes

Yes

No

[`AWS::SecurityHub::FindingAggregator`](../templatereference/aws-resource-securityhub-findingaggregator.md)

Yes

Yes

No

[`AWS::SecurityHub::Hub`](../templatereference/aws-resource-securityhub-hub.md)

Yes

Yes

No

[`AWS::SecurityHub::HubV2`](../templatereference/aws-resource-securityhub-hubv2.md)

Yes

Yes

No

[`AWS::SecurityHub::Insight`](../templatereference/aws-resource-securityhub-insight.md)

Yes

Yes

No

[`AWS::SecurityHub::OrganizationConfiguration`](../templatereference/aws-resource-securityhub-organizationconfiguration.md)

Yes

Yes

No

[`AWS::SecurityHub::PolicyAssociation`](../templatereference/aws-resource-securityhub-policyassociation.md)

Yes

Yes

No

[`AWS::SecurityHub::ProductSubscription`](../templatereference/aws-resource-securityhub-productsubscription.md)

Yes

Yes

No

[`AWS::SecurityHub::SecurityControl`](../templatereference/aws-resource-securityhub-securitycontrol.md)

Yes

Yes

No

[`AWS::SecurityHub::Standard`](../templatereference/aws-resource-securityhub-standard.md)

Yes

Yes

No

[`AWS::SecurityLake::AwsLogSource`](../templatereference/aws-resource-securitylake-awslogsource.md)

Yes

Yes

No

[`AWS::SecurityLake::DataLake`](../templatereference/aws-resource-securitylake-datalake.md)

Yes

Yes

No

[`AWS::SecurityLake::Subscriber`](../templatereference/aws-resource-securitylake-subscriber.md)

Yes

Yes

No

[`AWS::SecurityLake::SubscriberNotification`](../templatereference/aws-resource-securitylake-subscribernotification.md)

Yes

Yes

No

[`AWS::ServiceCatalog::CloudFormationProvisionedProduct`](../templatereference/aws-resource-servicecatalog-cloudformationprovisionedproduct.md)

Yes

No

No

[`AWS::ServiceCatalog::LaunchNotificationConstraint`](../templatereference/aws-resource-servicecatalog-launchnotificationconstraint.md)

Yes

Yes

No

[`AWS::ServiceCatalog::LaunchRoleConstraint`](../templatereference/aws-resource-servicecatalog-launchroleconstraint.md)

Yes

Yes

No

[`AWS::ServiceCatalog::LaunchTemplateConstraint`](../templatereference/aws-resource-servicecatalog-launchtemplateconstraint.md)

Yes

Yes

No

[`AWS::ServiceCatalog::Portfolio`](../templatereference/aws-resource-servicecatalog-portfolio.md)

Yes

Yes

No

[`AWS::ServiceCatalog::PortfolioPrincipalAssociation`](../templatereference/aws-resource-servicecatalog-portfolioprincipalassociation.md)

Yes

Yes

No

[`AWS::ServiceCatalog::PortfolioProductAssociation`](../templatereference/aws-resource-servicecatalog-portfolioproductassociation.md)

Yes

Yes

No

[`AWS::ServiceCatalog::PortfolioShare`](../templatereference/aws-resource-servicecatalog-portfolioshare.md)

Yes

No

No

[`AWS::ServiceCatalog::ResourceUpdateConstraint`](../templatereference/aws-resource-servicecatalog-resourceupdateconstraint.md)

Yes

Yes

No

[`AWS::ServiceCatalog::ServiceAction`](../templatereference/aws-resource-servicecatalog-serviceaction.md)

Yes

Yes

Yes

[`AWS::ServiceCatalog::ServiceActionAssociation`](../templatereference/aws-resource-servicecatalog-serviceactionassociation.md)

Yes

Yes

No

[`AWS::ServiceCatalog::StackSetConstraint`](../templatereference/aws-resource-servicecatalog-stacksetconstraint.md)

Yes

Yes

No

[`AWS::ServiceCatalog::TagOption`](../templatereference/aws-resource-servicecatalog-tagoption.md)

Yes

Yes

No

[`AWS::ServiceCatalog::TagOptionAssociation`](../templatereference/aws-resource-servicecatalog-tagoptionassociation.md)

Yes

Yes

No

[`AWS::ServiceCatalogAppRegistry::Application`](../templatereference/aws-resource-servicecatalogappregistry-application.md)

Yes

Yes

Yes

[`AWS::ServiceCatalogAppRegistry::AttributeGroup`](../templatereference/aws-resource-servicecatalogappregistry-attributegroup.md)

Yes

Yes

Yes

[`AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation`](../templatereference/aws-resource-servicecatalogappregistry-attributegroupassociation.md)

Yes

No

No

[`AWS::ServiceCatalogAppRegistry::ResourceAssociation`](../templatereference/aws-resource-servicecatalogappregistry-resourceassociation.md)

Yes

No

No

[`AWS::ServiceDiscovery::Service`](../templatereference/aws-resource-servicediscovery-service.md)

Yes

Yes

No

[`AWS::Shield::DRTAccess`](../templatereference/aws-resource-shield-drtaccess.md)

Yes

Yes

No

[`AWS::Shield::ProactiveEngagement`](../templatereference/aws-resource-shield-proactiveengagement.md)

Yes

Yes

No

[`AWS::Shield::Protection`](../templatereference/aws-resource-shield-protection.md)

Yes

Yes

No

[`AWS::Shield::ProtectionGroup`](../templatereference/aws-resource-shield-protectiongroup.md)

Yes

Yes

No

[`AWS::Signer::ProfilePermission`](../templatereference/aws-resource-signer-profilepermission.md)

Yes

Yes

No

[`AWS::Signer::SigningProfile`](../templatereference/aws-resource-signer-signingprofile.md)

Yes

Yes

Yes

[`AWS::SimSpaceWeaver::Simulation`](../templatereference/aws-resource-simspaceweaver-simulation.md)

Yes

Yes

Yes

[`AWS::StepFunctions::Activity`](../templatereference/aws-resource-stepfunctions-activity.md)

Yes

Yes

No

[`AWS::StepFunctions::StateMachine`](../templatereference/aws-resource-stepfunctions-statemachine.md)

Yes

Yes

Yes

[`AWS::StepFunctions::StateMachineAlias`](../templatereference/aws-resource-stepfunctions-statemachinealias.md)

Yes

Yes

Yes

[`AWS::StepFunctions::StateMachineVersion`](../templatereference/aws-resource-stepfunctions-statemachineversion.md)

Yes

Yes

Yes

[`AWS::SupportApp::AccountAlias`](../templatereference/aws-resource-supportapp-accountalias.md)

Yes

Yes

Yes

[`AWS::SupportApp::SlackChannelConfiguration`](../templatereference/aws-resource-supportapp-slackchannelconfiguration.md)

Yes

Yes

Yes

[`AWS::SupportApp::SlackWorkspaceConfiguration`](../templatereference/aws-resource-supportapp-slackworkspaceconfiguration.md)

Yes

Yes

Yes

[`AWS::Synthetics::Canary`](../templatereference/aws-resource-synthetics-canary.md)

Yes

Yes

Yes

[`AWS::Synthetics::Group`](../templatereference/aws-resource-synthetics-group.md)

Yes

Yes

Yes

[`AWS::SystemsManagerSAP::Application`](../templatereference/aws-resource-systemsmanagersap-application.md)

Yes

Yes

No

[`AWS::Timestream::Database`](../templatereference/aws-resource-timestream-database.md)

Yes

Yes

Yes

[`AWS::Timestream::InfluxDBCluster`](../templatereference/aws-resource-timestream-influxdbcluster.md)

Yes

Yes

No

[`AWS::Timestream::InfluxDBInstance`](../templatereference/aws-resource-timestream-influxdbinstance.md)

Yes

Yes

No

[`AWS::Timestream::ScheduledQuery`](../templatereference/aws-resource-timestream-scheduledquery.md)

Yes

Yes

No

[`AWS::Timestream::Table`](../templatereference/aws-resource-timestream-table.md)

Yes

Yes

Yes

[`AWS::Transfer::Agreement`](../templatereference/aws-resource-transfer-agreement.md)

Yes

Yes

Yes

[`AWS::Transfer::Certificate`](../templatereference/aws-resource-transfer-certificate.md)

Yes

Yes

Yes

[`AWS::Transfer::Connector`](../templatereference/aws-resource-transfer-connector.md)

Yes

Yes

Yes

[`AWS::Transfer::Profile`](../templatereference/aws-resource-transfer-profile.md)

Yes

Yes

Yes

[`AWS::Transfer::Server`](../templatereference/aws-resource-transfer-server.md)

Yes

Yes

No

[`AWS::Transfer::User`](../templatereference/aws-resource-transfer-user.md)

Yes

Yes

No

[`AWS::Transfer::WebApp`](../templatereference/aws-resource-transfer-webapp.md)

Yes

Yes

No

[`AWS::Transfer::Workflow`](../templatereference/aws-resource-transfer-workflow.md)

Yes

Yes

Yes

[`AWS::VerifiedPermissions::IdentitySource`](../templatereference/aws-resource-verifiedpermissions-identitysource.md)

Yes

Yes

Yes

[`AWS::VerifiedPermissions::Policy`](../templatereference/aws-resource-verifiedpermissions-policy.md)

Yes

Yes

Yes

[`AWS::VerifiedPermissions::PolicyStore`](../templatereference/aws-resource-verifiedpermissions-policystore.md)

Yes

Yes

Yes

[`AWS::VerifiedPermissions::PolicyTemplate`](../templatereference/aws-resource-verifiedpermissions-policytemplate.md)

Yes

Yes

Yes

[`AWS::VoiceID::Domain`](../templatereference/aws-resource-voiceid-domain.md)

Yes

Yes

No

[`AWS::VpcLattice::AccessLogSubscription`](../templatereference/aws-resource-vpclattice-accesslogsubscription.md)

Yes

Yes

Yes

[`AWS::VpcLattice::AuthPolicy`](../templatereference/aws-resource-vpclattice-authpolicy.md)

Yes

Yes

No

[`AWS::VpcLattice::DomainVerification`](../templatereference/aws-resource-vpclattice-domainverification.md)

Yes

Yes

No

[`AWS::VpcLattice::Listener`](../templatereference/aws-resource-vpclattice-listener.md)

Yes

Yes

Yes

[`AWS::VpcLattice::ResourceConfiguration`](../templatereference/aws-resource-vpclattice-resourceconfiguration.md)

Yes

Yes

No

[`AWS::VpcLattice::ResourceGateway`](../templatereference/aws-resource-vpclattice-resourcegateway.md)

Yes

Yes

No

[`AWS::VpcLattice::ResourcePolicy`](../templatereference/aws-resource-vpclattice-resourcepolicy.md)

Yes

Yes

No

[`AWS::VpcLattice::Rule`](../templatereference/aws-resource-vpclattice-rule.md)

Yes

Yes

Yes

[`AWS::VpcLattice::Service`](../templatereference/aws-resource-vpclattice-service.md)

Yes

Yes

Yes

[`AWS::VpcLattice::ServiceNetwork`](../templatereference/aws-resource-vpclattice-servicenetwork.md)

Yes

Yes

Yes

[`AWS::VpcLattice::ServiceNetworkResourceAssociation`](../templatereference/aws-resource-vpclattice-servicenetworkresourceassociation.md)

Yes

Yes

No

[`AWS::VpcLattice::ServiceNetworkServiceAssociation`](../templatereference/aws-resource-vpclattice-servicenetworkserviceassociation.md)

Yes

Yes

Yes

[`AWS::VpcLattice::ServiceNetworkVpcAssociation`](../templatereference/aws-resource-vpclattice-servicenetworkvpcassociation.md)

Yes

Yes

Yes

[`AWS::VpcLattice::TargetGroup`](../templatereference/aws-resource-vpclattice-targetgroup.md)

Yes

Yes

Yes

[`AWS::WAFv2::IPSet`](../templatereference/aws-resource-wafv2-ipset.md)

Yes

Yes

Yes

[`AWS::WAFv2::LoggingConfiguration`](../templatereference/aws-resource-wafv2-loggingconfiguration.md)

Yes

Yes

Yes

[`AWS::WAFv2::RegexPatternSet`](../templatereference/aws-resource-wafv2-regexpatternset.md)

Yes

Yes

Yes

[`AWS::WAFv2::RuleGroup`](../templatereference/aws-resource-wafv2-rulegroup.md)

Yes

Yes

Yes

[`AWS::WAFv2::WebACL`](../templatereference/aws-resource-wafv2-webacl.md)

Yes

Yes

Yes

[`AWS::WAFv2::WebACLAssociation`](../templatereference/aws-resource-wafv2-webaclassociation.md)

Yes

Yes

No

[`AWS::Wisdom::AIAgent`](../templatereference/aws-resource-wisdom-aiagent.md)

Yes

Yes

No

[`AWS::Wisdom::AIAgentVersion`](../templatereference/aws-resource-wisdom-aiagentversion.md)

Yes

Yes

No

[`AWS::Wisdom::AIGuardrail`](../templatereference/aws-resource-wisdom-aiguardrail.md)

Yes

Yes

No

[`AWS::Wisdom::AIGuardrailVersion`](../templatereference/aws-resource-wisdom-aiguardrailversion.md)

Yes

Yes

No

[`AWS::Wisdom::AIPrompt`](../templatereference/aws-resource-wisdom-aiprompt.md)

Yes

Yes

No

[`AWS::Wisdom::AIPromptVersion`](../templatereference/aws-resource-wisdom-aipromptversion.md)

Yes

Yes

No

[`AWS::Wisdom::Assistant`](../templatereference/aws-resource-wisdom-assistant.md)

Yes

Yes

Yes

[`AWS::Wisdom::AssistantAssociation`](../templatereference/aws-resource-wisdom-assistantassociation.md)

Yes

Yes

Yes

[`AWS::Wisdom::KnowledgeBase`](../templatereference/aws-resource-wisdom-knowledgebase.md)

Yes

Yes

Yes

[`AWS::Wisdom::MessageTemplate`](../templatereference/aws-resource-wisdom-messagetemplate.md)

Yes

Yes

No

[`AWS::Wisdom::MessageTemplateVersion`](../templatereference/aws-resource-wisdom-messagetemplateversion.md)

Yes

Yes

No

[`AWS::Wisdom::QuickResponse`](../templatereference/aws-resource-wisdom-quickresponse.md)

Yes

Yes

No

[`AWS::WorkSpaces::ConnectionAlias`](../templatereference/aws-resource-workspaces-connectionalias.md)

Yes

Yes

No

[`AWS::WorkSpaces::Workspace`](../templatereference/aws-resource-workspaces-workspace.md)

Yes

Yes

No

[`AWS::WorkSpaces::WorkspacesPool`](../templatereference/aws-resource-workspaces-workspacespool.md)

Yes

Yes

No

[`AWS::WorkSpacesThinClient::Environment`](../templatereference/aws-resource-workspacesthinclient-environment.md)

Yes

Yes

No

[`AWS::WorkSpacesWeb::BrowserSettings`](../templatereference/aws-resource-workspacesweb-browsersettings.md)

Yes

Yes

Yes

[`AWS::WorkSpacesWeb::DataProtectionSettings`](../templatereference/aws-resource-workspacesweb-dataprotectionsettings.md)

Yes

Yes

No

[`AWS::WorkSpacesWeb::IdentityProvider`](../templatereference/aws-resource-workspacesweb-identityprovider.md)

Yes

Yes

Yes

[`AWS::WorkSpacesWeb::IpAccessSettings`](../templatereference/aws-resource-workspacesweb-ipaccesssettings.md)

Yes

Yes

Yes

[`AWS::WorkSpacesWeb::NetworkSettings`](../templatereference/aws-resource-workspacesweb-networksettings.md)

Yes

Yes

Yes

[`AWS::WorkSpacesWeb::Portal`](../templatereference/aws-resource-workspacesweb-portal.md)

Yes

Yes

Yes

[`AWS::WorkSpacesWeb::SessionLogger`](../templatereference/aws-resource-workspacesweb-sessionlogger.md)

Yes

Yes

No

[`AWS::WorkSpacesWeb::TrustStore`](../templatereference/aws-resource-workspacesweb-truststore.md)

Yes

Yes

Yes

[`AWS::WorkSpacesWeb::UserAccessLoggingSettings`](../templatereference/aws-resource-workspacesweb-useraccessloggingsettings.md)

Yes

Yes

Yes

[`AWS::WorkSpacesWeb::UserSettings`](../templatereference/aws-resource-workspacesweb-usersettings.md)

Yes

Yes

Yes

[`AWS::WorkspacesInstances::Volume`](../templatereference/aws-resource-workspacesinstances-volume.md)

Yes

Yes

No

[`AWS::WorkspacesInstances::VolumeAssociation`](../templatereference/aws-resource-workspacesinstances-volumeassociation.md)

Yes

Yes

No

[`AWS::WorkspacesInstances::WorkspaceInstance`](../templatereference/aws-resource-workspacesinstances-workspaceinstance.md)

Yes

Yes

No

[`AWS::XRay::Group`](../templatereference/aws-resource-xray-group.md)

Yes

Yes

Yes

[`AWS::XRay::ResourcePolicy`](../templatereference/aws-resource-xray-resourcepolicy.md)

Yes

Yes

Yes

[`AWS::XRay::SamplingRule`](../templatereference/aws-resource-xray-samplingrule.md)

Yes

Yes

No

[`AWS::XRay::TransactionSearchConfig`](../templatereference/aws-resource-xray-transactionsearchconfig.md)

Yes

Yes

No

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Stack refactoring

Use quick-create links to create CloudFormation stacks

All content copied from https://docs.aws.amazon.com/.
