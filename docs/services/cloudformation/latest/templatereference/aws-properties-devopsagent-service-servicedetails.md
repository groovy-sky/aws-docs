This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service ServiceDetails

Service-specific configuration details provided during registration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dynatrace" : DynatraceServiceDetails,
  "GitLab" : GitLabDetails,
  "MCPServer" : MCPServerDetails,
  "MCPServerNewRelic" : NewRelicServiceDetails,
  "MCPServerSplunk" : MCPServerSplunkDetails,
  "ServiceNow" : ServiceNowServiceDetails
}

```

### YAML

```yaml

  Dynatrace:
    DynatraceServiceDetails
  GitLab:
    GitLabDetails
  MCPServer:
    MCPServerDetails
  MCPServerNewRelic:
    NewRelicServiceDetails
  MCPServerSplunk:
    MCPServerSplunkDetails
  ServiceNow:
    ServiceNowServiceDetails

```

## Properties

`Dynatrace`

Dynatrace service configuration.

_Required_: No

_Type_: [DynatraceServiceDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devopsagent-service-dynatraceservicedetails.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GitLab`

GitLab service configuration.

_Required_: No

_Type_: [GitLabDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devopsagent-service-gitlabdetails.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MCPServer`

Custom MCP server configuration.

_Required_: No

_Type_: [MCPServerDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devopsagent-service-mcpserverdetails.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MCPServerNewRelic`

New Relic MCP server configuration.

_Required_: No

_Type_: [NewRelicServiceDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devopsagent-service-newrelicservicedetails.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MCPServerSplunk`

Splunk MCP server configuration.

_Required_: No

_Type_: [MCPServerSplunkDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devopsagent-service-mcpserversplunkdetails.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceNow`

ServiceNow service configuration.

_Required_: No

_Type_: [ServiceNowServiceDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devopsagent-service-servicenowservicedetails.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RegisteredServiceNowDetails

ServiceNowAuthorizationConfig
