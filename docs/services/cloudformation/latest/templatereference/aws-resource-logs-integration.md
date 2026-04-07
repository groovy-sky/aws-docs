This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Integration

Creates an integration between CloudWatch Logs and another service in this account.
Currently, only integrations with OpenSearch Service are supported, and currently you can have
only one integration in your account.

Integrating with OpenSearch Service makes it possible for you to create curated vended
logs dashboards, powered by OpenSearch Service analytics. For more information, see [Vended log\
dashboards powered by Amazon OpenSearch Service](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-OpenSearch-Dashboards.html).

You can use this operation only to create a new integration. You can't modify an existing
integration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::Integration",
  "Properties" : {
      "IntegrationName" : String,
      "IntegrationType" : String,
      "ResourceConfig" : ResourceConfig
    }
}

```

### YAML

```yaml

Type: AWS::Logs::Integration
Properties:
  IntegrationName: String
  IntegrationType: String
  ResourceConfig:
    ResourceConfig

```

## Properties

`IntegrationName`

The name of this integration.

_Required_: Yes

_Type_: String

_Pattern_: `[\.\-_/#A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IntegrationType`

The type of integration. Integrations with OpenSearch Service have the type
`OPENSEARCH`.

_Required_: Yes

_Type_: String

_Allowed values_: `OPENSEARCH`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceConfig`

This structure contains configuration details about an integration between CloudWatch
Logs and another entity.

_Required_: Yes

_Type_: [ResourceConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-integration-resourceconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`IntegrationStatus`

The current status of this integration.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

OpenSearchResourceConfig
