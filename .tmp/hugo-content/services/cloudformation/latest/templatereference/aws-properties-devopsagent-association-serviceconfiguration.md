This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association ServiceConfiguration

The configuration that directs how Agent Space interacts with the given service. You can specify only one
configuration type per association.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Aws" : AWSConfiguration,
  "Dynatrace" : DynatraceConfiguration,
  "EventChannel" : EventChannelConfiguration,
  "GitHub" : GitHubConfiguration,
  "GitLab" : GitLabConfiguration,
  "MCPServer" : MCPServerConfiguration,
  "MCPServerDatadog" : MCPServerDatadogConfiguration,
  "MCPServerNewRelic" : MCPServerNewRelicConfiguration,
  "MCPServerSplunk" : MCPServerSplunkConfiguration,
  "ServiceNow" : ServiceNowConfiguration,
  "Slack" : SlackConfiguration,
  "SourceAws" : SourceAwsConfiguration
}

```

### YAML

```yaml

  Aws:
    AWSConfiguration
  Dynatrace:
    DynatraceConfiguration
  EventChannel:
    EventChannelConfiguration
  GitHub:
    GitHubConfiguration
  GitLab:
    GitLabConfiguration
  MCPServer:
    MCPServerConfiguration
  MCPServerDatadog:
    MCPServerDatadogConfiguration
  MCPServerNewRelic:
    MCPServerNewRelicConfiguration
  MCPServerSplunk:
    MCPServerSplunkConfiguration
  ServiceNow:
    ServiceNowConfiguration
  Slack:
    SlackConfiguration
  SourceAws:
    SourceAwsConfiguration

```

## Properties

`Aws`

Configuration for AWS monitor account integration. Specifies the account ID, assumable role ARN,
and resources to be monitored in the primary monitoring account.

_Required_: No

_Type_: [AWSConfiguration](aws-properties-devopsagent-association-awsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dynatrace`

Configuration for Dynatrace monitoring integration. Specifies the environment ID, resources to monitor, and
webhook settings to enable the Agent Space to access Dynatrace metrics, traces, and logs.

_Required_: No

_Type_: [DynatraceConfiguration](aws-properties-devopsagent-association-dynatraceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventChannel`

Configuration for Event Channel integration. Specifies webhook settings to enable the Agent Space to receive and
process real-time events from external systems.

_Required_: No

_Type_: [EventChannelConfiguration](aws-properties-devopsagent-association-eventchannelconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GitHub`

Configuration for GitHub repository integration. Specifies the repository name, repository ID, owner, and owner
type to enable the Agent Space to access code, pull requests, and issues.

_Required_: No

_Type_: [GitHubConfiguration](aws-properties-devopsagent-association-githubconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GitLab`

Configuration for GitLab project integration. Specifies the project ID, project path, instance identifier, and
webhook settings to enable the Agent Space to access code, merge requests, and issues.

_Required_: No

_Type_: [GitLabConfiguration](aws-properties-devopsagent-association-gitlabconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MCPServer`

Configuration for custom MCP (Model Context Protocol) server integration. Specifies the server name, endpoint
URL, available tools, description, and webhook settings to enable the Agent Space to interact with custom MCP
servers.

_Required_: No

_Type_: [MCPServerConfiguration](aws-properties-devopsagent-association-mcpserverconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MCPServerDatadog`

Configuration for Datadog MCP server integration. Specifies the server name, endpoint URL, optional description,
and webhook settings to enable the Agent Space to query metrics, traces, and logs from Datadog.

_Required_: No

_Type_: [MCPServerDatadogConfiguration](aws-properties-devopsagent-association-mcpserverdatadogconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MCPServerNewRelic`

Configuration for New Relic MCP server integration. Specifies the New Relic account ID and MCP endpoint URL to
enable the Agent Space to query metrics, traces, and logs from New Relic.

_Required_: No

_Type_: [MCPServerNewRelicConfiguration](aws-properties-devopsagent-association-mcpservernewrelicconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MCPServerSplunk`

Configuration for Splunk MCP server integration. Specifies the server name, endpoint URL, optional description,
and webhook settings to enable the Agent Space to query logs, metrics, and events from Splunk.

_Required_: No

_Type_: [MCPServerSplunkConfiguration](aws-properties-devopsagent-association-mcpserversplunkconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNow`

Configuration for ServiceNow instance integration. Specifies the instance URL, instance ID, and webhook settings
to enable the Agent Space to create, update, and manage ServiceNow incidents and change requests.

_Required_: No

_Type_: [ServiceNowConfiguration](aws-properties-devopsagent-association-servicenowconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slack`

Configuration for Slack workspace integration. Specifies the workspace ID, workspace name, and transmission
targets to enable the Agent Space to send notifications to designated Slack channels.

_Required_: No

_Type_: [SlackConfiguration](aws-properties-devopsagent-association-slackconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceAws`

Configuration for AWS source account integration. Specifies the account ID, assumable role ARN,
and resources to be monitored in the source account.

_Required_: No

_Type_: [SourceAwsConfiguration](aws-properties-devopsagent-association-sourceawsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MCPServerSplunkConfiguration

ServiceNowConfiguration

All content copied from https://docs.aws.amazon.com/.
