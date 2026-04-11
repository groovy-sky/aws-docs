This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::ConnectorV2 Provider

The third-party provider detail for a service configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JiraCloud" : JiraCloudProviderConfiguration,
  "ServiceNow" : ServiceNowProviderConfiguration
}

```

### YAML

```yaml

  JiraCloud:
    JiraCloudProviderConfiguration
  ServiceNow:
    ServiceNowProviderConfiguration

```

## Properties

`JiraCloud`

Details about a Jira Cloud integration.

_Required_: No

_Type_: [JiraCloudProviderConfiguration](aws-properties-securityhub-connectorv2-jiracloudproviderconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNow`

Details about a ServiceNow ITSM integration.

_Required_: No

_Type_: [ServiceNowProviderConfiguration](aws-properties-securityhub-connectorv2-servicenowproviderconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JiraCloudProviderConfiguration

ServiceNowProviderConfiguration

All content copied from https://docs.aws.amazon.com/.
